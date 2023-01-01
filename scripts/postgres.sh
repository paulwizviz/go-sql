#!/bin/bash

export PSQL_CLI_IMAGE=learn-sql/psqlcmd:current
export PSQL_CLI_CONTAINER=psqlcli
export NETWORK=learn-sql_postgres

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
    docker run --network=${NETWORK} -v ${PWD}/db/psql/sql:/opt/sql/ -v ${PWD}/db/psql/scripts:/opt/scripts -w /opt -it --rm ${PSQL_CLI_IMAGE} /bin/bash
}

function network(){
    local state=$1
    case $state in
        "start")
            docker-compose -f ./deployment/postgres/docker-compose.yml up
            ;;
        "stop")
            docker-compose -f ./deployment/postgres/docker-compose.yml down
            ;;
        *)
            echo "network [ start | stop ]"
            ;;
    esac
}

case $COMMAND in
    "clean")
        rm -rf ./tmp
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
    network  start and stop
"
        ;;
esac