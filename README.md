# k8s_workshop

Workshop for kubernetes

## Demo Setup

- Install Minikube (this includes kubectl, a hypervisor and minikube itself): https://kubernetes.io/docs/tasks/tools/install-minikube/

  - On mac: curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64
sudo install minikube-darwin-amd64 /usr/local/bin/minikube

- Install docker (https://docs.docker.com/docker-for-mac/install/)

- Try to start minikube: `minikube start --vm-driver=hyperkit`
  - Hyperkit may fail due to mac permission issue. Open System Preferences → Security & Privacy → Privacy and allow changes from Oracle

These all take a bit as you have to download a lot of things

## Build Service/Container

- cd to service root (e.g. `cd ./services/counterd`)
- `GOOS=linux GOARCH=amd64 go build -o build/linux/main`
- `docker build -t localhost:5000/k8s_workshop/counterd:1 .`

## Create K8s resources

- `eval $(minikube docker-env)`
- `kubectl create deployment counterd-dep --image=localhost:5000/k8s_workshop/counterd:1`
- `kubectl expose deployment counterd-dep --type=LoadBalancer --port=8080`
- `minikube service counterd-dep`

## Increase Number of Pods
- `kubectl scale deployment.v1.apps/counterd-dep --replicas=3`

## Update image

- `GOOS=linux GOARCH=amd64 go build -o build/linux/main`
- `docker build -t localhost:5000/k8s_workshop/counterd:2 .`
- `kubectl set image deployment/counterd-dep counterd=localhost:5000/k8s_workshop/counterd:2`
