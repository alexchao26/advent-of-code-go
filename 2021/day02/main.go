package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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

	ans := day2(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func day2(input string, part int) int {
	var horiz, depth int
	var aim int

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		dir := parts[0]
		dist := cast.ToInt(parts[1])

		if part == 1 {
			switch dir {
			case "down":
				depth += dist
			case "up":
				depth -= dist
			case "forward":
				horiz += dist
			}
		} else {
			switch dir {
			case "down":
				aim += dist
			case "up":
				aim -= dist
			case "forward":
				horiz += dist
				depth += aim * dist
			}
		}
	}

	return horiz * depth
}
