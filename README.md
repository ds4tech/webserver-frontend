# Webserver finesse-frontend

Status of latest build from branches:<br/> 
main:
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/ds4tech/finesse-frontend/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/ds4tech/finesse-frontend/tree/main)
<br/>
dev:
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/ds4tech/finesse-frontend/tree/dev.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/ds4tech/finesse-frontend/tree/dev)


## PreReq
```
export CALCULATOR_URL="http://localhost:8888"
```
### Google Account
```
gcloud iam workload-identity-pools create github-actions-pool \
--location="global" \
--description="The pool to authenticate GitHub actions." \
--display-name="GitHub Actions Pool"
```

```
gcloud iam workload-identity-pools providers create-oidc github-actions-oidc --workload-identity-pool="github-actions-pool" --issuer-uri="https://token.actions.ds4tech.com/" --attribute-mapping="google.subject=assertion.sub,attribute.repository=assertion.repository,attribute.repository_owner=assertion.repository_owner,attribute.branch=assertion.sub.extract('/heads/{branch}/')" \
--location=global \
--attribute-condition="assertion.repository_owner=='sec-mik'"
```

Create Service Account
```
gcloud iam service-accounts create finesse-frontend-sa --display-name="Finesse Application Service Account" --description="manages the application resources"

gcloud iam service-accounts create networking-sa --display-name="Networking Service Account" --description="manages the networking resources
```

Add policy binding
```
gcloud iam service-accounts add-iam-policy-binding finesse-frontend-sa@production-secmik.iam.gserviceaccount.com \
  --role="roles/iam.workloadIdentityUser" \
--member="principalSet://iam.googleapis.com/projects/987654321/locations/global/workloadIdentityPools/github-actions-pool/attribute.repository/sec-mik/finesse-frontend-repo"
```
## Simple Web-server application written in Go lang.

1. [Introduction](#intro)
2. [Build](#build) <br>
   2.1. [Exec](#build.exe) <br>
   2.2. [Docker](#build.docker)
3. [Deploy](#deploy) <br>
 3.1. [Kubernetes](#deploy.k8s) <br>
4. [Usage](#usage)
5. [Continous Integration](#ci)


## Introduction <a name="intro"></a>

Simple Webserver Go project:<a name="intro"></a>
API:
- /-/health - returns server version 
- echo - /api/echo?text=foo --> returns a JSON object with the key "text

Main page shows form which allows to input values which are sent to calculator webservice.

## BUILD <a name="build"></a>

### Executable <a name="build.exe"></a>
```
go build -o webserver cmd/main.go 
./webserver
```

http://localhost:9090/

### Docker container <a name="build.docker"></a>
```
docker build . -t ds4tech/finesse-frontend:0.0.1
docker run -it --rm -p 9090:9090 --name finesse-frontend ds4tech/finesse-frontend:0.0.1
```

http://localhost:9090/

## DEPLOY <a name="deploy"></a>

### Kubernetes <a name="deploy.k8s"></a>
```
kubectl apply -f deployment/kubernetes/manifest.yaml
kubectl port-forward svc/finesse-frontend 9090
```

http://localhost:9090/

### Helm <a name="deploy.k8s"></a>
```
helm install finesse-frontend ./deployment/helm/charts/finesse-frontend
kubectl port-forward svc/finesse-frontend 9090
```

http://localhost:9090/

## USEAGE <a name="usage"></a>

1. Echo
```
curl -X GET "http://localhost:9090/api/echo?text=testingJson"
```

## Continous Integration <a name="ci"></a>
Pipeline script written in yaml file for Circle CI is placed in [build/ci directory](https://github.com/ds4tech/finesse-frontend/blob/dev/.circleci/config.yml).  <br>
