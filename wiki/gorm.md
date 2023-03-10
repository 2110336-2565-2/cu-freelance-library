# About Gorm Repository
Gorm repository is the instance that has the code set for simple use case of gorm and can be extends by insert the scope

# Getting Start

## Initialize

```go
repo := gosdk.NewGormRepository(*GormDB)
```

## Configuration
### Parameters

| name    | description              |
|---------|--------------------------|
| Gorm DB | the client of the gorm   |

### Return

| name | description                  | example |
|------|------------------------------|---------|
| repo | the gorm repository instance |         |

## Types

```go
type Entity interface {
	TableName() string
}
```

| name        | description                     |
|-------------|---------------------------------|
| TableName() | name of the table of the entity |

## Scopes

### Pagination

the gorm scope for pagination query

```go
if err := r.db.GetDB().
    Scopes(gosdk.Pagination[Entity](entityList, metadata, gormDB, ...Scope)).
    Find(&entityList).
    Error; err != nil {
    return err
}

metadata.ItemCount = len(*entityList)
metadata.CalItemPerPage()

return nil
```

#### Parameters
| name             | description              | example |
|------------------|--------------------------|---------|
| entityList       | list of entities         |         |
| metadata         | pagination metadata      |         |
| gormDB           | gorm client              |         |
| Scope            | extends scope (optional) |         |


**Example Basic**

```go
if err := r.db.GetDB().
    Scopes(gosdk.Pagination[Entity](entity, metadata, r.db.GetDB())).
    Find(&entity).
    Error; err != nil {
    return err
}

metadata.ItemCount = len(*entity)
metadata.CalItemPerPage()

return nil
```

**Example with Scope**

```go
if err := r.db.GetDB().
    Preload("Relationship")
    Scopes(gosdk.Pagination[Entity](entity, metadata, r.db.GetDB(), func(db *gorm.DB) *gorm.DB{
        return db.Where("something = ?", something)	
    })).
    Find(&entity, "something = ?", something).
    Error; err != nil {
    return err
}

metadata.ItemCount = len(*entity)
metadata.CalItemPerPage()

return nil
```

## Usage

### GetDB
return the gorm db instance

```go
db = gorm.GetDB()
```

#### Returns
| name             | description              | example |
|------------------|--------------------------|---------|
| gormDB           | gorm client              |         |

### FindAll

findAll with pagination

```go
var entityList []Entity

if err := repo.FindOne(metadata, &entityList, ...scope); err != nil{
	// handle error
}
```

#### Parameters
| name             | description              | example |
|------------------|--------------------------|---------|
| entityList       | list of entities         |         |
| metadata         | pagination metadata      |         |
| Scope            | extends scope (optional) |         |


### FindOne

findOne entity

```go
entity := Entity{}

if err := repo.FindOne(id, &entity, ...scope); err != nil{
	// handle error
}
```

#### Parameters
| name   | description                   | example |
|--------|-------------------------------|---------|
| id     | id of entity                  |         |
| entity | empty entity for receive data |         |
| Scope  | extends scope (optional)      |         |

**Example with Scope**

```go
	if err := repo.FindOne(id, &entity, func(db *gorm.DB)*gorm.DB{
	    return db.Preload("Relationship").Where("something = ?")	
    }).
    Error; err != nil {
		return err
	}
```

### Create

create entity

```go
entity := Entity{}

if err := repo.Create(&entity, ...scope); err != nil{
	// handle error
}
```

#### Parameters
| name   | description              | example |
|--------|--------------------------|---------|
| entity | entity with data         |         |
| Scope  | extends scope (optional) |         |

### Create

update entity

```go
entity := Entity{}

if err := repo.Update(id, &entity, ...scope); err != nil{
	// handle error
}
```

#### Parameters
| name   | description              | example |
|--------|--------------------------|---------|
| id     | id of entity             |         |
| entity | entity with data         |         |
| Scope  | extends scope (optional) |         |

### Delete

delete entity

```go
entity := Entity{}

if err := repo.Delete(id, &entity, ...scope); err != nil{
	// handle error
}
```

#### Parameters
| name   | description              | example |
|--------|--------------------------|---------|
| id     | id of entity             |         |
| entity | entity with data         |         |
| Scope  | extends scope (optional) |         |
