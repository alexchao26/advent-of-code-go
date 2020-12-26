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

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	var codeChars, stringChars int
	for _, line := range strings.Split(input, "\n") {
		codeChars += len(line)

		for i := 1; i < len(line)-1; i++ {
			switch line[i] {
			case '\\':
				nextChar := line[i+1]
				if nextChar == '\\' || nextChar == '"' {
					i++ // skip an extra character
				} else if nextChar == 'x' {
					i += 3 // skip 2 extra chars
				}
			}
			stringChars++
		}
	}

	return codeChars - stringChars
}

func part2(input string) int {
	var encodedLen, originalLen int
	for _, line := range strings.Split(input, "\n") {
		originalLen += len(line)
		encodedLen += 2 // outer quotes
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '"', '\\':
				encodedLen += 2
			default:
				encodedLen++
			}
		}
	}
	return encodedLen - originalLen
}
