package handler

import (
	errorreturn "github.com/CRL-Studio/AuthServer/src/errorReturn"
	"github.com/CRL-Studio/AuthServer/src/service"
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
)

//LoginHandler to do Login
func LoginHandler(ctx iris.Context) {
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

	result, err := service.Login(params)

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn("Internal", 436, "", &err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}

//LogoutHandler to do Login
func LogoutHandler(ctx iris.Context) {
	type input struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
	}
	params := &input{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
	}
	result, err := service.Logout(params)

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn("Internal", 436, "", &err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}
