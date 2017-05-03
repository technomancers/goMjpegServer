# Go Mjpeg Server

This package houses two different packages related to eachother.

1. `mjpeg` package is used to create an HTTP server for a stream of JPEG images (may or may not be in MJPEG format).
2. `piCamera` is a nice interface for `raspivid` command. It does everything from stdOut so no files are created.