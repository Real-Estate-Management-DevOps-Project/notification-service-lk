BINARY=notification-service

.PHONY: all build run docker

all: build

build:
	go build -o $(BINARY) ./

run: build
	./$(BINARY)

docker:
	docker build -t notification-service:local .
