export VERSION=$(cat version.txt)
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o grpc_logdriver_linux_arm64
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grpc_logdriver_linux_amd64

# cleanup
docker plugin disable -f kubaraczkowski/grpc_logdriver
docker plugin rm -f kubaraczkowski/grpc_logdriver

# build
docker build -t rootfsimage .
id=$(docker create rootfsimage true) 
mkdir -p rootfs/run/docker/plugins
docker export "$id" | tar -x -C rootfs
docker rm -vf "$id"
docker rmi rootfsimage
docker plugin rm kubaraczkowski/grpc_logdriver
docker plugin create kubaraczkowski/grpc_logdriver .
rm -rf rootfs
# push
docker plugin push --disable-content-trust kubaraczkowski/grpc_logdriver:${VERSION}
docker plugin push --disable-content-trust kubaraczkowski/grpc_logdriver:latest
# install
docker plugin enable kubaraczkowski/grpc_logdriver