package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand"
	"strings"

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

func part1(input string) int {
	graph := parseInput(input)

	// if brute forcing, can check if the separated nodes are able to traverse to each other
	// which would indicate that it has broken an internal (to a group) edge, and not separated
	// the overall group...

	// brute forcing 3k edges for my input with 3 traversals for each...
	// for n = total edges
	// O(n^4), triple nested for loop with traversal check in each
	// O(3k^4) = 81,000,000,000,000 way too slow

	// strategy actually used:
	// pick 200 random pairs of nodes, traverse between the two of them
	// the most traversed nodes will likely be the bridges
	// a similar strategy would be to BFS traverse from a few randomly selected nodes to the furthest
	// possible node, the furthest depth away is likely in the other group (assuming a kind input)
	// and therefore those paths can be used to tabulate the most trafficked edges aka 3 bridges
	var allNodes []string
	for n := range graph {
		allNodes = append(allNodes, n)
	}

	// for a small (example) graph just pick every node
	// for a large (actual input) graph, pick 200 nodes
	pairsToPick := min(200, len(graph)*(len(graph)-1))

	traversedPairs := map[string]bool{}
	timesEdgeTraversed := map[string]int{}

	for len(traversedPairs) < pairsToPick {
		i1 := rand.Intn(len(allNodes))
		i2 := rand.Intn(len(allNodes))
		if i1 == i2 {
			continue
		}

		n1 := allNodes[i1]
		n2 := allNodes[i2]

		randomPairName := sortedEdgeName(n1, n2)

		if traversedPairs[randomPairName] {
			continue
		}
		traversedPairs[randomPairName] = true

		path := findShortestPath(graph, n1, n2)

		for i := 1; i < len(path); i++ {
			timesEdgeTraversed[sortedEdgeName(path[i-1], path[i])]++
		}
	}

	threeMostTraffickedEdges := getThreeMostTraffickedEdges(timesEdgeTraversed)

	// remove edge
	for _, edge := range threeMostTraffickedEdges {
		nodes := strings.Split(edge, " ")

		graph[nodes[0]] = removeElementFromSlice(graph[nodes[0]], nodes[1])
		graph[nodes[1]] = removeElementFromSlice(graph[nodes[1]], nodes[0])
	}

	sizes := []int{}
	for _, node := range strings.Split(threeMostTraffickedEdges[0], " ") {
		sizes = append(sizes, getGroupSize(graph, node, map[string]bool{}))
	}

	if len(sizes) != 2 {
		panic("expected two groups")
	}

	return sizes[0] * sizes[1]
}

func sortedEdgeName(node1, node2 string) string {
	lower := min(node1, node2)
	higher := max(node1, node2)
	return fmt.Sprintf("%v %v", lower, higher)
}

func findShortestPath(graph map[string][]string, start, end string) []string {
	type dfsNode struct {
		current   string
		pathSoFar []string
	}
	seen := map[string]bool{}
	queue := []dfsNode{
		{
			current:   start,
			pathSoFar: []string{start},
		},
	}

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		if seen[popped.current] {
			continue
		}
		seen[popped.current] = true

		if popped.current == end {
			return popped.pathSoFar
		}

		for _, neighbor := range graph[popped.current] {
			if seen[neighbor] {
				continue
			}
			nextNode := dfsNode{
				current:   neighbor,
				pathSoFar: append([]string{}, popped.pathSoFar...), // deep copy
			}
			nextNode.pathSoFar = append(nextNode.pathSoFar, neighbor)

			queue = append(queue, nextNode)
		}
	}

	panic("expect return from loop")
}

func getThreeMostTraffickedEdges(timesEdgeTraversed map[string]int) []string {
	ans := []string{}

	for len(ans) < 3 {
		var bestEdge string
		var bestCount int

		for edge, count := range timesEdgeTraversed {
			if count > bestCount {
				bestCount = count
				bestEdge = edge
			}
		}

		ans = append(ans, bestEdge)
		delete(timesEdgeTraversed, bestEdge)
	}

	return ans
}

func removeElementFromSlice(sli []string, ele string) []string {
	for i, n := range sli {
		if n == ele {
			sli[len(sli)-1], sli[i] = sli[i], sli[len(sli)-1]
			sli = sli[:len(sli)-1]

			return sli
		}
	}
	panic("element not found")
}

func getGroupSize(graph map[string][]string, node string, seen map[string]bool) int {
	if seen[node] {
		return 0
	}
	size := 1
	seen[node] = true
	for _, neighbor := range graph[node] {
		size += getGroupSize(graph, neighbor, seen)
	}
	return size
}

func part2(input string) string {
	return "happiness"
}

func parseInput(input string) (graph map[string][]string) {
	graph = map[string][]string{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		for _, node := range strings.Split(parts[1], " ") {
			graph[parts[0]] = append(graph[parts[0]], node)
			graph[node] = append(graph[node], parts[0])
		}
	}

	return graph
}
