FROM golang:1.24-alpine

RUN apk add --no-cache git
RUN go install github.com/aethiopicuschan/nocjk/cmd/nocjk@latest

ENV PATH="$PATH:/go/bin"
ENTRYPOINT ["nocjk"]
