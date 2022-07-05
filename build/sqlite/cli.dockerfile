ARG GO_VER
ARG OS_VER

# Go builder
FROM golang:${GO_VER} as builder

WORKDIR /opt

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

COPY ./cmd ./cmd

RUN go mod download && \
    go build -o ./build/bin/sqlitecmd ./cmd/sqlite/cli

# App container
FROM ${OS_VER}

RUN apt-get update && apt-get install -y libc6-dev

COPY --from=builder /opt/build/bin/sqlitecmd /usr/local/bin/sqlitecmd 