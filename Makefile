GOPATH:=$(shell go env GOPATH)

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

local:
	go run .