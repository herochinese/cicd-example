
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
Configure Spinnaker to be able to publicly accessible, and integrate with GitHub, Travis CI, ECR, Amazon EKS, etc.

### Public access to Spinnaker
```
kubectl edit svc spin-deck -n spinnaker
#type: LoadBalancer
#kubeclt describe svc spin-deck -n spinnaker|grep "LoadBalancer Ingress"
#loadBalancerIP:

kubectl edit svc spin-gate -n spinnaker
#type: LoadBalancer
#kubeclt describe svc spin-deck -n spinnaker|grep "LoadBalancer Ingress"
#loadBalancerIP: ...

hal config security ui edit \
    --override-base-url http://ac6fece6c68f811e98cb106e0fd571f9-1964482335.us-west-2.elb.amazonaws.com:9000

hal config security api edit \
    --override-base-url http://ac275d6cb68f811e9b39002fecc04dad-144154736.us-west-2.elb.amazonaws.com:8084

#hal deploy apply
#Add file - gate.yml at ~/.hal/default/service-settings/ if not working
#overrideBaseUrl: http://ac275d6cb68f811e9b39002fecc04dad-144154736.us-west-2.elb.amazonaws.com:8084
```

### GitHub
```
TOKEN=?
TOKEN_FILE=./git.token2
ARTIFACT_ACCOUNT_NAME=cc4i-github-artifact-account
echo $TOKEN > $TOKEN_FILE
hal config features edit --artifacts true
hal config artifact github enable

hal config artifact github account add $ARTIFACT_ACCOUNT_NAME \
    --token-file $TOKEN_FILE

#hal deploy apply
```

### Travis CI
```
```

### Amazon ECR
```
```


# Scenarios

## Configure Application for Simple Deployment
Setup a simple application with pipeline to experience Spinnaker.

## Configure Application for Blue/Green Deployment

## Configure Application for Canary Deployment

## Configure Application for Multiple Cloud

Deploy application to multiple cloud providers, in this case we'll use AWS and GCP as our targets. It's an excellent solution for continuous deployment on hybrid cloud infrastructure.

# Summary & Conclusion
