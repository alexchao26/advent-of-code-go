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

	ans := decompressLength(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

// well...... this is gross.......
func decompressLength(in string, part int) int {
	var decompressedLen int
	for i := 0; i < len(in); {
		switch in[i] {
		case '(':
			// find index of closing paren, then find total length of substring
			relativeCloseIndex := strings.Index(in[i:], ")")
			closeIndex := relativeCloseIndex + i

			var copyLen, repeat int
			fmt.Sscanf(in[i:closeIndex+1], "(%dx%d)", &copyLen, &repeat)

			substring := in[closeIndex+1 : closeIndex+1+copyLen]
			patternLength := len(substring)
			if part == 2 {
				patternLength = decompressLength(substring, 2)
			}
			decompressedLen += patternLength * repeat
			// jump the closed paren (+1) the length of the substring from THIS
			// function call
			i = closeIndex + 1 + len(substring)
		default:
			decompressedLen++
			i++
		}
	}
	return decompressedLen
}
