package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := memoryReallocation(util.ReadFile("./input.txt"), 1)
		fmt.Println("Output:", ans)
	} else {
		ans := memoryReallocation(util.ReadFile("./input.txt"), 2)
		fmt.Println("Output:", ans)
	}
}

func memoryReallocation(input string, part int) int {
	banks := parseInput(input)

	// [16]int arrays are comparable by values so can be used as map keys
	seenBanks := map[[16]int]int{banks: 0}
	var cycles int
	for {
		// find largest bank
		var index, maxVal int
		for i, val := range banks {
			if val > maxVal {
				index = i
				maxVal = val
			}
		}

		// run a cycle
		blocksToDistribute := banks[index]
		banks[index] = 0
		// unoptimized but works just fine
		for i := index + 1; blocksToDistribute > 0; i++ {
			if blocksToDistribute == 0 {
				break
			}
			banks[i%16]++
			blocksToDistribute--
		}
		cycles++

		// check if this set of banks has been seen before, if so return here
		if val, ok := seenBanks[banks]; ok {
			if part == 1 {
				return cycles
			}
			// for part 2 take the difference with cycles when it was last seen
			return cycles - val
		}
		// set the number of cycles that correspond to this state of banks
		seenBanks[banks] = cycles
	}

	panic("should resolve in for loop")
}

func parseInput(input string) (ans [16]int) {
	nums := strings.Split(input, "\t")
	for i, num := range nums {
		ans[i] = cast.ToInt(num)
	}
	return ans
}
