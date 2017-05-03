.DEFAULT_GOAL := all
DistDir = dist
Hosts = pi
CONN = pi@raspberrypi.local

$(DistDir):
	mkdir -p $(DistDir)

$(DistDir)/goMjpegServer_pi: $(DistDir) FORCE
	env GOOS=linux GOARCH=arm GOARM=7 go build -tags pi -a -o $@ .

.PHONY: FORCE
FORCE:

.PHONY: all
all: $(addprefix $(DistDir)/goMjpegServer_, $(Hosts))

.PHONY: deploy
deploy: $(DistDir)/goMjpegServer_pi FORCE
	scp $(DistDir)/goMjpegServer_pi $(CONN):goMjpegServer

.PHONY: clean
clean: FORCE
	rm -rf $(DistDir)
	go clean -r