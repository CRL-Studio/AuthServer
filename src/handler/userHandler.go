package handler

import (
	errorreturn "github.com/CRL-Studio/AuthServer/src/errorReturn"
	"github.com/CRL-Studio/AuthServer/src/service"
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
)

//CreateHandler to do
func CreateHandler(ctx iris.Context) {
	defer func() {
		if r := recover(); r != nil {
			//logger
		}
	}()
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
		output := errorreturn.Error(errorreturn.GetErrorReturn("Internal", 400, "", &err))
		failed(ctx, output)
		return
	}

	result, err := service.CreateUser(params)

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn("Internal", 436, "", &err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}

//CreateCheckHandler to do
func CreateCheckHandler(ctx iris.Context) {
	type input struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
	}
	params := &input{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
	}
	result, err := service.CreateUserCheck(params)

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn("Internal", 436, "", &err))
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
	result, err := service.UserInfo(params)

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn("Internal", 436, "", &err))
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
	result, err := service.UpdateUserInfo(params)

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn("Internal", 436, "", &err))
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

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn("Internal", 436, "", &err))
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

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn("Internal", 436, "", &err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}
