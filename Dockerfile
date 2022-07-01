FROM golang:1.18.3 AS builder

RUN mkdir -p /go/src/ms.api

# Access private repository using token
ARG ACCESS_TOKEN
RUN git config --global url."https://${ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /go/src/ms.api

COPY . .

RUN GIT_TERMINAL_PROMPT=1 \
    GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go build -v --installsuffix cgo --ldflags="-s" -o ms.api
FROM alpine:latest

# convert build-arg to env variables
ARG DATABASE_URL
ARG PORT
ARG ONBOARDING_SERVICE_URL
ARG VERIFY_SERVICE_URL
ARG AUTH_SERVICE_URL
ARG ONFIDO_SERVICE_URL
ARG CDD_SERVICE_URL
ARG SERVICE_NAME
ARG ENVIRONMENT

RUN apk add --no-cache tzdata
ENV TZ Africa/Lagos
ENV DATABASE_URL $DATABASE_URL
ENV PORT $PORT
ENV ONBOARDING_SERVICE_URL $ONBOARDING_SERVICE_URL
ENV VERIFY_SERVICE_URL $VERIFY_SERVICE_URL
ENV AUTH_SERVICE_URL $AUTH_SERVICE_URL
ENV ONFIDO_SERVICE_URL $ONFIDO_SERVICE_URL
ENV CDD_SERVICE_URL $CDD_SERVICE_URL
ENV SERVICE_NAME $SERVICE_NAME
ENV ENVIRONMENT $ENVIRONMENT

RUN mkdir -p /svc/
RUN touch /svc/.env
COPY --from=builder /go/src/ms.api/ms.api /svc/

WORKDIR /svc/

CMD ["./ms.api"]
