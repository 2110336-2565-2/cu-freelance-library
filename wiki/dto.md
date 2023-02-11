# About DTO
Data Transfer Object is the object use for represent the attribute between the service

# Getting Start

## Types

### QueryResult
The base query result entity for opensearch

#### Structure
The base entity is contains the essential attributes that entity should have

```go
type QueryResult struct {
    Took    uint  `json:"took"`
    Timeout bool  `json:"timeout"`
    Shards  Shard `json:"_shards"`
}
```

#### Usage
When you want to define new entity you need to embed this entity

```go
type NewQueryResult struct{
	gosdk.QueryResult
	// other fields
}
```


### Shard
The stats of shards

#### Structure 
The value of the statistic of shard

```go
type Shard struct {
    Total      uint `json:"total"`
    Successful uint `json:"successful"`
    Skipped    uint `json:"skipped"`
    Failed     uint `json:"failed"`
}
```