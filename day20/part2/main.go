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
		fmt.Println(grid[i])
	}

	dijkstra := MakeDijkstraGrid(grid)
	fmt.Println("initial dijkstra queue", dijkstra.queue)

}

// Dijkstra struct stores the 2D grid of nodes and a queue of next points to check
// and a portal map to add jumps to the queue
type Dijkstra struct {
	rawInput          [][]string // stores the raw 2D grid to be passed into a Layer factory
	layers            []*Layer
	queue             [][3]int
	portalMap         map[string][2]int // map to pass to Layer factory
	startCoordinates  [2]int            // coordinates on first layer
	finishCoordinates [2]int            // coordinates on first layer
}

// Layer is a single layer of the 3D maze
// assuming the maze grows downwards
type Layer struct {
	grid       [][]*Node
	layerIndex int               // ground level == 0, one down == 1, two down == 2, etc.
	portalMap  map[string][3]int // jumpLayer, jumpRow, jumpCol
}

// Node data type is custom built for this algo, i.e. also stores if this is a portal cell
type Node struct {
	value           string
	distance        int
	portalName      string // <portalCharacters><row><col>, will be used to jump to other indexes
	jumpCoordinates [2]int // coordinates of its paired portal (if applicable)
}

// MakeDijkstraGrid does just that
func MakeDijkstraGrid(inputGrid [][]string) *Dijkstra {
	dijkstra := Dijkstra{
		rawInput:          inputGrid,
		layers:            []*Layer{},
		queue:             [][3]int{},
		portalMap:         map[string][2]int{},
		startCoordinates:  [2]int{},
		finishCoordinates: [2]int{},
	}
	portalMapHelper := make(map[string][2]int)

	grid := make([][]*Node, len(inputGrid)-4)
	// iterate starting at 2,2 to skip the top and left and end len-2 to skip bottom & right
	for row := 2; row < len(inputGrid)-2; row++ {
		grid[row-2] = make([]*Node, len(inputGrid[0])-4)
		for col := 2; col < len(inputGrid[0])-2; col++ {
			// make a node for each cell
			switch value := inputGrid[row][col]; value {
			case "#": // wall
				grid[row-2][col-2] = &Node{"#", bigSafeInt, "", [2]int{0, 0}}
				// if this is a hallway node, use a helper function to determine if there this is a portal
			case ".": // hallway
				hallwayNode := &Node{
					value:           ".",
					distance:        bigSafeInt,
					portalName:      "",
					jumpCoordinates: [2]int{0, 0},
				}
				portalName := getPortalName(inputGrid, row, col)
				if len(portalName) != 0 {
					// assign portal name for this node
					hallwayNode.portalName = portalName

					// generatine the portal maps for each node is a pain...
					// if this is portal's pair hasn't been found yet (i.e. equal to zero value of [2]int), add it to a map
					if pairedPortal := portalMapHelper[portalName]; pairedPortal == [2]int{0, 0} {
						portalMapHelper[portalName] = [2]int{row - 2, col - 2}
					} else {
						// else it has been found, set the jumpCoordinates on this node to pair's coords
						hallwayNode.jumpCoordinates = pairedPortal
						// set its pair's jumpCoordinates to this node's coords
						grid[pairedPortal[0]][pairedPortal[1]].jumpCoordinates = [2]int{row - 2, col - 2}
					}
				}
				grid[row-2][col-2] = hallwayNode
				// if it is AA, update the distance of this node to zero, initialize queue
				if portalName == "AA" {
					// !! unused
					dijkstra.startCoordinates = [2]int{row - 2, col - 2}

					hallwayNode.distance = 0
					dijkstra.queue = [][2]int{
						[2]int{row - 2, col - 2},
					}
				}
				// if end portal, set finish coordinates
				if portalName == "ZZ" {
					dijkstra.finishCoordinates = [2]int{row - 2, col - 2}
				}
			}
		}
	}

	// set grid field
	dijkstra.grid = grid

	return &dijkstra
}

// AddLayer will add a new layer to the dijkstra layers slice
//
func (dijkstra *Dijkstra) AddLayer() (layerCount int) {
	return len(dijkstra.layers)
}

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

// returns true if the queue is empty OR the ZZ portal has been reached
func (dijkstra *Dijkstra) handleFrontOfQueue() (done bool) {
	// dRow := [4]int{0, 0, -1, 1}
	// dCol := [4]int{-1, 1, 0, 0}

	// row, col := dijkstra.queue[0][0], dijkstra.queue[0][1]
	// currentNode := dijkstra.grid[row][col]

	// if currentNode.portalName == "ZZ" {
	// 	return true
	// }

	// for i := 0; i < 4; i++ {
	// 	nextRow, nextCol := row+dRow[i], col+dCol[i]
	// 	isInbounds := nextRow >= 0 && nextRow < len(dijkstra.grid) && nextCol >= 0 && nextCol < len(dijkstra.grid[0])
	// 	if isInbounds {
	// 		// if the nextNode is a hallway & has not been traveled to yet
	// 		if nextNode := dijkstra.grid[nextRow][nextCol]; nextNode != nil && nextNode.value == "." && nextNode.distance == bigSafeInt {
	// 			// update the distance of the nextNode
	// 			nextNode.distance = currentNode.distance + 1
	// 			// add its coordinates to the queue
	// 			dijkstra.queue = append(dijkstra.queue, [2]int{nextRow, nextCol})
	// 		}
	// 	}
	// }

	// // check if a portal jump is possible!
	// if currentNode.portalName != "" {
	// 	// find coordinates to jump to and the node itself
	// 	jumpRow := currentNode.jumpCoordinates[0]
	// 	jumpCol := currentNode.jumpCoordinates[1]
	// 	jumpNode := dijkstra.grid[jumpRow][jumpCol]

	// 	// update distance
	// 	jumpNode.distance = currentNode.distance + 1
	// 	// add to queue
	// 	dijkstra.queue = append(dijkstra.queue, currentNode.jumpCoordinates)
	// }

	// // dequeue, return true if queue is now empty
	// dijkstra.queue = dijkstra.queue[1:]
	// if len(dijkstra.queue) == 0 {
	// 	return true
	// }
	// return false
}
