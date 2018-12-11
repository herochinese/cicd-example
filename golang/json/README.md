# Golang Example Project

It's demo project by Golang with docker definition, which is going to build image & push into ECR.

##Build

1. Build & test.

```
make
make run

```

2. Build minimal image.

```
export STACK_NAME=jsd
export BASE_STACK_NAME=dvdemo
export SERVICE_NAME=buzz
export ENV=dev

./build_image.sh

```

3. Configure task definition & service in ECS cluster. Run CloudFormation template to deploy into ECS.

```
export AWS_ACCESS_KEY_ID=
export AWS_SECRET_ACCESS_KEY=
export AWS_DEFAULT_REGION=
export AWS_ACCOUNT_ID=

./setup_ecs.sh

```


4. Push image into ECR.

```
./push_image.sh

```


5. Deploy it into ECS cluster.

```
./deploy.sh

```

##Task

##Inspection
