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
	races := parseInputPart1(input)
	ans := 1

	for _, r := range races {
		ans *= wayToWinRace(r)
	}

	return ans
}

func wayToWinRace(r race) int {
	// probably fast enough to do this naively for part 1...
	// Time:        40     92     97     90
	// total: 319
	// but could binary search this

	var ans int
	for chargeTime := 1; chargeTime < r.time; chargeTime++ {
		// velocity (chargeTime) * remaining time
		dist := chargeTime * (r.time - chargeTime)
		if dist > r.distance {
			ans++
		}
	}

	return ans
}

func part2(input string) int {
	combinedRace := parseInputPart2(input)

	// binary search this?
	// but the left and right and middle could all fail if the distribution is
	// F F P P P P F F F F F F F F F F F F
	// F = Fail, P = Pass
	// and it is some kind of (UN-CENTERED) bell curve distribution
	// could test 1 by 1 from left and right of time bound...

	// unfortunately you can just brute force this.
	return wayToWinRace(combinedRace)

	// all the optimizations might fall apart given unkind inputs, the example
	// has a distribution with many passes in the middle:
	// F P P P P P P P P P P P P P P P P P P P P P P P P P P P P P P P P P P F
	// but an unkind one could have F P F F F F F F F F F F F F F F F which could
	// create a linear time complexity anyways...

	// maybe there's some kind of search where you divide the input times until
	// you find a passing time, so in half, then quarters, then eights, etc
	// once you find a passing time you have a reduced window to search
	// but if the entire entry is 0 or 1 pass and the rest fails, then it
	// degrades to linear in the worst case scenario anyways, so sanitized or
	// expected inputs do go a long way
}

type race struct {
	time, distance int
}

func parseInputPart1(input string) (ans []race) {
	parts := strings.Split(input, "\n")
	timeParts := strings.Split(parts[0], " ")
	distParts := strings.Split(parts[1], " ")

	numRegexp := regexp.MustCompile("[0-9]+")

	var t, d int
	for t < len(timeParts) && d < len(distParts) {
		for !numRegexp.MatchString(timeParts[t]) {
			t++
		}
		for !numRegexp.MatchString(distParts[d]) {
			d++
		}
		ans = append(ans, race{
			time:     cast.ToInt(timeParts[t]),
			distance: cast.ToInt(distParts[d]),
		})
		t++
		d++
	}

	return ans
}

func parseInputPart2(input string) race {
	parts := strings.Split(input, "\n")
	timeLine := parts[0]
	distLine := parts[1]

	nonNums := regexp.MustCompile("[^0-9]+")
	timeLine = nonNums.ReplaceAllString(timeLine, "")
	distLine = nonNums.ReplaceAllString(distLine, "")

	return race{
		time:     cast.ToInt(timeLine),
		distance: cast.ToInt(distLine),
	}
}
