FROM    	golang:1.13-alpine3.10
ENV     	GO111MODULE on
COPY		server.go /home/
COPY		datasets/normal_data/* /home/
WORKDIR		/home
ENTRYPOINT	["go", "run", "server.go"]
