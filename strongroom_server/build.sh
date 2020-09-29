GO111MODULE=on  # Enable module mode
PATH="$PATH:$(go env GOPATH)/bin"
protoc -I proto/ proto/*.proto --go_out=plugins=grpc:proto --go_opt=paths=source_relative
