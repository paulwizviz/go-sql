#!/bin/bash

export SQLITE_CLI_IMAGE=go-db/sqlitecmd:current
export SQLITE_CLI_CONTAINER=sqlitecli

COMMAND="$1"

function build(){
    docker-compose -f ./build/sqlite/builder.yml build
}

function clean(){
    docker rmi -f ${SQLITE_CLI_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

case $COMMAND in
    "build")
        build
        ;;
    "shell")
        docker run --name ${SQLITE_CLI_CONTAINER} -w /opt -it --rm ${SQLITE_CLI_IMAGE} /bin/bash
        ;;
    "clean")
        clean
        ;;
    *)
        echo "$0 [ build | shell | clean ]"
        ;;
esac