# go-grpc-server-quotes

Basic example implementation of gRPC: serving up insightful quotes

## Dependencies
* [Go](https://golang.org/)
* [PostgreSQL](https://www.postgresql.org/)
  * check `db/` directory for schema and seeds initialization
* `.env` file with the following variable definitions (check `env.sample` for examples)
  * `PORT`: TCP port the gRPC server will listen on
  * `DATABASE_URL`: URL to the application database
  * `TEST_DATABASE_URL`: URL to the test database

## Execution
Invoke the following shell command from the repository root directory:
```go run main.go```

## Testing
Invoke the following shell command from the repository root directory:
```go run queries/quote_test.go```