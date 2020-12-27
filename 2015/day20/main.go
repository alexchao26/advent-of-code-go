package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/alexchao26/advent-of-code-go/cast"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := infiniteElvesAndHouses(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func infiniteElvesAndHouses(input string, part int) int {
	targetNum := cast.ToInt(input)

	for house := 1; house < math.MaxInt32; house++ {
		var gifts int
		for _, factor := range getFactors(house) {
			if part == 1 {
				gifts += factor * 10
			} else if part == 2 && house/factor <= 50 {
				// for part 2, ensure that this is the 50th or less house that
				// the elf has visited before adding gifts
				gifts += factor * 11
			}
		}
		if gifts >= targetNum {
			return house
		}
	}

	panic("expect return from loop")
}

func getFactors(num int) []int {
	var factors []int
	sqrt := int(math.Sqrt(float64(num)))
	for i := 1; i <= sqrt; i++ {
		if num%i == 0 {
			factors = append(factors, i, num/i)
		}
	}
	return factors
}
