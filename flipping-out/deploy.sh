#!/bin/bash
set -e

# Add Helm repos
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add go-feature-flag https://charts.gofeatureflag.org/
helm repo update

# Install Mongo
helm install mongo bitnami/mongodb -f values/mongo-values.yaml

# Install Kafka
helm install kafka bitnami/kafka -f values/kafka-values.yaml

# Install GoFF Relay
helm install relay-proxy go-feature-flag/relay-proxy -f values/relay-values.yaml
