package main

import (
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

// brute force it...
// assume only one valid answer
func main() {
	input := util.ReadFile("../input.txt")
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if sameChars := getSameCharacters(lines[i], lines[j]); sameChars != "" {
				fmt.Println("common letters are", sameChars)
				return
			}
		}
	}
}

func getSameCharacters(str1, str2 string) string {
	var mismatchSeen bool
	var sameChars string
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			sameChars += string(str1[i])
		} else if mismatchSeen {
			// if a mismatch has already been seen, then it's 2 characters off
			// return an empty string
			return ""
		} else {
			mismatchSeen = true
		}
	}
	return sameChars
}
