package main

import (
	"adventofcode/util"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	input := util.ReadFile("../input.txt")
	linesSli := strings.Split(input, "\n")

	// generate four quadrants
	quad1, quad2, quad3, quad4 := generateQuadrants(linesSli)

	// handle an individual grid, sum answers together to print
	one := handleGrid(quad1)
	two := handleGrid(quad2)
	three := handleGrid(quad3)
	four := handleGrid(quad4)

	fmt.Printf("Sum for four robots is %v\n", one+two+three+four)
}

// Helper function to do all the things to a grid...
func handleGrid(grid [][]string) int {
	removeDoorsWithoutKeys(grid)

	keysToCoordinates := getCoordinatesOfKeys(grid)

	graph := MakeGraph()

	for key, sourceCoordinates := range keysToCoordinates {
		dijkstraGrid := MakeDijkstraGrid(grid, sourceCoordinates)

		// step through grid until complete, handleFrontOfQueue returns true when the queue is empty...
		for !dijkstraGrid.handleFrontOfQueue() {
		}
		// small space optimization, overwrite the dijkstra queue b/c the underlying array is alive
		dijkstraGrid.queue = nil

		// update graph for all edges from `key` to all other keys
		for otherKey, destinationCoordinates := range keysToCoordinates {
			if key != otherKey && otherKey != "@" {
				row, col := destinationCoordinates[0], destinationCoordinates[1]
				// pass the key being pathed FROM and the dijkstraNode (found on the grid) of the key found
				graph.AddEdges(key, dijkstraGrid.grid[row][col])
			}
		}
	}

	// then concurrently write back to the result chan
	return graph.dfsMinmumDistance()
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

/****************************************************************************************
DIJKSTRA GRID Code
*****************************************************************************************/

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

/****************************************************************************************
GRAPH Code
*****************************************************************************************/

// Graph stores the edges between keys by storing all paths from a particular key to all other keys
// all cells in this 2D MAP will be a dijkstraNode because those contain all the needed data
//   such as distance from last key, the keysNeeded to take this path
type Graph struct {
	keysToKeys    map[string]map[string]*dijkstraNode
	allKeysNeeded map[string]bool
}

// MakeGraph initializes a Graph instance and returns a pointer to it
func MakeGraph() *Graph {
	return &Graph{
		keysToKeys:    make(map[string]map[string]*dijkstraNode),
		allKeysNeeded: make(map[string]bool),
	}
}

// AddEdges takes in the key being pathed FROM, and the dijkstraNode of a key pathed TO
// and adds that edge to the graph
func (graph *Graph) AddEdges(startKey string, destinationNode *dijkstraNode) {
	// add startKey to the allKeysNeeded hashmap
	graph.allKeysNeeded[startKey] = true

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

		cacheKey := makeCacheKey(entry, keysFound, graph.allKeysNeeded)

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
func makeCacheKey(entry string, keysFound map[string]bool, allKeysNeeded map[string]bool) string {
	cacheKey := entry
	// generate sorted list of the keys that have NOT been found yet
	var needToFind []string
	for key := range allKeysNeeded {
		if !keysFound[key] {
			needToFind = append(needToFind, key)
		}
	}
	sort.Strings(needToFind)

	return cacheKey + strings.Join(needToFind, "")
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

/****************************************************************************************
INPUT SANITIZATION FUNCTIONS
*****************************************************************************************/

/* Quadrant numbers :)  mathematical because why not
	     II  |   I
	         |
	 --------+---------
	         |
      III  |  IV
					 |       */
// This function is gross but it does what it says it does... returns four quadrants
func generateQuadrants(linesSli []string) ([][]string, [][]string, [][]string, [][]string) {
	quad1 := make([][]string, len(linesSli))
	quad2 := make([][]string, len(linesSli))
	quad3 := make([][]string, len(linesSli))
	quad4 := make([][]string, len(linesSli))
	for i, line := range linesSli {
		quad1[i] = strings.Split(line, "")
		quad2[i] = strings.Split(line, "")
		quad3[i] = strings.Split(line, "")
		quad4[i] = strings.Split(line, "")
	}

	// really only need the "@" coordinates, but might as well use this function that I alreade have..
	keyToCoordinates := getCoordinatesOfKeys(quad1)

	originRow, originCol := keyToCoordinates["@"][0], keyToCoordinates["@"][1]
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			row := originRow - 1 + i
			col := originCol - 1 + j
			if (row+col)%2 == 0 {
				quad1[row][col] = "@"
				quad2[row][col] = "@"
				quad3[row][col] = "@"
				quad4[row][col] = "@"
			} else {
				quad1[row][col] = "#"
				quad2[row][col] = "#"
				quad3[row][col] = "#"
				quad4[row][col] = "#"
			}
		}
	}
	// replace origin with wall
	quad1[originRow][originCol] = "#"
	quad2[originRow][originCol] = "#"
	quad3[originRow][originCol] = "#"
	quad4[originRow][originCol] = "#"

	// rescope quadrants to point to correct scope of underlying arrays
	quad1 = quad1[:originRow+1]
	quad2 = quad2[:originRow+1]
	for i := range quad1 {
		quad1[i] = quad1[i][originCol:]
		quad2[i] = quad2[i][:originCol+1]
	}
	quad3 = quad3[originRow:]
	quad4 = quad4[originRow:]
	for i := range quad3 {
		quad3[i] = quad3[i][:originCol+1]
		quad4[i] = quad4[i][originCol:]
	}

	// return the four quads
	return quad1, quad2, quad3, quad4
}

// NOTE this is a BIG assumption that will not work on all inputs, in fact it fails on a lot of the
// note   examples... :(
// removes doors from a quadrant if the associated key is not in the quadrant
func removeDoorsWithoutKeys(quadrant [][]string) {
	// put all the keys in this quadrant in a hashmap (along with walls, floors and the origin "@")
	valuesToKeep := map[string]bool{"#": true, ".": true, "@": true}
	for _, rowSli := range quadrant {
		for _, cell := range rowSli {
			if ascii := int(cell[0] - 'a'); ascii >= 0 && ascii < 26 {
				valuesToKeep[string(ascii+'a')] = true
			}
		}
	}

	// iterate through the quadrant again, this time removing any capital letters
	// who's lower case representation is NOT in the valuesToKeep hashmap
	for row, rowSli := range quadrant {
		for col, cell := range rowSli {
			if lower := strings.ToLower(cell); !valuesToKeep[lower] {
				quadrant[row][col] = "." // replace with an empty hallway
			}
		}
	}
}
