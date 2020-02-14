FROM    golang:1.13-alpine3.10

ENV     GO111MODULE on

RUN     apk update && apk upgrade && apk add vim curl git make bash && \
	go get -u github.com/gin-gonic/gin && \
	go get github.com/prometheus/client_golang/prometheus && go get github.com/prometheus/client_golang/prometheus/promauto && \
	go get github.com/prometheus/client_golang/prometheus/promhttp && go get github.com/streadway/amqp

WORKDIR	/go/src
VOLUME  [ "/go/pkg", "/go/pkg"  ]
