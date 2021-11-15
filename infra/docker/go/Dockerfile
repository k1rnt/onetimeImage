FROM golang:1.17-alpine

ENV LANG C.UTF-8

RUN apk add --no-cache \
        alpine-sdk \
        git

RUN go get github.com/pilu/fresh
RUN go get github.com/labstack/echo/v4

WORKDIR /server
COPY ./server .
ENV GO111MODULE=on

CMD ["fresh"]