package route

import (
	"fmt"

	"github.com/CRL-Studio/AuthServer/src/handler"
	"github.com/CRL-Studio/AuthServer/src/middleware"
	"github.com/CRL-Studio/AuthServer/src/utils/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

var addr = fmt.Sprintf("%v:%v", config.Get("server.host"), config.Get("server.port"))

// Run maps the routing path and keeps listening for request
func Run() {
	app := iris.New()

	//CORS
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CorsMiddleware)

	app.Get("/", hero.Handler(handler.HelloHandler))

	v1 := app.Party("/v1", hero.Handler(middleware.AuthMiddleware))
	{
		//登入
		v1.Post("/login", hero.Handler(handler.LoginHandler))
		//登出
		v1.Get("/logout", hero.Handler(handler.LogoutHandler))
		user := v1.Party("/user")
		{
			//創建新帳號
			user.Post("/create", hero.Handler(handler.CreateHandler))
			//創建新帳號-憑證碼驗證
			user.Post("/create-verification", hero.Handler(handler.CreateCheckHandler))
			//取得帳號資訊
			user.Get("/info", hero.Handler(handler.UserInfoHandler))
			//修改帳號資訊
			user.Patch("/update-info", hero.Handler(handler.UpdateUserInfoHandler))
			//修改密碼
			user.Patch("/update-password", hero.Handler(handler.UpdatePasswordHandler))
			//重設密碼
			user.Patch("/reset-password", hero.Handler(handler.ResetPasswordHandler))
		}
	}

	app.Run(iris.Addr(addr))
}
