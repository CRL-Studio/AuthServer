package handler

import (
	gormdao "github.com/CRL-Studio/AuthServer/src/dao/gorm"
	userdao "github.com/CRL-Studio/AuthServer/src/dao/gorm/userDao"
	errorreturn "github.com/CRL-Studio/AuthServer/src/errorReturn"
	"github.com/CRL-Studio/AuthServer/src/service"
	"github.com/CRL-Studio/AuthServer/src/utils/glossary"
	"github.com/CRL-Studio/AuthServer/src/utils/hash"
	"github.com/CRL-Studio/AuthServer/src/utils/logger"
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
)

//LoginHandler to do Login
func LoginHandler(ctx iris.Context) {
	tx := gormdao.DB()
	log := logger.Log()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
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
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusBadRequest, "帳號或密碼未填寫", &err))
		failed(ctx, output)
		return
	}

	operator := userdao.GetByAccount(tx, params.Account)

	if operator == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusUnauthorized, "該帳號不存在", nil))
		failed(ctx, output)
		return
	}

	if !hash.Verify(params.Password, operator.Password) {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusUnauthorized, "密碼錯誤", nil))
		failed(ctx, output)
		return
	}

	if operator.Verification == false {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusAccepted, "該帳號未驗證", nil))
		failed(ctx, output)
	}

	result, err := service.Login(params)

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", &err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}

//LogoutHandler to do Login
func LogoutHandler(ctx iris.Context) {
	tx := gormdao.DB()
	log := logger.Log()
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		}
	}()
	type input struct {
		Account string `valid:"required"`
	}
	params := &input{
		Account: ctx.URLParam("Account"),
	}

	operator := userdao.GetByAccount(tx, params.Account)

	if operator == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusUnauthorized, "該帳號不存在", nil))
		failed(ctx, output)
		return
	}

	result, err := service.Logout(params)

	if err == nil {
		output := errorreturn.Error(errorreturn.GetErrorReturn(glossary.ErrorTypeInternal, iris.StatusInternalServerError, "", &err))
		failed(ctx, output)
		return
	}

	success(ctx, result)
	return
}
