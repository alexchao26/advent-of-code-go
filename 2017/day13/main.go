package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

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
	freqs := parseFrequencies(input)

	var severity int
	for _, freq := range freqs {
		// depthIndex is also equivalent to the number of picoseconds that have
		// elapsed when we reach a particular depth of the firewall
		depthIndex := freq[0]

		// frequency to zero index is how many picoseconds it takes the scanner
		// to return to the zero index position
		frequencyToZeroIndex := (freq[1] - 1) * 2

		// if frequency is evenly divisible by the time elapsed, then we're
		// "caught" by this scanner, add to severity
		if depthIndex%frequencyToZeroIndex == 0 {
			severity += depthIndex * freq[1]
		}
	}

	return severity
}

func part2(input string) int {
	freqs := parseFrequencies(input)

	// same logic to part 1, but add a delay to the time elapsed and return which
	// delay leads to not getting caught by a scanner
	for delay := 0; delay < math.MaxInt32; delay++ {
		var gotCaught bool
		for _, freq := range freqs {
			depthIndex := freq[0]
			frequencyToZeroIndex := (freq[1] - 1) * 2
			// add delay to the time depthIndex, as it takes longer to get there
			if (depthIndex+delay)%frequencyToZeroIndex == 0 {
				gotCaught = true
			}
		}
		if !gotCaught {
			return delay
		}
	}

	panic("loop ended, increase limit?")
}

func parseFrequencies(input string) [][2]int {
	lines := strings.Split(input, "\n")
	var freqs [][2]int // depth, range
	for _, l := range lines {
		// depth is equivalent to the index in the firewall
		// range can be used to calculate the frequency in which a scanner
		// returns to index zero
		var depth, rng int
		fmt.Sscanf(l, "%d: %d", &depth, &rng)
		freqs = append(freqs, [2]int{depth, rng})
	}
	return freqs
}
