package main

import (
	"flag"
	"fmt"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := notQuiteLisp(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func notQuiteLisp(input string, part int) int {
	var level int
	for i, r := range input {
		if r == '(' {
			level++
		} else {
			level--
		}

		if part == 2 && level == -1 {
			return i + 1
		}
	}

	return level
}
