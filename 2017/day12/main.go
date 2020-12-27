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

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	graph := makeGraphFromInput(input)

	var count int
	for k := range graph {
		if dfsCanReachTarget(graph, k, 0, map[int]bool{}) {
			count++
		}
	}

	return count
}

func part2(input string) int {
	graph := makeGraphFromInput(input)

	allKeys := []int{}
	for k := range graph {
		allKeys = append(allKeys, k)
	}

	var groupCount int
	// nodes that have been added to a group (that has been counted)
	hasBeenGrouped := map[int]bool{}

	for target := range graph {
		if !hasBeenGrouped[target] {
			// iterate through all graph nodes and check if they can be reached
			for k := range graph {
				// performance optimization: skip nodes that are already grouped
				if k != target && !hasBeenGrouped[k] {
					// if this group can reach the target, they're part of that group
					if dfsCanReachTarget(graph, k, target, map[int]bool{}) {
						hasBeenGrouped[k] = true
					}
				}
			}
			groupCount++
		}
	}

	return groupCount
}

func dfsCanReachTarget(graph map[int][]int, entry int, target int, visited map[int]bool) bool {
	// break infinite loops
	if visited[entry] {
		return false
	}
	visited[entry] = true

	for _, child := range graph[entry] {
		if child == target || dfsCanReachTarget(graph, child, target, visited) {
			return true
		}
	}
	// default to returning false
	return false
}

func makeGraphFromInput(input string) map[int][]int {
	lines := strings.Split(input, "\n")
	graph := make(map[int][]int, len(lines))
	for _, l := range lines {
		parts := strings.Split(l, " <-> ")
		ID := cast.ToInt(parts[0])
		for _, child := range strings.Split(parts[1], ", ") {
			graph[ID] = append(graph[ID], cast.ToInt(child))
		}
	}
	return graph
}
