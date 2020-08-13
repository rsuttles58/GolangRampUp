package reverse

//Reverse takes a string and returns the list in reverse order
func Reverse(str string) string {

	revStr := []rune(str)

	for ch1, ch2 := 0, len(revStr)-1; ch1 < ch2; ch1, ch2 = ch1+1, ch2-1 {
		//Swap letters, starting first with last
		revStr[ch1], revStr[ch2] = revStr[ch2], revStr[ch1]
	}

	return string(revStr)
}
