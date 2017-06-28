FROM golang:1.8-alpine

# install git and mercurial
RUN apk add --update git mercurial && rm -rf /var/cache/apk/*
# install make
RUN apk add --update bash make && rm -rf /var/cache/apk/*
# install glide
RUN go get github.com/Masterminds/glide && go install github.com/Masterminds/glide

# install protobuf
RUN apk add --update protobuf && rm -rf /var/cache/apk/*
RUN go get github.com/golang/protobuf/proto github.com/golang/protobuf/protoc-gen-go
# install grpc
RUN go get google.golang.org/grpc

# install dependencies
ADD glide.yaml glide.lock /go/src/golang-learn/
RUN cd /go/src/golang-learn && glide install

# copy configuration
ADD ./config /etc/config

# add source code
ADD . /go/src/golang-learn/
WORKDIR /go/src/golang-learn/

# copy and build the project
# ADD . /go/src/golang-learn/
# RUN cd /go/src/golang-learn && glide install
# WORKDIR /go/src/golang-learn/

EXPOSE 8080
ENTRYPOINT ["go"]
