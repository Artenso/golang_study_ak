#!/bin/bash
#chmod +x tidy.sh

if [ -z "$1" ]; then
 echo "Необходимо указать имя модуля"
 exit 1
fi

go mod init $1
go get github.com/yuin/goldmark
go mod tidy