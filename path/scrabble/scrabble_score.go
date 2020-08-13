//Package scrabble implements a function that scores words in a game of scrabble.
package scrabble

import (
	"unicode"
)

//Score returns a score based on the input string provided
func Score(word string) int {

	var score int = 0

	for _, p := range word {
		l := unicode.ToLower(p)
		switch l {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}

	return score
}
