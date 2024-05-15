#!/bin/bash

service="message-service"
dockerhub_username="thenoobgrammer"
dockerhub_repo_base="pickside-service"

echo "Logging into Docker Hub..."
docker login --username $dockerhub_username

echo "Pushing $service to hub"

TAG="latest"

target_repo="${dockerhub_username}/${dockerhub_repo_base}:${service}-${TAG}"

docker tag ${service}:${TAG} $target_repo

if docker push $target_repo; then
    echo "Successfully pushed $service:$TAG to $target_repo"
    echo "Removing both images..."
    docker rmi -f ${service}:${TAG}
    docker rmi -f ${target_repo}
    echo "Successfully removed"
else
    echo "Failed to push $service:$TAG to $target_repo"
fi

cd -