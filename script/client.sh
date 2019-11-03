#!/bin/bash

cd $(dirname $0)

if [ ! -e "../client" ]; then
  go build -o ../client ../cmd/client.go
fi

../client