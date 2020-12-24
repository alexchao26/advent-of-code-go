package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := timingIsEverything(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func timingIsEverything(input string, part int) int {
	discs := parseInput(input)

	if part == 2 {
		discs = append(discs, &disc{
			number:    len(discs) + 1,
			positions: 11,
			starting:  0,
		})
	}

	t := 0
	for {
		var capsuleCollides bool
		for _, d := range discs {
			// some math equation for position must equal zero for the capsule to pass through each disc
			timeSinceDrop := d.number
			position := d.starting + t + timeSinceDrop
			position %= d.positions
			if position != 0 {
				capsuleCollides = true
			}
		}
		if !capsuleCollides {
			break
		}
		t++
	}

	return t
}

type disc struct {
	number    int
	positions int
	starting  int
}

func parseInput(input string) []*disc {
	var discs []*disc
	for _, l := range strings.Split(input, "\n") {
		d := disc{}
		fmt.Sscanf(l, "Disc #%d has %d positions; at time=0, it is at position %d.",
			&d.number, &d.positions, &d.starting)
		discs = append(discs, &d)
	}
	return discs
}
