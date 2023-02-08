FROM golang:alpine

RUN apk add curl

COPY . /web

WORKDIR /web

RUN go mod init myweb && go mod tidy
#RUN go mod tidy
CMD go run /web/cmd/main.go