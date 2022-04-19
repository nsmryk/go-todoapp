FROM golang:latest

WORKDIR /go/src/todo-app

COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git