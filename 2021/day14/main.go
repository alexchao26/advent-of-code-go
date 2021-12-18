package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

// naive/brute force solution that does not work for part 2 :(
func part1(input string) int {
	template, rules := parseInput(input)

	for step := 0; step < 10; step++ {
		sp := strings.Split(template, "")
		var next []string
		for i := 0; i < len(sp)-1; i++ {
			two := sp[i] + sp[i+1]
			rule := rules[two]
			next = append(next, sp[i])
			next = append(next, rule)
		}
		next = append(next, sp[len(sp)-1])

		template = strings.Join(next, "")
	}

	// most common element minus least common element
	most, least := 0, math.MaxInt32
	count := map[rune]int{}
	for _, r := range template {
		count[r]++
	}
	for _, ct := range count {
		if ct > most {
			most = ct
		}
		if ct < least {
			least = ct
		}
	}

	return most - least
}

func part2(input string) int {
	template, rules := parseInput(input)

	initialCount := map[string]int{}
	for _, r := range template {
		initialCount[string(r)]++
	}

	addlCountAfter40Steps := grow(template, 40, rules, map[string]map[string]int{})
	// have to add the initial template back in because grow only accounts for characters that are
	// added in ADDITION to the passed in template
	for k, v := range initialCount {
		addlCountAfter40Steps[k] += v
	}

	// most common element minus least common element
	most, least := 0, math.MaxInt64
	for _, v := range addlCountAfter40Steps {
		most = mathy.MaxInt(most, v)
		least = mathy.MinInt(least, v)
	}
	return most - least
}

func grow(template string, stepsLeft int, rules map[string]string, memo map[string]map[string]int) (addlCounts map[string]int) {
	addlCounts = map[string]int{}

	if stepsLeft == 0 {
		return addlCounts
	}

	key := fmt.Sprint(template, stepsLeft)
	if res, ok := memo[key]; ok {
		return res
	}

	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		between := rules[pair]
		addlCounts[between]++

		// get the additional characters for recursing just this three character section
		// calling grow will memoize that result to eliminate duplicate work
		recurse := grow(pair[:1]+between+pair[1:], stepsLeft-1, rules, memo)

		// merge those additional characters into this (parent function call) addlCounts map
		for k, v := range recurse {
			addlCounts[k] += v
		}
	}

	// store it, return it
	memo[key] = addlCounts
	return addlCounts
}

func parseInput(input string) (template string, rules map[string]string) {
	parts := strings.Split(input, "\n\n")
	template = parts[0]

	rules = map[string]string{}
	for _, line := range strings.Split(parts[1], "\n") {
		sp := strings.Split(line, " -> ")
		rules[sp[0]] = sp[1]
	}
	return template, rules
}
