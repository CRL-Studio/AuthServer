package userdao

import (
	"github.com/CRL-Studio/AuthServer/src/models"
	"github.com/jinzhu/gorm"
)

const table = "user"

type model struct {
	ID      int64  `gorm:"column:id; primary_key"`
	UUID    string `gorm:"column:uuid; unique_index"`
	Account string `gorm:"column:account; unique_index"`
}

// QueryModel is used for possible query column
type QueryModel struct {
	UUID    string
	Account string
}

// New a row
func New(tx *gorm.DB, user *models.User) {
	err := tx.Table(table).
		Create(&models.User{
			UUID:    user.UUID,
			Account: user.Account,
		}).Error

	if err != nil {
		panic(err)
	}
}

// Modify a row
func Modify(tx *gorm.DB, user *models.User) {
	attrs := map[string]interface{}{}

	err := tx.Table(table).
		Model(models.User{}).
		Where("uuid = ?", user.UUID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

// Delete a row
func Delete(tx *gorm.DB, user *models.User) {
	err := tx.Table(table).
		Where("uuid = ?", user.UUID).
		Delete(models.User{}).Error

	if err != nil {
		panic(err)
	}
}

// GetByUUID return a record found by uuid (after mapping)
func GetByUUID(tx *gorm.DB, uuid string) *models.User {
	var result models.User
	err := tx.Table(table).
		Where("user.uuid = ?", uuid).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}

	return mapping(tx, &result)
}

// GetByAccount return a record found by account (after mapping)
func GetByAccount(tx *gorm.DB, account string) *models.User {
	var result models.User
	err := tx.Table(table).
		Where("user.account = ?", account).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}

	return mapping(tx, &result)
}

// Get return a record as raw-data-form
func Get(tx *gorm.DB, query *QueryModel) *models.User {
	result := models.User{}
	err := tx.Table(table).
		Scopes(queryChain(query)).
		Scan(&result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}

	return &result
}

func queryChain(query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(UserUUIDEqualScope(query.UUID)).
			Scopes(UserAccountEqualScope(query.Account))
	}
}

func mapping(tx *gorm.DB, model *models.User) *models.User {
	return &models.User{
		ID:   model.ID,
		UUID: model.UUID,
	}
}
