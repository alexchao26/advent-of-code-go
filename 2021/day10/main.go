package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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

// 1470/1321
func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := syntaxScoring(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func syntaxScoring(input string, part int) int {
	var lines [][]string
	for _, l := range strings.Split(input, "\n") {
		lines = append(lines, strings.Split(l, ""))
	}

	closedToOpens := map[string]string{
		")": "(",
		"]": "[",
		">": "<",
		"}": "{",
	}

	// for part 1
	illegalScores := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	// for part 2
	autoCompleteScores := map[string]int{
		// needed to invert these to match what was left on the stack
		// alternatively could have used the closedToOpens map
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	// balanced parens
	var syntaxErrorScore int // part 1
	var part2Scores []int    // part 2
	for _, line := range lines {
		var stack []string
		var isCorrupted bool
		for _, char := range line {
			// part1: assign score for first illegal character
			opp, ok := closedToOpens[char]
			if !ok {
				stack = append(stack, char)
			} else {
				if len(stack) == 0 {
					panic("empty stack")
				}
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != opp {
					syntaxErrorScore += illegalScores[char]
					isCorrupted = true
				}
			}
		}

		// part 2: calculate score of the string needed to make the string valid
		if !isCorrupted {
			// stack contains all unmatched chars... match them in reverse order
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				score *= 5
				score += autoCompleteScores[stack[i]]
			}
			part2Scores = append(part2Scores, score)
		}
	}

	if part == 1 {
		return syntaxErrorScore
	}

	/// sort and return middle one, always an odd number of scores so just divide by 2 and rely on integer division
	sort.Ints(part2Scores)
	return part2Scores[len(part2Scores)/2]
}
