package main

import (
	"flag"
	"fmt"
	"strconv"
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

// literal implementation for part 1
func part1(input string) int {
	comp := newComputerFromInput(input)
	// prime until it gets a valid receive command, then return last outputted num
	return comp.countMulsRun()
}

// referenced this heavily: https://www.reddit.com/r/adventofcode/comments/7lms6p/2017_day_23_solutions/drnh5sx?utm_source=share&utm_medium=web2x&context=3
func part2(input string) int {
	b := 81
	c := 81

	b = b*100 + 100000
	c = b + 17000
	var h int
	for {
		f := 1
		// effectively a prime number checker.
		for d := 2; d*d <= b; d++ {
			if b%d == 0 {
				f = 0
				break
			}
		}

		if f == 0 {
			h++
		}
		if b == c {
			break
		}
		b += 17
	}

	return h
}

type computer struct {
	instructions [][]string
	pointer      int
	registers    map[string]int
	output       []int
}

func newComputerFromInput(input string) *computer {
	comp := &computer{registers: map[string]int{}}
	for _, line := range strings.Split(input, "\n") {
		comp.instructions = append(comp.instructions, strings.Split(line, " "))
	}

	return comp
}

func (c *computer) countMulsRun() (mulsRun int) {
	for c.pointer < len(c.instructions) {
		inst := c.instructions[c.pointer]
		valX := inst[1]
		var valFromY int
		if val, err := strconv.Atoi(inst[2]); err != nil {
			// if there is an error parsing to an integer, value at index 1 is a register
			valFromY = c.registers[inst[2]]
		} else {
			valFromY = val
		}

		switch inst[0] {
		case "set":
			c.registers[valX] = valFromY
			c.pointer++
		case "sub":
			c.registers[valX] -= valFromY
			c.pointer++
		case "mul":
			c.registers[valX] *= valFromY
			c.pointer++
			mulsRun++
		case "jnz":
			var parsedX int
			if num, err := strconv.Atoi(valX); err != nil {
				// err converting, not a number
				parsedX = c.registers[valX]
			} else {
				// no error then a number was parsed
				parsedX = num
			}
			if parsedX != 0 {
				c.pointer += valFromY
			} else {
				c.pointer++
			}
		default:
			panic("unhandled operator " + inst[0])
		}
	}
	return mulsRun
}
