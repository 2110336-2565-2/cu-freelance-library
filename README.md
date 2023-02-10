# CU Freelance Library

## Stacks
- golang

## Getting Start
These instructions will guide you to download the library

### Installing
`go get https://github.com/2110336-2565-2/cu-freelance-library@latest`

1. Clone the project from [CU Freelance Library](https://github.com/2110336-2565-2/cu-freelance-library)
2. Download this library by `go get github.com/2110336-2565-2/cu-freelance-library`

### Testing
1. Run `go test -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./...` or `make test`

## Services

### Logger

The logger service method usage

**Debug**
Print the debug level log
```go
loggerService.
    Debug().
    Keyword("key1", "val1").
    Keyword("key2", "val2").
    Msg("message")
```

**Info**
Print the info level log
```go
loggerService.
    Info().
    Keyword("key1", "val1").
    Keyword("key2", "val2").
    Msg("message")
```

**Warn**
Print the warning level log
```go
loggerService.
    Warn().
    Keyword("key1", "val1").
    Keyword("key2", "val2").
    Msg("message")
```

**Error**
Print the error level log and send sentry issue
```go
loggerService.
    Error().
    Keyword("key1", "val1").
    Keyword("key2", "val2").
    Msg("message")
```

**Fatal**
Print the fatal level log and send sentry issue
```go
loggerService.
    Fatal().
    Keyword("key1", "val1").
    Keyword("key2", "val2").
    Msg("message")
```
