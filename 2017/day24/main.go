package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := magneticMoat(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func magneticMoat(input string, part int) int {
	edges := getEdges(input)

	// start the bridge with a zero
	bridge := [][2]int{
		{0, 0},
	}

	usedEdges := map[[2]int]bool{}
	for _, edge := range edges {
		usedEdges[edge] = false
	}

	bestStrength, longestBridge := backtrackBridge(bridge, usedEdges)
	if part == 1 {
		return bestStrength
	}
	return calcStrengthOfBridge(longestBridge)
}

// backtracking algo that returns strongest bridge
func backtrackBridge(bridge [][2]int, usedEdges map[[2]int]bool) (bestStrength int, longestBridge [][2]int) {
	lastVal := bridge[len(bridge)-1][1]
	for edge, isUsed := range usedEdges {
		// skip edges that are marked as used
		if !isUsed {
			clonedEdge := edge
			// swap the edge vals if the first doesn't match
			if clonedEdge[0] != lastVal {
				clonedEdge[0], clonedEdge[1] = clonedEdge[1], clonedEdge[0]
			}
			// if match is found, add it to bridge, mark as used
			// add to strength, recurse
			if clonedEdge[0] == lastVal {
				bridge = append(bridge, clonedEdge)
				usedEdges[edge] = true
				strength := clonedEdge[0] + clonedEdge[1]

				// recurse and bestStrength and longestLength
				subStrength, subLongestBridge := backtrackBridge(bridge, usedEdges)

				strength += subStrength

				// if current bridge is longest (or wins tiebreak) set the longestBridge
				if len(bridge) > len(longestBridge) ||
					(len(bridge) == len(longestBridge) &&
						calcStrengthOfBridge(bridge) > calcStrengthOfBridge(longestBridge)) {
					// use this hacky append to create a copy of the bridge slice
					// otherwise appends could modify the underlying array
					longestBridge = append([][2]int{}, bridge...)
				}
				// also check if a recursive call had the longest bridge, update longest
				if len(subLongestBridge) > len(longestBridge) ||
					(len(subLongestBridge) == len(longestBridge) &&
						calcStrengthOfBridge(subLongestBridge) > calcStrengthOfBridge(longestBridge)) {
					longestBridge = append([][2]int{}, subLongestBridge...)
				}

				// backtrack
				usedEdges[edge] = false
				bridge = bridge[:len(bridge)-1]
				if strength > bestStrength {
					bestStrength = strength
				}
			}
		}
	}

	return bestStrength, longestBridge
}

func calcStrengthOfBridge(bridge [][2]int) int {
	var sum int
	for _, edge := range bridge {
		sum += edge[0] + edge[1]
	}
	return sum
}

func getEdges(input string) (edges [][2]int) {
	for _, line := range strings.Split(input, "\n") {
		var pair [2]int
		fmt.Sscanf(line, "%d/%d", &pair[0], &pair[1])
		edges = append(edges, pair)
	}
	return edges
}
