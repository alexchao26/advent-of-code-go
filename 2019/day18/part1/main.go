package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	input := util.ReadFile("../input.txt")
	linesSli := strings.Split(input, "\n")

	grid := make([][]string, len(linesSli)) // string might be excessive, but it will be easier to look at
	for i, line := range linesSli {
		grid[i] = strings.Split(line, "")
	}

	// get all the coordinates of keys. will be used as starting points for dijkstra's searches
	keyToCoordinates := getCoordinatesOfKeys(grid)

	// initialize Graph
	graph := MakeGraph()

	// for every key, generate a new dijkstra grid where all distances are set to MAX SAFE INT, and the key's distance is set to 0
	for key, sourceCoordinates := range keyToCoordinates {
		// initialize a dijkstra grid for each key
		dijkstraGrid := MakeDijkstraGrid(grid, sourceCoordinates)

		// step through grid until complete, handleFrontOfQueue returns true when the queue is empty...
		for !dijkstraGrid.handleFrontOfQueue() {
		}
		// small space optimization, overwrite the dijkstra queue b/c the underlying array is alive
		dijkstraGrid.queue = nil

		// update graph for all edges from `key` to all other keys
		for destinationKey, destinationCoordinates := range keyToCoordinates {
			if key != destinationKey && destinationKey != "@" {
				row, col := destinationCoordinates[0], destinationCoordinates[1]
				// pass the key being pathed FROM and the dijkstraNode (found on the grid) of the key found
				graph.AddEdges(key, dijkstraGrid.grid[row][col])
			}
		}
	}

	dfsStartTimestamp := time.Now()
	// first off (recursive, memoized) dfs method on graph that finds minimum distance
	fmt.Printf("DFS result: %v\n\nFinished in %v\n\n", graph.dfsMinmumDistance(), time.Since(dfsStartTimestamp))
}

// get a map of the coordinates of points of interests
func getCoordinatesOfKeys(grid [][]string) map[string][2]int {
	pointsOfInterest := make(map[string][2]int)
	for row, rowSli := range grid {
		for col, cell := range rowSli {
			switch {
			case cell == "@":
				pointsOfInterest["@"] = [2]int{row, col}
			// typecase cell to its ASCII value
			case int(cell[0]) >= 'a' && int(cell[0]) <= 'z':
				pointsOfInterest[cell] = [2]int{row, col}
			}
		}
	}
	return pointsOfInterest
}

/********************************************************************************************
DIJKSTRA GRID Code
*********************************************************************************************/

// DijkstraGrid stores the grid itself and also the needed queue to traverse through the grid
type DijkstraGrid struct {
	grid  [][]*dijkstraNode
	queue [][2]int // coordinates to be traversed next
}

// dijkstraNode is each cell within a DijkstraGrid.grid
type dijkstraNode struct {
	value      string          // string of the floor type (key, door, floor?, wall)
	distance   int             // distance from the source
	keysFound  map[string]bool // keys that have been run into
	keysNeeded map[string]bool // keys needed to get to this node, i.e. all doors passed thorugh
	seen       bool
}

// MakeDijkstraGrid initializes a 2D grid of dijkstraNodes with distances initialized to the max safe integer
func MakeDijkstraGrid(grid [][]string, startCoords [2]int) *DijkstraGrid {
	finalGrid := make([][]*dijkstraNode, len(grid))
	startKey := grid[startCoords[0]][startCoords[1]]
	for row, rowSli := range grid {
		// initialize the row's slice in the finalGrid
		finalGrid[row] = make([]*dijkstraNode, len(grid[0]))
		for col, cellString := range rowSli {
			finalGrid[row][col] = &dijkstraNode{
				value:      cellString,
				distance:   math.MaxInt32,                   // maximum safe integer, effectively 2^31 - 1
				keysFound:  map[string]bool{startKey: true}, // initialize with the starting key
				keysNeeded: make(map[string]bool),           // empty map for now
				seen:       false,                           // initialize as false
			}
			// remove the "@" key because it's not a key... it's just the starting point
			delete(finalGrid[row][col].keysFound, "@")
		}
	}

	// set properties of starting coordinate
	// distance = 0
	finalGrid[startCoords[0]][startCoords[1]].distance = 0
	finalGrid[startCoords[0]][startCoords[1]].seen = true
	queue := [][2]int{
		[2]int{startCoords[0], startCoords[1]},
	}
	return &DijkstraGrid{finalGrid, queue}
}

var dRow [4]int = [4]int{0, 0, -1, 1}
var dCol [4]int = [4]int{-1, 1, 0, 0}

// handleFrontOfQueue does just that, returns a true if queue is empty upon completion
func (dijk *DijkstraGrid) handleFrontOfQueue() (queueIsEmpty bool) {
	row, col := dijk.queue[0][0], dijk.queue[0][1]
	cell := dijk.grid[row][col]

	// mark as seen
	cell.seen = true

	// if key is found, add to cell details
	if ascii := int(cell.value[0]) - 'a'; ascii >= 0 && ascii < 26 {
		cell.keysFound[cell.value] = true
	}

	// if door is found, add to keysNeeded as a lowercase (easier comparison later)
	if ascii := int(cell.value[0]) - 'A'; ascii >= 0 && ascii < 26 {
		cell.keysNeeded[string('a'+ascii)] = true
	}

	// push neighbors into queue IF not already seen and not walls
	for i := 0; i < 4; i++ {
		neighborRow, neighborCol := row+dRow[i], col+dCol[i]
		neighborCell := dijk.grid[neighborRow][neighborCol]
		if !neighborCell.seen && neighborCell.value != "#" {
			// add to queue
			dijk.queue = append(dijk.queue, [2]int{neighborRow, neighborCol})

			// increment its distance by 1
			neighborCell.distance = cell.distance + 1

			// NOTE manually copying these, otherwise they'll be pointing to the same underlying map pointer
			// copy keysFound (still carrying the same keys)
			for key := range cell.keysFound {
				neighborCell.keysFound[key] = true
			}

			// copy keysNeeded as well
			for key := range cell.keysNeeded {
				neighborCell.keysNeeded[key] = true
			}
		}
	}

	// dequeue by rescoping the slice
	dijk.queue = dijk.queue[1:]

	// if queue is empty, there are no more paths to generate, return a true
	if len(dijk.queue) == 0 {
		return true
	}
	return false
}

/********************************************************************************************
GRAPH Code
*********************************************************************************************/

// Graph stores the edges between keys by storing all paths from a particular key to all other keys
// all cells in this 2D MAP will be a dijkstraNode because those contain all the needed data
//   such as distance from last key, the keysNeeded to take this path
type Graph struct {
	keysToKeys map[string]map[string]*dijkstraNode
}

// MakeGraph initializes a Graph instance and returns a pointer to it
func MakeGraph() *Graph {
	return &Graph{
		keysToKeys: make(map[string]map[string]*dijkstraNode),
	}
}

// AddEdges takes in the key being pathed FROM, and the dijkstraNode of a key pathed TO
// and adds that edge to the graph
func (graph *Graph) AddEdges(startKey string, destinationNode *dijkstraNode) {
	// initialize this "row" of the map if it does not exist yet
	if graph.keysToKeys[startKey] == nil {
		graph.keysToKeys[startKey] = make(map[string]*dijkstraNode)
	}
	// add the destination node to the graph
	keyFound := destinationNode.value
	graph.keysToKeys[startKey][keyFound] = destinationNode
}

func (graph *Graph) dfsMinmumDistance() int {
	// recursive function that dives through graph
	var traverse func(string, map[string]bool) int

	// Implement cache for traverse function
	// caches the "<entry key>:<ordered keys found>" to resulting distance
	cache := map[string]int{}

	// NOTE helper function that leverages the above cache
	// Traverse should return the minimum distance to a completion for these inputs
	// inputs are: 1. the entry point (the key to generate a path FROM)
	//             2. the keys that have been found so far
	traverse = func(entry string, keysFound map[string]bool) int {
		shortestFromThisNode := math.MaxInt32

		cacheKey := makeCacheKey(entry, keysFound, len(graph.keysToKeys)-1)

		// cache hit
		if cacheVal, found := cache[cacheKey]; found {
			return cacheVal
		}

		// base case: all keys found, no more distance to be traveled, i.e. zero
		if len(keysFound) == len(graph.keysToKeys)-1 {
			return 0
		}

		// iterate over all possible destination nodes.
		nextEdges := graph.keysToKeys[entry]
		for nextKey, node := range nextEdges {
			// only traverse if key has not been found yet AND all needed keys have been collected already
			if !keysFound[nextKey] && haveNeededKeys(keysFound, node.keysNeeded) {
				// add the nextKey
				keysFound[nextKey] = true

				// update the shortestFromThisNode if applicable
				distanceToEnd := node.distance + traverse(nextKey, keysFound)
				if distanceToEnd < shortestFromThisNode {
					shortestFromThisNode = distanceToEnd
				}

				// backtrack - remove the key
				delete(keysFound, nextKey)
			}
		}

		cache[cacheKey] = shortestFromThisNode
		return shortestFromThisNode
	}

	// fire off initially
	return traverse("@", map[string]bool{})
}

// helper function to generate a cache string key
// cache string is of the form entryKey:allKeysToFind
func makeCacheKey(entry string, keysFound map[string]bool, totalKeys int) string {
	// TODO: move this into the graph? to avoid making it 910238190382 times
	allKeys := make([]string, totalKeys)
	for i := 0; i < totalKeys; i++ {
		allKeys[i] = string(int('a') + i)
	}

	cacheKey := entry
	// generate
	for i := 0; i < totalKeys; i++ {
		char := string(int('a') + i)
		if !keysFound[char] {
			cacheKey += char
		}
	}

	return cacheKey
}

// simple helper function to make sure that all keysNeeded are in keysFound
func haveNeededKeys(keysFound, keysNeeded map[string]bool) bool {
	for key := range keysNeeded {
		if !keysFound[key] {
			return false
		}
	}
	return true
}
