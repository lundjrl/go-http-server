FROM golang:latest

RUN mkdir /app
WORKDIR /app

ENV TERM=xterm-256color
ENV COLORTERM=truecolor
ENV GO111MODULE=on

RUN go install github.com/githubnemo/CompileDaemon@latest

ADD . /app

RUN go install github.com/gofiber/fiber/v2

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
