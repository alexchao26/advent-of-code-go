package main

import (
	"flag"
	"fmt"
	"strings"

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
	opcodeComputer := parseInput(input)

	for !opcodeComputer.tick() {
		// instruction 28 of my input is the only one that accesses register zero
		// it is comparing reg 0 to reg 5. so to break out of loops asap, set
		// reg 0 to the value found in reg 5 when 28 is hit for the first time
		if opcodeComputer.registers[opcodeComputer.instructionPointer] == 28 {
			break
		}
	}

	return opcodeComputer.registers[5]
}

func part2(input string) int {
	opcodeComputer := parseInput(input)

	// similar idea for part 2 but now we need to find the previous state of
	// register 5 when register 5 REPEATS itself. this is a brute force solution
	// using a map to store previous reg5 values, and stores the previous reg5
	var lastReg5 int
	comparedRegister5s := map[int]bool{}
	for !opcodeComputer.tick() {
		if opcodeComputer.registers[opcodeComputer.instructionPointer] == 28 {
			reg5 := opcodeComputer.registers[5]
			if comparedRegister5s[reg5] {
				break
			}
			comparedRegister5s[reg5] = true
			lastReg5 = reg5
		}
	}

	return lastReg5
}

type opcodeComputer struct {
	instructions       []instruction
	registers          [6]int
	instructionPointer int // an index the stores the index for which instruction to run
}
type instruction struct {
	name      string
	abcValues [3]int
}

// literal opcode computer implementation, unoptimized
func (o *opcodeComputer) tick() (done bool) {
	if o.registers[o.instructionPointer] >= len(o.instructions) {
		fmt.Println("Out of range instruction, terminating...")
		return true
	}
	instIndex := o.registers[o.instructionPointer]
	inst := o.instructions[instIndex]

	// fmt.Println(strings.Repeat(" ", instIndex) + strconv.Itoa(instIndex))

	opcodeFunc := opcodeNamesToFuncs[inst.name]

	o.registers = opcodeFunc(o.registers, inst.abcValues)

	// increment value @ instructionPointer, validate that it's still in range
	o.registers[o.instructionPointer]++

	if o.registers[o.instructionPointer] >= len(o.instructions) {
		return true
	}

	return false
}

func parseInput(input string) opcodeComputer {
	lines := strings.Split(input, "\n")

	var instructionPointer int
	fmt.Sscanf(lines[0], "#ip %d", &instructionPointer)

	var instructions []instruction
	for _, l := range lines[1:] {
		var inst instruction
		fmt.Sscanf(l, "%4s %d %d %d", &inst.name, &inst.abcValues[0], &inst.abcValues[1], &inst.abcValues[2])
		instructions = append(instructions, inst)
	}

	return opcodeComputer{
		instructions:       instructions,
		instructionPointer: instructionPointer,
	}
}

var opcodeNamesToFuncs = map[string]opcodeFunc{
	"addr": addr, "addi": addi,
	"mulr": mulr, "muli": muli,
	"banr": banr, "bani": bani,
	"borr": borr, "bori": bori,
	"setr": setr, "seti": seti,
	"gtir": gtir, "gtri": gtri, "gtrr": gtrr,
	"eqir": eqir, "eqri": eqri, "eqrr": eqrr,
}

type opcodeFunc func([6]int, [3]int) [6]int

func addr(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = registers[abcValues[0]] + registers[abcValues[1]]
	return registers
}

func addi(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = registers[abcValues[0]] + abcValues[1]
	return registers
}
func mulr(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = registers[abcValues[0]] * registers[abcValues[1]]
	return registers
}
func muli(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = registers[abcValues[0]] * abcValues[1]
	return registers
}
func banr(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = registers[abcValues[0]] & registers[abcValues[1]]
	return registers
}
func bani(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = registers[abcValues[0]] & abcValues[1]
	return registers
}
func borr(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = registers[abcValues[0]] | registers[abcValues[1]]
	return registers
}
func bori(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = registers[abcValues[0]] | abcValues[1]
	return registers
}
func setr(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = registers[abcValues[0]]
	return registers
}
func seti(registers [6]int, abcValues [3]int) [6]int {
	registers[abcValues[2]] = abcValues[0]
	return registers
}
func gtir(registers [6]int, abcValues [3]int) [6]int {
	if abcValues[0] > registers[abcValues[1]] {
		registers[abcValues[2]] = 1
	} else {
		registers[abcValues[2]] = 0
	}
	return registers
}
func gtri(registers [6]int, abcValues [3]int) [6]int {
	if registers[abcValues[0]] > abcValues[1] {
		registers[abcValues[2]] = 1
	} else {
		registers[abcValues[2]] = 0
	}
	return registers
}
func gtrr(registers [6]int, abcValues [3]int) [6]int {
	if registers[abcValues[0]] > registers[abcValues[1]] {
		registers[abcValues[2]] = 1
	} else {
		registers[abcValues[2]] = 0
	}
	return registers
}
func eqir(registers [6]int, abcValues [3]int) [6]int {
	if abcValues[0] == registers[abcValues[1]] {
		registers[abcValues[2]] = 1
	} else {
		registers[abcValues[2]] = 0
	}
	return registers
}
func eqri(registers [6]int, abcValues [3]int) [6]int {
	if registers[abcValues[0]] == abcValues[1] {
		registers[abcValues[2]] = 1
	} else {
		registers[abcValues[2]] = 0
	}
	return registers
}
func eqrr(registers [6]int, abcValues [3]int) [6]int {
	if registers[abcValues[0]] == registers[abcValues[1]] {
		registers[abcValues[2]] = 1
	} else {
		registers[abcValues[2]] = 0
	}
	return registers
}
