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