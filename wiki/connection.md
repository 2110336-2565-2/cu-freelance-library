# About Service Connection
Service connection is the function use for initialize the connection to database or others service

## Supported Service Connection
Current support total 4 services

| name       | description                                 |
|------------|---------------------------------------------|
| PostgreSQL | The opensource SQL database                 |
| Redis      | The very useful key-value database          |
| RabbitMQ   | The message broker                          |
| Opensearch | The database that very efficient for search |
| Jaeger     | The distributed tracing service             |

# Getting Start

## Usage
### PostgreSQL

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

### Redis

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

### RabbitMQ

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

### Opensearch

return `*opensearch.Client` when successfully

```go
openserachClient, err := gosdk.InitOpenSearhClient(OpensearchConfig, Debug)
if err != nil {
    // handle error
}
```

**Parameters**

| name              | description       |
|-------------------|-------------------|
| Opensearch Config | Opensearch config |
| Debug             | Enable debug mode |


**Configuration**

```go
type OpensearchConfig struct {
   Host               string `mapstructure:"host"`
   Username           string `mapstructure:"username"`
   Password           string `mapstructure:"password"`
   InsecureSkipVerify bool   `mapstructure:"skip-ssl"`
}
```
| name               | description                                                   | example                |
|--------------------|---------------------------------------------------------------|------------------------|
| Host               | The host of the Opensearch in format ` http://hostname:port ` | https://localhost:9200 |
| Username           | Username of Opensearch                                        | admin                  |
| Password           | Password of Opensearch                                        | admin                  |
| InsecureSkipVerify | Skip verify SSL                                               | true                   |

### JaegerConnection

return `*tracesdk.TracerProvider` when successfully

```go
traceProvider, err := gosdk.InitJaegerTracerProvider(JaegerConfig)
if err != nil {
    // handle error
}
```

**Parameters**

| name          | description   |
|---------------|---------------|
| Jaeger Config | Jaeger config |


**Configuration**

```go
type JaegerConfig struct {
    Host        string `mapstructure:"host"`
    Environment string `mapstructure:"env"`
    ServiceName string `mapstructure:"service-name"`
}
```
| name              | description                                               | example                |
|-------------------|-----------------------------------------------------------|------------------------|
| Host              | The host of the Jaeger in format ` http://hostname:port ` | http://localhost:14268 |
| Environment       | Environment of current service                            | local                  |
| ServiceName       | The name of service                                       | gateway                |
