# CI/CD Trial Project

This project is designed for practicing various of CI/CD processes on top of AWS.


# AWS Native

## CodePipeline
AWS CodePipeline is a fully managed continuous delivery service that helps you automate your release pipelines for fast and reliable application and infrastructure updates. CodePipeline automates the build, test, and deploy phases of your release process every time there is a code change, based on the release model you define.

![](./docs/codepipeline.png)

https://aws.amazon.com/codepipeline/

### Example with EKS

# Third Party

## GoCD
GoCD is easily model and visualize complex workflows, which is a free and open-source Continuous Integration and Continuous Delivery system.

![](./docs/gocd.svg)

https://www.gocd.org/

### Example with ECS

<kbd>1. Infrastructure codes</kbd>

Checkout scheduled jobs from GoCD API call and provision GoCD agent accordingly. Each job will be assigned to a container, once job is done and then close container as needed.

<kbd>2. Example project</kbd>






## Jenkins  
