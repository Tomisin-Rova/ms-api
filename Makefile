GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto :
	./scripts/mage -d ./scripts genProto

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
	./scripts/generate.sh