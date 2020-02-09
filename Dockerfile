FROM    	golang:1.13-alpine3.10
ENV     	GO111MODULE on
RUN     	apk update && apk upgrade && go get -u github.com/gin-gonic/gin
COPY		server.go $GOPATH
ENTRYPOINT	["go", "run", "server.go"]
