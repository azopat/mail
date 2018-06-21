#!/bin/bash

set -x 
set -u
set -e

namespace=$2
image=$1
app=$3

IMAGE=${image} envsubst < ci/deploy/${app}-deployment.yaml.tmpl > ci/deploy/${app}-deployment.yaml
kubectl apply -f ci/deploy/${app}-deployment.yaml --namespace=${namespace}
