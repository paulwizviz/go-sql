#!/bin/bash

export MYSQL_CLI_IMAGE=learn-sql/mysqlcmd:current
export MYSQL_CLI_CONTAINER=mysqlcli
export NETWORK=learn-sql_mysql

COMMAND=$1
SUBCMD=$2

function client(){
    local cmd=$1
    case $cmd in
        "cli")
            docker run --network=${NETWORK}  -it --rm -w /opt ${MYSQL_CLI_IMAGE} /bin/bash
            ;;
        *)
            echo "client cli"
            ;;
    esac
}

function image(){
    local cmd=$1
    case $cmd in
        "build")
            docker-compose -f ./build/mysql/builder.yml build
            ;;
        "clean")
            docker rmi -f ${MYSQL_CLI_IMAGE}
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
        *)
            echo "image [ build | image ]"
            ;;
    esac
}

function network(){
    local cmd=$1
    case $cmd in
        "clean")
            docker-compose -f ./deployment/mysql/docker-compose.yml down
            rm -rf ./tmp/mysql
            ;;
        "start")
            docker-compose -f ./deployment/mysql/docker-compose.yml up
            ;;
        "stop")
            docker-compose -f ./deployment/mysql/docker-compose.yml down
            ;;
        *)
            echo "network [ clean |start | stop ]"
            ;;
    esac
}

case $COMMAND in
    "clean")
        network stop
        network clean
        image clean
        ;;
    "client")
        client $SUBCMD
        ;;
    "network")
        network $SUBCMD
        ;;
    "image")
        image $SUBCMD
        ;;
    *)
        echo "$0 network"
        ;;
esac