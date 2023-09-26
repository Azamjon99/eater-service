package utils

import (
	"github.com/Azamjon99/eater-service/src/domain/address/models"
	"gorm.io/gorm"
)

func Sort(sort string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		orderBy := "created_at DESC"
		if sort == models.SortByCreatedAtAsc {
			orderBy = "created_at ASC"
		}
		return db.Order(orderBy)
	}
}
