## Getting started

```sh
# Starting API Server
$ go run api/

# Curling the service
curl -i localhost:8080/todos
curl -i localhost:8080/todos-create

# Run analytics process
$ go run analytics/ # it needs both main.go & analytics.go to run
```