package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	// packet starts w/ 4 characters that are all different
	for i := 0; i+4 <= len(input); i++ {
		if allDifferentLetters(input[i : i+4]) {
			return i + 4
		}
	}

	return -1
}

// lazy but easier than sliding window...
func allDifferentLetters(str string) bool {
	// if len(str) != 4 {
	// 	panic(fmt.Sprintf("invalid length %q", str))
	// }
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}

func part2(input string) int {
	// wow super lazy but fast to write... ok
	for i := 0; i+14 <= len(input); i++ {
		if allDifferentLetters(input[i : i+14]) {
			return i + 14
		}
	}

	return -1
}
