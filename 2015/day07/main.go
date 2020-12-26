package main

import (
	"flag"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := someAssemblyRequired(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func someAssemblyRequired(input string, part int) int {
	wireToRule := map[string]string{}

	// generate graph of wires to their source rule
	for _, inst := range strings.Split(input, "\n") {
		parts := strings.Split(inst, " -> ")
		wireToRule[parts[1]] = parts[0]
	}

	aSignal := memoDFS(wireToRule, "a", map[string]int{})
	if part == 1 {
		return aSignal
	}

	// for part 2, override the value sent to wire b, then get output to a again
	wireToRule["b"] = cast.ToString(aSignal)
	return memoDFS(wireToRule, "a", map[string]int{})
}

func memoDFS(graph map[string]string, entry string, memo map[string]int) int {
	if memoVal, ok := memo[entry]; ok {
		return memoVal
	}

	// if it's a number, return the casted value
	if regexp.MustCompile("[0-9]").MatchString(entry) {
		return cast.ToInt(entry)
	}

	sourceRule := graph[entry]
	parts := strings.Split(sourceRule, " ")

	var result int
	switch {
	case len(parts) == 1:
		result = memoDFS(graph, parts[0], memo)
	case parts[0] == "NOT":
		start := memoDFS(graph, parts[1], memo)
		result = (math.MaxUint16) ^ start
	case parts[1] == "AND":
		result = memoDFS(graph, parts[0], memo) & memoDFS(graph, parts[2], memo)
	case parts[1] == "OR":
		result = memoDFS(graph, parts[0], memo) | memoDFS(graph, parts[2], memo)
	case parts[1] == "LSHIFT":
		result = memoDFS(graph, parts[0], memo) << memoDFS(graph, parts[2], memo)
	case parts[1] == "RSHIFT":
		result = memoDFS(graph, parts[0], memo) >> memoDFS(graph, parts[2], memo)
	}

	memo[entry] = result
	return result
}
