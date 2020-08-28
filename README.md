# ms.api
API Gateway - Go Version. So as to maintain existing source code in other repo.


```
.
|-- Dockerfile
|-- Makefile
|-- README.md
|-- _lab
|-- api
|   |-- Mutation.resolvers.go
|   |-- Query.resolvers.go
|   |-- generated
|   |   `-- generated.go
|   `-- resolver.go
|-- cache
|   |-- cache.go
|   `-- redis
|       |-- redis.go
|       `-- redis_test.go
|-- config
|   `-- config.go
|-- go.mod
|-- go.sum
|-- gqlgen.yml
|-- libs
|   |-- DateTime
|   |   |-- date_time.go
|   |   `-- date_time_test.go
|   |-- ObjectID
|   |   |-- object_id.go
|   |   `-- object_id_test.go
|   `-- sessions
|       |-- interface.go
|       |-- session.go
|       `-- session_manager.go
|-- main.go
|-- middlewares
|   |-- middlewares.go
|   |-- protected-calls.go
|   `-- protected.go
|-- models
|   `-- models.go
|-- protos
|   `-- kyc.proto
|-- routes
|   `-- routes.go
|-- schemas
|   |-- Mutation.graphql
|   |-- Query.graphql
|   `-- Shared.graphql
|-- scripts
|   |-- generate.sh
|   |-- mage
|   `-- protogen.go
|-- server
|   `-- graphql.go
|-- services
| ...

```
## How to start 

```shell script
# To generate proto
# Note: Ensure protogen dependencies is on your local machine visit: 
# https://developers.google.com/protocol-buffers

$ make proto 
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