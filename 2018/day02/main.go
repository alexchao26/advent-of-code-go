package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	var twos, threes int
	for _, boxID := range strings.Split(input, "\n") {
		charCounts := getCharCount(boxID)
		for _, v := range charCounts {
			if v == 2 {
				twos++
				break
			}
		}
		for _, v := range charCounts {
			if v == 3 {
				threes++
			}
		}
	}

	return twos * threes
}

func getCharCount(box string) map[rune]int {
	chars := make(map[rune]int)
	for _, c := range box {
		chars[c]++
	}
	return chars
}

func part2(input string) string {
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if sameChars := getSameCharacters(lines[i], lines[j]); sameChars != "" {
				return sameChars
			}
		}
	}
	return ""
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
