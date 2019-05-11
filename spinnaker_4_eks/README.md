
# Continuous delivery with Spinnaker

## Why continuous delivery

## Description

Spinnaker is one of popular tooling for continuous deployment/delivery. The following experiment of CI/CD process will leverage the power of Spinnaker to address some key scenarios, mainly focus on containerized applications on Kubernetes with Kubernetes Provider V2.

## Prerequisite

### 1.Initial EKS on AWS

Setup three different EKS clusters for test, staging and production environments. Spinnaker will be installed on test environment due to experimental purpose.

> [Initial EKS Cluster on AWS ... ...](EKS.md)

### 2.Initial GKE on GCP

### 3.Install Spinnaker on EKS

Install Halyard client ands Spinnaker on EKS cluster, as well as configure essential properties for following trials.

> [Install Spinnaker on EKS ... ...](Spinnaker.md)

### 4.Configure Spinnaker

Configure Spinnaker to be able to access publicly and integrate with GitHub, Travis CI, ECR, Amazon EKS, S3, etc.

> [Configure Spinnaker ... ...](SpinnakerConfig.md)

## Scenarios

### 1.Configure Applications for Simple Deployment

Setup a simple application with pipeline to experience Spinnaker. Here is screen shot of first example. Easy!!!
![Here is first example](../docs/spinnaker-simple-example.png)

High level process:

- [ ] Create an application
- [ ] Define the infrastructure the service will run on
- [ ] Create a pipeline
- [ ] Run your pipeline to deploy your service

For practice, and to see some sample deployment scenarios, check out : https://www.spinnaker.io/guides/user/get-started/

### 2.Configure Applications for Blue/Green Deployment

> [Blue/Green Deployment ... ...](BlueGreen.md)

![Triggered by Travis](../docs/ab-travis-trigger.png)

Note: *Blue/Green deployment was controlled by manifest on Kubernetes Provider V2.*

### 3.Configure Applications for Canary Deployment

### 4.Configure Applications for Multiple Cloud

Deploy application to multiple cloud providers, in this case we'll use AWS and GCP as our targets. It's an excellent solution for continuous deployment on hybrid cloud infrastructure.

## Summary & Conclusion
