FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-rabbitmq-raja/producer

WORKDIR /go/src/github.com/moemoe89/go-rabbitmq-raja/producer

COPY . /go/src/github.com/moemoe89/go-rabbitmq-raja/producer

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/producer
