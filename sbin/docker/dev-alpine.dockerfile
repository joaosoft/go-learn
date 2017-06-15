FROM golang:1.8-alpine

# install git and mercurial
RUN apk add --update git mercurial && rm -rf /var/cache/apk/*
# install make
RUN apk add --update bash make && rm -rf /var/cache/apk/*
# install glide
RUN go get github.com/Masterminds/glide && go install github.com/Masterminds/glide

# copy configuration
ADD ./config /etc/config

# copy and build the project
WORKDIR /go/src/golang-learn/
VOLUME /go/src/golang-learn/
EXPOSE 8080
ENTRYPOINT ["go"]
