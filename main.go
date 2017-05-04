package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/technomancers/goMjpegServer/mjpeg"
	"github.com/technomancers/goMjpegServer/piCamera"
)

const listen = ":8080"

const defaultTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8" />
		<title>{{.Title}}</title>
	</head>
	<body>
		<h1>Live Pi</h1>
		<img src="{{.Stream}}" />
	</body>
</html>`

func main() {
	t, err := template.New("homePage").Parse(defaultTemplate)
	if err != nil {
		log.Fatalf("Could not parse the template: %v\n", err)
	}
	data := struct {
		Title  string
		Stream string
	}{
		Title:  "MJPEG Server",
		Stream: "/mjpeg",
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	piArgs := piCamera.NewArgs()
	piArgs.Mode = 7
	piCamera, err := piCamera.New(ctx, piArgs)
	if err != nil {
		log.Fatalf("Could not create an instance of PiCamera: %v\n", err)
	}
	if err = piCamera.Start(); err != nil {
		log.Fatalf("Could not start the PiCamera: %v\n", err)
	}
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not render the template"))
		}
	})
	serveMux.Handle("/mjpeg", mjpeg.New(piCamera))
	log.Printf("Start listen on %s\n", listen)
	if err := http.ListenAndServe(listen, serveMux); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
