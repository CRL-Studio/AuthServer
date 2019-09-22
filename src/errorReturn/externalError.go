package errorreturn

// ExternalError to define the enternal error
func ExternalError(code int, message string, err *error) (msg string) {
	switch code {
	case 400:
		msg = "E參數錯誤" + message
	case 436:
		msg = "E問號錯誤"
	case 500:
		msg = "E錯誤"
	default:
		msg = "E無預期錯誤"
	}
	return
}
