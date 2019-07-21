FROM golang:latest
MAINTAINER Loren Hale <loren.hale@gmail.com>

COPY . /backend-api
WORKDIR /backend-api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o backend-api

EXPOSE 4000

ENTRYPOINT ["./backend-api"]
