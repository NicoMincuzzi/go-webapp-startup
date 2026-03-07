# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run Commands

```bash
# Build
go build -o build/main ./cmd/*.go

# Run
go run ./cmd/*.go
# Server starts on port 3030

# Test all
go test -v ./cmd/...

# Test a single test
go test -v ./cmd/ -run TestSingleParam

# Docker build & run
docker build -t go-webapp .
docker run -p 3030:3030 go-webapp
```

## Architecture

This is a Go web application using the **Gin** framework (v1.6.3) with **testify** for test assertions. Go version: 1.16.

All application code lives in `cmd/`:
- `main.go` — Entry point, sets up Gin router and route definitions
- `status.go` — Health check endpoint (`GET /status`)
- `path_params.go` — Path parameter handlers (`GET /user/:name`, `GET /user/:name/*action`) with `User` struct
- `query_params.go` — Query parameter handler (`GET /user` with query string)
- `model_binding_validation.go` — JSON body binding/validation (`POST /user`) with `Login` struct

Tests are co-located with source files in `cmd/` (e.g., `status_test.go`). Tests use `httptest.NewRecorder()` to simulate HTTP requests against isolated Gin routers.

The Dockerfile uses a multi-stage build: `golang:1.16.2-alpine3.13` for building, `alpine:3.13` for runtime. Exposes port 3030.

Dependencies are vendored in `vendor/`. The directories `api/`, `configs/`, `deployments/`, `gen/`, `internal/`, `pkg/`, `test/` are empty placeholders following Go project layout conventions.
