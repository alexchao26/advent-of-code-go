package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
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
	particles := parseInput(input)

	sort.Slice(particles, func(i, j int) bool {
		pI, pJ := particles[i], particles[j]
		// prioritize acceleration, smaller accelerations will stay closest to origin
		if pI.acceleration != pJ.acceleration {
			return sumAbs(pI.acceleration) < sumAbs(pJ.acceleration)
		}

		// then prioritize velocity
		if pI.velocity != pJ.velocity {
			return sumAbs(pI.velocity) < sumAbs(pJ.velocity)
		}

		// hopefully none of these are equal...
		if pI.positions == pJ.positions {
			panic("equal positions")
		}
		return sumAbs(pI.positions) < sumAbs(pJ.positions)
	})

	return particles[0].index
}

func sumAbs(nums [3]int) int {
	return mathy.AbsInt(nums[0]) + mathy.AbsInt(nums[1]) + mathy.AbsInt(nums[2])
}

func part2(input string) int {
	particles := parseInput(input)

	// this is a very literal solution and just run a large-ish number of times...
	// there is a mathematical way to determine position coordinates for a given
	// time, and then iterate from t = 0 -> large number... and remove like that
	// but I'm not sure if it's actually any more efficient
	for i := 0; i < math.MaxInt8; i++ {
		particles = tick(particles)
		particles = removeCollisions(particles)
	}
	return len(particles)
}

type particle struct {
	index        int
	positions    [3]int
	velocity     [3]int
	acceleration [3]int
}

func tick(particles []particle) []particle {
	var nextState []particle
	for _, p := range particles {
		for i, acc := range p.acceleration {
			p.velocity[i] += acc
		}
		for i, vel := range p.velocity {
			p.positions[i] += vel
		}
		nextState = append(nextState, p)
	}
	return nextState
}

func removeCollisions(particles []particle) []particle {
	set := map[[3]int]int{}
	for _, p := range particles {
		set[p.positions]++
	}

	var nextState []particle
	for _, p := range particles {
		if count, ok := set[p.positions]; ok && count == 1 {
			nextState = append(nextState, p)
		}
	}
	return nextState
}

func parseInput(input string) (particles []particle) {
	for i, line := range strings.Split(input, "\n") {
		p := particle{index: i}
		fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.positions[0],
			&p.positions[1],
			&p.positions[2],
			&p.velocity[0],
			&p.velocity[1],
			&p.velocity[2],
			&p.acceleration[0],
			&p.acceleration[1],
			&p.acceleration[2],
		)
		particles = append(particles, p)
	}
	return particles
}
