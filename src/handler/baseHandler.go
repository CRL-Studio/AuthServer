package handler

import (
	errorreturn "github.com/CRL-Studio/AuthServer/src/errorReturn"
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
)

func init() {
	govalidator.CustomTypeTagMap.Set("password", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		hasSmall := govalidator.Matches(i.(string), "^.*[a-z]{1,}.*$")
		hasNumber := govalidator.Matches(i.(string), "^.*[0-9]{1,}.*$")
		checkLength := govalidator.StringLength(i.(string), "6", "16")
		return checkLength && hasSmall && hasNumber
	}))
}

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
