# Builder
FROM whatwewant/builder-go:v1.22-1 as builder

WORKDIR /build

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 \
  go build \
  -trimpath \
  -ldflags '-w -s -buildid=' \
  -v -o scheduler ./cmd/scheduler

# Server
FROM whatwewant/alpine:v3.17-1

LABEL MAINTAINER="Zero<tobewhatwewant@gmail.com>"

LABEL org.opencontainers.image.source="https://github.com/go-idp/scheduler"

ARG VERSION=latest

ENV VERSION=${VERSION}

COPY --from=builder /build/scheduler /bin

RUN scheduler --version

CMD scheduler server
