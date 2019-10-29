package errorreturn

// InternalError to define the internal error
func InternalError(code int, message string) (msg string) {
	switch code {
	case 400:
		msg = "參數錯誤" + message
	case 436:
		msg = "問號錯誤"
	case 500:
		msg = "伺服器錯誤"
	default:
		msg = "無預期錯誤"
	}
	return
}
