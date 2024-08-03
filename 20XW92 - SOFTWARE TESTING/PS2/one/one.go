package one

import (
	"math"
	"strings"
)

func CountChar(ip string) (vCount, tCount int) {
	ip = strings.ToLower(ip)
	for _, chr := range ip {
		if (chr >= 'a') && (chr <= 'z') && (tCount < math.MaxInt) {
			tCount += 1
			if (chr == 'a') || (chr == 'e') || (chr == 'i') || (chr == 'o') || (chr == 'u') {
				vCount += 1
			}
		}
	}
	return
}
