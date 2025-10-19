FROM golang:alpine

RUN apk add build-base

RUN mkdir /app
WORKDIR /app

ADD . /app

RUN go mod tidy
RUN go build main.go

ENTRYPOINT ["main"]
