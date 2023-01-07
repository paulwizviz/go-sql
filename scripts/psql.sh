#!/bin/bash

export PSQL_CLI_IMAGE=learn-sql/psqlcmd:current
export PSQL_CLI_CONTAINER=psqlcli
export NETWORK=learn-sql_psql

COMMAND="$1"
SUBCOMMAND="$2"

function image(){
    local cmd="$1"
    case $cmd in
        "build")
            docker-compose -f ./build/psql/builder.yml build
            ;;
        "clean")
            docker rmi -f ${PSQL_CLI_IMAGE}
            docker rmi -f $(docker images --filter "dangling=true" -q)
            ;;
        *)
            echo "image [ build | clean]"
            ;;
    esac
}

function client(){
    local cmd=$1
    case $cmd in
        "cli")
            docker run --network=${NETWORK} -v ${PWD}/db/psql/config/.pgpass:/root/.pgpass -v ${PWD}/db/psql/sql:/opt/sql/ -v ${PWD}/db/psql/scripts:/opt/scripts -w /opt -it --rm ${PSQL_CLI_IMAGE} /bin/bash
            ;;
        *)
            echo "client [ cli ]"
            ;;
    esac
}

function network(){
    local cmd=$1
    case $cmd in
        "clean")
            docker-compose -f ./deployment/postgres/docker-compose.yml down
            rm -rf ./tmp
            ;;
        "start")
            docker-compose -f ./deployment/postgres/docker-compose.yml up
            ;;
        "stop")
            docker-compose -f ./deployment/postgres/docker-compose.yml down
            ;;
        *)
            echo "network [ clean | start | stop ]"
            ;;
    esac
}

case $COMMAND in
    "client")
        client $SUBCOMMAND
        ;;
    "clean")
        network stop
        network clean
        image clean
        ;;
    "image")
        image $SUBCOMMAND
        ;;
    "network")
        network $SUBCOMMAND
        ;;
    *)
        echo "$0 <commands>

commands:  
    clean    project setting
    image    build and clean
    network  clean, start and stop
"
        ;;
esac