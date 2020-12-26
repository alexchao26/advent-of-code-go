package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	var totalSqFt int
	for _, line := range strings.Split(input, "\n") {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%dx%dx%d", &x, &y, &z)
		if err != nil {
			panic(err)
		}
		totalSqFt += x * y * 2
		totalSqFt += x * z * 2
		totalSqFt += z * y * 2
		totalSqFt += mathutil.MinInt(x*y, y*z, x*z) // slack in wrapping paper...
	}

	return totalSqFt
}

func part2(input string) int {
	var totalLen int
	for _, line := range strings.Split(input, "\n") {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%dx%dx%d", &x, &y, &z)
		if err != nil {
			panic(err)
		}
		cubic := x * y * z
		totalLen += cubic
		sides := []int{
			2 * (x + y),
			2 * (y + z),
			2 * (x + z),
		}
		totalLen += mathutil.MinInt(sides...)
	}
	return totalLen
}
