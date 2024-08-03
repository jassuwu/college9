package two

import (
	"strings"
)

func ValidateInput(fname, address, email string) (isValid bool, output string) {
	valid := true
	flag := true

	if len(fname) < 4 || len(fname) > 30 {
		output += "Invalid first name\n"
		valid = false
	}
	if len(address) < 4 || len(address) > 100 {
		output += "Invalid address length\n"
		valid = false
	}
	if len(email) < 8 || len(email) > 100 {
		output += "Invalid Email length\n"
		flag = false
		valid = false
	}
	if flag {
		if !strings.Contains(email, ".") || !strings.Contains(email, "@") {
			output += "Email must contain . and @ characters\n"
			valid = false
		}
	}
	if valid {
		output += "First name: " + fname + " \t Address: " + address + " \t Email: " + email
	}
	return valid, output
}
