package handler

import (
	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
	userdao "github.com/CRL-Studio/AuthServer/src/dao/gorm/userDao"
	redisverificationdao "github.com/CRL-Studio/AuthServer/src/dao/redis/redisVerificationDao"
	errorreturn "github.com/CRL-Studio/AuthServer/src/errorReturn"
	"github.com/CRL-Studio/AuthServer/src/service"
	"github.com/CRL-Studio/AuthServer/src/utils/glossary"
	"github.com/CRL-Studio/AuthServer/src/utils/logger"
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
)

var log = logger.Log()

//CreateHandler to do
func CreateHandler(ctx iris.Context) {
	//parameters
	type input struct {
		Account  string `valid:"required"`
		Password string `valid:"required, password"`
		Name     string `valid:"required"`
		Email    string `valid:"email"`
	}
	params := &input{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
		Name:     ctx.FormValue("Name"),
		Email:    ctx.FormValue("Email"),
	}

	// validate
	if _, err := govalidator.ValidateStruct(params); err != nil {
		errorOutput := err.Error()
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "", errorOutput))
		failed(ctx, output)
		return
	}

	//check
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()
	account := userdao.Get(tx, &userdao.QueryModel{Account: params.Account})
	if account != nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "該帳號已被使用", ""))
		failed(ctx, output)
		return
	}

	//logic
	result, err := service.CreateUser(params)
	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	//Success
	success(ctx, result)
	return
}

//VerificationCheckHandler to do
func VerificationCheckHandler(ctx iris.Context) {
	//log := logger.Log()
	//parameters
	type input struct {
		Account      string `valid:"required"`
		Verification string `valid:"required"`
	}

	params := &input{
		Account:      ctx.FormValue("Account"),
		Verification: ctx.FormValue("Verification"),
	}

	// validate
	if _, err := govalidator.ValidateStruct(params); err != nil {
		errorOutput := err.Error()
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "", errorOutput))
		failed(ctx, output)
		return
	}

	//check
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()
	account := userdao.Get(tx, &userdao.QueryModel{Account: params.Account})
	if account == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "找不到此帳號", ""))
		failed(ctx, output)
		return
	}
	redisKey := "Ver" + params.Account
	verificationcode, _ := redisverificationdao.Get(redisKey)
	if verificationcode == "" {
		log.Error("找不到" + params.Account + "驗證碼")
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "找不到驗證碼", ""))
		failed(ctx, output)
		return
	}

	//logic
	result, err := service.UserVerificationCheck(params)

	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	//Success
	success(ctx, result)
	return
}

//ResendVerificationHandler to do
func ResendVerificationHandler(ctx iris.Context) {
	//log := logger.Log()
	//parameters
	type input struct {
		Account string `valid:"required"`
	}

	params := &input{
		Account: ctx.FormValue("Account"),
	}

	// validate
	if _, err := govalidator.ValidateStruct(params); err != nil {
		errorOutput := err.Error()
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "", errorOutput))
		failed(ctx, output)
		return
	}

	//check
	tx := gormdao.DB()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()
	account := userdao.Get(tx, &userdao.QueryModel{Account: params.Account})
	if account == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "找不到此帳號", ""))
		failed(ctx, output)
		return
	}

	//logic
	result, err := service.ResendVerificationCode(params)
	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	//Success
	success(ctx, result)
	return
}

//UserInfoHandler to do
func UserInfoHandler(ctx iris.Context) {
	// parameters
	type input struct {
		Account string `valid:"required"`
	}
	params := &input{
		Account: ctx.FormValue("Account"),
	}

	// validate
	if _, err := govalidator.ValidateStruct(params); err != nil {
		errorOutput := err.Error()
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "", errorOutput))
		failed(ctx, output)
		return
	}
	//check
	tx := gormdao.DB()
	account := userdao.Get(tx, &userdao.QueryModel{Account: params.Account})
	if account == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "找不到此帳號", ""))
		failed(ctx, output)
		return
	}

	//logic
	result, err := service.UserInfo(params)
	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	//Success
	success(ctx, result)
	return
}

//UpdateUserInfoHandler to do
func UpdateUserInfoHandler(ctx iris.Context) {
	// parameters
	type input struct {
		Account string `valid:"required"`
		Name    string `valid:"-"`
	}
	params := &input{
		Account: ctx.FormValue("Account"),
		Name:    ctx.FormValue("Name"),
	}

	// validate
	if _, err := govalidator.ValidateStruct(params); err != nil {
		errorOutput := err.Error()
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "", errorOutput))
		failed(ctx, output)
		return
	}

	//check
	tx := gormdao.DB()
	account := userdao.Get(tx, &userdao.QueryModel{Account: params.Account})
	if account == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "找不到此帳號", ""))
		failed(ctx, output)
		return
	}

	//logic
	result, err := service.UpdateUserInfo(params)
	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	//success
	success(ctx, result)
	return
}

//UpdatePasswordHandler to do
func UpdatePasswordHandler(ctx iris.Context) {
	// parameters
	type input struct {
		Account     string `valid:"required"`
		Password    string `valid:"required"`
		NewPassword string `valid:"required"`
	}
	params := &input{
		Account:     ctx.FormValue("Account"),
		Password:    ctx.FormValue("Password"),
		NewPassword: ctx.FormValue("NewPassword"),
	}

	// validate
	if _, err := govalidator.ValidateStruct(params); err != nil {
		errorOutput := err.Error()
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "", errorOutput))
		failed(ctx, output)
		return
	}

	//check
	tx := gormdao.DB()
	account := userdao.Get(tx, &userdao.QueryModel{Account: params.Account})
	if account == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "找不到此帳號", ""))
		failed(ctx, output)
		return
	}

	//logic
	result, err := service.UpdatePassword(params)
	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	//success
	success(ctx, result)
	return
}

//ResetPasswordHandler to do
func ResetPasswordHandler(ctx iris.Context) {
	// parameters
	type input struct {
		Account string `valid:"required"`
		Email   string `valid:"required"`
	}
	params := &input{
		Account: ctx.FormValue("Account"),
		Email:   ctx.FormValue("Email"),
	}

	// validate
	if _, err := govalidator.ValidateStruct(params); err != nil {
		errorOutput := err.Error()
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "", errorOutput))
		failed(ctx, output)
		return
	}

	// check
	tx := gormdao.DB()
	user := userdao.Get(tx, &userdao.QueryModel{Account: params.Account})
	if user == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "找不到此帳號", ""))
		failed(ctx, output)
		return
	}
	if user.Email != params.Email {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "該會員驗證信箱錯誤", ""))
		failed(ctx, output)
		return
	}

	// logic
	result, err := service.ResetPassword(params)
	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	// success
	success(ctx, result)
	return
}
