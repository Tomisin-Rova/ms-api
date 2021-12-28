GOPATH:=$(shell go env GOPATH)
PULSAR_CERT:=cert+data

GO_SOURCES := $(wildcard *.go)
GO_SOURCES += $(shell find . -type f -name "*.go")

GOFMT ?= gofmt -s

ifeq ($(filter $(TAGS_SPLIT),bindata),bindata)
	GO_SOURCES += $(BINDATA_DEST)
endif

GO_SOURCES_OWN := $(filter-out vendor/%, $(GO_SOURCES))
PROTO_SRC_DIR := ${PWD}/protos
PROTO_DST_DIR := ${PWD}/../
environment ?= $(ENVIRONMENT)

.PHONY: proto
proto:
	@ if [[ ! -r "../zebra/protos" ]]; \
	then \
	  	echo "Zebra repository doesn't exist locally."; \
	  	git clone git@github.com:roava/zebra.git ../zebra; \
	fi; \
	echo "Copying proto files..."; \
	cp -v ../zebra/protos/*.proto ./protos/; \
	echo "Generating proto..."; \
	find ./protos -type f -name "*.proto" -print0 | xargs -0 sed -i '' -e 's/go_package = "protos/go_package = "ms.api\/protos/g';\
	protoc -I=${PROTO_SRC_DIR} --go_out=plugins=grpc:${PROTO_DST_DIR} ${PROTO_SRC_DIR}/*.proto; \

.PHONY: build
build: proto
	go build -o srv *.go

.PHONY: schema
schema:
	@ if [[ ! -r "../zebra/graphql" ]]; \
	then \
		echo "Zebra repository doesn't exist locally."; \
		git clone git@github.com:roava/zebra.git ../zebra; \
  	fi; \
  	echo "Copying schema files..."; \
	cp -v ../zebra/graphql/* ./graph/schemas; \
	echo "Generating graphql..."; \
	go run github.com/99designs/gqlgen generate; \

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: proto gen-mocks
	docker build . -t ms.notify:alpine

gen-mocks: proto
	mockgen -source=./protos/pb/account/account.pb.go -destination=./mocks/account_mock.go -package=mocks
	mockgen -source=./protos/pb/auth/auth.pb.go -destination=./mocks/auth_mock.go -package=mocks
	mockgen -source=./protos/pb/customer/customer.pb.go -destination=./mocks/customer_mock.go -package=mocks
	mockgen -source=./protos/pb/onboarding/onboarding.pb.go -destination=./mocks/onboarding_mock.go -package=mocks
	mockgen -source=./protos/pb/payment/payment.pb.go -destination=./mocks/payment_mock.go -package=mocks
	mockgen -source=./protos/pb/pricing/pricing.pb.go -destination=./mocks/pricing_mock.go -package=mocks
	mockgen -source=./protos/pb/verification/verification.pb.go -destination=./mocks/verification_mock.go -package=mocks
	go generate ./...

local: update-dependencies schema proto gen-mocks lint test
	go fmt ./...
	go mod tidy
	@echo "Running service with '${environment}' environment set..."; \
	(export ENVIRONMENT=${environment}; go run main.go)	

update-dependencies:
	GOSUMDB=off go get -u github.com/roava/zebra@master

tools:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/kisielk/errcheck
	go get golang.org/x/lint/golint
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html
	go get github.com/tools/godep
	go get github.com/mitchellh/gox

lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		export BINARY="golangci-lint"; \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.31.0; \
	fi
	golangci-lint run --timeout 5m

vet:
	go vet -v ./...

fmt:
	gofmt -w .

fmt-check:
	@diff=$$($(GOFMT) -d $(GO_SOURCES_OWN)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

errors:
	errcheck -ignoretests -blank ./...

coverage:
	gocov test ./... > $(CURDIR)/coverage.out 2>/dev/null
	gocov report $(CURDIR)/coverage.out
	if test -z "$$CI"; then \
	  gocov-html $(CURDIR)/coverage.out > $(CURDIR)/coverage.html; \
	  if which open &>/dev/null; then \
	    open $(CURDIR)/coverage.html; \
	  fi; \
	fi

docker-pulsar:
	docker run -d \
      -p 6650:6650 \
      -p 8080:8080 \
      --mount source=pulsardata,target=/pulsar/data \
      --mount source=pulsarconf,target=/pulsar/conf \
      --name pulsar-standalone \
      apachepulsar/pulsar:2.7.2 \
      bin/pulsar standalone

# 2.7.2

docker-mongo:
	docker run -d  \
	  -p 27018:27017 \
	  --env MONGO_INITDB_ROOT_USERNAME=root \
	  --env MONGO_INITDB_ROOT_PASSWORD=root \
	  --env MONGO_INITDB_DATABASE=roava \
	  mongo:4.4.10

build-local:
	docker build -t ms.api --build-arg ACCESS_TOKEN=${GITHUB_TOKEN} .
	docker tag ms.api localhost:15000/ms.api
	docker push localhost:15000/ms.api
