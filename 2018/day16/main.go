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
	examples, _ := parseInput(input) // ignore instructions for part 1

	var threePlusBehaviors int
	for _, e := range examples {
		var matches int
		for _, opcodeFunc := range opcodeNamesToFuncs {
			before, instructions, after := e[0], e[1], e[2]
			if after == opcodeFunc(before, instructions) {
				matches++
			}
		}
		if matches >= 3 {
			threePlusBehaviors++
		}
	}

	return threePlusBehaviors
}

func part2(input string) int {
	examples, instructions := parseInput(input) // ignore instructions for part 1

	// for each num, names that it _COULD_ be
	opCodeNumToNameGraph := map[int]map[string]bool{}

	for _, e := range examples {
		for name, opcodeFunc := range opcodeNamesToFuncs {
			before, instructions, after := e[0], e[1], e[2]
			if after == opcodeFunc(before, instructions) {
				if opCodeNumToNameGraph[instructions[0]] == nil {
					opCodeNumToNameGraph[instructions[0]] = map[string]bool{}
				}
				opCodeNumToNameGraph[instructions[0]][name] = true
			}
		}
	}

	derivedOpcodeNumToFunc := map[int]opcodeFunc{}
	for len(derivedOpcodeNumToFunc) < 16 {
		for num, edges := range opCodeNumToNameGraph {
			if len(edges) == 1 {
				for name := range edges { // only way to get the one val out of a map?
					derivedOpcodeNumToFunc[num] = opcodeNamesToFuncs[name]

					// delete name from all other graph edges b/c it's settled
					for _, edges := range opCodeNumToNameGraph {
						delete(edges, name)
					}
				}

				// break to restart the main loop form the beginning
				break
			}
		}
	}

	// run all instructions
	var registers [4]int
	for _, inst := range instructions {
		opcodeFunc := derivedOpcodeNumToFunc[inst[0]]
		registers = opcodeFunc(registers, inst)
	}

	return registers[0]
}

func parseInput(input string) ([][3][4]int, [][4]int) {
	lines := strings.Split(input, "\n\n\n\n")

	inputExamples := lines[0]
	inputInstructions := lines[1]

	var examples [][3][4]int
	for _, e := range strings.Split(inputExamples, "\n\n") {
		var before, op, after [4]int
		fmt.Sscanf(e, "Before: [%d, %d, %d, %d]\n%d %d %d %d\nAfter:  [%d, %d, %d, %d]",
			&before[0], &before[1], &before[2], &before[3],
			&op[0], &op[1], &op[2], &op[3],
			&after[0], &after[1], &after[2], &after[3],
		)
		examples = append(examples, [3][4]int{before, op, after})
	}

	var instructions [][4]int
	for _, i := range strings.Split(inputInstructions, "\n") {
		var inst [4]int
		fmt.Sscanf(i, "%d %d %d %d", &inst[0], &inst[1], &inst[2], &inst[3])
		instructions = append(instructions, inst)
	}

	return examples, instructions
}

var opcodeNamesToFuncs = map[string]opcodeFunc{
	"addr": addr, "addi": addi,
	"multr": multr, "multi": multi,
	"banr": banr, "bani": bani,
	"borr": borr, "bori": bori,
	"setr": setr, "seti": seti,
	"gtir": gtir, "gtri": gtri, "gtrr": gtrr,
	"eqir": eqir, "eqri": eqri, "eqrr": eqrr,
}

type opcodeFunc func([4]int, [4]int) [4]int

func addr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] + registers[instructions[2]]
	return registers
}

func addi(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] + instructions[2]
	return registers
}
func multr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] * registers[instructions[2]]
	return registers
}
func multi(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] * instructions[2]
	return registers
}
func banr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] & registers[instructions[2]]
	return registers
}
func bani(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] & instructions[2]
	return registers
}
func borr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] | registers[instructions[2]]
	return registers
}
func bori(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] | instructions[2]
	return registers
}
func setr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]]
	return registers
}
func seti(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = instructions[1]
	return registers
}
func gtir(registers [4]int, instructions [4]int) [4]int {
	if instructions[1] > registers[instructions[2]] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}
func gtri(registers [4]int, instructions [4]int) [4]int {
	if registers[instructions[1]] > instructions[2] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}
func gtrr(registers [4]int, instructions [4]int) [4]int {
	if registers[instructions[1]] > registers[instructions[2]] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}
func eqir(registers [4]int, instructions [4]int) [4]int {
	if instructions[1] == registers[instructions[2]] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}
func eqri(registers [4]int, instructions [4]int) [4]int {
	if registers[instructions[1]] == instructions[2] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}
func eqrr(registers [4]int, instructions [4]int) [4]int {
	if registers[instructions[1]] == registers[instructions[2]] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}
