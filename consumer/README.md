# GO-RABBITMQ-RAJA #

Practicing RabbitMQ Using Golang with Go Mod as Programming Language, RabbitMQ as Messaging

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ go-rabbitmq-raja/
  |     |
  |     +--+ consumer
  |     |  |
  |     |  +--+ config/
  |     |  +--+ main.go
  |     |  +--+ ... any other source code
  |     +--+ producer
  |        |
  |        +--+ config/
  |        +--+ main.go
  |        +--+ ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Requirements

Go >= 1.11

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup RabbitMQ <https://www.rabbitmq.com/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> go-rabbitmq-raja
```

## Running Application
Make config file inside producer & consumer dir :
```
$ cp config-sample.json config.json
```
Change RabbitMQ address & collection based on your config :
```
amqp://guest:guest@rabbitmq:5672/
```
Build producer / consumer inside dir :
```
$ go build
```
Run producer / consumer inside dir :
```
$ go run main.go
```

## How to Run with Docker
Make config file inside producer & consumer dir for docker :
```
$ cp config-sample.json config.json
```
Change RabbitMQ address & collection based on your docker config :
```
amqp://guest:guest@rabbitmq:5672/
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```
Trigger the message by opening this url on browser :
```
localhost:8781
```