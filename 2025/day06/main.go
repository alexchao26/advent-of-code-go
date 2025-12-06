package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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
	problems := parseInputPart1(input)

	return scoreProblems(problems)
}

func scoreProblems(problems []problem) int {
	ans := 0

	for _, p := range problems {
		if p.op == "*" {
			s := 1
			for _, n := range p.nums {
				s *= n
			}
			ans += s
		} else if p.op == "+" {
			s := 0
			for _, n := range p.nums {
				s += n
			}
			ans += s
		} else {
			panic("unexpected op: " + p.op)
		}
	}

	return ans
}

func part2(input string) int {
	lines := strings.Split(input, "\n")
	grid := [][]string{}
	for _, l := range lines {
		grid = append(grid, strings.Split(l, ""))
	}

	problems := []problem{}
	for col := len(grid[0]) - 1; col >= 0; col-- {
		if len(problems) == 0 {
			problems = append(problems, problem{})
		}
		chunk := ""
		for row := 0; row < len(grid); row++ {
			// end of row reached, handle the chunk and potentially an operation
			if row == len(grid)-1 {
				// handle chunked number
				chunk = strings.TrimSpace(chunk)

				// skip empty columns
				if chunk == "" {
					continue
				}

				problems[len(problems)-1].nums = append(problems[len(problems)-1].nums, cast.ToInt(chunk))

				// handle op
				if grid[row][col] != " " {
					op := grid[row][col]
					problems[len(problems)-1].op = op

					// add new problem to begin chunking
					if col != 0 {
						problems = append(problems, problem{})
					}
				}
			} else {
				chunk += grid[row][col]
			}
		}
	}

	return scoreProblems(problems)
}

type problem struct {
	op   string
	nums []int
}

func parseInputPart1(input string) (ans []problem) {
	lines := strings.Split(input, "\n")

	var multipleSpaces = regexp.MustCompile(`\s{2,}`)

	// remove duplicate spaces and trim each line, then it's easy to split
	for i := range lines {
		lines[i] = multipleSpaces.ReplaceAllString(lines[i], " ")
		lines[i] = strings.TrimSpace(lines[i])
	}
	grid := [][]string{}

	for _, line := range lines {
		grid = append(grid, strings.Split(line, " "))
	}

	for c := range len(grid[0]) {
		p := problem{}
		for r := range len(grid) {
			if r == len(grid)-1 {
				p.op = grid[r][c]
			} else {
				p.nums = append(p.nums, cast.ToInt(grid[r][c]))
			}
		}
		ans = append(ans, p)
	}

	return ans
}
