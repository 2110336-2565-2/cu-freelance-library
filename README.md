# CU Freelance Library

## Stacks
- golang

## Getting Start
These instructions will guide you to download the library

### Installing
`go get https://github.com/2110336-2565-2/sec3-group15-cu-freelance-library@latest`

1. Clone the project from [CU Freelance Library](https://github.com/2110336-2565-2/sec3-group15-cu-freelance-library)
2. Download this library by `go get github.com/2110336-2565-2/sec3-group15-cu-freelance-library`

### Testing
1. Run `go test -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./...` or `make test`

## Services

### Logger

The logger service method usage

**Info**
Print the info level log
```go
    loggerService.Info("message", "key1", "val1", "key2", "val2",...)
```

**Debug**
Print the debug level log
```go
    loggerService.Info("message", "key1", "val1", "key2", "val2",...)
```

**Error**
Print the error level log and send sentry issue
```go
    loggerService.Info("message", error, "key1", "val1", "key2", "val2",...)
```
