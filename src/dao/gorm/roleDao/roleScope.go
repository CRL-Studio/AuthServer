package roledao

import (
	"github.com/jinzhu/gorm"
)

// RoleCodeEqualScope is
func RoleCodeEqualScope(code string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if code != "" {
			return db.Where("code = ?", code)
		}
		return db
	}
}

// RoleUUIDEqualScope is
func RoleUUIDEqualScope(uuid string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if uuid != "" {
			return db.Where("uuid = ?", uuid)
		}
		return db
	}
}
