
# Jenkins
Configure Jenkins on top of EKS and cooperate with Spinnaker.

# 1. Configure storage class for Jenkins
```
cat <<EoF > ./storage-class.yaml
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: gp2
  annotations:
    "storageclass.kubernetes.io/is-default-class": "true"
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
reclaimPolicy: Retain
mountOptions:
  - debug
EoF

kubectl apply -f ./storage-class.yaml
```
#2. Install Jenkins
```
helm install stable/jenkins --set rbac.install=true --name cicd

kubectl get pods -w
```

#3. Login into Jenkins
```
1. Get your 'admin' user password by running:
  printf $(kubectl get secret --namespace default my-release-jenkins -o jsonpath="{.data.jenkins-admin-password}" | base64 --decode);echo

2. Get the Jenkins URL to visit by running these commands in the same shell:
  NOTE: It may take a few minutes for the LoadBalancer IP to be available.
        You can watch the status of by running 'kubectl get svc --namespace default -w my-release-jenkins'
  export SERVICE_IP=$(kubectl get svc --namespace default my-release-jenkins --template "{{ range (index .status.loadBalancer.ingress 0) }}{{ . }}{{ end }}")
  echo http://$SERVICE_IP:8080/login
```
