FROM golang:1.16-alpine as builder

WORKDIR /go/src/bookshelf
COPY ./ /go/src/bookshelf

ENV GO111MODULE=on

ENV TZ Asia/Tokyo

RUN set -ex \
    && apk update \
    && apk add --no-cache git make tzdata \
    && go build -o ./bin/bookshelf main.go

CMD ["/go/src/bookshelf/bin/bookshelf"]
EXPOSE 8080
