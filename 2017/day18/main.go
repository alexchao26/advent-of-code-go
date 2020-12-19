package main

import (
	"flag"
	"fmt"
	"math"
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

func part1(input string) int {
	comp := newComputerFromInput(input, 0)
	// prime until it gets a valid receive command, then return last outputted num
	comp.step(noInput)
	return comp.output[len(comp.output)-1]
}

func part2(input string) int {
	program0 := newComputerFromInput(input, 0)
	program1 := newComputerFromInput(input, 1)

	// prime the computers
	program0.step(noInput)
	program1.step(noInput)

	var sentFrom1 int

	for len(program0.output)+len(program1.output) > 0 {
		// run outputs from program zero through program 1
		for len(program0.output) > 0 {
			v := program0.output[0]
			program0.output = program0.output[1:]
			program1.step(v)
		}
		// and vice versa
		for len(program1.output) > 0 {
			v := program1.output[0]
			program1.output = program1.output[1:]
			program0.step(v)
			sentFrom1++
		}
	}

	return sentFrom1
}

type computer struct {
	instructions [][]string
	pointer      int
	registers    map[string]int
	output       []int
}

func newComputerFromInput(input string, programID int) *computer {
	comp := &computer{registers: map[string]int{"p": programID}}
	for _, line := range strings.Split(input, "\n") {
		comp.instructions = append(comp.instructions, strings.Split(line, " "))
	}
	return comp
}

var noInput = math.MinInt16

func (c *computer) step(inputNum int) {
	for {
		inst := c.instructions[c.pointer]
		valX := inst[1]
		var valY int
		if len(inst) == 3 && inst[2] != "" {
			if val, err := strconv.Atoi(inst[2]); err != nil {
				// if there is an error parsing to an integer, value at index 1 is a register
				valY = c.registers[inst[2]]
			} else {
				valY = val
			}
		}

		switch inst[0] {
		case "snd":
			c.output = append(c.output, c.registers[valX])
			c.pointer++
		case "set":
			c.registers[valX] = valY
			c.pointer++
		case "add":
			c.registers[valX] += valY
			c.pointer++
		case "mul":
			c.registers[valX] *= valY
			c.pointer++
		case "mod":
			c.registers[valX] %= valY
			c.pointer++
		case "rcv":
			if inputNum == noInput {
				return
			}
			c.registers[valX] = inputNum
			inputNum = noInput
			c.pointer++
		case "jgz":
			var parsedX int
			if num, err := strconv.Atoi(valX); err != nil {
				// err converting, not a number
				parsedX = c.registers[valX]
			} else {
				// no error then a number was parsed
				parsedX = num
			}
			if parsedX > 0 {
				c.pointer += valY + len(c.instructions)
				c.pointer %= len(c.instructions)
			} else {
				c.pointer++
			}
		default:
			panic("unhandled operator " + inst[0])
		}
	}
}
