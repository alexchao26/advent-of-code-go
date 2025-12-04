package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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
	nums := strings.Split(input, "\n")

	ans := 0
	for _, num := range nums {
		largest := 0
		highestDigitSoFar := 0
		for i := range len(num) {
			char := num[i : i+1]
			digit := cast.ToInt(char)

			largest = max(largest, highestDigitSoFar*10+digit)
			highestDigitSoFar = max(highestDigitSoFar, digit)
		}
		ans += largest
	}

	return ans
}

func part2(input string) int {
	nums := strings.Split(input, "\n")

	ans := 0

	for _, num := range nums {
		digits := []int{}
		for _, d := range strings.Split(num, "") {
			digits = append(digits, cast.ToInt(d))
		}
		battery := make([]int, 12)
		for i := range len(battery) {
			battery[i] = digits[len(digits)-12+i]
		}

		// building from the back... start with 12 digits
		// iterate backwards from the 13th-to-last digit
		// if that is greater than battery[0], replaces battery[0], check if prev battery[0] should replace battery[1], etc etc
		//    if two digits are equal, continue checking (2 -> 1110 should become 2111 and need to evict final 0)
		// else if less than, stop checking
		for i := len(digits) - 12 - 1; i >= 0; i-- {
			digit := digits[i]
			for j := range battery {
				if digit >= battery[j] {
					digit, battery[j] = battery[j], digit
				} else {
					break
				}
			}
		}

		batteryTotal := 0
		for _, d := range battery {
			batteryTotal *= 10
			batteryTotal += d
		}
		ans += batteryTotal
	}

	return ans
}
