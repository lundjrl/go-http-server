FROM golang:alpine

RUN mkdir /app
WORKDIR /app

ADD . /app

RUN go mod tidy
RUN go build main.go

ENTRYPOINT ["main"]
