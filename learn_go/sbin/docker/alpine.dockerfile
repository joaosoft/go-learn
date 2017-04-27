FROM golang:1.8-alpine

# install git and mercurial
RUN apk add --update git mercurial && rm -rf /var/cache/apk/*
# install make
RUN apk add --update bash make && rm -rf /var/cache/apk/*
# install glide
RUN go get github.com/Masterminds/glide && go install github.com/Masterminds/glide




# Copy the local package files to the container's workspace.
ADD . /go/src/learn_go

WORKDIR /go/src/learn_go/
VOLUME /go/src/learn_go/

EXPOSE 8080
ENTRYPOINT ["go"]
