package utils

import "unicode"

func ValidateNumber(strArr string) bool {
	for _, element := range strArr {
		if !unicode.IsNumber(element) {
			return false
		}
	}
	return true
}
