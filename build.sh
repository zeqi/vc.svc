#!/bin/bash

pushd `dirname $0` > /dev/null

# PACKAGE_NAME="vc"
# REGISTORY="docker.io"
# REGISTORY_PATH="zeqi"
# SVC_NAME="vc.svc"
# TAG="latest"
# # IMAGE_NAME=$REGISTORY/$REGISTORY_PATH/$SVC_NAME:$TAG
# IMAGE_NAME=$REGISTORY_PATH/$SVC_NAME:$TAG
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o $PACKAGE_NAME
# docker rmi $IMAGE_NAME
# docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
# docker build -t $IMAGE_NAME .
# docker push $IMAGE_NAME

make build-linux-server
docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
make docker-build