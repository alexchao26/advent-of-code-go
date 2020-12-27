package main

import (
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	part1Ans, part2Ans := cookieScience(util.ReadFile("./input.txt"))
	fmt.Printf("Part1: %d\nPart2: %d\n", part1Ans, part2Ans)
}

func cookieScience(input string) (int, int) {
	lines := strings.Split(input, "\n")
	cookieVals := [][]int{}

	for _, line := range lines {
		var name string
		var cap, dur, fla, tex, cal int
		fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d",
			&name, &cap, &dur, &fla, &tex, &cal)

		// trim off colons, fmt.Sscanf parses between spaces/whitespace
		// names end up going unused...
		name = strings.Trim(name, ":")

		// cookie vals are simply all five values in order
		cookieVals = append(cookieVals, []int{cap, dur, fla, tex, cal})
	}

	var bestScore, best500CalScore int
	for ing1 := 0; ing1 < 100; ing1++ {
		for ing2 := 0; ing2 < 100; ing2++ {
			for ing3 := 0; ing3 < 100; ing3++ {
				ing4 := 100 - ing1 - ing2 - ing3

				cap := ing1*cookieVals[0][0] + ing2*cookieVals[1][0] + ing3*cookieVals[2][0] + ing4*cookieVals[3][0]
				dur := ing1*cookieVals[0][1] + ing2*cookieVals[1][1] + ing3*cookieVals[2][1] + ing4*cookieVals[3][1]
				fla := ing1*cookieVals[0][2] + ing2*cookieVals[1][2] + ing3*cookieVals[2][2] + ing4*cookieVals[3][2]
				tex := ing1*cookieVals[0][3] + ing2*cookieVals[1][3] + ing3*cookieVals[2][3] + ing4*cookieVals[3][3]

				cal := ing1*cookieVals[0][4] + ing2*cookieVals[1][4] + ing3*cookieVals[2][4] + ing4*cookieVals[3][4]

				// make negatives zero, without this two negative scores could
				// make a very large positive
				cap = mathutil.MaxInt(0, cap)
				dur = mathutil.MaxInt(0, dur)
				fla = mathutil.MaxInt(0, fla)
				tex = mathutil.MaxInt(0, tex)

				score := cap * dur * fla * tex

				if cal == 500 {
					best500CalScore = mathutil.MaxInt(best500CalScore, score)
				}
				bestScore = mathutil.MaxInt(bestScore, score)
			}
		}
	}

	return bestScore, best500CalScore
}
