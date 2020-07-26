//Package strand implements a function that converts DNA strands to RNA
package strand

import (
	"strings"
	"fmt"
)

//ToRNA provides the RNA transcription from DNA -> RNA
func ToRNA(dna string) string {

	rnaOutput := strings.Split(dna, "")

	for i := 0; i < len(rnaOutput); i++ {
		switch rnaOutput[i] {
		case "G":
			rnaOutput[i] = "C"
		case "C":
			rnaOutput[i] = "G"
		case "T":
			rnaOutput[i] = "A"
		case "A":
			rnaOutput[i] = "U"
		}
	}

	fmt.Println(rnaOutput)
	return strings.Join(rnaOutput, "")
}
