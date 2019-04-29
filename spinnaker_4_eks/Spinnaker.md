# Description

Deploy Spinnaker on EKS and configure the Spinnaker, such as GitHub, Travis CI, ECR, etc., so CI/CD demo could work smoothly.

# Prerequisites

# Processes

**#1** - Install Halyard on MacOS (or check it here for other operation system. https://www.spinnaker.io/setup/install/halyard/).

> Halyard is a command-line administration tool that manages the lifecycle of your Spinnaker deployment, including writing & validating your deployment’s configuration, deploying each of Spinnaker’s microservices, and updating the deployment.

```
curl -O https://raw.githubusercontent.com/spinnaker/halyard/master/install/macos/InstallHalyard.sh
sudo bash InstallHalyard.sh

#Check out installation
hal -v
#make sure 'hal' in your $PATH

```


**#2** - Choose cloud provider, obviously it's EKS. In this case, we're gonna use Kubernetes Provider V2 (Manifest Based)

```
# Create a kubernetes service account
CONTEXT=$(kubectl config current-context)

kubectl apply --context $CONTEXT \
    -f https://spinnaker.io/downloads/kubernetes/service-account.yml

TOKEN=$(kubectl get secret --context $CONTEXT \
   $(kubectl get serviceaccount spinnaker-service-account \
       --context $CONTEXT \
       -n spinnaker \
       -o jsonpath='{.secrets[0].name}') \
   -n spinnaker \
   -o jsonpath='{.data.token}' | base64 --decode)

kubectl config set-credentials ${CONTEXT}-token-user --token $TOKEN

kubectl config set-context $CONTEXT --user ${CONTEXT}-token-user


# Adding an account
hal config provider kubernetes enable

ACCOUNT=aqt-test-k8s-v2-account

hal config provider kubernetes account add $ACCOUNT \
    --provider-version v2 \
    --context $CONTEXT

hal config features edit --artifacts true    
```
>**Need to repeat above processes on different EKS cluster.**
>
> test, staging, production environments

**#3** - Distributed installation on Kubernetes. Distributed installations are for development orgs with large resource footprints, and for those who can’t afford downtime during Spinnaker updates. Spinnaker is deployed to a remote cloud, with each microservice deployed independently. Halyard creates a smaller, headless Spinnaker to update your Spinnaker and its microservices, ensuring zero-downtime updates.

>**Deploy Spinnaker inside the EKS cluster, which related to $ACCOUNT**

```
hal config deploy edit --type distributed --account-name $ACCOUNT

```



**#4** - Choose S3 a storage source means that Spinnaker will store all of its persistent data in a Bucket.

```
AWS_DEFAULT_REGION=us-west-2
AWS_ACCESS_KEY_ID=?

hal config storage s3 edit \
    --access-key-id $AWS_ACCESS_KEY_ID \
    --secret-access-key \
    --region $AWS_DEFAULT_REGION

hal config storage edit --type s3

```

**#5** - Pick up a version to deploy Spinnaker and connect the UI.

```
hal version list

#Choose a version
hal config version edit --version $(hal version latest -q)

hal deploy apply

#Make sure you choose kube config as original name "config"

```
