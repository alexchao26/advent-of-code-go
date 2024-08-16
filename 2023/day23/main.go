package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	grid := parseInput(input)

	// do not step on same tile twice, longest hike possible
	// standard backtrack?
	var startCol int
	for c := 0; c < len(grid[0]); c++ {
		if grid[0][c] == "." {
			startCol = c
			break
		}
	}

	return backtrackLongest(grid, 0, startCol, map[[2]int]bool{}, 0)
}

var slopes = map[string][2]int{
	">": {0, 1},
	"<": {0, -1},
	"v": {1, 0},
	"^": {-1, 0},
}

type node struct {
	row, col      int
	weightedEdges map[*node]int
}

func backtrackLongest(grid [][]string, row, col int, visited map[[2]int]bool, steps int) int {
	if row == len(grid)-1 && grid[row][col] == "." {
		return steps
	}

	if diff, ok := slopes[grid[row][col]]; ok {

		nextCoord := [2]int{row + diff[0], col + diff[1]}
		if visited[nextCoord] {
			return 0
		}

		visited[[2]int{row, col}] = true

		result := backtrackLongest(grid, row+diff[0], col+diff[1], visited, steps+1)

		visited[[2]int{row, col}] = false
		return result
	}

	best := 0

	for _, diff := range slopes {
		nextRow := row + diff[0]
		nextCol := col + diff[1]

		if nextRow < 0 || nextRow >= len(grid) ||
			nextCol < 0 || nextCol >= len(grid[0]) {
			continue
		}

		nextCoord := [2]int{nextRow, nextCol}

		if visited[nextCoord] {
			continue
		}

		if grid[nextRow][nextCol] != "#" {
			visited[[2]int{row, col}] = true

			result := backtrackLongest(grid, nextRow, nextCol, visited, steps+1)
			best = max(best, result)

			visited[[2]int{row, col}] = false
		}
	}

	return best
}

func part2(input string) int {
	grid := parseInput(input)

	var startCol int
	for c := 0; c < len(grid[0]); c++ {
		if grid[0][c] == "." {
			startCol = c
			break
		}
	}
	_ = startCol
	// reduce to a graph with weighted edges
	allNodes := map[[2]int]*node{}

	// just make all nodes
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == "#" {
				continue
			}
			allNodes[[2]int{r, c}] = &node{
				row:           r,
				col:           c,
				weightedEdges: map[*node]int{},
			}

		}
	}

	// connect all adjacent nodes and assign a weight of 1
	for coords, node := range allNodes {
		for _, diff := range slopes {
			nextCoord := [2]int{
				coords[0] + diff[0],
				coords[1] + diff[1],
			}

			if neighbor, ok := allNodes[nextCoord]; ok {
				node.weightedEdges[neighbor] = 1
				neighbor.weightedEdges[node] = 1
			}
		}
	}

	// reduce the graph by combining neighbors if there are exactly two
	for _, currentNode := range allNodes {
		if len(currentNode.weightedEdges) == 2 {
			twoNeighbors := []*node{}
			summedWeight := 0
			for neighborNode := range currentNode.weightedEdges {
				twoNeighbors = append(twoNeighbors, neighborNode)
				summedWeight += neighborNode.weightedEdges[currentNode]
			}

			delete(twoNeighbors[0].weightedEdges, currentNode)
			delete(twoNeighbors[1].weightedEdges, currentNode)
			twoNeighbors[0].weightedEdges[twoNeighbors[1]] = summedWeight
			twoNeighbors[1].weightedEdges[twoNeighbors[0]] = summedWeight

			// doesn't affect map iteration
			delete(allNodes, [2]int{currentNode.row, currentNode.col})
		}
	}

	// backtrack through graph again
	return backtrackThroughGraph(allNodes[[2]int{0, startCol}],
		map[*node]bool{}, 0, len(grid)-1)
}

func backtrackThroughGraph(currentNode *node, seen map[*node]bool,
	distance int, destinationRow int) int {

	// destination row is knowing that there is only one node that is on the
	// final row, so if we reach that depth we've reached the end
	if currentNode.row == destinationRow {
		return distance
	}

	best := 0
	seen[currentNode] = true

	for neighbor, weight := range currentNode.weightedEdges {
		if seen[neighbor] {
			continue
		}
		best = max(best,
			backtrackThroughGraph(neighbor, seen, distance+weight, destinationRow))
	}

	seen[currentNode] = false

	return best
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
