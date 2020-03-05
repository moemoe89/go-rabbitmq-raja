FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-rabbitmq-golang/consumer

WORKDIR /go/src/github.com/moemoe89/practicing-rabbitmq-golang/consumer

COPY . /go/src/github.com/moemoe89/practicing-rabbitmq-golang/consumer

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/consumer
