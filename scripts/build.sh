service="message-service" 

echo "Building $service"

cd "../" || {
    echo "Directory not found: $service"
}

TAG="latest"

if docker build --build-arg VERSION=$(cat .version) -t ${service}:${TAG} .; then
    echo "Successfully built $service:$TAG"
else
    echo "Failed to build $service:$TAG"
    cd -
fi

cd -
