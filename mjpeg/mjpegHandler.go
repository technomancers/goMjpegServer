/*Package mjpeg creates a motion jpeg handler from a camera*/
package mjpeg

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
)

//FrameGetter is any structure that can return a JPEG image in bytes.
//Ideally this would be a camera of some sort.
type FrameGetter interface {
	GetFrame() ([]byte, error)
}

//Mjpeg is a http.Handler used to server up Motion JPEG.
type Mjpeg struct {
	camera FrameGetter
}

//New creates a new MJPEG instance with the given FrameGetter.
func New(cam FrameGetter) *Mjpeg {
	return &Mjpeg{cam}
}

//ServerHTTP will use the camera in Mjpeg server it over the response.
//Sets all the appropriate headers to be able to stream a Mjpeg over http.
func (m *Mjpeg) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mimeWriter := multipart.NewWriter(w)
	contentType := fmt.Sprintf("multipart/x-mixed-replace;boundary=%s", mimeWriter.Boundary())
	w.Header().Add("Content-Type", contentType)
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		partHeader := make(textproto.MIMEHeader)
		partHeader.Add("Content-Type", "image/jpeg")

		partWriter, err := mimeWriter.CreatePart(partHeader)
		if err != nil {
			log.Printf("Could not create Part Writer: %v\n", err)
			break
		}

		img, err := m.camera.GetFrame()
		if err != nil {
			log.Printf("Could not get the next frame: %v\n", err)
			break
		}
		partHeader.Add("Content-Length", strconv.Itoa(len(img)))
		if _, err = io.Copy(partWriter, bytes.NewReader(img)); err != nil {
			log.Printf("Could not write the image to the response: %v\n", err)
			break
		}
	}
}
