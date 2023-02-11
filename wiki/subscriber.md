# About Subscriber
Subscriber is use for subscribe the topic for receive an event that was emitted from another service

# Getting Start
## Types

### MessageHandler

The handler of the subscriber, using the context for receive the message

```go
type MessageHandler func(ctx context.Context)
```

## Structure
There are 3 layers of subscriber package
- message handler
- subscribe
- subscriber management

```
subscriber management -> subscriber -> message handler
```

## Handler
The handler that handle the incoming events 

## Subscriber
The 

## Subscriber Management
The management service that use to register the subscriber for incoming event and handle the graceful shutdown
