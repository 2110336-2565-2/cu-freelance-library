# About Entity
The entity is the object that we interested in database

# Getting Start

## Types

### Base Entity
The fundamental entity

#### Structure
The base entity is contains the essential attributes that entity should have

```go
type Base struct {
	ID        *uuid.UUID     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp;autoCreateTime:nano"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp;autoUpdateTime:nano"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;type:timestamp"`
}
```

#### Hook
The `Base` entity will create new uuid everytime when save to database if the id is **blank**

```go
func (b *Base) BeforeCreate(_ *gorm.DB) error {
	if b.ID == nil {
		b.ID = UUIDAdr(uuid.New())
	}

	return nil
}
```

#### Usage
When you want to define new entity you need to embed this entity

```go
type NewEntity struct{
	gosdk.Base
	// other fields
}
```

### PaginationMetadata
The entity for collect the metadata of pagination

#### Structure
The metadata of pagination

```go
type PaginationMetadata struct {
    ItemsPerPage int
    ItemCount    int
    TotalItem    int
    CurrentPage  int
    TotalPage    int
}
```

#### Methods

##### GetOffset
The method for get the offset value

```go
offset := meta.GetOffset()
```

#### Return

| name   | description                                                               | example |
|--------|---------------------------------------------------------------------------|---------|
| offset | the offset value (calculate from ItemPerPage value and CurrentPage value) | 10      |

##### GetItemPerPage
The method for get the item per page value

```go
itemPerPage := meta.GetItemPerPage()
```

#### Return

| name         | description                                 | example |
|--------------|---------------------------------------------|---------|
| itemPerPage  | the item per page value (min: 10, max: 100) | 10      |


##### GetCurrentPage
The method for get the current page value

```go
currentPage := meta.GetItemPerPage()
```

#### Return

| name         | description                     | example |
|--------------|---------------------------------|--------|
| currentPage  | the current page value (min: 1) | 1      |


##### ToProto
Convert to proto type

```go
metaProto := meta.ToProto()
```

#### Return

| name       | description                            | example |
|------------|----------------------------------------|---------|
| metaProto  | metadata in `*pb.PaginationMetadata`   |         |
