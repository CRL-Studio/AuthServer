package handler

import (
	errorreturn "github.com/CRL-Studio/AuthServer/src/errorReturn"
	"github.com/CRL-Studio/AuthServer/src/service"
	"github.com/kataras/iris"
)

//LoginHandler to do Login
func LoginHandler(ctx iris.Context) {
	type input struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
	}
	params := &input{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
	}
	result, err := service.Login(params)

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn("Int", 436, "", err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}
