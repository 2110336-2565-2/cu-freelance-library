# About RabbitMQ
RabbitMQ SDK is the complete utils function for implement the publisher and subscriber

# Getting Start

## Initialization

```go
rabbitmq, err := gosdk.NewRabbitMQ(RabbitMQClient)
```

## Configuration
### Parameters

| name            | description                  |
|-----------------|------------------------------|
| RabbitMQ Client | the client  of the rabbitmq  |

## Usage

### Close
Close the rabbitmq channel connection

```go
if err := rabbitmq.Close(); err != nil {
    // handle error
}
```

### Publish
Publish the massage to RabbitMQ

```go
if err := rabbitmq.Publish(exchangeName, key, message); err != nil {
    // handle error
}
```

#### Parameters
| name         | description                                           | example                           |
|--------------|-------------------------------------------------------|-----------------------------------|
| exchangeName | the name of exchange that will publish the message to | exchange-1                        |
| key          | the binding key to queue                              | cufreelance.service.action.status |
| message      | raw data in `any` type                                |                                   |

### Create Exchange
Create exchange connection

```go
if err := rabbitmq.CreateExchange(
	name, 
	kind, 
	durable, 
	autoDelete, 
	internal, 
	noWait, 
	args
); err != nil {
    // handle error
}
```

#### Parameters
| name       | description                                | example    |
|------------|--------------------------------------------|------------|
| name       | the name of exchange                       | exchange-1 |
| kind       | the kind of exchange                       | topic      |
| durable    | flag durable (true/false)                  | true       |
| autoDelete | flag auto delete (true/false)              | false      |
| internal   | flag internal (true/false)                 | false      |
| noWait     | flag no wait (true/false)                  | false      |
| args       | more arguments in `map[string]interface{}` |            |


### Create Queue
Create the queue connection for subscriber
```go
queue, err := rabbitmq.CreateQueue(
	name, 
	durable, 
	autoDelete, 
	exclusive, 
	noWait, 
	args
)

if err != nil {
    // handle error
}
```

#### Parameters
| name       | description                                | example |
|------------|--------------------------------------------|---------|
| name       | name of queue                              | queue-1 |
| durable    | flag durable (true/false)                  | true    |
| autoDelete | flag auto delete (true/false)              | false   |
| exclusive  | flag exclusive (true/false)                | false   |
| noWait     | flag no wait (true/false)                  | false   |
| args       | more arguments in `map[string]interface{}` |         |


### Bind Queue With Exchange
Bind subscriber's queue with exchange

```go
if err := rabbitmq.BindQueueWithExchange(
	queueName, 
	key, 
	exchangeName, 
	noWait, 
	args
); err != nil {
    // handle error
}
```

#### Parameters
| name         | description                                | example                           |
|--------------|--------------------------------------------|-----------------------------------|
| queueName    | the name of queue                          | queue-1                           |
| key          | the binding key                            | cufreelance.service.action.status |
| exchangeName | the name of exchange                       | exchange-1                        |
| noWait       | flag no wait (true/false)                  | false                             |
| args         | more arguments in `map[string]interface{}` |                                   |


### Create Message Channel
Create message channel for subscriber

```go
if err := rabbitmq.CreateMessageChannel(
	queueName, 
	consumerName,
	autoAck,
    exclusive, 
    noLocal, 
	noWait, 
	args
); err != nil {
    // handle error
}
```

#### Parameters
| name         | description                                | example    |
|--------------|--------------------------------------------|------------|
| queueName    | the name of queue                          | queue-1    |
| consumerName | the name of consumer                       | consumer-1 |
| autoAck      | flag auto acknowledge (true/false)         | true       |
| exclusive    | flag exclusive (true/false)                | false      |
| noLocal      | flag no local (true/false)                 | false      |
| noWait       | flag no wait (true/false)                  | false      |
| args         | more arguments in `map[string]interface{}` |            |

### Listen
Create message channel for subscriber

```go
raw, err := rabbitmq.Listen()
if err != nil {
    // handle error
}
```

#### Return

| name      | description                                                                                        | example |
|-----------|----------------------------------------------------------------------------------------------------|---------|
| raw       | the raw data that return from message channel in `[]byte` (use json.Unmarshal to decode to struct) |         |
