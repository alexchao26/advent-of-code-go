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

// 180/1133
func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := step(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func step(input string, part int) int {
	state := make([]int, 9)
	for _, num := range strings.Split(input, ",") {
		daysLeft := cast.ToInt(num)
		state[daysLeft]++
	}

	// different number of days for part 1 vs part 2
	days := 80
	if part == 2 {
		days = 256
	}

	// count down to zero
	// then it makes a new lanternfish with a count of 8
	for d := 0; d < days; d++ {
		save := state[0]
		for i := 0; i < len(state)-1; i++ {
			state[i] = state[i+1]
		}
		state[8] = save
		state[6] += save
	}

	// count up final sum
	var sum int
	for _, n := range state {
		sum += n
	}
	return sum
}
