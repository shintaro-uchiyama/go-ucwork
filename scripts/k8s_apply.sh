#!/bin/zsh
minikube start --vm-driver=hyperkit
minikube addons enable ingress
minikube addons enable gcp-auth
gcloud secrets versions access latest --secret="mysql" | jq -r '.value' > deployments/kubernetes/mysql/password
kubectl apply -k ./deployments/kubernetes/mysql

gsutil cp gs://uchiyama-sandbox-secret/firebase-adminsdk.json.enc .
gcloud kms decrypt --ciphertext-file=firebase-adminsdk.json.enc --plaintext-file=firebase-adminsdk.json --location=asia-northeast1 --keyring=sandbox-pre-processing --key=sandbox-encrypt-service-account
export GOOGLE_APPLICATION_CREDENTIALS=./firebase-adminsdk.json
skaffold dev
