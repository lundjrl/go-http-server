FROM golang:alpine

RUN apk add build-base

RUN mkdir /app
WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

ADD . /app

RUN go mod tidy

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
 