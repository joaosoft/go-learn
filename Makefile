env:
	./sbin/do.sh compose up -d psql
	#./sbin/do.sh compose up -d nsqd
	#./sbin/do.sh compose up -d redis_master
	./sbin/do.sh compose up -d elasticsearch
	./sbin/do.sh compose up -d cassandra-1
	./sbin/do.sh compose up -d cassandra-2

buildhello:
	go build ./examples/1_helloworld

build:
	./sbin/do.sh buildmake

install:
	go install ./examples/1_hello_world

exec:
	./examples/1_helloworld/1_hello_world

clean:
	go clean ./examples/1_hello_world

start:
	docker-compose build
	docker build -f ${DOCKER_BASE_DOCKERFILE} -t ${DOCKER_REGISTRY}/${PROJECT_NAME}:base .
	docker-compose start
	docker-compose stop

stop:
	docker-compose stop
    