## Dockerizing the app
# docker build -t go-app .
# docker run -d -p 3333:3000 --name go-app-container go-app

## Deploy on Minikube Cluster
# kubectl create -f deployment.yaml
# kubectl get deployments
# kubectl get pods
# kubectl logs <podname>

## Exposing the Microservice
# kubectl expose deployment my-go-app --type=NodePort --name=go-app-svc --target-port=3000
# kubectl get svc
# minikube ip

## Testing Microservice
# Hit minikube-ip:nodePort from browser
