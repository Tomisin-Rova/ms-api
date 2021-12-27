<br />
<div align="center">
<h2 align="center">API</h2>
  <p align="center">
    The Roava GraphQL API sitting behind the platform’s chosen API gateway.
    <br />
    <a href="https://fcmbuk.atlassian.net/wiki/spaces/ROAV/pages/1046315011/Features">Features</a>
    ·
    <a href="https://fcmbuk.atlassian.net/wiki/spaces/ROAV/pages/486244390/api+graphql">Api</a>
    ·
    <a href="https://github.com/roava/zebra">Zebra</a>
  </p>
</div>

---

<!-- GETTING STARTED -->
## 🚀 Getting Started

### 🛠 Prerequisites

* 🐳 Make sure that your [docker](https://docs.docker.com/get-docker/) is installed and up-to-date.
* 🦦 Golang configured.

```sh
# Setup pulsar steam locally
$ make docker-pulsar
# Run mongodb cluster for DB support
$ make docker-mongo
```
---
### Run application locally

1. Setup needed keys inside _local.yml_

```yml
# Make sure that mongo & pulsar ports are the same here as in the docker-makefile command
SERVICE_NAME: "ms.api"
PORT: "2000"
mongodb_uri: "mongodb://127.0.0.1:27017"
pulsar_url: "pulsar://127.0.0.1:6650"
pulsar_cert: ""
HTTP_PORT: 8000

onboarding_service_url: ""
verification_service_url: ""
auth_service_url: ""
account_service_url: ""
customer_service_url: ""
payment_service_url: ""
pricing_service_url: ""
JWT_SECRETS: ""
redis_url: ""
redis_password: ""

bvn:
  key: ""
  url: ""

complyadvantage:
  signed_branch_id: ""
  key: ""
  risk_profile_id: ""
  url: ""

mambu:
  key: ""
  url: ""

onfido:
  token: ""
  application_id: ""

postcoder:
  key: ""

postmark:
  key: ""

twilio:
  sid: ""
  token: ""
```

2. Use those command to set up local environment.
```sh
$ export environment=local #or use dedicated environment name
$ make local
```

_It's possible to force the environment (without setting environment variable globally) by using the below command:_
```shell script
# Force environment
$ make local environment=local
```
---
###Pulsar
* Dockerized option:

```sh
# For local event testing via dockerized pulsar
# First enter your pulsar container
$ docker exec -it pulsar-standalone bash
$ cd bin
# Run test client-command
$ ./pulsar-client produce "io.roava.event.name" -s = -m '{"valid","payload"}'
```

* Local option:

```sh
# For local event testing
# Enter your local pulsar instance shell
$ cd bin
$ ./pulsar-client produce "io.roava.event.name" -s = -m '{"valid","payload"}'
```