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

	var ans int
	if part == 1 {
		ans = balanceBots(util.ReadFile("./input.txt"), []int{17, 61})
	} else {
		ans = balanceBots(util.ReadFile("./input.txt"), nil)
	}
	fmt.Println("Output:", ans)
}

func balanceBots(input string, part1CompareValues []int) int {
	botsMap, rules := parseInput(input)
	outputs := map[int]int{}

	// for loop conditional is for part 2. part 1 returns from inside the loop.
	for outputs[0] == 0 || outputs[1] == 0 || outputs[2] == 0 {
		for _, r := range rules {
			if len(botsMap[r.botID]) == 2 {
				sort.Ints(botsMap[r.botID])
				low, high := botsMap[r.botID][0], botsMap[r.botID][1]
				// part 1 return value
				if len(part1CompareValues) != 0 &&
					low == part1CompareValues[0] && high == part1CompareValues[1] {
					return r.botID
				}
				var outputIndex, receivingBot int
				if strings.Contains(r.lowRule, "output") {
					_, err := fmt.Sscanf(r.lowRule, "low to output %d", &outputIndex)
					if err != nil {
						panic(err)
					}
					outputs[outputIndex] = low
				} else {
					_, err := fmt.Sscanf(r.lowRule, "low to bot %d", &receivingBot)
					if err != nil {
						panic(err)
					}
					botsMap[receivingBot] = append(botsMap[receivingBot], low)
				}
				if strings.Contains(r.highRule, "output") {
					_, err := fmt.Sscanf(r.highRule, "high to output %d", &outputIndex)
					if err != nil {
						panic(err)
					}
					outputs[outputIndex] = high
				} else {
					_, err := fmt.Sscanf(r.highRule, "high to bot %d", &receivingBot)
					if err != nil {
						panic(err)
					}
					botsMap[receivingBot] = append(botsMap[receivingBot], high)
				}
				botsMap[r.botID] = []int{}
			}
		}
	}

	// part 2 output
	return outputs[0] * outputs[1] * outputs[2]
}

type rule struct {
	botID    int
	lowRule  string
	highRule string
}

func parseInput(input string) (botsMap map[int][]int, rules []rule) {
	botsMap = map[int][]int{}
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "value") {
			var value, botID int
			_, err := fmt.Sscanf(line, "value %d goes to bot %d", &value, &botID)
			if err != nil {
				panic(err)
			}
			botsMap[botID] = append(botsMap[botID], value)
		} else {
			parts := splitOn(line, []string{" gives ", " and "})
			r := rule{lowRule: parts[1], highRule: parts[2]}
			_, err := fmt.Sscanf(parts[0], "bot %d", &r.botID)
			if err != nil {
				panic(err)
			}
			rules = append(rules, r)
		}
	}
	return botsMap, rules
}

func splitOn(in string, cutset []string) []string {
	parts := strings.Split(in, cutset[0])
	cutset = cutset[1:]
	var done bool
	for !done && len(cutset) > 0 {
		divider := cutset[0]
		cutset = cutset[1:]
		var newParts []string
		for _, oldPart := range parts {
			newParts = append(newParts, strings.Split(oldPart, divider)...)
		}
		parts = newParts
	}
	return parts
}
