//Package twofer implements a simple function for saying "One for __, one for me".
package twofer

//ShareWith returns "One for X, one for me" if a name parameter is passed, otherwise provides "One for you, one for me."
func ShareWith(name string) string {

	if name == "" {
		return "One for you, one for me."
	}

	return "One for " + name + ", one for me."
}
