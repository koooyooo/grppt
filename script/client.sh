#!/bin/bash

cd $(dirname $0)

NAME="client"

if [ ! -e "../${NAME}" ]; then
  echo "build"
  go build -o ../${NAME} ../cmd/client.go
fi

../${NAME}