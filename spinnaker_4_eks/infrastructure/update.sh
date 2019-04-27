#!/bin/bash

set -ex

aws eks --region "${AWS_DEFAULT_REGION}" update-kubeconfig --name "${ENVIRONMENT_NAME}-test"
aws eks --region "${AWS_DEFAULT_REGION}" update-kubeconfig --name "${ENVIRONMENT_NAME}-staging"
aws eks --region "${AWS_DEFAULT_REGION}" update-kubeconfig --name "${ENVIRONMENT_NAME}-production"
