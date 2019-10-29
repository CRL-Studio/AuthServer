package errorreturn

import "github.com/kataras/iris"

// InternalError to define the internal error
func InternalError(code int, message string) (msg string) {
	switch code {
	case iris.StatusAccepted:
		msg = "請求有異：" + message
	case iris.StatusBadRequest:
		msg = "參數錯誤：" + message
	case iris.StatusUnauthorized:
		msg = "驗證錯誤：" + message
	case iris.StatusInternalServerError:
		msg = "伺服器錯誤"
	default:
		msg = "非預期錯誤"
	}
	return
}
