package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) string {
	stacks, steps := parseInput(input)

	for _, step := range steps {
		// move crates ONE AT A TIME
		for q := 0; q < step.qty; q++ {
			top := stacks[step.from][len(stacks[step.from])-1]
			stacks[step.to] = append(stacks[step.to], top)
			stacks[step.from] = stacks[step.from][:len(stacks[step.from])-1]
		}
	}

	ans := ""
	for _, stack := range stacks {
		ans += stack[len(stack)-1]
	}
	return ans
}

func part2(input string) string {
	stacks, steps := parseInput(input)

	for _, step := range steps {
		// move crates ONCE
		fromIndex := len(stacks[step.from]) - step.qty
		stacks[step.to] = append(stacks[step.to], stacks[step.from][fromIndex:]...)
		stacks[step.from] = stacks[step.from][:fromIndex]
	}

	ans := ""
	for _, stack := range stacks {
		ans += stack[len(stack)-1]
	}
	return ans
}

// move 4 from 3 to 1
type step struct {
	qty, from, to int
}

func (s step) String() string {
	return fmt.Sprintf("move %d from %d to %d", s.qty, s.from, s.to)
}

func parseInput(input string) ([][]string, []step) {
	parts := strings.Split(input, "\n\n")

	state := parts[0]
	oversized := [][]string{}
	for _, row := range strings.Split(state, "\n") {
		oversized = append(oversized, strings.Split(row, ""))
	}
	oRows, oCols := len(oversized), len(oversized[0])

	actual := [][]string{}

	for c := 0; c < oCols-1; c++ {
		if oversized[oRows-1][c] != " " {
			// hit a column with values... move up from here
			stack := []string{}
			for r := oRows - 2; r >= 0; r-- {
				char := oversized[r][c]
				if char != " " {
					stack = append(stack, char)
				}
			}
			actual = append(actual, stack)
		}
	}

	stepsRaw := parts[1]
	steps := []step{}
	for _, row := range strings.Split(stepsRaw, "\n") {
		inst := step{}
		_, err := fmt.Sscanf(row, "move %d from %d to %d", &inst.qty, &inst.from, &inst.to)
		if err != nil {
			panic(err)
		}
		// subtract one so they're zero indexed...
		inst.from--
		inst.to--
		steps = append(steps, inst)
	}

	return actual, steps
}
