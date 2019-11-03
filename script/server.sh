#!/bin/bash

cd $(dirname $0)

#if [ ! -e "../server" ]; then
#  go build -o ../server ../cmd/server.go
#fi
#
#../server

go run ../cmd/server.go
