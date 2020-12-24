package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := firewall(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func firewall(input string, part int) int {
	var allBlockedRanges [][2]int
	for _, line := range strings.Split(input, "\n") {
		var r [2]int
		fmt.Sscanf(line, "%d-%d", &r[0], &r[1])
		allBlockedRanges = append(allBlockedRanges, r)
	}
	sort.Slice(allBlockedRanges, func(i, j int) bool {
		if allBlockedRanges[i][0] != allBlockedRanges[j][0] {
			return allBlockedRanges[i][0] < allBlockedRanges[j][0]
		}
		return allBlockedRanges[i][1] < allBlockedRanges[j][1]
	})

	// merge allBlockedRanges
	merged := [][2]int{[2]int{}}
	for _, r := range allBlockedRanges {
		endOfLastRange := merged[len(merged)-1][1]
		if endOfLastRange >= r[0]-1 {
			merged[len(merged)-1][1] = mathutil.MaxInt(endOfLastRange, r[1])
		} else {
			merged = append(merged, r)
		}
	}

	if part == 1 {
		return merged[0][1] + 1
	}

	if merged[len(merged)-1][1] != math.MaxUint32 {
		merged = append(merged, [2]int{math.MaxUint32, 0})
	}

	var totalAllowed int
	for i := 1; i < len(merged); i++ {
		totalAllowed += merged[i][0] - merged[i-1][1] - 1
	}

	return totalAllowed
}
