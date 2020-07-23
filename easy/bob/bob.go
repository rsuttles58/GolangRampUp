//Package bob implements a function that provides a response to remark text.
package bob

import (
	"strings"
)

//Hey provides "Bob's" response to comments made to him
func Hey(remark string) string {
	var remarks = strings.TrimSpace(remark)

	var question = false
	var yelling = false
	var hasChars = false
	
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	if len(remarks) == 0 {
		return "Fine. Be that way!"
	}

	if strings.ContainsAny(remarks, letters){
		hasChars = true
	}
	//evaluate if remark and the uppercase version are the same to see if remark is yelling.
	if (strings.ToUpper(remarks) == remarks) && hasChars == true{
		yelling = true
	}
	
	if remarks[len(remarks)-1] == '?' {
		question = true
	}

	//evaluate what to respond based on booleans that are true
	switch {
		case question == true && yelling == true:
			return "Calm down, I know what I'm doing!"

		case question == true:
			return "Sure."

		case yelling == true:
			return "Whoa, chill out!"
		
		default:
			return "Whatever."
	}
}
