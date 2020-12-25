package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/algos"
	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var ans string
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"), "abcdefgh", "")
	} else {
		ans = part1(util.ReadFile("./input.txt"), "abcdefgh", "fbgdceah")
	}
	fmt.Println("Output:", ans)
}

func part1(input string, starting string, target string) string {
	runSteps := func(starting string) string {
		registers := strings.Split(starting, "")
		for _, inst := range strings.Split(input, "\n") {
			registers = modifySlice(inst, registers)
		}
		return strings.Join(registers, "")
	}

	// part 1
	if target == "" {
		return runSteps(starting)
	}

	for _, p := range algos.MakeStringPermutations(starting) {
		if runSteps(p) == target {
			return p
		}
	}

	panic("no perms matched")
}

func modifySlice(line string, sli []string) []string {
	switch {
	case strings.HasPrefix(line, "swap"):
		var i1, i2 int
		if strings.Contains(line, "position") {
			fmt.Sscanf(line, "swap position %d with position %d", &i1, &i2)
		} else {
			var c1, c2 string
			fmt.Sscanf(line, "swap letter %1s with letter %1s", &c1, &c2)
			i1, i2 = getIndex(sli, c1), getIndex(sli, c2)
		}
		sli[i1], sli[i2] = sli[i2], sli[i1]
	case strings.HasPrefix(line, "rotate"):
		var rightShift int
		parts := strings.Split(line, " ")
		if strings.Contains(line, "letter") {
			index := getIndex(sli, parts[6])
			if index >= 4 {
				index++
			}
			index++
			rightShift = index % len(sli)
		} else {
			// left or right
			rightShift = cast.ToInt(parts[2])
			if parts[1] == "left" {
				rightShift = len(sli) - rightShift
			}
		}
		// perform shift
		sli = append(sli[len(sli)-rightShift:], sli[:len(sli)-rightShift]...)
	case strings.HasPrefix(line, "reverse"):
		var i1, i2 int
		fmt.Sscanf(line, "reverse positions %d through %d", &i1, &i2)
		for i1 < i2 {
			sli[i1], sli[i2] = sli[i2], sli[i1]
			i1++
			i2--
		}
	case strings.HasPrefix(line, "move"):
		var i1, i2 int
		fmt.Sscanf(line, "move position %d to position %d", &i1, &i2)
		store := sli[i1]

		// remove char at i1
		copy(sli[i1:], sli[i1+1:])

		for i := len(sli) - 1; i >= i2+1; i-- {
			sli[i] = sli[i-1]
		}
		sli[i2] = store
	default:
		panic("unhandled instruction type: " + line)
	}

	// return modified slice
	return sli
}

func getIndex(letters []string, toFind string) int {
	for i, v := range letters {
		if v == toFind {
			return i
		}
	}
	return -1
}
