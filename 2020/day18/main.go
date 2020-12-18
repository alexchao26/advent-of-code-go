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
		total += doMaths(line)
	}

	return total
}

// ~1700 :(
func part2(input string) int {
	lines := parseInput(input)
	var total int

	for _, line := range lines {
		total += doMaths2(line)
	}

	return total
}

// 370?
func parseInput(input string) (ans [][]string) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		ans = append(ans, strings.Split(strings.ReplaceAll(l, " ", ""), ""))
	}
	return ans
}

func doMaths(input []string) int {
	var parensBalance int
	var sum int
	lastOperator := "+"
	var nested []string
	for _, val := range input {
		switch val {
		case "+", "*":
			if parensBalance == 0 {
				lastOperator = val
			} else {
				nested = append(nested, val)
			}
		case "(":
			if parensBalance != 0 {
				nested = append(nested, "(")
			}
			parensBalance++
		case ")":
			parensBalance--
			if parensBalance == 0 {
				if lastOperator == "+" {
					sum += doMaths(nested)
				} else {
					sum *= doMaths(nested)
				}
				nested = []string{}
			} else {
				nested = append(nested, ")")
			}
		default:
			if parensBalance == 0 {
				if lastOperator == "+" {
					sum += mathutil.StrToInt(val)
				} else {
					sum *= mathutil.StrToInt(val)
				}
			} else {
				nested = append(nested, val)
			}
		}
	}

	return sum
}

var numReg = regexp.MustCompile("[0-9]")

func isNum(str string) bool {
	return numReg.MatchString(str)
}

func doMaths2(input []string) int {
	var stack []int
	for i := 0; i < len(input); i++ {
		// iterate throuhg input
		// track open and close parens
		switch val := input[i]; val {
		case "(":
			stack = append(stack, i)
		case ")":
			// on close parens, pass a section tohandleFlat
			// then remove a bunch of shit from input,
			openIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if input[openIndex] == "+" || input[openIndex] == "*" {
				openIndex++
			}
			sliToFlatten := input[openIndex+1 : i] // will not include the closing paren

			flatVal := handleFlat(sliToFlatten)
			input[openIndex] = flatVal

			copy(input[openIndex+1:], input[i+1:])
			input = input[:len(input)-(i-openIndex)]
			i = openIndex
		default:
			// continue along
		}

	}

	return mathutil.StrToInt(handleFlat(input))
}

func handleFlat(input []string) string {
	for _, v := range input {
		if v == "(" || v == ")" {
			panic("unexpected paren in flat input")
		}
	}

	// handle all additions
	for i := 1; i < len(input)-1; i++ {
		if input[i] == "+" {
			toLeft := input[i-1]
			toRight := input[i+1]
			if numReg.MatchString(toLeft) && numReg.MatchString(toRight) {
				input[i-1] = mathutil.IntToStr(mathutil.StrToInt(toLeft) + mathutil.StrToInt(toRight))
				// remove two things
				for j := i; j < len(input)-2; j++ {
					input[j] = input[j+2]
				}
				input = input[:len(input)-2]
				i--
			}
		}
	}

	// then handle all multiplications
	for i := 1; i < len(input)-1; i++ {
		if input[i] == "*" {
			toLeft := input[i-1]
			toRight := input[i+1]
			if numReg.MatchString(toLeft) && numReg.MatchString(toRight) {
				input[i-1] = mathutil.IntToStr(mathutil.StrToInt(toLeft) * mathutil.StrToInt(toRight))
				// remove two things
				for j := i; j < len(input)-2; j++ {
					input[j] = input[j+2]
				}
				input = input[:len(input)-2]
				i--
			}
		}
	}

	return input[0]
}

// removes a particular number of elements from the middle of the slice
func splice(sli []string, startIndex, items int) []string {
	copy(sli[startIndex:], sli[startIndex+items:])
	sli = sli[:len(sli)-items]
	return sli
}
