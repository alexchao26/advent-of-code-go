package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

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
	steps, graph := parseInput(input)

	location := "AAA"
	numOfSteps := 0
	for location != "ZZZ" {
		dir := steps[numOfSteps%len(steps)]

		if dir == "L" {
			location = graph[location][0]
		} else {
			location = graph[location][1]
		}

		numOfSteps++
	}

	return numOfSteps
}

func part2(input string) int {
	steps, graph := parseInput(input)

	locations := []string{}
	for loc := range graph {
		if loc[2:3] == "A" {
			locations = append(locations, loc)
		}
	}

	/**
	brute force doesn't work... need to figure out cycle times of each starting location
	but they won't cycle just based on number of steps because of the weird L-R randomness

	so we can only rely on the "full cycle" of all steps before it loops

	- there are six starting locations

	NOTE: BIG assumptions based on KIND inputs
	- assume that the Z-end locations will sync EXACTLY at the end of a cycle of steps
	- after further analyzing logs of the end of each cycle, the entry point VERY kindly deposits us
	  	at the very start of a cycle that will eventually end in a Z-end location
		AAA -> MLM -> ... -> XKZ -> MLM -> ... -> XKZ -> MLM -> ... -> XKZ -> MLM
		and this holds true for all six locations in my input
		Therefore the cycles are not offset by a particular number of steps at the start to get to the cycle
		such as START --> LOC1 --> LOC2 --> Start -> A
											  ^		 |
											  |		 v
											  D <--  C
		this makes the maths fairly straight forward with just having to find the LCM (least common multiple)
		of all the cycle periods because that is when they will all sync up and land on a Z
	*/

	numOfSteps := 0

	locationCyclePeriods := []int{}
	for cycle := 0; len(locations) > 0; cycle++ {
		for _, dir := range steps {
			for i, loc := range locations {
				if dir == "L" {
					locations[i] = graph[loc][0]
				} else {
					locations[i] = graph[loc][1]
				}
			}
			numOfSteps++
		}

		// if any location is at a z-end at the end of a cycle, record the cycle time
		// to do the final maths at the end
		newLocations := []string{}
		for _, loc := range locations {
			if loc[2:3] == "Z" {
				locationCyclePeriods = append(locationCyclePeriods, numOfSteps)
			} else {
				newLocations = append(newLocations, loc)
			}
		}
		locations = newLocations
	}

	// combine all into an LCM (helper function added to mathy package)
	lcm := locationCyclePeriods[0]
	for i := 1; i < len(locationCyclePeriods); i++ {
		lcm = mathy.LeastCommonMultiple(lcm, locationCyclePeriods[i])
	}

	return lcm
}

func parseInput(input string) (steps []string, graph map[string][]string) {
	graph = map[string][]string{}

	parts := strings.Split(input, "\n\n")
	steps = strings.Split(parts[0], "")

	for _, line := range strings.Split(parts[1], "\n") {
		graph[line[0:3]] = []string{
			line[7:10],
			line[12:15],
		}
	}

	return steps, graph
}
