# About Redis Repository
Redis repository is the repository interface for using redis work on-top of [go-redis](https://github.com/redis/go-redis)

# Getting Start

## Initialization
Redis repository can be initialized by **NewRedisRepository** method

```go
repo := gosdk.NewRedisRepository(*RedisClient)
```

## Configuration
### Parameters

| name         | description                                 |
|--------------|---------------------------------------------|
| Redis Client | the client  of the redis for calling an API |

### Return

| name | description                   | example |
|------|-------------------------------|---------|
| repo | the redis repository instance |         |


## Usage

### SaveCache

```go
if err := repo.SaveCache(key, value, ttl); err != nil{
    // handle error
}
```

#### Parameters
| name               | description                     | example                 |
|--------------------|---------------------------------|-------------------------|
| key                | key of cache (must be `string`) | "key"                   |
| value              | value of cache (any type)       | "value", 1, &struct{}{} |
| ttl                | expiration time of cache        | 3600                    |

### SaveHashCache

```go
if err := repo.SaveHashCache(key, field, value, ttl); err != nil{
    // handle error
}
```

#### Parameters
| name  | description                            | example |
|-------|----------------------------------------|---------|
| key   | key of cache (must be `string`)        | "user"  |
| field | field of hash cache (must be `string`) | "name"  |
| value | value of hash cache (must be `string`) | "alice" |
| ttl   | expiration time of cache               | 3600    |

### SaveAllHashCache

```go
if err := repo.SaveAllHashCache(key, value, ttl); err != nil{
    // handle error
}
```

#### Parameters
| name  | description                                     | example                            |
|-------|-------------------------------------------------|------------------------------------|
| key   | key of cache (must be `string`)                 | "user"                             |
| value | map of hash cache (must be `map[string]string`) | map[string]string{"name": "alice"} |
| ttl   | expiration time of cache                        | 3600                               |

### GetCache

```go
type User struct{
	name string
}

result := User{}

if err := repo.GetCache(key, &result); err != nil{
    // handle error
}
```

#### Parameters
| name   | description                                       | example |
|--------|---------------------------------------------------|---------|
| key    | key of cache (must be `string`)                   | "user"  |
| result | the result point `struct{}` for receive the cache |         |

### GetHashCache

```go
value ,err := repo.GetHashCache(key, field)
if err != nil{
    // handle error
}
```

#### Parameters
| name   | description                             | example |
|--------|-----------------------------------------|---------|
| key    | key of cache (must be `string`)         | "user"  |
| field  | field of hash cache (must be `string`)  | "name"  |

#### Return
| name  | description                | example |
|-------|----------------------------|---------|
| value | value of cache in `string` | "alice" |

### GetAllHashCache

```go
values ,err := repo.GetAllHashCache(key, field)
if err != nil{
    // handle error
}
```

#### Parameters
| name   | description                             | example |
|--------|-----------------------------------------|---------|
| key    | key of cache (must be `string`)         | "user"  |
| field  | field of hash cache (must be `string`)  | "name"  |

#### Return
| name   | description                            | example                           |
|--------|----------------------------------------|-----------------------------------|
| values | values of cache in `map[string]string` | map[string]string{"name":"alice"} |

### RemoveCache

```go
if err := repo.RemoveCache(key); err != nil{
    // handle error
}
```

#### Parameters
| name               | description                     | example                 |
|--------------------|---------------------------------|-------------------------|
| key                | key of cache (must be `string`) | "key"                   |

### SetExpire

```go
if err := repo.SetExpire(key, ttl); err != nil{
    // handle error
}
```

#### Parameters
| name     | description                       | example   |
|----------|-----------------------------------|-----------|
| key      | key of cache (must be `string`)   | "key"     |
| ttl      | expiration time of cache          | 3600      |
