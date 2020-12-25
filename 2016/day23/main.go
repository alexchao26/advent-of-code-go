package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := assemblyComputer(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func assemblyComputer(input string, part int) int {
	instructions := strings.Split(input, "\n")
	registers := map[string]int{}
	var instIndex int

	registers["a"] = 7
	if part == 2 {
		registers["a"] = 12
	}

	for instIndex < len(instructions) {
		// uncomment this to print out the instruction set and see how it's changed
		// fmt.Println(instIndex)
		// for in, i := range instructions {
		// 	fmt.Println(in, i)
		// }
		// fmt.Println()

		inst := instructions[instIndex]
		parts := strings.Split(inst, " ")

		// My instruction list gets transformed into this
		// all jump instructions can be optimized, in particular the ones
		// tagged here b/c they jump over another jump instruction
		//   effectively becoming multiplication steps
		// 0 cpy a b
		// 1 dec b
		// 2 cpy a d
		// 3 cpy 0 a
		// 4 cpy b c
		// 5 inc a
		// 6 dec c
		// 7 jnz c -2
		// 8 dec d
		// 9 jnz d -5 // <-
		// 10 dec b
		// 11 cpy b c
		// 12 cpy c d
		// 13 dec d
		// 14 inc c
		// 15 jnz d
		// 16 tgl c
		// 17 cpy -16 c
		// 18 cpy 1 c
		// 19 cpy 89 c
		// 20 cpy 77 d
		// 21 inc a
		// 22 dec d
		// 23 jnz d
		// 24 dec c
		// 25 jnz c -5 // <-

		// Hard coded multiplication skippers
		if inst == "jnz d -5" && instructions[instIndex-1] == "dec d" &&
			instructions[instIndex-2] == "jnz c -2" {
			registers["a"] += registers["b"] * registers["d"]
			registers["c"] = 0
			registers["d"] = 0
		}

		if inst == "jnz c -5" && instructions[instIndex-1] == "dec c" {
			registers["a"] += 77 * registers["c"]
			registers["c"] = 0
			registers["d"] = 0
		}

		switch parts[0] {
		case "cpy":
			valX := parseValueOrRegister(registers, parts[1])
			registers[parts[2]] = valX
			instIndex++
		case "inc":
			registers[parts[1]]++
			instIndex++
		case "dec":
			registers[parts[1]]--
			instIndex++
		case "jnz":
			valX := parseValueOrRegister(registers, parts[1])
			if valX != 0 {
				valY := parseValueOrRegister(registers, parts[2])
				instIndex += valY
			} else {
				instIndex++
			}
		case "tgl":
			// valX is an offset
			valX := parseValueOrRegister(registers, parts[1])
			indexToMod := instIndex + valX
			if indexToMod < len(instructions) {
				instToModParts := strings.Split(instructions[indexToMod], " ")
				var newType string
				if len(instToModParts) == 2 {
					newType = "inc"
					if instToModParts[0] == "inc" {
						newType = "dec"
					}
				} else if len(instToModParts) == 3 {
					newType = "jnz"
					if instToModParts[0] == "jnz" {
						newType = "cpy"
					}
				}
				instToModParts[0] = newType
				instructions[indexToMod] = strings.Join(instToModParts, " ")
			}
			instIndex++
		default:
			panic("unhanded instruction type " + parts[0])
		}
	}

	return registers["a"]
}

func parseValueOrRegister(registers map[string]int, part string) int {
	if regexp.MustCompile("[0-9]").MatchString(part) {
		return cast.ToInt(part)
	}
	return registers[part]
}
