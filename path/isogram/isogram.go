//Package isogram implements a function that evaluates whether a string is an isogram
package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram checks whether the word parameter is an isogram
func IsIsogram(word string) bool {

	isoArr := make([]rune, len(word))
	strArr := []rune(word)

	for _, char := range strArr {
		if char == '-' || char == ' ' {
			continue
		}

		lowerChar := unicode.ToLower(char)
		if strings.ContainsRune(string(isoArr), lowerChar) {
			return false
		}

		isoArr = append(isoArr, lowerChar)

	}

	return true
}
