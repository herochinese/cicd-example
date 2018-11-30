#!/bin/bash

echo "SERVICE_NAME -> $SERVICE_NAME"
echo "ECS_STACK_NAME -> $ECS_STACK_NAME"
echo "APP_VER -> $APP_VER"
echo "AWS_ACCOUNT_ID -> $AWS_ACCOUNT_ID"
echo "AWS_DEFAULT_REGION -> $AWS_DEFAULT_REGION"

repo=$AWS_ACCOUNT_ID.dkr.ecr.us-east-1.amazonaws.com/$STACK_NAME-app-repo/$SERVICE_NAME

echo "Pushing image into ECR -> $repo"

$(aws ecr get-login --no-include-email --region us-east-1)
docker tag $SERVICE_NAME:latest $repo:$APP_VER
docker push $repo:$APP_VER

echo "Done"
