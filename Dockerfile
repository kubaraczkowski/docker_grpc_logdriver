FROM arm64v8/alpine

LABEL version=$VERSION
COPY ./grpc_logdriver_linux_arm64 /usr/bin/docker-log-driver_arm64
# COPY ./grpc_logdriver_linux_amd64 /usr/bin/docker-log-driver_amd64
RUN mkdir -p /run/docker/plugins/
RUN mkdir /log