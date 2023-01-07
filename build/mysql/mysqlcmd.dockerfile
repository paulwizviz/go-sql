ARG OS_VER

# App container
FROM ubuntu:${OS_VER}

RUN apt-get update && \
    apt-get install -y mysql-client

