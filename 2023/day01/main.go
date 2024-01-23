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
	var sum int
	for _, line := range strings.Split(input, "\n") {
		var tens, ones int
		for i := 0; i < len(line); i++ {
			if strings.ContainsAny(line[i:i+1], "0123456789") {
				tens = cast.ToInt(line[i : i+1])
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if strings.ContainsAny(line[i:i+1], "0123456789") {
				ones = cast.ToInt(line[i : i+1])
				break
			}
		}
		sum += tens*10 + ones
	}

	return sum
}

func part2(input string) int {
	prefixes := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
	for i := 0; i <= 9; i++ {
		prefixes[cast.ToString(i)] = i
	}

	var sum int
	for _, line := range strings.Split(input, "\n") {
		var first, last int

		for len(line) > 0 {
			for prefix, val := range prefixes {
				if doesStringHavePrefix(line, prefix) {
					if first == 0 {
						first = val
					}
					last = val
					break
				}
			}

			// shorten line
			line = line[1:]
		}

		sum += first*10 + last
	}

	return sum
}

func doesStringHavePrefix(str string, prefix string) bool {
	if len(str) < len(prefix) {
		return false
	}
	return str[:len(prefix)] == prefix
}
