#!/bin/bash

DEPLOY_STACK_NAME="$ECS_STACK_NAME-deploy"

echo "DEPLOY_STACK_NAME -> $DEPLOY_STACK_NAME"
echo "ECS_STACK_NAME -> $ECS_STACK_NAME"
echo "BASE_STACK_NAME -> $BASE_STACK_NAME"
echo "BUCKET_NAME -> $BUCKET_NAME"
echo "SERVICE_NAME -> $SERVICE_NAME"
echo "APP_VER -> $APP_VER"
echo "ENV -> $ENV"

echo "Deploying application into ECS ..."
aws cloudformation describe-stacks --stack-name $DEPLOY_STACK_NAME
isExist=$?

if [ $isExist -eq 0 ]
then

  echo "Createing new stack -> $DEPLOY_STACK_NAME"
  aws cloudformation create-stack --stack-name $DEPLOY_STACK_NAME \
    --template-url https://s3.amazonaws.com/$BUCKET_NAME/gocd-cf/ecs/app-main.yaml  \
    --parameters \
    ParameterKey=baseStackName,ParameterValue=$BASE_STACK_NAME \
    ParameterKey=ecsStackName,ParameterValue=$ECS_STACK_NAME \
    ParameterKey=env,ParameterValue=$ENV \
    ParameterKey=imageVersion,ParameterValue=$APP_VER \
    ParameterKey=serviceName,ParameterValue=$SERVICE_NAME

  aws cloudformation wait stack-create-complete --stack-name $DEPLOY_STACK_NAME

else

  echo "Updating new stack -> $DEPLOY_STACK_NAME"
  aws cloudformation update-stack --stack-name $DEPLOY_STACK_NAME \
    --template-url https://s3.amazonaws.com/$BUCKET_NAME/gocd-cf/ecs/app-main.yaml  \
    --parameters \
    ParameterKey=baseStackName,ParameterValue=$BASE_STACK_NAME \
    ParameterKey=ecsStackName,ParameterValue=$ECS_STACK_NAME \
    ParameterKey=env,ParameterValue=$ENV \
    ParameterKey=imageVersion,ParameterValue=$APP_VER \
    ParameterKey=serviceName,ParameterValue=$SERVICE_NAME

  aws cloudformation wait stack-update-complete --stack-name $DEPLOY_STACK_NAME

fi
echo "Done"
