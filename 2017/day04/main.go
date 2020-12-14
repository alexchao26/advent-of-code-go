package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

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
	phrases := parseInput(input)
	var valid int
	for _, phrase := range phrases {
		seen := map[string]bool{}
		valid++ // include by default, remove if it is not valid
		for _, word := range phrase {
			// if word is a duplicate, decrement valid & break out
			if seen[word] {
				valid--
				break
			}
			seen[word] = true
		}
	}

	return valid
}

func part2(input string) int {
	phrases := parseInput(input)
	var valid int
	for _, phrase := range phrases {
		seen := map[string]bool{}
		valid++ // include by default, remove if it is not valid
		for _, word := range phrase {
			sortedWord := orderCharacters(word)
			// if word is a duplicate, decrement valid & break out
			if seen[sortedWord] {
				valid--
				break
			}
			seen[sortedWord] = true
		}
	}

	return valid
}

func orderCharacters(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func parseInput(input string) (ans [][]string) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		ans = append(ans, strings.Split(l, " "))
	}
	return ans
}
