package handler

import (
	errorreturn "github.com/CRL-Studio/AuthServer/src/errorReturn"
	"github.com/kataras/iris"
)

// Success is the function to return Success By JSON
func success(ctx iris.Context, data interface{}) {
	ctx.JSON(iris.Map{
		"code":    0,
		"message": "Success",
		"data":    data,
	})
}

// Failed is the function to return Failed By JSON
func failed(ctx iris.Context, err *errorreturn.ErrorOutput) {
	ctx.JSON(iris.Map{
		"code":    err.Code(),
		"message": err.Error(),
		"data":    "",
	})
}
