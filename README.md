# 📙 Building microservice to list products using golang and grpc

Command to generate protobuf
```go
$ protoc -I . protos/product/product.proto --go_out=plugins=grpc:./application
```

### Create folder for config.json file
```bat
$ sudo mkdir /etc/ms-products
$ sudo touch /etc/ms-products/config.json
$ sudo cp ./config.json /etc/ms-products/config.json
$ sudo chmod 777 /etc/ms-products/config.json
```
if you changed the config.json file, use the command at the bottom to update the config.json file on your computer
```bat
$ sudo cp ./config.json /etc/ms-products/config.json
```

command to run the test
```go
$ go test ./... --cover
```

Command to generate test files
```go
$ go test -coverprofile cover.out 
$ go tool cover -html=cover.out -o cover.html
```

Create product table in PostgreSQL 
```sql
CREATE TABLE public.t_products (
  id serial4 NOT NULL,
  "name" varchar(100) NOT NULL,
  description varchar(255) NOT NULL,
  price numeric NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NULL,
  CONSTRAINT t_products_pk PRIMARY KEY (id)
);
```
