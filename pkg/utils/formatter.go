package utils

import (
	"fmt"
	"unicode"
)

func ValidateNumber(strArr string) bool {
	for _, element := range strArr {
		if !unicode.IsNumber(element) {
			return false
		}
	}
	return true
}

func ConvSliceToStr(strSlice []string) string {

	var fullString string
	for _, element := range strSlice {
		fullString = fmt.Sprintf("%v %v", fullString, element)
	}

	return fullString

}
