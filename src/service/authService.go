package service

import (
	"reflect"

	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
	userdao "github.com/CRL-Studio/AuthServer/src/dao/gorm/userDao"
	rediskeysetdao "github.com/CRL-Studio/AuthServer/src/dao/redis/redisKeySetDao"
	"github.com/CRL-Studio/AuthServer/src/utils/logger"
)

//Login is login
func Login(params interface{}) (result map[string]interface{}, outputError string) {
	tx := gormdao.DB()
	log := logger.Log()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
			outputError = "Unexpected Error"
		}
	}()

	value := reflect.ValueOf(params)
	operator := userdao.GetByAccount(tx, value.FieldByName("Account").String())

	result = map[string]interface{}{
		"Account": operator.Account,
		"Role":    operator.Role,
		"Score":   operator.Score,
	}
	return result, ""
}

//Logout is login
func Logout(params interface{}) (result map[string]interface{}, outputError string) {
	log := logger.Log()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
			outputError = "Unexpected Error"
		}
	}()

	value := reflect.ValueOf(params)

	if err := rediskeysetdao.Del(value.FieldByName("Account").String()); err != nil {
		panic(err)
	}

	return result, ""
}
