# Go Web Application Startup

[![CI](https://github.com/NicoMincuzzi/go-webapp-startup/actions/workflows/ci.yml/badge.svg)](https://github.com/NicoMincuzzi/go-webapp-startup/actions/workflows/ci.yml)
![Golang version](https://img.shields.io/badge/golang-1.16-9cf)
![GitHub repo size](https://img.shields.io/github/repo-size/NicoMincuzzi/go-webapp-startup)

A starter web application built with [Gin](https://github.com/gin-gonic/gin), demonstrating routing patterns including path parameters, query parameters, and JSON body binding with validation.

## Prerequisites

- [Go](https://golang.org/) 1.16+
- [Docker](https://www.docker.com/) (optional, for containerized builds)

## Getting Started

```bash
# Clone the repository
git clone https://github.com/NicoMincuzzi/go-webapp-startup.git
cd go-webapp-startup

# Run the application
go run ./cmd/*.go
# Server starts on http://localhost:3030
```

## API Endpoints

| Method | Path                   | Description                          |
|--------|------------------------|--------------------------------------|
| GET    | `/status`              | Health check — returns `"Healthy!"`  |
| GET    | `/user/:name`          | Get user by path parameter           |
| GET    | `/user/:name/*action`  | Get user with nested path parameters |
| GET    | `/user`                | Get user by query parameters (`firstname`, `lastname`, `nickname`) |
| POST   | `/user`                | Create user with JSON body validation |

### Example Requests

```bash
# Health check
curl http://localhost:3030/status

# Path parameter
curl http://localhost:3030/user/john

# Query parameters
curl "http://localhost:3030/user?firstname=John&lastname=Doe&nickname=jd"

# JSON body binding
curl -X POST http://localhost:3030/user \
  -H "Content-Type: application/json" \
  -d '{"user": "admin", "password": "secret"}'
```

## Build

```bash
# Build binary
go build -o build/main ./cmd/*.go

# Using Make
make build
```

## Testing

```bash
# Run all tests
go test -v ./cmd/...

# Run a specific test
go test -v ./cmd/ -run TestSingleParam

# Test with coverage
make coverage
```

## Docker

```bash
# Build image
docker build -t go-webapp .

# Run container
docker run -p 3030:3030 go-webapp
```

## Project Structure

```
.
├── cmd/                          # Application source code
│   ├── main.go                   # Entry point and route definitions
│   ├── status.go                 # Health check handler
│   ├── path_params.go            # Path parameter handlers
│   ├── query_params.go           # Query parameter handler
│   ├── model_binding_validation.go # JSON body binding and validation
│   └── *_test.go                 # Co-located test files
├── vendor/                       # Vendored dependencies
├── Dockerfile                    # Multi-stage build (alpine-based)
├── Makefile                      # Build, test, lint, and Docker targets
└── go.mod / go.sum               # Go module files
```

## Make Targets

Run `make help` to see all available targets, including:

- `make build` — Build the binary
- `make test` — Run tests
- `make coverage` — Run tests with coverage report
- `make lint` — Run all linters (Go, Dockerfile, YAML)
- `make docker-build` — Build Docker image
- `make vendor` — Vendor dependencies

## How to Contribute

Contributions are welcome! Submit bugs or suggest improvements by [opening an issue on GitHub](https://github.com/NicoMincuzzi/go-webapp-startup/issues).

## License

Distributed under the Apache-2.0 License. See the license file for details.
