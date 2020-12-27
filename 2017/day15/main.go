package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := duelingGenerators(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func duelingGenerators(input string, part int) int {
	values := parseInput(input)

	factors := []int{16807, 48271}
	divisor := 2147483647

	// set criteria for passing a value to judge, and rounds based on part num
	criteria := []int{1, 1}
	rounds := 40000000
	if part == 2 {
		rounds = 5000000
		criteria = []int{4, 8}
	}

	var judgeCount int
	for i := 0; i < rounds; i++ {
		for i, v := range values {
			values[i] = getNextValue(v, factors[i], divisor, criteria[i])
		}

		// do the back 16 line up ?
		// XOR them together, the last 16 should not have bits because they should
		// should be zero'ed out.
		// (XOR result) % (2^16) will be zero (i.e. no remainder) if back 16 match
		compareVal := values[0] ^ values[1]
		twoPow16 := 1 << 16
		if (compareVal % twoPow16) == 0 {
			judgeCount++
		}
	}

	return judgeCount
}

// iterate until a value that can be passed to the judge is reached
// for part 1 the criteria will be 1, so only one iteration is made
// for part 2 it will be 4 or 8, so the loop will run until a good value is reached
func getNextValue(value, factor, divisor, criteria int) int {
	// run at least once
	value *= factor
	value %= divisor

	for value%criteria != 0 {
		value *= factor
		value %= divisor
	}

	return value
}

func parseInput(input string) (ans []int) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		split := strings.Split(l, " starts with ")
		ans = append(ans, cast.ToInt(split[1]))
	}
	return ans
}
