#!/usr/bin/env bash

cd `dirname $0`

VER=$2
if [ "$2" == '' ]; then
  VER="v1.0.0"
fi

IMAGE_NAME="httpserver:$VER"
echo $IMAGE_NAME
case "$1" in
    build)
        docker build -t $IMAGE_NAME  -f Dockerfile .
        ;;

    push)
        docker push $IMAGE_NAME
        ;;
    *)
        echo $"Usage: $0 {build|push}"
        exit 1
esac
