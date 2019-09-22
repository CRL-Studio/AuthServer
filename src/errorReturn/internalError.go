package errorreturn

// InternalError to define the internal error
func InternalError(code int, message string, err *error) (msg string) {
	switch code {
	case 400:
		msg = "I參數錯誤" + message
	case 436:
		msg = "I問號錯誤"
	case 500:
		msg = "I錯誤"
	default:
		msg = "I無預期錯誤"
	}
	return
}
