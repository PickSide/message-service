#!/bin/bash

service="message-service"
port=8083

echo "Running $service"

docker.exe run -d -e VAULT_TOKEN=${VAULT_TOKEN} -e VAULT_ADDRESS=${VAULT_ADDRESS} --name pickside-$service-dev -p $port:$port $service:latest
