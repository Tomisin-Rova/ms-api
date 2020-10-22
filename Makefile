GOPATH:=$(shell go env GOPATH)
PULSAR_CERT:=cert+data

.PHONY: proto
proto :
	#./libs/mage genProto
	# Add: option go_option = "pb/onfidoService"; to your proto. modify onfidoService to follow suite with your proto.
	protoc -I protos/ -I${PWD}/ --go_out=plugins=grpc:protos protos/*.proto

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

docker-compose:
	PULSAR_TLS_CERT="string"

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

