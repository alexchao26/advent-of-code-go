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

	ans := trickShot(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

// x = forward
// y = up/down

// starts at 0,0

// must be within target area (input) EVENTUALLY (aka on some step)

// find MAX Y position that can be reached for a valid initial velocity?

// ACTUAL: target area: x=88..125, y=-157..-103
// EXAMPLE: target area: x=20..30, y=-10..-5

func trickShot(input string, part int) int {
	// dummy state just to get some reasonable bounds for starting x velocities
	dummyState := newState(input, -1, -1)
	// hypothesis: the starting x velocity will be positive, (likely) under xmax+1 because then it
	// will pass the entire box within a single "step"
	// search bounds... (binary?) <- ended up being a trap
	leftXVel, rightXVel := 0, dummyState.xmax+1

	var highestY int                     // part 1
	var totalValidStartingVelocities int // part 2

	for xVel := leftXVel; xVel <= rightXVel; xVel++ {
		// brute force the starting y velocities :/
		for yVel := -1000; yVel <= 1000; yVel++ {
			st := newState(input, xVel, yVel)
			maybeHigherY, inBox := st.run()
			if inBox {
				highestY = mathy.MaxInt(highestY, maybeHigherY)

				// part2
				totalValidStartingVelocities++
			}
		}
	}

	if part == 1 {
		return highestY
	}

	return totalValidStartingVelocities
}

type state struct {
	xmin, xmax, ymin, ymax int
	x, y                   int
	xvel, yvel             int
	highestY               int
}

func newState(input string, startingXVel, startingYVel int) *state {
	//target area: x=88..125, y=-157..-103
	var xmin, xmax, ymin, ymax int
	n, err := fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	if n != 4 || err != nil {
		panic(fmt.Sprintf("%d read, want 4; error? %s", n, err))
	}

	return &state{
		xmin: xmin,
		xmax: xmax,
		ymin: ymin,
		ymax: ymax,
		xvel: startingXVel,
		yvel: startingYVel,
		// zero values handle the rest
	}
}

func (s *state) step() (reached, done bool) {
	// The probe's x position increases by its x velocity.
	s.x += s.xvel
	// The probe's y position increases by its y velocity.
	s.y += s.yvel

	s.highestY = mathy.MaxInt(s.highestY, s.y)

	// the probe's x velocity changes by 1 toward the value 0, stays zero if 0
	if s.xvel > 0 {
		s.xvel--
	} else if s.xvel < 0 {
		s.xvel++
	}

	// the probe's y velocity decreases by 1.
	s.yvel--

	// check if within bounds of (x|y)(min|max)
	if s.x >= s.xmin && s.x <= s.xmax && s.y >= s.ymin && s.y <= s.ymax {
		return true, true
	}

	// done overshot to the right
	if s.x > s.xmax {
		return false, true
	}

	// done: undershot to the left and no velocity to get more right
	if s.xvel == 0 && s.x < s.xmin {
		return false, true
	}

	// done: y is getting lower and lower and will never recover
	if s.y < s.ymin {
		return false, true
	}

	// indicate not reached but also not done, so call again
	return false, false
}

func (s *state) run() (maxY int, inBox bool) {
	var reached, done bool
	for !done {
		reached, done = s.step()
		if reached {
			return s.highestY, true
		}
	}
	return -1, false
}
