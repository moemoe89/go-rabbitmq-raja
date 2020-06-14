FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-rabbitmq-raja/consumer

WORKDIR /go/src/github.com/moemoe89/go-rabbitmq-raja/consumer

COPY . /go/src/github.com/moemoe89/go-rabbitmq-raja/consumer

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/consumer
