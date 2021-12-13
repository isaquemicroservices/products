# Building microservice to list products using golang and grpc

Command to generate protobuf
```go
$ protoc -I . protos/product/product.proto --go_out=plugins=grpc:./application
```

Set environment variables
```go
$ export DB_HOST=localhost DB_USER=postgres DB_PASS=? DB_PORT=5432 DB_NAME=postgres DB_DRIVER=postgres
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
