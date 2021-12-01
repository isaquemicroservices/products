# Building microservice to list products of one ecommerce using golang and grpc


### Command for generate protobuf
```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product.proto```