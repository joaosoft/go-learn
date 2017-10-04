# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD ./0_docker /go/src/teste

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install teste

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/teste

# Document that the service listens on port 8081.
EXPOSE 8081


# HOW TO RUN
#$ docker build -t teste .
#$ docker run --publish 6060:8081 --name test --rm teste
# -> http://localhost:6060/
