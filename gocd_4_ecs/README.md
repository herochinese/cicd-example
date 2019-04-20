# CloudFormation Templates & correlated scripts

## Description

CloudFormation templates include two major parts, one is for provisioning CI/CD environments, other is for  provisioning ECR/ECS with dynamic port mapping.

## CI/CD
CI/CD is built on GoCD, which is very popular CI/CD tool with great ecosystem and community. GoCD is typical control server and worker nodes model, which could be construct based on containers. Currently it's working on a group of EC2 instances with auto scaling, it'll be moving to container based structure.

The whole processes are correlated by awscli commands, could be optimized further.

It's also provision a bastion server for future needs.

```
#CI/CD
x-main.yaml

```

## ECR/ECS

Basically each individual microservice has their own load balancer, image repository, cluster, service and task definition.

ECS cluster is running on a group of EC2 instances across multi-AZ, which is attached with auto scaling policy and targeting average CPU utilization.

```
#ECR/ECS
ecs/app-ecs.yaml

#Dploy
ecs/app-main.yaml

```

## Network
The network environment includes dev, test, qual and prod, each one of them has its own VPC with related network resource, such as NAT, public and private subnets, routing tables, etc. CIDR is 10.0.0.0/16, but those subnets are not overlay.

```
Note:
1. Templates is required at least 3 availability zone.
2. All domain names were followed global definition and not working on regions in China.
3. All EC2 instances are installed SSM agent and CloudWatch Agent.


```
