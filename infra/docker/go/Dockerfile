FROM golang:1.17-alpine

ENV LANG C.UTF-8
ENV GO111MODULE=on

WORKDIR $GOPATH/src

RUN apk add --no-cache \
        alpine-sdk \
        git
RUN go get -u github.com/cosmtrek/air

CMD [ "air", "-c", ".air.toml" ]