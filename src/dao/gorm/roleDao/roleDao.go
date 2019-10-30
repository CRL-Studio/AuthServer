package roledao

import (
	"time"

	"github.com/CRL-Studio/AuthServer/src/models"

	"github.com/jinzhu/gorm"
)

const (
	table = "role"
)

type pattern struct {
	ID        int64      `gorm:"column:id; primary_key"`
	UUID      string     `gorm:"column:uuid; unique_index"`
	Name      string     `gorm:"column:name; unique_index"`
	Code      string     `gorm:"column:code; unique_index"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	CreatedBy string     `gorm:"column:created_by"`
	UpdatedBy string     `gorm:"column:updated_by"`
}

//QueryModel is used for possible query column
type QueryModel struct {
	UUID string
	Code string
}

// New a record
func New(tx *gorm.DB, role *models.Role) {
	err := tx.Table(table).
		Create(&pattern{
			UUID:      role.UUID,
			Name:      role.Name,
			Code:      role.Code,
			CreatedBy: "Seeder",
		}).Error

	if err != nil {
		panic(err)
	}
}

// GetByUUID get a record by uuid
func GetByUUID(tx *gorm.DB, uuid string) *models.Role {
	result := &pattern{}
	err := tx.Table(table).
		Where("uuid = ?", uuid).
		Scan(result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return mapping(result)
}

// GetByCode get a record by code
func GetByCode(tx *gorm.DB, code string) *models.Role {
	result := &pattern{}
	err := tx.Table(table).
		Where("code = ?", code).
		Scan(result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return mapping(result)
}

// GetAll get all roles
func GetAll(tx *gorm.DB) []models.Role {
	var rows []pattern
	err := tx.Table(table).
		Scan(&rows).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	if err != nil {
		panic(err)
	}

	var result []models.Role
	for _, row := range rows {
		temp := mapping(&row)
		result = append(result, *temp)
	}
	return result
}

// Modify a record
func Modify(tx *gorm.DB, role *models.Role) {
	attrs := map[string]interface{}{
		"name":       role.Name,
		"code":       role.Code,
		"updated_by": "Seeder",
	}
	err := tx.Table(table).
		Model(pattern{}).
		Where("uuid = ?", role.UUID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

// Delete a record
func Delete(tx *gorm.DB, role *models.Role) {
	attrs := map[string]interface{}{
		"deleted_by": "Seeder",
		"deleted_at": time.Now(),
	}
	err := tx.Table(table).
		Model(pattern{}).
		Where("uuid = ?", role.UUID).
		Updates(attrs).Error

	if err != nil {
		panic(err)
	}
}

//Get is only one
func Get(tx *gorm.DB, query *QueryModel) *models.Role {
	result := &pattern{}
	err := tx.Table(table).
		Scopes(queryChain(tx, query)).
		Scan(result).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil
	}

	if err != nil {
		panic(err)
	}
	return mapping(result)
}

func queryChain(tx *gorm.DB, query *QueryModel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Scopes(RoleUUIDEqualScope(query.UUID)).
			Scopes(RoleCodeEqualScope(query.Code))
	}
}

func mapping(pattern *pattern) *models.Role {
	return &models.Role{
		ID:        pattern.ID,
		UUID:      pattern.UUID,
		Name:      pattern.Name,
		Code:      pattern.Code,
		CreatedAt: pattern.CreatedAt,
		CreatedBy: pattern.CreatedBy,
		UpdatedAt: pattern.UpdatedAt,
		UpdatedBy: pattern.UpdatedBy,
	}
}
