# Testing

## Commands

### Running Tests

- `go test .` - Run all tests in the current directory.
- `go test -v .` - Run all tests in the current directory and print the output of each test.
- `go test ./...` - Run all tests in the current directory and subdirectories.
- `go test -v ./...` - Run all tests in the current directory and subdirectories and print the output of each test.

### Coverage

- `go test -cover` - Run all tests in the current directory and print the coverage.
- `go test -coverprofile=coverage.out` - Run all tests in the current directory and write the coverage to a file.
- `go tool cover -html=coverage.out` - Open the coverage file in a browser.

### Benchmarking

- `go test -bench .` - Run all benchmarks in the current directory.
- `go test -bench ./...` - Run all benchmarks in the current directory and subdirectories.

#### Benchmarking Flags

- `-benchmem` - Print memory allocations for benchmarks.
- `-benchtime=5s` - Run benchmarks for 5 seconds.
- `-count=3` - Run benchmarks 3 times.
- `-run=^$` - Run benchmarks that match the regular expression.

### Fuzzing

- `go test -fuzz=.` - Run fuzz tests in the current directory.

#### Fuzzing Flags

- `-fuzz=Function` - Run fuzz tests for a specific function.
- `-fuzztime=5s` - Run fuzz tests for 5 seconds.
- `-run=^$` - Run fuzz tests that match the regular expression.
