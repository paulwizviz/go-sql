#!/bin/bash

export PSQL_CLI_IMAGE=learn-sql/psqlcmd:current
export PSQL_CLI_CONTAINER=psqlcli
export NETWORK=learn-sql_postgres

COMMAND="$1"

function clean(){
    docker rmi -f ${PSQL_CLI_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
    rm -rf ./tmp
}

case $COMMAND in
    "build")
         docker-compose -f ./build/builder.yml build psqlcmd
        ;;
    "clean")
        clean
        ;;
    "cli")
        docker run --network=${NETWORK} -w /opt -it --rm ${PSQL_CLI_IMAGE} /bin/bash
        ;;
    "start")
        docker-compose -f ./deployment/postgres/docker-compose.yml up
        ;;
    "stop")
        docker-compose -f ./deployment/postgres/docker-compose.yml down
        ;;
    *)
        echo "$0 [cli | clean | start | stop]"
        ;;
esac