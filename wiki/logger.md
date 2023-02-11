# About Logger Service

## What is logger
Logger is the service that print the system log when has the incoming event for monitoring purpose

## Why need logger
We need logger for push the system logged for monitoring when the service do some event or the service got an error, so we can tracking the source of problem

# Getting Start

You can initialize the logger by **NewLogger** method

```go
loggerService := gosdk.NewLogger("service name")
```

**Parameters**

| name         | description                             |
|--------------|-----------------------------------------|
| service name | name of the service that logger existed |


## Usage

CU Freelance Library provided a following method

### Integrate with sentry
The logger can create the sentry issue when log level is **more than or equal to error** if you set the sentry DSN

```go
if err := loggerService.SetSentryDSN("sentry dsn"); err != nil{
	return nil, err
}
```

**Parameters**

| name       | description       |
|------------|-------------------|
| sentry dsn | DSN of the sentry |

### Rename
You can rename the logger with **SetName** method

```go
loggerService.SetName("new name"); err != nil{
	return nil, err
}
```

**Parameters**

| name         | description                             |
|--------------|-----------------------------------------|
| service name | name of the service that logger existed |

### Debug Log

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

### Info Log

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

### Warning Log

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

### Error Log

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

### Fatal Log
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

