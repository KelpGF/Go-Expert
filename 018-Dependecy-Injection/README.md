# Dependency Injection

## Uber FX

[Uber FX](https://github.com/uber-go/fx) is a dependency injection framework for Go. Your app will be started by this library, and it will manage the lifecycle of your application.

## Google Wire

[Google Wire](https://github.com/google/wire) is a compile-time dependency injection framework for Go. It generates code that satisfies the dependencies of your application.

You won't change your entire app to works with Wire, you will just change the entry point of your application to use the generated code.

### Installation

```bash
go install github.com/google/wire/cmd/wire@latest
```

### Example

```go
//go:build wireinject
// +build wireinject

func NewProductUseCase() *product.ProductUseCase {
 wire.Build(
  product.NewProductRepository,
  product.NewProductUseCase,
 )

 return &product.ProductUseCase{}
}
```

### Commands

- generate: `wire`
- run project: `go run main.go wire_gen.go`
