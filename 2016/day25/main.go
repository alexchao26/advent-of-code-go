package main

import (
	"flag"
	"fmt"
	"math"
	"regexp"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := part1(util.ReadFile("./input.txt"))
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	for i := 0; i < math.MaxInt32; i++ {
		if assemblyComputer(input, i) {
			return i
		}
	}
	return -1
}

func assemblyComputer(input string, registerAInitialValue int) bool {
	pattern := regexp.MustCompile("^(01){1,}$")
	var outputs string

	var a, b, d int

	a = registerAInitialValue
	d = a + 2532
	for {
		a = d
		for a != 0 {
			b = a % 2
			a /= 2
			outputs += cast.ToString(b)
			if len(outputs)%2 == 0 {
				if !pattern.MatchString(outputs) {
					return false
				} else if len(outputs) > 100 {
					return true
				}
			}
		}
	}

	panic("should return from loop")
}
