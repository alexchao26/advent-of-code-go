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

func part1(input string) string {
	graph := parseInput(input)

	allNames := map[string]bool{}
	for k := range graph {
		allNames[k] = true
	}

	// iterate through all graph edges and remove any dependants from allNames
	// map b/c that cannot be at the bottom of the stack
	for _, node := range graph {
		for _, name := range node.edges {
			delete(allNames, name)
		}
	}

	if len(allNames) != 1 {
		panic("Expected one name left, got" + cast.ToString(len(allNames)))
	}

	// have to iterate over graph to get remaining name
	var nameAtBottom string
	for name := range allNames {
		nameAtBottom = name
	}
	return nameAtBottom
}

// NOTE this is not a fully generalized solution, there are graphs where this
// could fail. Although I suspect AoC inputs are tailored to make this solution
// sufficient
func part2(input string) int {
	graph := parseInput(input)
	currentNode := part1(input)
	weightCalculator := memoCalcWeight(graph)

	var siblings []string
	// should not run for more than the number of nodes in the graph
	for indexToExit := 0; indexToExit < len(graph); indexToExit++ {
		// store dependents which have a particular weight
		weightToDependents := map[int][]string{}
		for _, dependentName := range graph[currentNode].edges {
			weight := weightCalculator(dependentName)
			weightToDependents[weight] = append(weightToDependents[weight], dependentName)
		}

		// one of the dependents has a different weight than the others, dive into it
		if len(weightToDependents) > 1 {
			// store its siblings b/c IF the next loop finds that all dependents
			// are an equal weight, we need to compare the current set of siblings
			siblings = graph[currentNode].edges

			// find the node to dive into
			for _, names := range weightToDependents {
				if len(names) == 1 {
					currentNode = names[0]
				}
			}
		} else if len(weightToDependents) == 1 {
			// if dependents all have the same weight, this node is the problem.
			// compare to its siblings to determine what its weight should be
			currentWeight := weightCalculator(currentNode)
			for _, sib := range siblings {
				if sib != currentNode {
					desiredWeight := weightCalculator(sib)
					// apply diff to current node's weight and return it
					return graph[currentNode].weight - (currentWeight - desiredWeight)
				}
			}
		} else {
			panic("unhandled case, weightToDependents == 0")
		}
	}

	panic("something's wrong in the loop...")
}

// memoized: calculate the weight, including children, of a given node name
func memoCalcWeight(graph map[string]graphNode) func(string) int {
	memo := make(map[string]int, len(graph))

	var closureFunc func(string) int
	closureFunc = func(rootName string) int {
		// return from memo if possible
		if wt, ok := memo[rootName]; ok {
			return wt
		}

		// otherwise calculate & set in memo
		sum := graph[rootName].weight
		for _, dependent := range graph[rootName].edges {
			sum += closureFunc(dependent)
		}
		memo[rootName] = sum

		return sum
	}
	return closureFunc
}

type graphNode struct {
	name   string // unused but useful for debugging
	weight int
	edges  []string
}

func parseInput(input string) map[string]graphNode {
	lines := strings.Split(input, "\n")
	graph := make(map[string]graphNode, len(lines))
	for _, l := range lines {
		parts := strings.Split(l, " -> ")

		leftParts := strings.Split(parts[0], " ")
		name := leftParts[0]
		weight := cast.ToInt(leftParts[1][1 : len(leftParts[1])-1])

		var edges []string
		if len(parts) == 2 {
			edges = strings.Split(parts[1], ", ")
			// assumption confirmed that all nodes have zero or >= 2 edges
			if len(edges) == 1 {
				panic("Assumed no nodes have exactly one dependant, but got 1 for node: " + name)
			}
		}
		graph[name] = graphNode{
			name:   name,
			weight: weight,
			edges:  edges,
		}
	}

	return graph
}
