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

// ~1200
func part1(input string) int {
	graph, messages := parseInput(input)

	fillInGraph(graph, 0)

	var matchRuleZero int
	for _, m := range messages {
		for _, opt := range graph[0].resolved {
			if m == opt {
				matchRuleZero++
				break
			}
		}
	}

	return matchRuleZero
}

// 1018 @ 1:19AM
func part2(input string) int {
	graph, messages := parseInput(input)

	// these are the dependencies of rules 8 & 11 (which are the only dependencies
	// of rule 0). fill in the resolved fields for these graph nodes
	// b/c these definitions are not circular (downstream of changes), they will
	// not infinite loop
	fillInGraph(graph, 42)
	fillInGraph(graph, 31)

	// generate regexp strings that will be used to match against rules 8 and 11
	// and ultimate match rule 0
	part42 := fmt.Sprintf("(%s)", strings.Join(graph[42].resolved, "|"))
	part31 := fmt.Sprintf("(%s)", strings.Join(graph[31].resolved, "|"))

	// rule 8 is essentially 1 or more instances of rule 42
	rule8String := fmt.Sprintf("(%s)+", part42)

	// note: i'm unaware of how to make two regexp portions have the same number
	// of segments, so I made this helper function that changes that number
	// then I just run it ten times because that should be large enough to cover
	// any of the rules in the input...
	makeRegexp := func(num int) *regexp.Regexp {
		// rule 11 is an equal number of 42 and 31 rules
		return regexp.MustCompile(fmt.Sprintf("^%s%s{%d}%s{%d}$", rule8String, part42, num, part31, num))
	}

	var matchRuleZero int
	for _, m := range messages {
		for i := 1; i < 10; i++ {
			pattern := makeRegexp(i)
			if pattern.MatchString(m) {
				matchRuleZero++
				break
			}
		}
	}

	return matchRuleZero
}

func fillInGraph(graph map[int]*rule, entry int) []string {
	if len(graph[entry].resolved) != 0 {
		// return a copy of resolved otherwise there's all kinds of side effect errors
		return append([]string{}, graph[entry].resolved...)
	}

	// iterate through options, resolve children and append resolved paths
	// for the current entry point
	for _, option := range graph[entry].options {
		// this will be all permutations generated from recursive calls to fillInGraph
		// Note: there's probably a cleaner algorithm to do this kind of perm generation...
		resolved := []string{""}
		for _, entryPoint := range option {
			nestedResolveVals := fillInGraph(graph, entryPoint)
			var newResolved []string
			for _, nextPiece := range nestedResolveVals {
				for _, prev := range resolved {
					newResolved = append(newResolved, prev+nextPiece)
				}
			}
			resolved = newResolved
		}
		graph[entry].resolved = append(graph[entry].resolved, resolved...)
	}

	return graph[entry].resolved
}

type rule struct {
	resolved []string
	options  [][]int
}

// Stringer interface for debugging
func (r rule) String() string {
	var ans string
	ans += fmt.Sprintf("OPTIONS:  %v\n", r.options)
	ans += fmt.Sprintf(" RESOLVED: %v\n", r.resolved)
	return ans
}

func parseInput(input string) (rules map[int]*rule, messages []string) {
	parts := strings.Split(input, "\n\n")

	rules = map[int]*rule{}
	for _, r := range strings.Split(parts[0], "\n") {
		if regexp.MustCompile("[a-z]").MatchString(r) {
			var num int
			var char string
			fmt.Sscanf(r, "%d: \"%1s\"", &num, &char)
			rules[num] = &rule{resolved: []string{char}}
		} else {
			split := strings.Split(r, ": ")
			key := mathutil.StrToInt(split[0])
			newRule := rule{}
			for _, ruleNums := range strings.Split(split[1], " | ") {
				nums := strings.Split(ruleNums, " ")
				var option []int
				for _, n := range nums {
					option = append(option, mathutil.StrToInt(n))
				}
				newRule.options = append(newRule.options, option)
			}
			rules[key] = &newRule
		}
	}

	messages = strings.Split(parts[1], "\n")

	return rules, messages
}
