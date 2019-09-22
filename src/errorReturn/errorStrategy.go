package errorreturn

import (
	"strings"
)

// ErrorStrategy to check the error function
type ErrorStrategy interface {
	errorReturn() *ErrorOutput
}

// ErrorInput to make the error input
type ErrorInput struct {
	code      int
	message   string
	err       *error
	errorType string
}

// ErrorOutput to make the error output
type ErrorOutput struct {
	code    int
	message string
}

func (errorInput ErrorInput) errorReturn() *ErrorOutput {
	var errorOutput ErrorOutput
	if strings.EqualFold(errorInput.errorType, "Internal") == true {
		//logger
	} else if strings.EqualFold(errorInput.errorType, "EXternal") == true {
		//logger
	} else {
		//logger
	}
	errorOutput.code = errorInput.code
	errorOutput.message = errorInput.message
	return &errorOutput
}

// Error to choose a Error Return
func Error(strategy ErrorStrategy) *ErrorOutput {
	if strategy != nil {
		return strategy.errorReturn()
	}
	return strategy.errorReturn()
}

// GetErrorReturn to choose the error
func GetErrorReturn(errorType string, code int, message string, err *error) ErrorStrategy {
	var msg string
	switch errorType {
	case "Internal":
		msg = InternalError(code, message, err)
	case "External":
		msg = ExternalError(code, message, err)
	default:
		code = 999
		msg = "Unexpected Error"
	}
	return &ErrorInput{code: code, message: msg, err: err, errorType: errorType}
}

// Error return error message
func (err *ErrorOutput) Error() string {
	return err.message
}

// Code return error code
func (err *ErrorOutput) Code() int {
	return err.code
}
