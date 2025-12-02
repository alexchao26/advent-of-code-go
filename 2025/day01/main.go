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
	instructions := strings.Split(input, "\n")

	val := 50
	zeroes := 0
	for _, inst := range instructions {
		dir := inst[:1]
		num := cast.ToInt(inst[1:])

		if dir == "R" {
			val += num
			val %= 100
		} else if dir == "L" {
			val -= num
			for val < 0 {
				val += 100
			}
			val %= 100 // is this even necessary?
		} else {
			panic(fmt.Sprintf("unexpected dir: %q", dir))
		}
		if val == 0 {
			zeroes++
		}
	}

	return zeroes
}

func part2(input string) int {
	val := 50
	zeroes := 0
	for _, inst := range strings.Split(input, "\n") {
		dir := inst[:1]
		num := cast.ToInt(inst[1:])

		if dir == "R" {
			val += num
			zeroes += val / 100
			val %= 100
		} else if dir == "L" {
			wasZero := val == 0
			val -= num

			// special case when lands on zero because for loop will undercount by 1
			landedOnZero := val%100 == 0
			if landedOnZero {
				zeroes++
			}

			for val < 0 {
				val += 100
				zeroes++
			}

			// if it started at zero and turned left, the for loop above will over count by 1
			// 0 -> L5 = -5 + 100 (zeroes++ incorrectly)
			if wasZero {
				zeroes--
			}
		} else {
			panic(fmt.Sprintf("unexpected dir: %q", dir))
		}
	}

	return zeroes
}
