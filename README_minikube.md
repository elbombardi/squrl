# Steps to deploy and test the application on minikube
## 1. Setup & start minikube
* Follow the install steps in : https://minikube.sigs.k8s.io/docs/start/

* Start minikube with a static ip address : 
```bash
minikube start --static-ip=192.168.200.200
```

* Add the following lines to your /etc/hosts file : 
```
192.168.200.200 squrl.local
192.168.200.200 api.squrl.local
```

* Enable the ingress addon : 
```bash
minikube addons enable ingress
```

* Point your shell to minikube's docker-daemon: 
```bash
eval $(minikube -p minikube docker-env)
```

* Build the docker images : 
```bash
make docker-build
```

## 2. Apply the manifests with kubectl
```bash
kubectl apply -f deployments/minikube
```

## 3. Check the status of the pods and services
```bash
kubectl get all 
```

## 4. See the logs of the pods
* For the API service : 
```bash
kubectl logs  -f deployment.apps/squrl-api-deployment
```

* For the redirection service :
```bash 
kubectl logs -f deployment.apps/squrl-redirection-deployment
```

## 5. Connect to the database with dbeaver
* Download dbeaver : https://dbeaver.io/download/
* Connect to the database with the following parameters : 
    * Host : squrl.local
    * Port : 30001
    * Database : postgres
    * User : postgres
    * Password : postgres

## 5. Test the application
Use the API server : http://api.squrl.local/v1/docs

Note : the admin password is : **admin**

## 6 Reapply the manifests
If you want to reapply the manifests, you can use the following command : 
```bash
kubectl apply -f deployments/minikube --force
```

## 7. Cleanup
This command will clean up the environement, delete the pods, services, deployments, volumes, etc.

```bash
miniube delete
```