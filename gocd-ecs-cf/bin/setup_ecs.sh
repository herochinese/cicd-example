#!/bin/bash


echo "ECS_STACK_NAME -> $ECS_STACK_NAME"
echo "BUCKET_NAME -> $BUCKET_NAME"
echo "SERVICE_NAME -> $SERVICE_NAME"

echo "Seting ECS  environment ..."
aws cloudformation describe-stacks --stack-name $ECS_STACK_NAME
isExist=$?

if [ $isExist -ne 0 ]
then

  echo "Createing new stack -> $ECS_STACK_NAME"
  aws cloudformation create-stack --stack-name $ECS_STACK_NAME \
    --template-url https://s3.amazonaws.com/$BUCKET_NAME/gocd-cf/ecs/app-ecs.yaml  \
    --parameters \
    ParameterKey=serviceName,ParameterValue=$SERVICE_NAME

  aws cloudformation wait stack-create-complete --stack-name $ECS_STACK_NAME

else

  echo "Updating new stack -> $ECS_STACK_NAME"
  aws cloudformation update-stack --stack-name $ECS_STACK_NAME \
    --template-url https://s3.amazonaws.com/$BUCKET_NAME/gocd-cf/ecs/app-ecs.yaml  \
    --parameters \
    ParameterKey=serviceName,ParameterValue=$SERVICE_NAME
  aws cloudformation wait stack-update-complete --stack-name $ECS_STACK_NAME

fi
echo "Done"
