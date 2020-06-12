FROM golang:alpine

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/SND1231/user-service

RUN apk update \
 && apk add vim curl git

COPY ./src .
RUN git clone https://github.com/SND1231/proto-files.git \
 && mkdir proto \
 && cp -pr ./proto-files/user/* /go/src/github.com/SND1231/user-service/proto
RUN go get -u github.com/dgrijalva/jwt-go \
 && go get -u github.com/golang/protobuf \
 && go get -u github.com/gorilla/handlers \
 && go get -u github.com/grpc-ecosystem/grpc-gateway \
 && go get -u github.com/jinzhu/gorm \
 && go get -u golang.org/x/crypto \
 && go get -u golang.org/x/net \
 && go get -u google.golang.org/appengine \
 && go get -u google.golang.org/grpc \
 && go get -u github.com/stretchr/testify/assert

ENV DOCKERIZE_VERSION v0.6.0

RUN apt-get update && apt-get install -y wget \
  && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz
  