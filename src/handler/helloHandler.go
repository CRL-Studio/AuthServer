package handler

import (
	errorreturn "github.com/CRL-Studio/AuthServer/src/errorReturn"
	"github.com/kataras/iris"
)

// HelloHandler return hello message
func HelloHandler(ctx iris.Context) {

	output := errorreturn.Error(errorreturn.GetErrorReturn("Int", 436, "", nil))
	failed(ctx, output)
	return
	/*
		success(ctx, "hello")

		return
	*/
}
