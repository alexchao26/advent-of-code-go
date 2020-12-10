package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	recipesWanted := parseInput(input)

	recipes := []int{3, 7}
	elf1, elf2 := 0, 1

	for len(recipes) < recipesWanted+10 {
		recipes, elf1, elf2 = step(recipes, elf1, elf2)
		// fmt.Println("new recipes", recipes)
		// fmt.Println("new elf indices", elf1, elf2)
	}

	ans := 0
	for i := 0; i < 10; i++ {
		ans *= 10
		ans += recipes[i+recipesWanted]
	}

	return ans
}

func part2(input string) int {
	patternToFind := parseInput(input)

	recipes := []int{3, 7}
	elf1, elf2 := 0, 1

	var recipesToLeft int
	for recipesToLeft == 0 {
		recipes, elf1, elf2 = step(recipes, elf1, elf2)

		// check patterns
		recipesToLeft = patternMatch(recipes, patternToFind)
	}

	return recipesToLeft
}

func parseInput(input string) int {
	lines := strings.Split(input, "\n")

	return mathutil.StrToInt(lines[0])
}

func step(recipes []int, elf1, elf2 int) ([]int, int, int) {
	recipe1, recipe2 := recipes[elf1], recipes[elf2]
	newRecipe := recipe1 + recipe2

	// add new recipes onto slice
	if newRecipe >= 10 {
		recipes = append(recipes, 1)
	}
	recipes = append(recipes, newRecipe%10)

	// get new elf indices
	elf1 += recipe1 + 1
	elf2 += recipe2 + 1
	elf1 %= len(recipes)
	elf2 %= len(recipes)

	return recipes, elf1, elf2
}

func patternMatch(recipes []int, patternToFind int) int {
	patternLength := len(strconv.Itoa(patternToFind))
	// not enough recipes to compare
	if len(recipes) < patternLength {
		return 0
	}

	// check last two recipes in recipes slice (because it will grow by at most
	// 2 recipes per step
	var pattern int
	// get first pattern
	for i := len(recipes) - patternLength; i < len(recipes); i++ {
		pattern *= 10
		pattern += recipes[i]
	}

	if pattern == patternToFind {
		return len(recipes) - patternLength
	}

	// check second pattern
	// removes one's digit
	pattern /= 10
	// add new 10^6 digit
	if index := len(recipes) - patternLength - 1; index >= 0 {
		pattern += 100000 * recipes[index]
	}

	if pattern == patternToFind {
		return len(recipes) - patternLength - 1
	}

	// return zero when no pattern is found
	return 0
}
