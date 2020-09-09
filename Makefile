GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto :
	./libs/mage genProto

.PHONY: schema
schema: proto
	./libs/mage genSchema

.PHONY: build
build: proto
	go build -o srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t ms-api:alpine

local:
	go run .

docker-pulsar:
	docker run -d \
      -p 6650:6650 \
      -p 8080:8080 \
      --mount source=pulsardata,target=/var/pulsar/data \
      --mount source=pulsarconf,target=/var/pulsar/conf \
      apachepulsar/pulsar:2.6.1 \
      bin/pulsar standalone

docker-mongo:
	docker run -d  \
	  -p 27017:27017 \
	  --env MONGO_INITDB_ROOT_USERNAME=root \
	  --env MONGO_INITDB_ROOT_PASSWORD=root \
	  --env MONGO_INITDB_DATABASE=roava \
	  mongo:4.2.9
