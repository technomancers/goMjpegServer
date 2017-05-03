/*Package piCamera is a simple wrapper for raspivid.*/
package piCamera

import (
	"context"
)

//PiCamera creates a way for code to be able to pull images from the camera live.
//You must start eh raspivid explicitly first but once started, PiCamera will have the latest image available to view.
//
//PiCamera is thread safe so many calls to GetFrame() will not break.
//Be careful as the more calls to GetFrame() the slower GetFrame() may become due to all the read locks.
type PiCamera struct {
}

//New creates an instance of PiCamera.
//Width and Height are for the image size.
//ctx is the parent context. If nil a background context will be created.
//
//This creates the command raspivid with the appropriate settings.
//The stdErr of the command is redirected to os.Stderr so that one may see why the command may have failed.
func New(parentCtx context.Context, width, height int) (*PiCamera, error) {

	return &PiCamera{}, nil
}

//Start raspivid in the background.
//Also logs the PID to os.stdOut.
func (pc *PiCamera) Start() error {
	return nil
}

//GetFrame returns the latest frame from raspivid.
//If there is no frame available it will throw an error.
func (pc *PiCamera) GetFrame() ([]byte, error) {
	return nil, nil
}

//Stop the raspivid command.
//Safely stop all the commands and routines with this.
func (pc *PiCamera) Stop() {
}
