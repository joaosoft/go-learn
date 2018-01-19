env:
	docker-compose up -d psql
	docker-compose compose up -d nsqd
	docker-compose compose up -d nsqadmin
	docker-compose compose up -d redis_master
	docker-compose compose up -d elasticsearch
	docker-compose compose up -d cassandra-1
	docker-compose compose up -d cassandra-2

start:
	docker-compose build
	docker build -f ${DOCKER_BASE_DOCKERFILE} -t ${DOCKER_REGISTRY}/${PROJECT_NAME}:base .
	docker-compose start
	docker-compose stop

stop:
	docker-compose stop


buildhello:
	go build ./examples/1_hello_world

installhello:
	go install ./examples/1_hello_world

exechello:
	./examples/1_hello_world

cleanhello:
	go clean ./examples/1_hello_world