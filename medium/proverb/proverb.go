//Package proverb implements a function to create a custom proverb based on the input parameter.
package proverb

//Proverb iterates over the string array parameter and returns a proverb utilizing those values
func Proverb(rhyme []string) []string {

	var proverbArr []string

	for i := 0; i < len(rhyme); i++ {

		if rhyme[i] == rhyme[len(rhyme)-1] {
			proverbArr = append(proverbArr, "And all for the want of a "+rhyme[0]+".")
		} else {
			proverbArr = append(proverbArr, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
		}

	}

	return proverbArr
}
