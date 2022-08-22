export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=./internal/driver \
       --go_opt=paths=source_relative \
       --go-grpc_out=./internal/driver \
       --go-grpc_opt=paths=source_relative \
       --proto_path=./api \
       ./api/*.proto
