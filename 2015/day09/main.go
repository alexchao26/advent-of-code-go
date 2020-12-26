package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	min, max := travelingSalesman(util.ReadFile("./input.txt"))
	fmt.Printf("Part1: %d\nPart2: %d\n", min, max)
}

func travelingSalesman(input string) (int, int) {
	graph := newGraphFromInput(input)

	min := math.MaxInt32
	max := 0
	for k := range graph {
		dfsMin, dfsMax := dfsTotalDistance(graph, k, map[string]bool{k: true})
		min = mathutil.MinInt(min, dfsMin)
		max = mathutil.MaxInt(max, dfsMax)
	}

	return min, max
}

func dfsTotalDistance(graph mapGraph, entry string, visited map[string]bool) (min, max int) {
	// if all nodes have been visited, return a zero length
	if len(visited) == len(graph) {
		return 0, 0
	}

	minDistance := math.MaxInt32
	maxDistance := 0

	for k := range graph {
		if !visited[k] {
			visited[k] = true

			weight := graph[entry][k]
			minRecurse, maxRecurse := dfsTotalDistance(graph, k, visited)
			minDistance = mathutil.MinInt(minDistance, weight+minRecurse)
			maxDistance = mathutil.MaxInt(maxDistance, weight+maxRecurse)

			// backtrack
			// delete to so length of visited is accurate
			delete(visited, k)
		}
	}

	return minDistance, maxDistance
}

type mapGraph map[string]map[string]int

func newGraphFromInput(input string) (graph mapGraph) {
	graph = make(mapGraph)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		start, end := parts[0], parts[2]
		weight := cast.ToInt(parts[4])

		// ensure nested map exists
		if graph[start] == nil {
			graph[start] = make(map[string]int)
		}
		if graph[end] == nil {
			graph[end] = make(map[string]int)
		}

		// set weight in both directions
		graph[start][end] = weight
		graph[end][start] = weight
	}
	return graph
}
