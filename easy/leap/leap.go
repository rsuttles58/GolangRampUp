
//Package leap implements a function return true or false depending on if the year is a leap year or not.
package leap

//IsLeapYear returns true or false based on the year passed to the function
func IsLeapYear(year int) bool {

	//if the year is evenly divided by 4 and not also evenly divided by 100, return true.
	if year % 4 == 0 && year % 100 != 0 {
		return true
	}

	//if the year is evenly divided by 4 and also divided by 100 and 400, return true.
	if year % 4 == 0 && year % 100 == 0 && year % 400 == 0 {
		return true
	}

	//If neither of those, return false
	return false
}
