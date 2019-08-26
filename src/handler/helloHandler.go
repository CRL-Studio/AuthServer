package handler

import "github.com/kataras/iris"

// HelloHandler return hello message
func HelloHandler(ctx iris.Context) {

	ctx.JSON(iris.Map{
		"code":    0,
		"message": "success",
	})

	return
}
