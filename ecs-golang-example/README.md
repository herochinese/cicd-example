# Golang Example Project

It's demo project by Golang with docker definition, which is going to build image & push into ECR.

##Build

1. Local building & testing.
```
make
make run
```

2. Build minimal image.

```
export STACK_NAME=jsd
export SERVICE_NAME=buzz
docker build -t $APP_NAME -f docker/Dockerfile .

```

3. Configure task definition & service in ECS cluster. Run CloudFormation template to deploy into ECS.
```
cd cicd-example/gocd-ecs-cf/ecs
aws cloudformation create-stack --stack-name myEcDeploy \
  --template-url https://s3.amazonaws.com/cloudformation-box/gocd-cf/ecs/app-ecs.yaml  \
  --parameters

aws cloudformation wait stack-exists --stack-name jsd

```


4. Push image into ECR

```
export APP_VER=9.0.1
$(aws ecr get-login --no-include-email --region us-east-1)
docker tag $SERVICE_NAME:latest 000000000000.dkr.ecr.us-east-1.amazonaws.com/$STACK_NAME-app-repo/$SERVICE_NAME:$APP_VER
docker push 000000000000.dkr.ecr.us-east-1.amazonaws.com/$STACK_NAME-app-repo/$SERVICE_NAME:$APP_VER

```
Note:
000000000000 -> <account-id>
us-east-1 -> <region>

5. Deploy
```
aws cloudformation describe-stacks --stack-name jsd

if [ $? -eq 0 ] do
then
  aws cloudformation describe-stacks --stack-name jsd

else
  aws cloudformation create-stack --stack-name myEcDeploy \
    --template-url https://s3.amazonaws.com/cloudformation-box/gocd-cf/ecs/app-ecs.yaml  \
    --parameters
fi


```

##Task

##Inspection
