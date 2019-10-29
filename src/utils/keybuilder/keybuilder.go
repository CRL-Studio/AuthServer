package keybuilder

const (
	jti          = "JTI"
	refreshToken = "REFRESH_TOKEN"
	secret       = "SECRET"
	set          = "SET"
	permission   = "PERMISSION"
	notify       = "NOTIFY"
)

// Jti builds a key looks like {account}-JTI
func Jti(account string) string {
	return dash(
		account,
		jti,
	)
}

// RefreshToken builds a key looks like {account}-REFRESH_TOKEN
func RefreshToken(account string) string {
	return dash(
		account,
		refreshToken,
	)
}

// Secret builds a key looks like {account}-SECRET
func Secret(account string) string {
	return dash(
		account,
		secret,
	)
}

// KeySet builds a key looks like {prefix}-SET
func KeySet(prefix string) string {
	return dash(
		prefix,
		set,
	)
}

// RolePermission builds a key looks like {role}-{method}-PERMISSION
func RolePermission(role, method string) string {
	return dash(
		role,
		method,
		permission,
	)
}

// Notify builds a key looks like NOTIFY
func Notify(orderNo string) string {
	return dash(
		orderNo,
		notify,
	)
}

func dash(params ...string) string {
	len := len(params)

	if len == 0 {
		return ""
	}

	var result = params[0]
	for i := 1; i < len; i++ {
		result += "-" + params[i]
	}
	return result
}
