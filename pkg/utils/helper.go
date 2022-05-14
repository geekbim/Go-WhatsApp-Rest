package utils

import "strconv"

func StringToInt(value string) (conv int) {
	if value != "" {
		conv, _ = strconv.Atoi(value)
	}
	if value == "" {
		conv = 0
	}
	return
}
