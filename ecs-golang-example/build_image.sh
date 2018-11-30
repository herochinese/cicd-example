#!/bin/bash

APP_VER=$GO_PIPELINE_LABEL

echo "SERVICE_NAME -> $SERVICE_NAME"
echo "APP_VER -> $APP_VER"
echo `which docker`
echo `pwd`

docker build --build-arg ver=$APP_VER -t $SERVICE_NAME -f docker/Dockerfile .
