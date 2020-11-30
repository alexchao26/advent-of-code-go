package main

import (
	"errors"
	"flag"
	"fmt"
	"math"

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
	var err error
	for err != ErrDone {
		input, err = step(input)
	}
	return len(input)
}

func part2(input string) int {
	// find all units (lowercased) that can be removed
	units := map[byte]bool{}
	for i := 0; i < len(input); i++ {
		if input[i] < byte(96) {
			units[input[i]] = true
		} else {
			units[input[i]-ASCIIOffset] = true
		}
	}

	lowestResult := math.MaxInt16
	for u := range units {
		// remove all unit & its corresponding capital/lowercase version
		newStr := removeUnit(input, u)

		// run new string through part1 and update lowestResult if possible
		if result := part1(newStr); result < lowestResult {
			lowestResult = result
		}
	}

	return lowestResult
}

const ASCIIOffset = byte('a' - 'A')

var ErrDone = errors.New("no-op")

func step(units string) (string, error) {
	offset := byte('a' - 'A')

	for i := 1; i < len(units); i++ {
		if units[i-1]-units[i] == offset || units[i-1]-units[i] == -offset {
			// remove units i-1 and i
			newStr := units[:i-1] + units[i+1:]
			return newStr, nil
		}
	}
	return units, ErrDone
}

func removeUnit(input string, unit byte) string {
	var newStr string
	for i := 0; i < len(input); i++ {
		if input[i] == unit || input[i] == unit+ASCIIOffset {
			continue
		}
		newStr += string(input[i])
	}
	return newStr
}
