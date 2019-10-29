package service

import (
	"fmt"
	"reflect"

	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
	userdao "github.com/CRL-Studio/AuthServer/src/dao/gorm/userDao"
	"github.com/CRL-Studio/AuthServer/src/models"
	"github.com/CRL-Studio/AuthServer/src/utils/glossary"
	"github.com/CRL-Studio/AuthServer/src/utils/hash"
	"github.com/CRL-Studio/AuthServer/src/utils/logger"
	uuid "github.com/satori/go.uuid"
)

//CreateUser is
func CreateUser(params interface{}) (result map[string]interface{}, err error) {
	tx := gormdao.DB().Begin()
	log := logger.Log()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error(r)
		}
	}()

	value := reflect.ValueOf(params)

	user := &models.User{}
	user.UUID = uuid.NewV4().String()
	user.Role = &models.Role{UUID: glossary.RoleMember}
	user.Account = value.FieldByName("Account").String()
	user.Password = hash.New(value.FieldByName("Password").String())
	user.Name = value.FieldByName("Name").String()
	user.Email = value.FieldByName("Email").String()
	user.Score = 0
	user.Status = glossary.StatusEnabled
	user.Verification = false
	user.CreatedBy = value.FieldByName("Account").String()

	userdao.New(tx, user)

	return result, nil
}

//CreateUserCheck is
func CreateUserCheck(params interface{}) (result map[string]interface{}, err error) {
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	value := reflect.ValueOf(params)

	fmt.Println(value)
	fmt.Println(tx)

	return result, nil
}

//UserInfo is
func UserInfo(params interface{}) (result map[string]interface{}, err error) {
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	value := reflect.ValueOf(params)

	fmt.Println(value)
	fmt.Println(tx)

	return result, nil
}

//UpdateUserInfo is
func UpdateUserInfo(params interface{}) (result map[string]interface{}, err error) {
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	value := reflect.ValueOf(params)

	fmt.Println(value)
	fmt.Println(tx)

	return result, nil
}

//UpdatePassword is
func UpdatePassword(params interface{}) (result map[string]interface{}, err error) {
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	value := reflect.ValueOf(params)

	fmt.Println(value)
	fmt.Println(tx)

	return result, nil
}

//ResetPassword is
func ResetPassword(params interface{}) (result map[string]interface{}, err error) {
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	value := reflect.ValueOf(params)

	fmt.Println(value)
	fmt.Println(tx)

	return result, nil
}
