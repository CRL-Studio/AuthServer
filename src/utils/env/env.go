package env

import (
	"os"
)

// Set is to set Email_Account&Password
func Set() {
	os.Setenv("Gmail_Account", "studycrl@gmail.com")
	os.Setenv("Gmail_Password", "xxxxx")
}
