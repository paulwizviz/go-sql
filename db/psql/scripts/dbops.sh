#!/bin/bash

COMMAND="$1"

message="$0 [create]"

case $COMMAND in
    "create")
        createdb newdb
        ;;
    *)
        echo $message
        ;;
esac