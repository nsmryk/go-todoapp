FROM golang:1.17-alpine as server-build

WORKDIR /go/src/todo-app

COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git