package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := calcRegisters(util.ReadFile("./input.txt"), part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func calcRegisters(input string, part int) int {
	instructions := parseInput(input)

	registers := map[string]int{}
	var highestEverRegister int // for part 2
	for _, inst := range instructions {
		registerVal := registers[inst.conditional[0]]
		compareVal := cast.ToInt(inst.conditional[2])
		var conditionalResult bool
		switch inst.conditional[1] {
		case "==":
			conditionalResult = registerVal == compareVal
		case "!=":
			conditionalResult = registerVal != compareVal
		case "<":
			conditionalResult = registerVal < compareVal
		case ">":
			conditionalResult = registerVal > compareVal
		case "<=":
			conditionalResult = registerVal <= compareVal
		case ">=":
			conditionalResult = registerVal >= compareVal
		default:
			panic("unhandled operator")
		}
		if conditionalResult {
			registers[inst.registerName] += inst.diff
		}
		highestEverRegister = mathy.MaxInt(highestEverRegister, registers[inst.registerName])
	}

	largestFinalRegister := -math.MaxInt32
	for _, v := range registers {
		largestFinalRegister = mathy.MaxInt(largestFinalRegister, v)
	}

	if part == 1 {
		return largestFinalRegister
	}
	return highestEverRegister
}

type instruction struct {
	registerName string    // register to modify
	diff         int       // if instruction is a dec, multiply by -1
	conditional  [3]string // contains all 3 parts @ end of each instruction
}
type compareFunc func(value int) bool

func parseInput(input string) []instruction {
	lines := strings.Split(input, "\n")
	var instructions []instruction

	for _, l := range lines {
		parts := strings.Split(l, " ")
		if len(parts) != 7 {
			panic("parts len to seven")
		}
		inst := instruction{
			registerName: parts[0],
			diff:         cast.ToInt(parts[2]),
			conditional:  [3]string{parts[4], parts[5], parts[6]},
		}
		if parts[1] == "dec" {
			inst.diff *= -1
		}

		instructions = append(instructions, inst)
	}

	return instructions
}
