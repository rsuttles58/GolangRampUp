//Package acronym implements a function that creates abbreviations from string input
package acronym

import (
	"strings"
)

//Abbreviate creates abbreviations from string input
func Abbreviate(s string) string {

	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'"
	abrOutput := strings.Split(s, "")
	var returnS string

	//Warning: I feel dirty for writing this.  This feels like the Tom Hanks raft on CastAway "flimsy" of code.
	for i := 0; i < len(abrOutput); i++ {
		if strings.ContainsAny(abrOutput[i], chars) && i == 0 || strings.ContainsAny(abrOutput[i], chars) &&
			!strings.ContainsAny(abrOutput[i-1], chars) {
			returnS += strings.ToUpper(abrOutput[i])
		}
	}

	return returnS
}
