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
export APP_NAME=jsd
docker build -t $APP_NAME -f docker/Dockerfile .

```

3. Push image into ECR

```
export APP_VER=9.0.1
$(aws ecr get-login --no-include-email --region us-east-1)
docker tag jsd:latest 530820415924.dkr.ecr.us-east-1.amazonaws.com/jsd-app-repo/jsd:$APP_VER
docker push 530820415924.dkr.ecr.us-east-1.amazonaws.com/jsd-app-repo/jsd:$APP_VER

```

4. Configure task definition & service in ECS cluster. Run CloudFormation template to deploy into ECS. 

##Task

##Inspection
