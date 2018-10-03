env:
	docker-compose up -d psql
	docker-compose up -d nsqd
	docker-compose up -d nsqadmin
	docker-compose up -d redis_master
	docker-compose up -d elasticsearch
	docker-compose up -d cassandra-1
	docker-compose up -d cassandra-2
	docker-compose up -d mysql

start:
	docker-compose build
	docker build -f ${DOCKER_BASE_DOCKERFILE} -t ${DOCKER_REGISTRY}/${PROJECT_NAME}:base .
	docker-compose start
	docker-compose stop

stop:
	docker-compose stop

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*

mocks:
	@echo "=== Running mockery generator ==="
	CGO_ENABLED=0 /Users/joaoribeiro/workspace/go/bin/mockery -all -dir 28_tests/6_mock/domain -outpkg mocks -output 28_tests/6_mock/gen