# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY *.go ./

CMD [ "go", "run", "lb.go" ]
