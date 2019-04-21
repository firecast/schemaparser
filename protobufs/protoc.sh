protoc types/thing/product.proto -I. --go_out=:$GOPATH/src
protoc -I. --go_out=plugins=grpc:$GOPATH/src services.proto