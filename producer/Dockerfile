FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/practicing-rabbitmq-golang/producer

WORKDIR /go/src/github.com/moemoe89/practicing-rabbitmq-golang/producer

COPY . /go/src/github.com/moemoe89/practicing-rabbitmq-golang/producer

RUN go mod download
RUN go install

ENTRYPOINT /go/bin/producer
