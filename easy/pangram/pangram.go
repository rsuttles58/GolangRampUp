package pangram

import (
	"strings"
)

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

const otherChars = ",._-!?\""

func IsPangram(word string) bool {

	for _, char := range letters {
		if strings.Contains(otherChars, char) {
			continue
		}

		if !strings.Contains(word, strings.ToLower(char)) {
			return false
		}
	}

	return true
}
