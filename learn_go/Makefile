build:
	go build ./examples/1_helloworld

install:
	go install ./examples/1_helloworld

exec:
	./examples/1_helloworld/1_helloworld

clean:
	go clean ./examples/1_helloworld

start:
	docker-compose build
	docker-compose start
	docker-compose stop

stop:
	docker-compose stop