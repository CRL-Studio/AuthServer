package auth

var (
	account string
)

// Set store token claim
func Set(params map[string]interface{}) {
	account = params["account"].(string)
}
