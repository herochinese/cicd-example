
# Description

Initial multiple EKS cluster to host Spinnaker and demo applications.


# Prerequisites

- Python/2.7.15+
- aws-cli/1.16.130+
- eksctl/0.1.29+
- kubectl/v1.12.7+
- aws-iam-authenticator

# Processes
#1 - Configure environments for following steps.
```
export AWS_DEFAULT_REGION=us-west-2
export AWS_PROFILE=default
export ENVIRONMENT_NAME=aqt
```
#2 - Setup infrastructure for network.
```
./infrastructure/vpc.sh

```

#3 - Setup three EKS cluster for test, staging and production environment. Spinnaker is going to deploy to test environment due to frugality.

```
export APPLICATION_NAME=aqt
export PRIVATE_SUBNETS="subnet-0572cb7833d677df2,subnet-0e631dd66b6fb8f3e"
export PUBLIC_SUBNETS="subnet-0bd2063eb1187a262,subnet-0cf599fe173fe9401"

./infrastructure/eks.sh

```

#4 - Update Kubernetes config file as per EKS clusters.
```
./infrastructure/update.sh

kubectl config get-contexts
kubectl config use-context <my-cluster-name>
kubectl get svc

```
