
> draft@

# Description

Spinnaker is one of popular tooling for continuous deployment/delivery. The following experiment of CI/CD process will leverage the power of Spinnaker to address some key scenarios, mainly focus on containerize applications on Kubernetes.



# Prerequisite

## Initial EKS on AWS

Setup three different EKS clusters for test, staging and production environments. Spinnaker will be installed on test environment due to experimental purpose.

> [Initial EKS Cluster on AWS ... ...](EKS.md)

## Initial GKE on GCP

## Install Spinnaker on EKS

Install Halyard client ands Spinnaker on EKS cluster, as well as configure essential properties for following trials.

> [Install Spinnaker on EKS ... ...](Spinnaker.md)

## Configure Spinnaker
Configure Spinnaker to be able to publicly accessible, and integrate with GitHub, Travis CI, ECR, EKS, etc.

# Scenarios

## Configure Application for Simple Deployment
Setup a simple application with pipeline to experience Spinnaker.

## Configure Application for Blue/Green Deployment

## Configure Application for Canary Deployment

## Configure Application for Multiple Cloud

Deploy application to multiple cloud providers, in this case we'll use AWS and GCP as our targets. It's an excellent solution for continuous deployment on hybrid cloud infrastructure.

# Summary & Conclusion
