package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
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

// 392
func part1(input string) int {
	parsed := parseInput(input)

	var total int
	for _, line := range parsed {
		// sum up this line, addressing things inside parens first...
		total += doMaths(line, calcFlatSlicePart1)
	}

	return total
}

// ~1700 :(
func part2(input string) int {
	lines := parseInput(input)
	var total int

	for _, line := range lines {
		total += doMaths(line, calcFlatSlicePart2)
	}

	return total
}

func parseInput(input string) (ans [][]string) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		// I got lucky that they were all single digit numbers...
		ans = append(ans, strings.Split(strings.ReplaceAll(l, " ", ""), ""))
	}
	return ans
}

func doMaths(input []string, flatteningFunc func([]string) string) int {
	var stackOpenIndices []int
	var stackFlattened []string
	for i := 0; i < len(input); i++ {
		// iterate through input, always append onto the flattened stack
		// track open paren indices (in the flattened stack)
		// on closing parens, use the top of the stackOpenIndices to flatten
		// the most recent set of values/operations within parens, and replace
		// their opening paren with the flattneed value
		stackFlattened = append(stackFlattened, input[i])
		switch input[i] {
		case "(":
			stackOpenIndices = append(stackOpenIndices, len(stackFlattened)-1)
		case ")":
			// on close parens, pass a section tohandleFlat
			// then remove a bunch of shit from input,
			openIndex := stackOpenIndices[len(stackOpenIndices)-1]
			stackOpenIndices = stackOpenIndices[:len(stackOpenIndices)-1]

			// do not include leading or trailing paren
			sliToFlatten := stackFlattened[openIndex+1 : len(stackFlattened)-1]
			stackFlattened[openIndex] = flatteningFunc(sliToFlatten)

			// remove the values that were flattened off the top of the stack
			stackFlattened = stackFlattened[:openIndex+1]
		}

	}
	// slice should now be flat
	return mathutil.StrToInt(flatteningFunc(stackFlattened))
}

func calcFlatSlicePart1(input []string) string {
	for _, v := range input {
		if v == "(" || v == ")" {
			panic(fmt.Sprintf("unexpected paren in flat input, %v", input))
		}
	}

	result := mathutil.StrToInt(input[0])

	for i := range input {
		if i+2 < len(input) {
			switch input[i+1] {
			case "+":
				result += mathutil.StrToInt(input[i+2])
			case "*":
				result *= mathutil.StrToInt(input[i+2])
			}
		}
	}

	return mathutil.IntToStr(result)
}

func calcFlatSlicePart2(input []string) string {
	for _, v := range input {
		if v == "(" || v == ")" {
			panic(fmt.Sprintf("unexpected paren in flat input, %v", input))
		}
	}

	// handle all additions
	for i := 1; i < len(input)-1; i++ {
		if input[i] == "+" {
			toLeft := input[i-1]
			toRight := input[i+1]
			if isNum(toLeft) && isNum(toRight) {
				input[i-1] = addStrings(toLeft, toRight)
				input = splice(input, i, 2)
				i--
			}
		}
	}

	// then handle all multiplications
	for i := 1; i < len(input)-1; i++ {
		if input[i] == "*" {
			toLeft := input[i-1]
			toRight := input[i+1]
			if isNum(toLeft) && isNum(toRight) {
				input[i-1] = multiplyStrings(toLeft, toRight)
				input = splice(input, i, 2)
				i--
			}
		}
	}

	return input[0]
}

var numReg = regexp.MustCompile("[0-9]")

func isNum(str string) bool {
	return numReg.MatchString(str)
}

func addStrings(strs ...string) string {
	var sum int
	for _, str := range strs {
		sum += mathutil.StrToInt(str)
	}
	return mathutil.IntToStr(sum)
}
func multiplyStrings(strs ...string) string {
	sum := 1
	for _, str := range strs {
		sum *= mathutil.StrToInt(str)
	}
	return mathutil.IntToStr(sum)
}

// removes a particular number of elements from the middle of the slice
func splice(sli []string, startIndex, items int) []string {
	copy(sli[startIndex:], sli[startIndex+items:])
	sli = sli[:len(sli)-items]
	return sli
}
