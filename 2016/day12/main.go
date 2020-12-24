package main

import (
	"flag"
	"fmt"
	"strconv"
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
	registers := map[string]int{} // a b c d = 0
	var instIndex int

	if part == 2 {
		registers["c"] = 1
	}

	for instIndex < len(instructions) {
		parts := strings.Split(instructions[instIndex], " ")

		switch parts[0] {
		case "cpy":
			valX, err := strconv.Atoi(parts[1])
			if err != nil {
				valX = registers[parts[1]]
			}
			registers[parts[2]] = valX
			instIndex++
		case "inc":
			registers[parts[1]]++
			instIndex++
		case "dec":
			registers[parts[1]]--
			instIndex++
		case "jnz":
			valX, err := strconv.Atoi(parts[1])
			if err != nil {
				valX = registers[parts[1]]
			}
			if valX != 0 {
				instIndex += cast.ToInt(parts[2])
			} else {
				instIndex++
			}
		}
	}

	return registers["a"]
}
