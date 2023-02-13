package gosdk

import (
	"gorm.io/gorm"
	"math"
)

type Entity interface {
	TableName() string
}

func Pagination[T Entity](value *[]T, meta *PaginationMetadata, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalItems int64
	db.Model(&value).Count(&totalItems)

	meta.TotalItem = int(totalItems)
	totalPages := math.Ceil(float64(totalItems) / float64(meta.GetItemPerPage()))
	meta.TotalPage = int(totalPages)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(meta.GetOffset()).Limit(meta.ItemsPerPage)
	}
}

func FindOneByID[T Entity](id string, entity T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			First(entity, "id = ?", id)
	}
}

func UpdateWithoutResult[T Entity](id string, entity T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where(id, "id = ?", id).
			Updates(&entity)
	}
}

func UpdateByIDWithResult[T Entity](id string, entity T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where(id, "id = ?", id).
			Updates(&entity).
			First(&entity, "id = ?", id)
	}
}

func DeleteWithResult[T Entity](id string, entity T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			First(&entity, "id = ?", id).
			Delete(&entity, "id = ?", id)
	}
}

func DeleteWithoutResult[T Entity](id string, entity T) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Delete(&entity, "id = ?", id)
	}
}
