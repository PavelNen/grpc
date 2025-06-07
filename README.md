# gRPC Go Learning Project

## Description
This is a simple educational project for learning gRPC with Go. The project demonstrates basic gRPC concepts and implementation.

## Prerequisites
- Go 1.16 or higher
- Protocol Buffers compiler (protoc)
- Go plugins for protoc:
  - protoc-gen-go
  - protoc-gen-go-grpc

## Installation

1. Install Protocol Buffers compiler:
   ```bash
   # For macOS
   brew install protobuf
   ```

2. Install Go plugins for protoc:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. Install project dependencies:
   ```bash
   go mod download
   ```

## Usage

1. Generate protobuf code:
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/*.proto
   ```

2. Run the server:
   ```bash
   go run cmd/server/main.go
   ```

3. Run the client:
   ```bash
   go run cmd/client/main.go
   ```

## Project Structure
```
.
├── cmd/                    # Application entry points
│   ├── client/            # Client application
│   └── server/            # Server application
├── proto/                 # Protocol buffer definitions
├── internal/              # Private application code
│   ├── service/          # Service implementations
│   └── repository/       # Data access layer
└── pkg/                   # Public library code
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.