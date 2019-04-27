#!/bin/bash

set -ex

eksctl create cluster \
  --region "${AWS_DEFAULT_REGION}" \
  --name "${ENVIRONMENT_NAME}-test" \
  --vpc-private-subnets="${PRIVATE_SUBNETS}" \
  --vpc-public-subnets="${PUBLIC_SUBNETS}" \
  --version 1.12 \
  --nodes-min 2 \
  --nodes-max 4 \
  --nodes 2 \
  --auto-kubeconfig \
  --full-ecr-access \
  --appmesh-access

eksctl create cluster \
  --region "${AWS_DEFAULT_REGION}" \
  --name "${ENVIRONMENT_NAME}-staging" \
  --vpc-private-subnets="${PRIVATE_SUBNETS}" \
  --vpc-public-subnets="${PUBLIC_SUBNETS}" \
  --version 1.12 \
  --nodes-min 2 \
  --nodes-max 4 \
  --nodes 2 \
  --auto-kubeconfig \
  --full-ecr-access \
  --appmesh-access

eksctl create cluster \
  --region "${AWS_DEFAULT_REGION}" \
  --name "${ENVIRONMENT_NAME}-production" \
  --vpc-private-subnets="${PRIVATE_SUBNETS}" \
  --vpc-public-subnets="${PUBLIC_SUBNETS}" \
  --version 1.12 \
  --nodes-min 2 \
  --nodes-max 4 \
  --nodes 2 \
  --auto-kubeconfig \
  --full-ecr-access \
  --appmesh-access
