FROM    	golang:1.13-alpine3.10
ENV     	GO111MODULE on
RUN     	apk update && apk upgrade
COPY		server.go $GOPATH
COPY		datasets/ $GOPATH
ENTRYPOINT	["go", "run", "server.go"]
