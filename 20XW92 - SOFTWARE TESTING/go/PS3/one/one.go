package one

func ValidPasswordLength(passwd string) bool {
	if len(passwd) >= 6 && len(passwd) <= 12 {
		return true
	}
	return false
}
