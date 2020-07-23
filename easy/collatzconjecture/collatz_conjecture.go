//Package collatzconjecture implements a function to provide number of steps for number to equal 1:
//https://en.wikipedia.org/wiki/Collatz_conjecture
package collatzconjecture

import (
	"fmt"
)

//CollatzConjecture evaluates how many steps are needed for the num to equal 1 based on starting input parameter
func CollatzConjecture(n int) (int, error) {

	var counter = 0
	var num = n

	//evaluate if num is 1.  If true, return 0.
	if num == 1 {
		return 0, nil
	}

	if num < 0 {
		return 0, fmt.Errorf("Num %d cannot be less than 0", num)
	}

	if num == 0 {
		return 0, fmt.Errorf("num %d cannot be zero", num)
	}

	for {
		if num == 1 {
			break
		}

		counter++

		if num%2 == 0 {
			num = num / 2
		} else if num%2 == 1 {
			num = num*3 + 1
		}
	}

	return counter, nil
}
