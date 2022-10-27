FROM golang:alpine

RUN apk add curl

COPY src /web

WORKDIR /web

RUN go mod init myweb && go mod tidy
#RUN go mod tidy
CMD go run /web/server.go
