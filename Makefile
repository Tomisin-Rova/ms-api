GOPATH:=$(shell go env GOPATH)
PULSAR_CERT:=cert+data

GO_SOURCES := $(wildcard *.go)
GO_SOURCES += $(shell find . -type f -name "*.go")

GOFMT ?= gofmt -s

ifeq ($(filter $(TAGS_SPLIT),bindata),bindata)
	GO_SOURCES += $(BINDATA_DEST)
endif

GO_SOURCES_OWN := $(filter-out vendor/%, $(GO_SOURCES))


.PHONY: proto
proto:
	@ # NOTE, to generate the protos, you have to have your local, vendor folder because of the magefile
	@	if [[ ! -r "../zebra/protos/ms.api" ]]; \
		then \
			echo "Make sure the zebra project exists."; \
		else \
			echo "Copying proto files..."; \
			cp -r ../zebra/protos/ms.api/* ./protos; \
			./libs/mage genProto; \
		fi
	
.PHONY: build
build: proto
	go build -o srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: proto gen-mocks
	docker build . -t ms.notify:alpine

gen-mocks:
	mockery --name=IdentityServiceClient --recursive
	mockery --name=OnBoardingServiceClient --recursive
	mockery --name=PersonServiceClient --recursive
	mockery --name=PaymentServiceClient --recursive
	mockery --name=CddServiceClient --recursive
	mockery --name=AuthServiceClient --recursive
	mockery --name=AccountServiceClient --recursive
	mockery --name=ProductServiceClient --recursive
	mockery --name=VerifyServiceClient --recursive
	mockery --name=PricingServiceClient --recursive
	mockery --name=Preloader --recursive
	go generate ./...

local: schema proto gen-mocks lint test
	go fmt ./...
	go mod tidy
	go run main.go

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

.PHONY: schema
schema: proto
	./libs/mage genSchema

docker-pulsar:
	docker run -d \
      -p 6650:6650 \
      -p 8080:8080 \
      --mount source=pulsardata,target=/pulsar/data \
      --mount source=pulsarconf,target=/pulsar/conf \
      --name pulsar-standalone \
      apachepulsar/pulsar:2.6.1 \
      bin/pulsar standalone

docker-mongo:
	docker run -d  \
	  -p 27017:27017 \
	  --env MONGO_INITDB_ROOT_USERNAME=root \
	  --env MONGO_INITDB_ROOT_PASSWORD=root \
	  --env MONGO_INITDB_DATABASE=roava \
	  mongo:4.2.9

build-local:
	docker build -t ms.api --build-arg ACCESS_TOKEN=${GITHUB_TOKEN} .
	docker tag ms.api localhost:15000/ms.api
	docker push localhost:15000/ms.api
