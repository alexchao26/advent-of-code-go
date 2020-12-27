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

	var ans int
	if part == 1 {
		ans = rambunctiousRecitation(util.ReadFile("./input.txt"), 2020)
	} else {
		// brute force, takes ~5seconds to run
		ans = rambunctiousRecitation(util.ReadFile("./input.txt"), 30000000)
	}
	fmt.Println("Output:", ans)
}

func rambunctiousRecitation(input string, turnToReturn int) int {
	var startingNums []int
	for _, num := range strings.Split(input, ",") {
		startingNums = append(startingNums, cast.ToInt(num))
	}

	said := map[int][]int{}
	var numSaidLast int
	// populate the starting map with the numbers from the input
	for i, num := range startingNums {
		said[num] = append(said[num], i+1)
		numSaidLast = num
	}

	// then picking up from the next turn (len of input), continue until the
	// desired turn is reached
	for i := len(startingNums); i <= turnToReturn; i++ {
		indexSlice := said[numSaidLast]
		// if the length of the slice of incides is 1, that means it was only
		// called once. Say zero, add to the zero slice of indices
		if len(indexSlice) == 1 {
			numSaidLast = 0
			said[0] = append(said[0], i)
		} else {
			// otherwise determine the diff between the last 2 times it was said
			length := len(indexSlice)
			numSaidLast = indexSlice[length-1] - indexSlice[length-2]
			// add this index i to the indices slice for the number that is said
			said[numSaidLast] = append(said[numSaidLast], i)
		}
	}

	return numSaidLast
}
