export VERSION=$(cat version.txt)
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o grpc_logdriver_linux_arm64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grpc_logdriver_linux_amd64

DOCKER_BUILDKIT=1 docker buildx build \
  --platform linux/arm64 \
  --push \
  --build-arg=VERSION=${VERSION} \
  -t kubaraczkowski/grpc_logdriver:${VERSION} \
  -t kubaraczkowski/grpc_logdriver:latest \
  -f Dockerfile .
