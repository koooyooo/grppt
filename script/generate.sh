#!/bin/bash

cd $(dirname $0)/..

# brew install protobuf
protoc pb/grppt.proto --go_out=plugins=grpc:.