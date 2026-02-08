FROM golang:1.25-alpine AS build
RUN apk add --no-cache git ca-certificates
WORKDIR /src
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/nocjk ./cmd/nocjk

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=build /out/nocjk /usr/local/bin/nocjk
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
