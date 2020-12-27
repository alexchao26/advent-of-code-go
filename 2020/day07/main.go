package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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
	graph := parseInput(input)
	delete(graph, "shiny gold") // shouldn't be allowed in this graph just in case

	var outers []string
	// traverse edges until a "shiny gold bag" is found as a content?
	for bag := range graph {
		if dfsFindShinyGold(graph, bag) {
			outers = append(outers, bag)
		}
	}

	return len(outers)
}

func part2(input string) int {
	graph := parseInput(input)

	// subtract one for the shiny gold bag counting itself
	return countSubbags(graph, "shiny gold") - 1
}

// not a fun input to parse through, there has to be a better way to do this in go...
func parseInput(input string) map[string]map[string]int {
	graph := map[string]map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " contain ")
		color := split[0][:strings.Index(split[0], " bags")]
		graph[color] = map[string]int{}

		for _, content := range strings.Split(split[1], ", ") {
			if content == "no other bags." {
				continue
			}
			parts := strings.Split(content, " ")
			graph[color][parts[1]+" "+parts[2]] = cast.ToInt(parts[0])
		}
	}
	return graph
}

func dfsFindShinyGold(graph map[string]map[string]int, entry string) bool {
	// shiny gold is contained within this bag, return true & collapse call stack
	if _, ok := graph[entry]["shiny gold"]; ok {
		return true
	}

	// otherwise traverse into its edges and if it returns true, send down callstack
	for subBags := range graph[entry] {
		if dfsFindShinyGold(graph, subBags) {
			return true
		}
	}

	// otherwise shiny gold was not reached, just return false
	return false
}

func countSubbags(graph map[string]map[string]int, entry string) int {
	// count itself
	bags := 1

	// traverse into its subbags and add time its count
	for subBag, count := range graph[entry] {
		bags += count * countSubbags(graph, subBag)
	}

	return bags
}
