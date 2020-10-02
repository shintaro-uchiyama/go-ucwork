#!/bin/zsh
gcloud secrets versions access latest --secret="mysql" | jq -r '.value' > deployments/kubernetes/mysql/password
kubectl apply -k ./deployments/kubernetes/mysql
