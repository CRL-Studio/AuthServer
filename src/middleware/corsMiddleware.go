package middleware

import (
	"github.com/CRL-Studio/AuthServer/src/utils/config"
	"github.com/kataras/iris"
)

// CorsMiddleware enalbe cors
func CorsMiddleware(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", config.Get("http.header.Access-Control-Allow-Origin").(string))
	ctx.Header("Access-Control-Allow-Headers", config.Get("http.header.Access-Control-Allow-Headers").(string))
	ctx.Header("Access-Control-Allow-Methods", config.Get("http.header.Access-Control-Allow-Methods").(string))

	if ctx.Method() == iris.MethodOptions {
		ctx.StatusCode(iris.StatusOK)
		return
	}

	ctx.Next()
}
