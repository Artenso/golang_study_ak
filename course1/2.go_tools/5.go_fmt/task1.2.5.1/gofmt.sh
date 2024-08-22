#!/bin/bash
# Проверяем, что передан аргумент
if [ -z "$1" ]; then
    echo "Usage: gofmt.sh $1"
    exit 1
fi

go fmt $1