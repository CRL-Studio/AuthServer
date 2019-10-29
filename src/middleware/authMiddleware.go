package middleware

import (
	"regexp"

	"github.com/kataras/iris"
)

var (
	whiteList = []string{
		`/v1/login`,
	}
)

// AuthMiddleware validate token before handler
func AuthMiddleware(ctx iris.Context) {
	for _, path := range whiteList {
		pathRegexp := regexp.MustCompile(path)
		if pathRegexp.MatchString(ctx.Path()) {
			ctx.Next()
			return
		}
	}
	/*
		if ok := checkToken(ctx); !ok {
			ctx.JSON(iris.Map{
				"code":    438,
				"message": "Access Token Error",
				"data":    []string{},
			})
			return
		}
	*/
	ctx.Next()

}

/*
func checkToken(ctx iris.Context) bool {
	tokenString, err := jwtmiddleware.FromAuthHeader(ctx)
	if err != nil {
		return false
	}

	claim, err := token.Parse(tokenString)
	if err != nil {
		return false
	}

	auth.Set(claim)
	ctx.Values().Set("account", claim["account"])

	return true
}
*/
