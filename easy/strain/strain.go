package strain

//Ints is a type for collections of integers
type Ints []int

//Lists is a type for collections of integer slices
type Lists [][]int

//Strings is a type for collections of integers
type Strings []string

//Keep creates a new collection based on filter criteria
func (ints Ints) Keep(filter func(int) bool) Ints {
	var keepers Ints
	for _, i := range ints {
		if filter(i) {
			keepers = append(keepers, i)
		}
	}
	return keepers
}

//Discard creates a new collection based on filter criteria
func (ints Ints) Discard(filter func(int) bool) Ints {
	var keepers Ints
	for _, i := range ints {
		if !filter(i) {
			keepers = append(keepers, i)
		}
	}
	return keepers
}

//Keep creates a new collection of slices based on filter criteria
func (list Lists) Keep(filter func([]int) bool) Lists {
	var keepers Lists
	for _, list := range list {
		if filter(list) {
			keepers = append(keepers, list)
		}
	}

	return keepers
}

//Keep creates a new collection of strings based on filter criteria
func (str Strings) Keep(filter func(string) bool) Strings {
	var returnStr Strings
	for _, str := range str {
		if filter(str) {
			returnStr = append(returnStr, str)
		}
	}
	return returnStr
}
