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

## Table of contents
[Logger](#Logger)

## Services

### Logger

The logger service for system log

#### Initialize
You can initialize the logger by **NewService** method 

```go
loggerService := gosdk.NewLogger("service name")
```

##### Parameters
| name         | description                             |
|--------------|-----------------------------------------|
| service name | name of the service that logger existed |


#### Integrate with sentry
The logger can create the sentry issue when log level is more than or equal to error if you set the sentry DSN

```go
if err := loggerService.SetSentryDSN("sentry dsn"); err != nil{
	return nil, err
}
```

##### Parameters
| name       | description       |
|------------|-------------------|
| sentry dsn | DSN of the sentry |

#### Rename
You can rename the logger with **SetName** method

```go
loggerService.SetName("new name"); err != nil{
	return nil, err
}
```

##### Parameters
| name         | description                             |
|--------------|-----------------------------------------|
| service name | name of the service that logger existed |


#### Usage

**Debug**
Print the debug level log
```go
loggerService.
    Debug().
    Keyword("key1", "val1").
    Keyword("key2", "val2").
    Msg("message")
```
Output
```shell
{"level":"debug","ts":1676061990.7172852,"caller":"cu-freelance-library/logger.go:89","msg":"message","service":"test","key1":"val1","key2":"val2"}
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

Output
```shell
{"level":"info","ts":1676061990.7172852,"caller":"cu-freelance-library/logger.go:89","msg":"message","service":"test","key1":"val1","key2":"val2"}
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
Output
```shell
{"level":"warn","ts":1676061990.7172852,"caller":"cu-freelance-library/logger.go:89","msg":"message","service":"test","key1":"val1","key2":"val2"}
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
Output
```shell
{"level":"error","ts":1676061990.7172852,"caller":"cu-freelance-library/logger.go:89","msg":"message","service":"test","key1":"val1","key2":"val2"}
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
Output
```shell
{"level":"fatal","ts":1676061990.7172852,"caller":"cu-freelance-library/logger.go:89","msg":"message","service":"test","key1":"val1","key2":"val2"}
```
