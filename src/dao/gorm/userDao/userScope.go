package userdao

import (
	"github.com/jinzhu/gorm"
)

// UserUUIDEqualScope is
func UserUUIDEqualScope(uuid string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if uuid != "" {
			return db.Where("user.uuid = ?", uuid)
		}
		return db
	}
}

// UserAccountEqualScope is
func UserAccountEqualScope(account string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if account != "" {
			return db.Where("user.account = ?", account)
		}
		return db
	}
}
