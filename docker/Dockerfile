FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /go/src/github.com/SND1231/user_service

RUN apk update \
 && apk add vim curl git \
 && go get -u github.com/dgrijalva/jwt-go \
 && go get -u github.com/golang/protobuf \
 && go get -u github.com/gorilla/handlers \
 && go get -u github.com/grpc-ecosystem/grpc-gateway \
 && go get -u github.com/jinzhu/gorm \
 && go get -u golang.org/x/crypto \
 && go get -u golang.org/x/net \
 && go get -u google.golang.org/appengine \
 && go get -u google.golang.org/grpc
