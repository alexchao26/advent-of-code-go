package algos

import "strings"

// SplitStringOn is like strings.Split but takes in a slice of strings that are
// all used as dividers in the incoming string
func SplitStringOn(in string, cutset []string) []string {
	parts := strings.Split(in, cutset[0])
	cutset = cutset[1:]
	var done bool
	for !done && len(cutset) > 0 {
		divider := cutset[0]
		cutset = cutset[1:]
		var newParts []string
		for _, oldPart := range parts {
			newParts = append(newParts, strings.Split(oldPart, divider)...)
		}
		parts = newParts
	}
	return parts
}
