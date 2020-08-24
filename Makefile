GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto :
	protoc -I proto/ -I${PWD}/ --go_out=plugins=grpc:proto proto/*.proto

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

.PHONY: schema
schema:
	./generate.sh