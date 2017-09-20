# Go Mjpeg Server

This package houses two different packages related to eachother.

1. `mjpeg` package is used to create an HTTP server for a stream of JPEG images (may or may not be in MJPEG format).
2. `piCamera` is a nice interface for `raspivid` command. It does everything from stdOut so no files are created.

## Installation

```sh
go get github.com/technomancers/goMjpegServer
```

`main.go` is an example project that consumes both mjpeg and piCamera (please use [technomancers/piCamera](https://github.com/technomancers/piCamera) instead of the one in this package as I am deprecating the one in this repository).

When you run the application `goMjpegServer` (if you have `$GOPATH/bin` in your path it should be installed) it will start a mjpeg server on [http://localhost:8080](http://localhost:8080). Navigate to it and you should see that your camera is now capturing and displaying to webpage.
