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
	}

	return opcodeComputer.registers[0]
}

func part2(input string) int {
	opcodeComputer := parseInput(input)

	opcodeComputer.registers[0] = 1

	for !opcodeComputer.tick() {
	}

	return opcodeComputer.registers[0]
}

// after deeply studying how the intcode cycles, in order to optimize it and
// run in any reasonable amount of time. It's become clear that the answer is
// just the sum of all the factors of the number that is generated after a few
// steps of the opcode computer running
func part2Cheeky(input string) int {
	computer := parseInput(input)
	computer.registers[0] = 1
	for i := 0; i < 20; i++ {
		computer.tick()
	}

	numberToFactorize := computer.registers[2] // this index varies based on inputs

	var ans int
	for i := 1; i <= numberToFactorize; i++ {
		if numberToFactorize%i == 0 {
			ans += i
		}
	}

	return ans
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

func (o *opcodeComputer) tick() (done bool) {
	// custom logic for the repetitive behavior
	ipValue := o.registers[o.instructionPointer]
	if ipValue == 4 {
		for o.registers[4] == 4 {
			if o.registers[1] >= o.registers[2] {
				break
			}
			if o.registers[5] == o.registers[2] {
				o.registers[0] += o.registers[1]
			}

			o.registers[3]++
			reg3 := o.registers[3]
			if reg3 > o.registers[2] {
				o.registers[1]++
				// escape hatch
				if o.registers[1] > o.registers[2] {
					o.registers[4] *= o.registers[4]
					o.registers[4]++
					break
				}
				o.registers[3] = 1
				o.registers[5] = o.registers[1]
			} else if o.registers[2]%o.registers[1] == 0 {
				// increase registers 5 to hit o.reg5 == o.reg2
				// side of 5 = 1 x 3
				o.registers[5] = o.registers[2]
				o.registers[3] = o.registers[2]
			} else {
				o.registers[3] = o.registers[2]
			}
		}
	}

	if o.registers[o.instructionPointer] >= len(o.instructions) {
		return true
	}

	inst := o.instructions[o.registers[o.instructionPointer]]

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

	// for i, v := range instructions {
	// 	fmt.Println(i, v)
	// }
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
