#!/bin/bash

# pushd `dirname $0` > /dev/null

export GOPATH=/go:$GOPATH
# go build -o vc
/go/src/vc.svc/vc-svc --selector=static --server_address=0.0.0.0:8080 --broker_address=0.0.0.0:10001 --registry=kubernetes
# /go/src/vc.svc/vc --server_address=0.0.0.0:8080