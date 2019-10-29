package service

import (
	"fmt"
	"reflect"

	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
)

//CreateUser is 
func CreateUser(params interface{}) (result map[string]interface{}, err error) {
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