FROM golang:latest

RUN mkdir /app
WORKDIR /app
ADD . /app

# RUN apk add --no-cache git

ENV GO111MODULE=on

RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install github.com/gofiber/fiber/v2

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main