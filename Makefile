env:
    docker-compose up -d psql
    docker-compose up -d nsqd
    docker-compose up -d redis_master
    docker-compose up -d elasticsearch

buildhello:
    go build ./examples/1_helloworld

build:
    ./sbin/do.sh build

install:
    go install ./examples/1_helloworld

exec:
    ./examples/1_helloworld/1_helloworld

clean:
    go clean ./examples/1_helloworld

start:
    docker-compose build
    docker build -f ${DOCKER_BASE_DOCKERFILE} -t ${DOCKER_REGISTRY}/${PROJECT_NAME}:base .
    docker-compose start
    docker-compose stop

stop:
    docker-compose stop
