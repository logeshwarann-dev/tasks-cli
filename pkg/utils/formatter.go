package utils

import (
	"fmt"
	"log"
	"os"
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

func ClearContentsOfFile(filePath string) error {
	if err := os.Truncate(filePath, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
		return err
	}
	return nil
}
