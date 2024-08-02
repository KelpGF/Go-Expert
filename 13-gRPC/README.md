# gRPC

## What is gRPC?

gRPC is a free and open-source framework developed by Google that allows for the creation of high-performance, language-agnostic web services. It is based on the HTTP/2 protocol and uses Protocol Buffers as its interface definition language.

## When should I use gRPC?

- Ideal for microservices architectures.
- Streaming bidirectional using HTTP/2.
- Can be used in mobile, web, and cloud applications.

## How to use with Go?

1. Install Protocol Buffers compiler:

```bash
apt-get install protobuf-compiler
```

2. Install Go protocol buffers plugin to generate protobuf types:

```bash
go get github.com/golang/protobuf/protoc-gen-go@x.x
```

3. Install Go gRPC plugin to generate interfaces for gRPC communication:

```bash
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@x.x
```

4. Create a proto file and generate the Go code:

```bash
protoc --go_out=. --go-grpc_out=. path/to/protofile/
```

5. Install Evans to test your gRPC services:

```bash
go install github.com/ktr0731/evans@latest

evans -r repl --host hostname -p serverport
```
