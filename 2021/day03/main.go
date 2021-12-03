package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

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
	var gamma, epsilon string
	binaries := strings.Split(input, "\n")
	for i := 0; i < len(binaries[0]); i++ {
		var zeroes, ones int
		for _, b := range binaries {
			if b[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}

		// gamma is the most common bit at each index (accross all bits)
		// epsilon is the opposite
		if zeroes > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"

		}
	}

	// multiply together for final ans
	// actual answer in decimal, not binary

	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(e * g)
}

func part2(input string) int {
	// filtering values until one remains

	// consider just first bit of each number
	//   only keep numbers for the correct bit criteria
	// stop when only one number is left, otherwise continue onto the next bit
	nums := strings.Split(input, "\n")
	// assume this will work and len(nums) will eventually hit 1
	for i := 0; len(nums) > 1; i++ {
		var zeroes, ones int
		for _, n := range nums {
			if n[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		// bit criteria:
		//   oxygen: most common value of the bit, 1s win ties
		//   CO2:    opposite, 0s win ties
		keepChar := "1" // 1s win ties for oxygen
		if zeroes > ones {
			keepChar = "0"
		}

		var newNums []string
		for _, n := range nums {
			if string(n[i]) == keepChar {
				newNums = append(newNums, n)
			}
		}
		nums = newNums
	}
	oxygen := nums[0]

	// copy pasta for co2
	nums = strings.Split(input, "\n")
	// assume this will work and len(nums) will eventually hit 1
	for i := 0; len(nums) > 1; i++ {
		var zeroes, ones int
		for _, n := range nums {
			if n[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		// bit criteria:
		//   oxygen: most common value of the bit, 1s win ties
		//   CO2:    opposite, 0s win ties
		keepChar := "0" // 0s win ties for co2
		if ones < zeroes {
			keepChar = "1"
		}

		var newNums []string
		for _, n := range nums {
			if string(n[i]) == keepChar {
				newNums = append(newNums, n)
			}
		}
		nums = newNums
	}
	co2 := nums[0]

	// multiplying the oxygen generator rating by the CO2 scrubber rating.
	o, _ := strconv.ParseInt(oxygen, 2, 64)
	c, _ := strconv.ParseInt(co2, 2, 64)

	return int(c * o)
}
