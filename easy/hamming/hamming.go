//Package hamming implements a function for evaluating the Hamming distance of two strings
package hamming

import (
	"fmt"
)

//Distance evaluates the number of differences between two strings
func Distance(a, b string) (int, error) {
	var countDif = 0

	if len(a) != len(b) {
		return 0, fmt.Errorf("Strings %s, %s are not equal length", a, b)
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			countDif++
		}
	}

	return countDif, nil
}
