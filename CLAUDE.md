# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a simple Go project that demonstrates building and deploying with Buildkite CI/CD using Docker. The project contains a basic "Hello" application that accepts command-line arguments.

## Build Commands

Build the application:
```bash
go build -o dist/hello ./hello
```

Run the built binary:
```bash
./hello/hello [name]
# or after building:
./dist/hello [name]
```

Run directly without building:
```bash
go run ./hello/hello.go [name]
```

Run tests:
```bash
cd hello && go test -v
```

Run a single test:
```bash
cd hello && go test -v -run TestGreet
```

## Project Structure

- `hello/hello.go` - Main application entry point. Simple CLI tool that accepts names as command-line arguments and prints a greeting. Contains a `Greet(name string)` function that generates greeting messages.
- `hello/hello_test.go` - Unit tests for the greeting functionality
- `.buildkite/pipeline.yml` - Buildkite CI/CD pipeline configuration
- `go.mod` - Go module definition (module path: `golang.org/x/example`)

## Buildkite Pipeline

The Buildkite pipeline (`.buildkite/pipeline.yml`) includes:

1. **Build step**: Compiles the Go application using Docker (golang:1.18.0 image) and uploads the binary as an artifact
2. **Block step**: Interactive prompt that asks for user input (user's name) before proceeding

The pipeline uses the Docker plugin (v5.13.0) to ensure consistent build environments.

## Go Version

- Module requires Go 1.15+
- Buildkite pipeline uses golang:1.18.0 Docker image for builds
