ARG GO_VER
ARG OS_VER

# Go builder
FROM golang:${GO_VER} as builder

WORKDIR /opt

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

COPY ./cmd ./cmd

RUN go mod download && \
    go build -o ./build/bin/ex1 ./cmd/sqlite/ex1 && \
    go build -o ./build/bin/ex2 ./cmd/sqlite/ex2

# Ubuntu container
FROM ubuntu:${OS_VER}

RUN apt-get update && apt-get install -y libc6-dev

COPY --from=builder /opt/build/bin/ex1 /usr/local/bin/ex1
COPY --from=builder /opt/build/bin/ex2 /usr/local/bin/ex2 