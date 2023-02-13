package gosdk

import (
	"gorm.io/gorm"
	"math"
)

func Pagination(value any, meta *PaginationMetadata, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalItems int64
	db.Model(&value).Count(&totalItems)

	meta.TotalItem = int(totalItems)
	totalPages := math.Ceil(float64(totalItems) / float64(meta.GetItemPerPage()))
	meta.TotalPage = int(totalPages)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(meta.GetOffset()).Limit(meta.ItemsPerPage)
	}
}

type Entity interface {
	TableName() string
}

type GormRepository[T Entity] interface {
	FindAll(metadata *PaginationMetadata, entities *[]T, scope ...func(db *gorm.DB) *gorm.DB) error
	FindOne(id string, entity T, scope ...func(db *gorm.DB) *gorm.DB) error
	Create(entity T, scope ...func(db *gorm.DB) *gorm.DB) error
	Update(id string, entity T, scope ...func(db *gorm.DB) *gorm.DB) error
	Delete(id string, entity T, scope ...func(db *gorm.DB) *gorm.DB) error
}

type gormRepository[T Entity] struct {
	db *gorm.DB
}

func NewGormRepository[T Entity](db *gorm.DB) GormRepository[T] {
	return &gormRepository[T]{
		db: db,
	}
}

func (r *gormRepository[T]) FindAll(metadata *PaginationMetadata, entities *[]T, scope ...func(db *gorm.DB) *gorm.DB) error {
	if err := r.db.
		Scopes(scope...).
		Scopes(Pagination(&entities, metadata, r.db)).
		Find(&entities).
		Error; err != nil {
		return err
	}

	metadata.ItemCount = len(*entities)
	return nil
}

func (r *gormRepository[T]) FindOne(id string, entity T, scope ...func(db *gorm.DB) *gorm.DB) error {
	return r.db.
		Scopes(scope...).
		First(entity, "id = ?", id).
		Error
}

func (r *gormRepository[T]) Create(entity T, scope ...func(db *gorm.DB) *gorm.DB) error {
	return r.db.
		Scopes(scope...).
		Create(entity).
		Error
}

func (r *gormRepository[T]) Update(id string, entity T, scope ...func(db *gorm.DB) *gorm.DB) error {
	return r.db.
		Scopes(scope...).
		Where(id, "id = ?", id).
		Updates(&entity).
		First(&entity, "id = ?", id).
		Error
}

func (r *gormRepository[T]) Delete(id string, entity T, scope ...func(db *gorm.DB) *gorm.DB) error {
	return r.db.
		Scopes(scope...).
		First(&entity, "id = ?", id).
		Delete(&entity).
		Error
}
