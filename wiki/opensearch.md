# About Opensearch Repository
Opensearch repository is the interface repository for query the data from opensearch database.

# Getting Start

## Initialization
Opensearch repository can be initialize by **NewOpenSearchRepository** method with the **OpensearchDocumentAble** entity

```go
repo := gosdk.NewOpenSearchRepository[*portfolio.Portfolio](*OpensearchClient)
```

## Configuration
### Parameters

| name              | description                                      |
|-------------------|--------------------------------------------------|
| Opensearch Client | the client  of the opensearch for calling an API |

### Return

| name | description                        | example |
|------|------------------------------------|---------|
| repo | the opensearch repository instance |         |

## Types

### Suggestion

The field that use to search for autocomplete suggestion

```go
type Suggestion struct {
    Input  []string `json:"input" mapstructure:"input"`
    Weight int      `json:"weight" mapstructure:"weight"`
}
```

| name   | description                                           |
|--------|-------------------------------------------------------|
| input  | name of the inputs that use as keyword for suggestion |
| weight | the boots use for calculate the search score          |

### OpensearchDocumentAble

The type of struct that can use with OpensearchRepository

```go
type OpenSearchDocumentAble interface {
    ToDoc() any
    GetID() string
}
```

## Usage

### Create Index
This library provided the method to create an index for the documents

```go
if err := repo.CreateIndex(indexName, indexJsonRaw); err != nil {
    // handle error
}
```

#### Parameters
| name         | description                               | example    |
|--------------|-------------------------------------------|------------|
| indexName    | the name of index that you want to create | my-index-1 |
| indexJsonRaw | the json raw data in `[]byte`             |            |

### Insert
This library provided the method to insert a document

```go
if err := repo.Insert(indexName, docId, docData); err != nil {
    // handle error
}
```

#### Parameters
| name      | description                                        | example    |
|-----------|----------------------------------------------------|------------|
| indexName | the name of index that you want to insert document | my-index-1 |
| docId     | define id of the documentation                     | 1          |
| docData   | raw data in `[]byte`                               |            |

### Insert Bulk
This library provided the method to insert the bulk of documents

```go
if err := repo.InsertBulk(indexName, docDataList); err != nil {
// handle error
}
```

#### Parameters
| name        | description                                                                              | example    |
|-------------|------------------------------------------------------------------------------------------|------------|
| indexName   | the name of index                                                                        | my-index-1 |
| docDataList | the slice of docData that match with the type that you define when initialize repository |            |

### Search
This library provided the method for searching documents in database

```go
if err := repo.Search(indexName, *request, *result, *paginationMetadata); err != nil {
    // handle error
}
```

#### Parameters
| name               | description                                                                | example    |
|--------------------|----------------------------------------------------------------------------|------------|
| indexName          | the name of index                                                          | my-index-1 |
| request            | the search request for convert to JSON in format `*map[string]interface{}` |            |
| result             | the `*map[string]interface{}` for search result                            |            |
| paginationMetadata | the `PaginationMetadata` from `Base Entity`                                |            |


### Suggestion
This library provided the method for searching documents in database

```go
if err := repo.Suggest(indexName, *request, *result); err != nil {
    // handle error
}
```

#### Parameters
| name               | description                                                                 | example    |
|--------------------|-----------------------------------------------------------------------------|------------|
| indexName          | the name of index                                                           | my-index-1 |
| request            | the suggest request for convert to JSON in format `*map[string]interface{}` |            |
| result             | the `*map[string]interface{}` for suggest result                            |            |
