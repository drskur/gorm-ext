package gorm_ext

import "github.com/jinzhu/gorm"

func calcOffset(page int, limit int) int {
	return limit * (page - 1)
}

func Page(db *gorm.DB, page int, limit int) *gorm.DB {
	offset := calcOffset(page, limit)

	return db.Limit(limit).Offset(offset)
}