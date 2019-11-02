package userdao

import (
	"time"

	roledao "github.com/CRL-Studio/AuthServer/src/dao/gorm/roleDao"
	"github.com/CRL-Studio/AuthServer/src/models"
	"github.com/jinzhu/gorm"
)

const table = "user"

type pattern struct {
	ID           int64      `gorm:"column:id; primary_key"`
	UUID         string     `gorm:"column:uuid; unique_index"`
	RoleUUID     string     `gorm:"column:role_uuid"`
	Account      string     `gorm:"column:account; unique_index"`
	Password     string     `gorm:"column:password"`
	Name         string     `gorm:"column:name"`
	Email        string     `gorm:"column:email; unique_index"`
	Status       string     `gorm:"column:status; default:'enabled'"`
	Score        int        `gorm:"column:score"`
	Verification bool       `gorm:"column:verification"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
	CreatedBy    string     `gorm:"column:created_by"`
	UpdatedBy    string     `gorm:"column:updated_by"`
}

// QueryModel is used for possible query column
type QueryModel struct {
	UUID    string
	Account string
}

// New a row
func New(tx *gorm.DB, user *models.User) {
	err := tx.Table(table).
		Create(&pattern{
			UUID:         user.UUID,
			RoleUUID:     user.Role.UUID,
			Account:      user.Account,
			Password:     user.Password,
			Name:         user.Name,
			Email:        user.Email,
			Status:       user.Status,
			Score:        user.Score,
			Verification: user.Verification,
			CreatedBy:    user.UpdatedBy,
			CreatedAt:    time.Now(),
		}).Error

	if err != nil {
		panic(err)
	}
}

// Modify a row
func Modify(tx *gorm.DB, user *models.User) {
	attrs := map[string]interface{}{
		"Name": user.Name,
	}

	err := tx.Table(table).
		Model(models.User{}).
		Where("uuid = ?", user.UUID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

// ModifyVerification a row
func ModifyVerification(tx *gorm.DB, user *models.User) {
	attrs := map[string]interface{}{
		"verification": user.Verification,
	}

	err := tx.Table(table).
		Model(models.User{}).
		Where("uuid = ?", user.UUID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

// ModifyPassword a row
func ModifyPassword(tx *gorm.DB, user *models.User) {
	attrs := map[string]interface{}{
		"password": user.Password,
	}

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
	result := &pattern{}
	err := tx.Table(table).
		Where("user.uuid = ?", uuid).
		Scan(result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}

	return mapping(tx, result)
}

// GetByAccount return a record found by account (after mapping)
func GetByAccount(tx *gorm.DB, account string) *models.User {
	result := &pattern{}
	err := tx.Table(table).
		Where("user.account = ?", account).
		Scan(result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}

	return mapping(tx, result)
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

func mapping(tx *gorm.DB, pattern *pattern) *models.User {
	return &models.User{
		ID:           pattern.ID,
		UUID:         pattern.UUID,
		Account:      pattern.Account,
		Role:         roledao.GetByUUID(tx, pattern.RoleUUID),
		Password:     pattern.Password,
		Name:         pattern.Name,
		Email:        pattern.Email,
		Status:       pattern.Status,
		Score:        pattern.Score,
		Verification: pattern.Verification,
		CreatedBy:    pattern.UpdatedBy,
		CreatedAt:    time.Now(),
	}
}
