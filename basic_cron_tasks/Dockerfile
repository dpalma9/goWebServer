FROM golang:1.17.3-stretch

COPY src /web

WORKDIR /web

RUN go mod init main && go mod tidy && go mod download -x

CMD go run /web/server.go
