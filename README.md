<br />
<div align="center">
<h2 align="center">API</h2>
  <p align="center">
    Roava api service provides the graphql api to the clients. The api service connects with all other microservices to handle the mutations and queries. 
    <br />
    <a href="https://fcmbuk.atlassian.net/wiki/spaces/ROAV/pages/1046315011/Features">Features</a>
    ·
    <a href="https://fcmbuk.atlassian.net/wiki/spaces/ROAV/pages/486244390/api+graphql">Graphql API</a>
    ·
    <a href="https://github.com/roava/zebra">Zebra</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
      </ul>
    </li>
    <li><a href="#events">Events</a></li>
    <li><a href="#errors">Errors</a></li>
  </ol>
</details>

---

<!-- GETTING STARTED -->
## 🚀 Getting Started

### 🛠 Prerequisites

* 🐳 Make sure that your [docker](https://docs.docker.com/get-docker/) is installed and up-to-date.

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
SERVICE_NAME: "api"
PORT: "2000"
mongodb_uri: "mongodb://127.0.0.1:27017"
pulsar_url: "pulsar://127.0.0.1:6650"

# Microservices urls
onboarding_service_url: "127.0.0.1:<SERVICE_PORT>"
verification_service_url: "127.0.0.1:<SERVICE_PORT>"
auth_service_url: "127.0.0.1:<SERVICE_PORT>"
account_service_url: "127.0.0.1:<SERVICE_PORT>"
customer_service_url: "127.0.0.1:<SERVICE_PORT>"
payment_service_url: "127.0.0.1:<SERVICE_PORT>"
pricing_service_url: "127.0.0.1:<SERVICE_PORT>"

```

2. Use those command to setup local environment.
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

## ❌ Errors
| Name                        | Payload | Message                                                                                                                 |
|-----------------------------|---------|-------------------------------------------------------------------------------------------------------------------------|
| InvalidEmailError           | 1100    | invalid email address                                                                                                   |
| InvalidPhoneNumberException | 7010    | phone number is not valid                                                                                               |
| InvalidPassword             | 7010    | Your transaction password must have at least one number and at least one letter and must be at least 8-characters long. |
| InvalidPayeeDetails         | 7011    | Invalid payee account details                                                                                           |
| InvalidPaymentDetails       | 7012    | Invalid payment details                                                                                                 |
| InvalidPassCode             | 7011    | invalid pass code                                                                                                       |
| InternalErr                 | 7021    | failed to process the request, please try again later.                                                                  |
| ErrInvalidDateFormat        | 7007    | invalid date format. Date format must be dd/mm/yyyy                                                                     |
| ErrInvalidType              | 7008    | not a valid date                                                                                                        |
| ErrInvalidAge               | 7009    | minimum age requirement for using Roava is 18years                                                                      |
