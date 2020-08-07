package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

// large value to act as "infinite distance away" when setting up a dijkstra grid
const bigSafeInt = 1 << 30 // 2^30

func main() {
	input := util.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	grid := make([][]string, len(lines))
	for i, v := range lines {
		grid[i] = strings.Split(v, "")

		// * Uncomment to print the input
		// fmt.Println(i-2, grid[i])
	}

	dijkstra := MakeDijkstraRecursive(grid)

	for !dijkstra.handleFrontOfQueue() {
		// * Uncomment to watch the queue length grow
		// fmt.Println("  QUEUE LENGTH ", len(dijkstra.queue))
	}

	finalLayer, finalRow, finalCol := 0, dijkstra.finishCoordinates[0], dijkstra.finishCoordinates[1]

	fmt.Println("Distance to ZZ", dijkstra.layers[finalLayer].grid[finalRow][finalCol].distance)
}

// DijkstraRecursive struct stores the 2D grid of nodes and a queue of next points to check
// and a portal map to add jumps to the queue
type DijkstraRecursive struct {
	sanitizedGrid      [][]string // stores the raw 2D grid to be passed into a Layer factory
	layers             []*Layer
	queue              [][3]int
	outerPortalCoords  map[string][2]int // map portal name to coordinates on outer edge
	innerPortalCoords  map[string][2]int // map portal name to coordinates on inner edge
	mapCoordsToPortals map[[2]int]string // arrays pass by value, so this can be used as a key
	startCoordinates   [2]int            // coordinates on first layer
	finishCoordinates  [2]int            // coordinates on first layer
}

// Layer is a single layer of the 3D maze
// assuming the maze grows downwards
type Layer struct {
	grid       [][]*Node
	layerIndex int // ground level == 0, one down == 1, two down == 2, etc.
}

// Node data type is custom built for this algo, i.e. also stores if this is a portal cell
type Node struct {
	value           string
	distance        int
	portalName      string // <portalCharacters><row><col>, will be used to jump to other indexes
	jumpCoordinates [3]int // coordinates of its paired portal if applicable
}

// MakeDijkstraRecursive does just that
func MakeDijkstraRecursive(inputGrid [][]string) *DijkstraRecursive {
	dijkstra := DijkstraRecursive{
		sanitizedGrid:      make([][]string, len(inputGrid)-4), // preprocess the inputGrid to make future layer creation easier
		layers:             []*Layer{},
		queue:              make([][3]int, 1),
		outerPortalCoords:  map[string][2]int{},
		innerPortalCoords:  map[string][2]int{},
		mapCoordsToPortals: map[[2]int]string{},
	}

	for i := range dijkstra.sanitizedGrid {
		dijkstra.sanitizedGrid[i] = inputGrid[i+2][2 : len(inputGrid[0])-2]
	}

	// populate outer/innerPortalCoords, critical when generating Layers
	// populating maps of jump coordinates
	for row := 2; row < len(inputGrid)-2; row++ {
		for col := 2; col < len(inputGrid[0])-2; col++ {
			// if a hallway and portalName is not an empty string
			portalName := getPortalName(inputGrid, row, col)

			if inputGrid[row][col] == "." && portalName != "" {
				// add to map of coordinates to portal name
				dijkstra.mapCoordsToPortals[[2]int{row - 2, col - 2}] = portalName

				// add to outer or inner portal coords maps
				if onEdgeOfGrid(inputGrid, row-2, col-2) || onEdgeOfGrid(inputGrid, row+2, col+2) {
					dijkstra.outerPortalCoords[portalName] = [2]int{row - 2, col - 2}
				} else {
					dijkstra.innerPortalCoords[portalName] = [2]int{row - 2, col - 2}
				}

				// Initial and final portal detection
				if portalName == "AA" {
					dijkstra.startCoordinates = [2]int{row - 2, col - 2}
					dijkstra.queue[0] = [3]int{0, row - 2, col - 2}
				}
				if portalName == "ZZ" {
					dijkstra.finishCoordinates = [2]int{row - 2, col - 2}
				}
			}
		}
	}

	// create first layer
	dijkstra.AddLayer()

	return &dijkstra
}

// AddLayer will add a new layer to the dijkstra layers slice
// will be called as out of range layers are jumped to
func (dijkstra *DijkstraRecursive) AddLayer() (layerCount int) {
	sanitizedGrid := dijkstra.sanitizedGrid
	grid := make([][]*Node, len(sanitizedGrid))
	layerIndex := len(dijkstra.layers)

	// make copies of the outer/innerPortalMaps
	innerPortalCoords, outerPortalCoords := map[string][3]int{}, map[string][3]int{}
	// For all layers copy all outer portal coordinates except for AA and ZZ
	for key, val := range dijkstra.outerPortalCoords {
		// if jumping TO an outer portal, that means we're going DOWN a level
		// so increment the first coordinate
		outerPortalCoords[key] = [3]int{
			layerIndex + 1,
			val[0],
			val[1],
		}
	}

	// disallow jumping to an inner (lower) layer from layer0 because all outer edges are "blocked"
	if layerIndex != 0 {
		for key, val := range dijkstra.innerPortalCoords {
			innerPortalCoords[key] = [3]int{
				layerIndex - 1,
				val[0],
				val[1]}
		}
	}

	for row := 0; row < len(sanitizedGrid); row++ {
		grid[row] = make([]*Node, len(sanitizedGrid[0]))
		for col := 0; col < len(sanitizedGrid); col++ {
			switch value := sanitizedGrid[row][col]; value {
			case "#":
				grid[row][col] = &Node{"#", bigSafeInt, "", [3]int{0, 0, 0}}
			case ".":
				grid[row][col] = &Node{
					value:    ".",
					distance: bigSafeInt,
				}
				// get portal name and jump coord from maps if applicable
				portalName, found := dijkstra.mapCoordsToPortals[[2]int{row, col}]
				if found {
					// ! this may go unused
					grid[row][col].portalName = portalName

					// determine if inner or outer coordinates are the ones being jumped to
					if onEdgeOfGrid(sanitizedGrid, row, col) {
						grid[row][col].jumpCoordinates = innerPortalCoords[portalName]
					} else {
						grid[row][col].jumpCoordinates = outerPortalCoords[portalName]
					}
				}
				// set initial distance for AA cell
				if portalName == "AA" && layerIndex == 0 {
					grid[row][col].distance = 0
				}
			}
		}
	}

	dijkstra.layers = append(dijkstra.layers, &Layer{grid, layerIndex})
	return len(dijkstra.layers)
}

// dequeues a set of coordinates, enqueues any of its appropriate neighbors (including potential
// portals/jumps and adds layers if necessary)
// returns true if the queue is empty OR the ZZ portal on layer0 has been reached
func (dijkstra *DijkstraRecursive) handleFrontOfQueue() (done bool) {
	dRow := [4]int{0, 0, -1, 1}
	dCol := [4]int{-1, 1, 0, 0}

	dequeued := dijkstra.queue[0]
	layer, row, col := dequeued[0], dequeued[1], dequeued[2]
	currentNode := dijkstra.layers[layer].grid[row][col]

	// return out of the final node has been found!
	if currentNode.portalName == "ZZ" && layer == 0 {
		return true
	}

	// add layers on the same layer if they are hallways to traverse into
	currentLayersGrid := dijkstra.layers[layer].grid
	for i := 0; i < 4; i++ {
		nextRow, nextCol := row+dRow[i], col+dCol[i]
		isInbounds := nextRow >= 0 && nextRow < len(currentLayersGrid) && nextCol >= 0 && nextCol < len(currentLayersGrid[0])
		if isInbounds {
			// if the nextNode is a hallway & has not been traveled to yet
			if nextNode := currentLayersGrid[nextRow][nextCol]; nextNode != nil && nextNode.value == "." && nextNode.distance == bigSafeInt {
				// update the distance of the nextNode
				nextNode.distance = currentNode.distance + 1
				// add its coordinates to the queue, will always be on the same layer b/c this is NOT handling jumps
				dijkstra.queue = append(dijkstra.queue, [3]int{layer, nextRow, nextCol})
			}
		}
	}

	// check if a portal jump is possible!
	// also check if the jumpCoordinates are the zero value of a [3]int
	if currentNode.portalName != "" {
		// find coordinates to jump to and the node itself
		jumpLayer := currentNode.jumpCoordinates[0]
		jumpRow := currentNode.jumpCoordinates[1]
		jumpCol := currentNode.jumpCoordinates[2]

		// if jump is going to be on an out of range layer, fire off an AddLayer
		if jumpLayer == len(dijkstra.layers) {
			dijkstra.AddLayer()
		}

		// update distance of node being jumped to
		jumpNode := dijkstra.layers[jumpLayer].grid[jumpRow][jumpCol]
		jumpNode.distance = currentNode.distance + 1

		// add to queue
		dijkstra.queue = append(dijkstra.queue, currentNode.jumpCoordinates)
	}

	// dequeue, return true if queue is now empty
	dijkstra.queue = dijkstra.queue[1:]
	if len(dijkstra.queue) == 0 {
		fmt.Println("EMPTY QUEUE")
		return true
	}
	return false
}

/*************************************************************************************
*** SMALL HELPER FUNCTIONS
*************************************************************************************/

// helper function to run in 4 directions and see if any of them are a capital letter
// if that's true, then grab the portal name in that direction and return it (two char string)
func getPortalName(grid [][]string, row, col int) string {
	// NOTE I'm hard coding directions
	leftTwo := grid[row][col-2] + grid[row][col-1]
	rightTwo := grid[row][col+1] + grid[row][col+2]
	upTwo := grid[row-2][col] + grid[row-1][col]
	downTwo := grid[row+1][col] + grid[row+2][col]

	isPortalString := func(str string) bool {
		ascii1 := str[0] - 'A'
		ascii2 := str[1] - 'A'

		if ascii1 >= 0 && ascii1 < 26 && ascii2 >= 0 && ascii2 < 26 {
			return true
		}
		return false
	}

	// if both characters are capital letters
	switch {
	case isPortalString(leftTwo):
		return leftTwo
	case isPortalString(rightTwo):
		return rightTwo
	case isPortalString(upTwo):
		return upTwo
	case isPortalString(downTwo):
		return downTwo
	}

	return ""
}

func onEdgeOfGrid(grid [][]string, row, col int) bool {
	if row == 0 || col == 0 {
		return true
	}

	if row == len(grid)-1 || col == len(grid[0])-1 {
		return true
	}

	return false
}
