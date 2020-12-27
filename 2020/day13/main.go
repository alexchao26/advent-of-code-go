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
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	estimate, busses := parseInput(input)

	var busID, waited int
	for timeValue := estimate; busID == 0; timeValue++ {
		for _, inter := range busses {
			if timeValue%inter[1] == 0 {
				busID = inter[1]
				waited = timeValue - estimate
				break
			}
		}
	}

	return busID * waited
}

// i didn't come up with this myself... generally speaking i understand it...
func part2(input string) int {
	_, busses := parseInput(input)

	var timeValue int
	runningProduct := 1
	for _, bus := range busses {
		index, busID := bus[0], bus[1]
		// this for loop adjusts the time until the constaint for this bus is met
		// i.e. ensure (time + index) is divisible by the busID to ensure the bus arrives
		for (timeValue+index)%busID != 0 {
			// running product is used to increment because it will not affect
			// the modulo of any of the previously scheduled busses, we've found
			// the frequency to match them.
			// e.g. if busID: 5 & index: 2, min timeValue is 3 b/c (3+2)%5 == 0
			//      if the running product were 5, adding 5 means (8+2)%5 == 0
			//      and (3 + 5x + 2) % 5 == 0 for any x
			timeValue += runningProduct
		}
		runningProduct *= busID
	}

	return timeValue
}

// busses are [2]int{index, busID}, not the best way to parse stuff but it works
func parseInput(input string) (estimate int, busses [][2]int) {
	lines := strings.Split(input, "\n")
	estimate = cast.ToInt(lines[0])
	for index, busID := range strings.Split(lines[1], ",") {
		if busID != "x" {
			busses = append(busses, [2]int{index, cast.ToInt(busID)})
		}
	}
	return estimate, busses
}
