FROM golang:latest AS builder

RUN mkdir -p /go/src/ms.api

WORKDIR /go/src/ms.api

COPY . .

RUN GIT_TERMINAL_PROMPT=1 \
    GOARCH=amd64 \
    GOOS=linux \
    CGO_ENABLED=0 \
    go build -v --installsuffix cgo --ldflags="-s" -o ms.api
FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ Africa/Lagos

RUN mkdir -p /svc/
RUN touch /svc/.env
COPY --from=builder /go/src/ms.api/ms.api /svc/

WORKDIR /svc/

CMD ["./ms.api"]
