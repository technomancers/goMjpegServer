/*Package piCamera is a simple wrapper for raspivid.*/
package piCamera

import (
	"context"
	"io"
	"log"
	"os/exec"

	"bytes"
	"errors"
	"os"
	"strconv"
	"sync"
)

var jpgMagic = []byte{0xff, 0xd8, 0xff, 0xe0, 0x00, 0x10, 0x4a, 0x46, 0x49, 0x46} //This is covered here https://asecuritysite.com/forensics/jpeg

//PiCamera creates a way for code to be able to pull images from the camera live.
//You must start eh raspivid explicitly first but once started, PiCamera will have the latest image available to view.
//
//PiCamera is thread safe so many calls to GetFrame() will not break.
//Be careful as the more calls to GetFrame() the slower GetFrame() may become due to all the read locks.
type PiCamera struct {
	rwMutext  *sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
	command   *exec.Cmd
	stdOut    io.ReadCloser
	latestImg []byte
}

//New creates an instance of PiCamera.
//Width and Height are for the image size.
//ctx is the parent context. If nil a background context will be created.
//
//This creates the command raspivid with the appropriate settings.
//The stdErr of the command is redirected to os.Stderr so that one may see why the command may have failed.
func New(parentCtx context.Context, width, height int) (*PiCamera, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	if parentCtx == nil {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithCancel(parentCtx)
	}
	cmd := exec.CommandContext(ctx, "raspivid", "-cd", "MJPEG", "-t", "0", "-w", strconv.Itoa(width), "-h", strconv.Itoa(height), "-o", "-")
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return nil, err
	}
	cmd.Stderr = os.Stderr
	return &PiCamera{
		ctx:      ctx,
		cancel:   cancel,
		command:  cmd,
		stdOut:   stdOut,
		rwMutext: new(sync.RWMutex),
	}, nil
}

//Start raspivid in the background.
//Also logs the PID to os.stdOut.
func (pc *PiCamera) Start() error {
	err := pc.command.Start()
	go pc.updateLatest()
	log.Printf("PiCamera running PID: %d", pc.command.Process.Pid)
	return err
}

//GetFrame returns the latest frame from raspivid.
//If there is no frame available it will throw an error.
func (pc *PiCamera) GetFrame() ([]byte, error) {
	pc.rwMutext.RLock()
	defer pc.rwMutext.RUnlock()
	if pc.latestImg == nil {
		return nil, errors.New("Latest Image is empty")
	}
	return pc.latestImg, nil
}

//Stop the raspivid command.
//Safely stop all the commands and routines with this.
func (pc *PiCamera) Stop() {
	pc.cancel()
}

func (pc *PiCamera) updateLatest() {
	readBuff := make([]byte, 4096)         //Buffer of the currently read bytes (4 kilobytes)
	var work = make([]byte, len(jpgMagic)) //This is the currently working bytes in process data. Must be outside function as it can carry over calls
	var buffer = new(bytes.Buffer)         //The new image buffer. The one currently being processed
	for {
		select {
		case <-pc.ctx.Done():
			break
		default:
			n, err := pc.stdOut.Read(readBuff)
			if err != nil {
				log.Printf("Reading from raspivid stdOut error: %v", err)
				break
			}
			start := 0 //This is where the image starts if this data splits images
			//Read through the data
			for i := 0; i < n; i++ {
				//add byte to working bytes
				work = append(work[1:], readBuff[i])
				//If we are at the start of a new image
				if bytes.Compare(work, jpgMagic) == 0 {
					buffer.Write(readBuff[start:i]) //write what is left of the old image
					if buffer.Len() > 0 {
						end := buffer.Len() - len(jpgMagic) + 1 //figure out where the end of the previous image was
						image := buffer.Bytes()[:end]
						cpyImage := make([]byte, len(image))
						copy(cpyImage, image)
						rest := buffer.Bytes()[end:]
						//write the image to the latest
						pc.rwMutext.Lock()
						pc.latestImg = cpyImage
						pc.rwMutext.Unlock()
						buffer.Reset()     //Clear out the buffer
						buffer.Write(rest) //Include the partial image that was left back into the buffer
						start = i          //make sure to change I so that the rest of the readBuffer is cleared correctly
					}
				}
			}
			buffer.Write(readBuff[start:n]) //write to the buffer readBuffer depending on where the start was
		}
	}
}
