#!/bin/bash

echo "SERVICE_NAME -> $SERVICE_NAME"
echo `which docker`
echo `pwd`

docker build -t $SERVICE_NAME -f docker/Dockerfile .
