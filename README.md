# ms.api
API Gateway - Go Version. So as to maintain existing source code in other repo.

## Glossary:

1. `mage` -- A "makefile" for go. 
1. `generate.go`  Code file for scripts used by mage to generate models + proto

## How to start 

```shell script
# To generate proto
# Note: Ensure protogen dependencies is on your local machine visit: 
# https://developers.google.com/protocol-buffers

$ make proto 
```

```shell script
# To generate graphql schemas
# Note: see server/graph/schemas and also ./gqlgen.yml
$ make schema 
```

```shell script
# For local development

$ make local
```

```shell script
# For tests

$ make test 
```


```shell script
# For docker

$ make docker 
```


```shell script
# to build binary

$ make build 
```