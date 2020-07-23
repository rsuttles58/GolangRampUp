//Package raindrops implements a function similar to fizzbuzz
package raindrops

import (
	"strconv"
)

//Convert is a fizzbuzz relative
func Convert(num int) string {

	//short circuit logic that if the num doesn't match the 3/5/7 pattern
	if num % 3 != 0 && num % 5 != 0 && num % 7 != 0{
		return strconv.Itoa(num)
	}

	var returnStr string
	
	//if statements for 3, 5, 7
	if num % 3 == 0 {
		returnStr += "Pling"
	}

	if num % 5 == 0 {
		returnStr += "Plang"
	}

	if num % 7 == 0 {
		returnStr += "Plong"
	}

	return returnStr
}