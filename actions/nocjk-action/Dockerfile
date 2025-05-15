FROM golang:1.24-alpine

RUN apk add --no-cache git bash
RUN go install github.com/aethiopicuschan/nocjk/cmd/nocjk@latest

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENV PATH="$PATH:/go/bin"
