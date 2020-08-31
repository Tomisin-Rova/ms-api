GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto :
	./scripts/mage -d ./scripts genProto

.PHONY: schema
schema:
	./scripts/mage -d ./scripts genSchema

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