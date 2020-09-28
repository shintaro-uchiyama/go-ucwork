#!/bin/zsh
echo "password=$(gcloud secrets versions access latest --secret="mysql" | jq -r '.value')" > deployments/kubernetes/mysql/password.txt
kubectl apply -k ./deployments/kubernetes/mysql
kubectl apply -f ./deployments/kubernetes/mysql
