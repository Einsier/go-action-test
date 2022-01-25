FROM golang:1.17-alpine as builder
WORKDIR /root/go/src/github.com/einsier/go-action-test
COPY . /root/go/src/github.com/einsier/go-action-test
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go build -o app main.go

FROM alpine:latest
WORKDIR  /root/go/src/github.com/einsier/go-action-test
COPY --from=builder  /root/go/src/github.com/einsier/go-action-test/app .
EXPOSE 9090/udp
ENTRYPOINT ./app