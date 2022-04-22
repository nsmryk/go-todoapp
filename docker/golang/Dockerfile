FROM golang:alpine
RUN apk update
RUN apk add alpine-sdk
COPY ./go-app /app/todo-app
WORKDIR /app/todo-app
RUN go build -o main .