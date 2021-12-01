# Building microservice to list products of one ecommerce using golang and grpc


### Command for generate protobuf
```protoc -I ./protos/... file.proto --go_out=plugins=grpc:./application```