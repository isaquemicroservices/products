# Building microservice to list products of one ecommerce using golang and grpc

Command for generate protobuf
```cmd
$ protoc -I ./protos/... file.proto --go_out=plugins=grpc:./application
```

Set Environment Variables
```cmd
$ export DB_HOST=localhost DB_USER=postgres DB_PASS= DB_PORT=5432 DB_NAME=postgres DB_DRIVE=postgres
```

Command for running test
```cmd
$ go test ./...
```

Create products table on postgresql 
```sql
CREATE TABLE IF NOT EXISTS t_products (
  id             SERIAL NOT NULL,
  name           VARCHAR(100) NOT NULL,
  description    VARCHAR(255) NOT NULL,
  price          DECIMAL NOT NULL,
  created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at     TIMESTAMPTZ
);
```
