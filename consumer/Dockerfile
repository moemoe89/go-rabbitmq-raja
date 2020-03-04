FROM golang:1.13 as builder

WORKDIR /go/src/consumer
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/consumer/app /app

CMD ["/app"]