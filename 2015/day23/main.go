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

	ans := simpleAssemblyComputer(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func simpleAssemblyComputer(input string, part int) int {
	instructions := strings.Split(input, "\n")
	var index int

	registers := map[string]int{}
	if part == 2 {
		registers["a"] = 1
	}

	for index < len(instructions) {
		parts := strings.Split(instructions[index], " ")
		switch parts[0] {
		case "hlf":
			reg := parts[1]
			registers[reg] /= 2
			index++
		case "tpl":
			reg := parts[1]
			registers[reg] *= 3
			index++
		case "inc":
			reg := parts[1]
			registers[reg]++
			index++
		case "jmp":
			diff := cast.ToInt(parts[1])
			index += diff
		case "jie":
			reg := strings.Trim(parts[1], ",")
			diff := cast.ToInt(parts[2])
			if registers[reg]%2 == 0 {
				index += diff
			} else {
				index++
			}
		case "jio":
			reg := strings.Trim(parts[1], ",")
			diff := cast.ToInt(parts[2])
			if registers[reg] == 1 {
				index += diff
			} else {
				index++
			}
		default:
			panic("unhandled instruction type: " + parts[0])
		}
	}

	return registers["b"]
}
