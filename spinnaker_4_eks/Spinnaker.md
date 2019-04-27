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

```


**#3**
