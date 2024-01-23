package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/mathy"
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

func part1(input string) int {
	games := parseInput(input)

	var possibleIDSum int

	for _, g := range games {
		isPossible := true
		for _, step := range g.steps {
			// only 12 red cubes, 13 green cubes, and 14 blue cubes
			if step["red"] > 12 || step["green"] > 13 || step["blue"] > 14 {
				isPossible = false
				break
			}
		}

		if isPossible {
			possibleIDSum += g.id
		}
	}

	return possibleIDSum
}

func part2(input string) int {
	games := parseInput(input)

	var sum int

	for _, g := range games {
		lowestPossibleCount := map[string]int{}

		for _, step := range g.steps {
			lowestPossibleCount["red"] = mathy.MaxInt(lowestPossibleCount["red"], step["red"])
			lowestPossibleCount["green"] = mathy.MaxInt(lowestPossibleCount["green"], step["green"])
			lowestPossibleCount["blue"] = mathy.MaxInt(lowestPossibleCount["blue"], step["blue"])
		}

		sum += lowestPossibleCount["red"] * lowestPossibleCount["blue"] * lowestPossibleCount["green"]
	}

	return sum
}

type game struct {
	id    int
	steps []map[string]int
}

func parseInput(input string) (ans []game) {
	for i, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		g := game{
			id: i + 1,
		}

		for _, p := range strings.Split(parts[1], "; ") {
			step := map[string]int{}
			for _, group := range strings.Split(p, ", ") {
				numberColor := strings.Split(group, " ")
				if len(numberColor) != 2 {
					panic(fmt.Sprintf("group not in two pieces %q", group))
				}

				step[numberColor[1]] = cast.ToInt(numberColor[0])
			}
			g.steps = append(g.steps, step)
		}
		ans = append(ans, g)
	}
	return ans
}
