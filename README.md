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
1. [Logger](#Logger)
2. [Database Connection](#Connection)
3. [Subscriber](#Subscriber)
4. [Publisher](#Publisher)

## Services

### Logger

The logger service for system log

#### Initialize
You can initialize the logger by **NewService** method 

```go
loggerService := gosdk.NewLogger("service name")
```

**Parameters**

| name         | description                             |
|--------------|-----------------------------------------|
| service name | name of the service that logger existed |


#### Integrate with sentry
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

#### Rename
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

### Connection

The initialization connection of the database and other service

#### Supported Service Connection
Current support total 4 services

| name          | description                                 |
|---------------|---------------------------------------------|
| PostgreSQL    | The opensource SQL database                 |
| Redis         | The very useful key-value database          |
| RabbitMQ      | The message broker                          |
| Elasticsearch | The database that very efficient for search |

#### Usage
**PostgreSQL**

return `*gorm.DB` when successfully  

```go
db, err := gosdk.InitDatabase(PostgresDatabaseConfig, Debug)
if err != nil {
    // handle error
}
```

**Parameters**

| name                      | description        |
|---------------------------|--------------------|
| Postgres Database Config  | Postgres config    |
| Debug                     | Enable debug mode  |


**Configuration**

```go
type PostgresDatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"username"`
    Password string `mapstructure:"password"`
    Name     string `mapstructure:"name"`
    SSL      string `mapstructure:"ssl"`
}
```

| name     | description              | example   |
|----------|--------------------------|-----------|
| Host     | Hostname of the postgres | localhost | 
| Port     | Port of database         | 5432      |
| User     | Postgres username        | postgres  |
| Password | Postgres password        | root      |
| Name     | The database name        | postgres  |
| SSL      | SSL mode                 | disable   |

**Redis**

return `*redis.Client` when successfully

```go
 cache, err := gosdk.InitRedisConnect(RedisConfig)
 if err != nil {
    // handle error
 }
```

**Parameters**

| name          | description       |
|---------------|-------------------|
| Redis Config  | Redis config      |


**Configuration**

```go
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
```
| name     | description                                     | example        |
|----------|-------------------------------------------------|----------------|
| Host     | The host of the redis in format `hostname:port` | localhost:6379 |
| Password | Redis password                                  | password       |
| DB       | The database number                             | 0              |

**RabbitMQ**

return `*amqp.Connection` when successfully

```go
rabbitMQConn, err := gosdk.InitRabbitMQConnection(RabbitMQConfig)
if err != nil {
    // handle error
}
```

**Parameters**

| name            | description     |
|-----------------|-----------------|
| RabbitMQ Config | RabbitMQ config |


**Configuration**

```go
type RabbitMQConfig struct {
   Host  string `mapstructure:"host"`
   VHost string `mapstructure:"vhost"`
}
```
| name  | description                                                                  | example                            |
|-------|------------------------------------------------------------------------------|------------------------------------|
| Host  | The host of the rabbitmq in format `amqp://username:password@hostname:port/` | amqp://guest:guest@localhost:5672/ |
| VHost | The virtual host name (default is `/`)                                       | cufreelance                        |

**Elasticsearch (Default Client)**

return `*elasticsearch.Client` when successfully

```go
esDefaultClient, err := gosdk.InitElasticDefaultClient(ElasticsearchConfig, Debug)
if err != nil {
    // handle error
}
```

**Parameters**

| name                 | description          |
|----------------------|----------------------|
| Elasticsearch Config | Elasticsearch config |


**Configuration**

```go
type ElasticsearchConfig struct {
   Host     string `mapstructure:"host"`
   Username string `mapstructure:"username"`
   Password string `mapstructure:"password"`
}
```
| name     | description                                                      | example               |
|----------|------------------------------------------------------------------|-----------------------|
| Host     | The host of the Elasticsearch in format ` http://hostname:port ` | http://localhost:9200 |
| Username | Username of Elasticsearch                                        | elastic               |
| Password | Password of Elasticsearch                                        | admin                 |

**Elasticsearch (Default Typed Client)**

return `*elasticsearch.TypedClient` when successfully

```go
esDefaultClient, err := gosdk.InitElasticTypedClient(ElasticsearchConfig, Debug)
if err != nil {
    // handle error
}
```

**Parameters**

| name                 | description          |
|----------------------|----------------------|
| Elasticsearch Config | Elasticsearch config |


**Configuration**

```go
type ElasticsearchConfig struct {
   Host     string `mapstructure:"host"`
   Username string `mapstructure:"username"`
   Password string `mapstructure:"password"`
}
```
| name     | description                                                      | example               |
|----------|------------------------------------------------------------------|-----------------------|
| Host     | The host of the Elasticsearch in format ` http://hostname:port ` | http://localhost:9200 |
| Username | Username of Elasticsearch                                        | elastic               |
| Password | Password of Elasticsearch                                        | admin                 |

