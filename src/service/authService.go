package service

import (
	"fmt"
	"reflect"

	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
)

//Login
func Login(params interface{}) (result map[string]interface{}, err *error) {
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
