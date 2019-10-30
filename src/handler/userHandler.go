package handler

import (
	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
	userdao "github.com/CRL-Studio/AuthServer/src/dao/gorm/userDao"
	errorreturn "github.com/CRL-Studio/AuthServer/src/errorReturn"
	"github.com/CRL-Studio/AuthServer/src/service"
	"github.com/CRL-Studio/AuthServer/src/utils/glossary"
	"github.com/CRL-Studio/AuthServer/src/utils/logger"
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
)

//CreateHandler to do
func CreateHandler(ctx iris.Context) {
	tx := gormdao.DB()
	log := logger.Log()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()
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

	account := userdao.Get(tx, &userdao.QueryModel{Account: params.Account})
	if account != nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "該帳號已被使用", ""))
		failed(ctx, output)
		return
	}

	result, err := service.CreateUser(params)

	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}

//VerificationCheckHandler to do
func VerificationCheckHandler(ctx iris.Context) {
	tx := gormdao.DB()
	log := logger.Log()

	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()

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

	account := userdao.Get(tx, &userdao.QueryModel{Account: params.Account})
	if account == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "找不到此帳號", ""))
		failed(ctx, output)
		return
	}

	result, err := service.UserVerificationCheck(params)

	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}

//UserInfoHandler to do
func UserInfoHandler(ctx iris.Context) {
	type input struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
	}
	params := &input{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
	}

	// validate
	if _, err := govalidator.ValidateStruct(params); err != nil {
		errorOutput := err.Error()
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "", errorOutput))
		failed(ctx, output)
		return
	}

	result, err := service.UserInfo(params)

	if err != "" {

		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}

//UpdateUserInfoHandler to do
func UpdateUserInfoHandler(ctx iris.Context) {
	type input struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
	}
	params := &input{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
	}

	// validate
	if _, err := govalidator.ValidateStruct(params); err != nil {
		errorOutput := err.Error()
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "", errorOutput))
		failed(ctx, output)
		return
	}
	result, err := service.UpdateUserInfo(params)

	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}

//UpdatePasswordHandler to do
func UpdatePasswordHandler(ctx iris.Context) {
	type input struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
	}
	params := &input{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
	}
	result, err := service.UpdatePassword(params)

	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}

//ResetPasswordHandler to do
func ResetPasswordHandler(ctx iris.Context) {
	type input struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
	}
	params := &input{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
	}
	result, err := service.ResetPassword(params)

	if err != "" {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}
