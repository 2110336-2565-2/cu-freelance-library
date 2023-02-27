# About Jaeger
Jaeger SDK is the code for start tracing the service's operation in distributed system

# Getting Start

## Configuration

```go
type JaegerConfig struct {
    Host        string `mapstructure:"host"`
    Environment string `mapstructure:"env"`
    ServiceName string `mapstructure:"service-name"`
}
```
| name              | description                                             | example                |
|-------------------|---------------------------------------------------------|------------------------|
| Host              | The host of the Jaeger in format `http://hostname:port` | http://localhost:14268 |
| Environment       | Environment of current service                          | local                  |
| ServiceName       | The name of service                                     | gateway                |


## Usage

### Initialize
Initialize by calling `gosdk.SetupTracer`

```go
if err := gosdk.SetupTracer(*JaegerConfig, tracerName){
    // handle error	
}
```

#### Parameters

| name         | description          | example |
|--------------|----------------------|---------|
| JaegerConfig | Jaeger Configuration |         |
| tracerName   | name of the tracer   | gateway |

### StartTracer
Start to trace

```go
newCtx, span := gosdk.StartTracer(tracerName, spanName, ctx, opt...)
if span != nil {
    defer span.End()
}
```

#### Parameters
| name       | description                                  | example               |
|------------|----------------------------------------------|-----------------------|
| tracerName | name of tracer                               | verify-ticket-handler |
| spanName   | name of span                                 | verify-ticket         |
| ctx        | context to pass to another service with span |                       |
| opt        | span start option (optional)                 |                       |

#### Returns
| name    | description                 | example |
|---------|-----------------------------|---------|
| newCtx  | context from tracer         |         |
| span    | interface of span in tracer |         |

