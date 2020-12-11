package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/datastructures"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	depth, targetX, targetY := parseInput(input)

	regionTypeCalculator := memoRegionTypeCalculator(depth, targetX, targetY)
	var ans int

	for x := 0; x <= targetX; x++ {
		for y := 0; y <= targetY; y++ {
			riskLevel := int(regionTypeCalculator(x, y)) % 3
			ans += riskLevel
		}
	}

	return ans
}

func parseInput(input string) (int, int, int) {
	lines := strings.Split(input, "\n")
	var depth, targetX, targetY int
	_, err := fmt.Sscanf(lines[0], "depth: %d", &depth)
	if err != nil {
		panic("parsing depth from input" + err.Error())
	}
	_, err = fmt.Sscanf(lines[1], "target: %d,%d", &targetX, &targetY)
	if err != nil {
		panic("parsing targetX and targetY from input" + err.Error())
	}

	return depth, targetX, targetY
}

func memoErosionLevelCalculator(depth, targetX, targetY int) func(x, y int) int {
	// map to memoize results and prevent branch recursion
	xyToErosion := map[[2]int]int{}

	var closureGetErosionFunc func(x, y int) int
	closureGetErosionFunc = func(x, y int) int {
		coords := [2]int{x, y}
		if e, ok := xyToErosion[coords]; ok {
			return e
		}

		var geologicIndex int
		if coords == [2]int{0, 0} || coords == [2]int{targetX, targetY} {
			geologicIndex = 0
		} else if y == 0 {
			geologicIndex = x * 16807
		} else if x == 0 {
			geologicIndex = y * 48271
		} else {
			geologicIndex = closureGetErosionFunc(x-1, y) * closureGetErosionFunc(x, y-1)
		}
		erosionLevel := (geologicIndex + depth) % 20183

		xyToErosion[coords] = erosionLevel
		return erosionLevel
	}

	return closureGetErosionFunc
}

func memoRegionTypeCalculator(depth, targetX, targetY int) func(x, y int) regionType {
	erosionCalculator := memoErosionLevelCalculator(depth, targetX, targetY)

	xyToRegionType := map[[2]int]regionType{}
	var closureRegionFunc func(x, y int) regionType
	closureRegionFunc = func(x, y int) regionType {
		if x < 0 || y < 0 {
			return -1
		}
		coords := [2]int{x, y}
		if rt, ok := xyToRegionType[coords]; ok {
			return rt
		}
		erosion := erosionCalculator(x, y)

		rt := regionType(erosion % 3)
		xyToRegionType[coords] = rt
		return rt
	}

	return closureRegionFunc
}

func part2(input string) int {
	depth, targetX, targetY := parseInput(input)
	regionCalculator := memoRegionTypeCalculator(depth, targetX, targetY)

	heap := datastructures.NewMinHeap()
	firstNode := node{
		coords:     [2]int{0, 0},
		regionType: regionCalculator(0, 0),
		equipped:   torch,
		totalTime:  0,
	}
	heap.Add(firstNode)

	var eqCoordsToMinDist = map[equipmentType]map[[2]int]int{
		neither:      map[[2]int]int{},
		climbingGear: map[[2]int]int{},
		torch:        map[[2]int]int{},
	}

	var currentNode node
	for !(currentNode.coords[0] == targetX && currentNode.coords[1] == targetY) {
		currentNode = step(heap, eqCoordsToMinDist, regionCalculator)
	}

	if currentNode.regionType != rocky {
		panic("target must be rocky")
	}

	if currentNode.equipped != torch {
		finalTime := currentNode.totalTime + 7
		return finalTime
	}

	return currentNode.totalTime
}

type regionType int
type equipmentType int

const (
	rocky  regionType = iota // climbing gear or torch
	wet                      // neither or climbing gear
	narrow                   // neither or torch
)
const (
	neither equipmentType = iota
	climbingGear
	torch
)

type node struct {
	coords     [2]int
	equipped   equipmentType
	regionType regionType
	totalTime  int
}

func (n node) Value() int {
	return n.totalTime
}

func step(heap *datastructures.MinHeap, eqCoordsToMinDist map[equipmentType]map[[2]int]int, regionCalculator func(x, y int) regionType) node {
	// remove node from heap, this will be returned at the end
	minNodeInterface := heap.Remove()
	if minNodeInterface == nil {
		panic("Heap is empty, it shouldn't be empty...")
	}
	minNode, ok := minNodeInterface.(node)
	if !ok {
		panic("interface conversion error")
	}

	// if we've already visited these coordinates with this equipment, check
	// that the current time is
	t, ok := eqCoordsToMinDist[minNode.equipped][minNode.coords]
	if ok && t <= minNode.totalTime {
		return node{}
	}
	eqCoordsToMinDist[minNode.equipped][minNode.coords] = minNode.totalTime
	// fmt.Println("coords, region, equipped, time", minNode.coords, minNode.regionType, minNode.equipped, minNode.totalTime)
	// fmt.Println("    Steps", minNode.steps)
	// first check if movement is possible because this requires less time
	// add those steps to the heap
	nodesToTravelTo := getNextNodes(minNode, regionCalculator)
	// fmt.Println("nodes to travel to", nodesToTravelTo)
	for _, n := range nodesToTravelTo {
		heap.Add(n)
	}

	// then try to change equipment if possible
	newEq := getSwitchableEquipment(minNode)

	// find travelable-to nodes with new equipment on & append if applicable
	swappedEquipmentNode := node{
		coords:     minNode.coords,
		equipped:   newEq,
		regionType: minNode.regionType,
		totalTime:  minNode.totalTime + 7, // swapping takes 7 minutes
	}
	// in scope of for loop, just to be safe...
	nodesToTravelTo = getNextNodes(swappedEquipmentNode, regionCalculator)

	for _, n := range nodesToTravelTo {
		heap.Add(n)
	}

	return minNode
}

var directions = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

// returns a slice of nodes that represent traveling in the four directions from
// the currentNode
// includes adding 1 to the time
func getNextNodes(currentNode node, regionCalculator func(x, y int) regionType) []node {
	var nodesToTravelTo []node
	for _, d := range directions {
		nextX := currentNode.coords[0] + d[0]
		nextY := currentNode.coords[1] + d[1]
		nextCoord := [2]int{nextX, nextY}
		nextRegionType := regionCalculator(nextX, nextY)

		// ensure next region is not out of range
		if nextRegionType != -1 {
			// check if it can be traveled to
			if canTravelTo(nextRegionType, currentNode.equipped) {
				nodesToTravelTo = append(nodesToTravelTo, node{
					coords:     nextCoord,
					equipped:   currentNode.equipped,
					regionType: nextRegionType,
					totalTime:  currentNode.totalTime + 1,
				})
			}
		}
	}

	return nodesToTravelTo
}

var mapRegionsToToolsMap = map[regionType]map[equipmentType]bool{
	rocky: map[equipmentType]bool{
		climbingGear: true,
		torch:        true,
	},
	wet: map[equipmentType]bool{
		climbingGear: true,
		neither:      true,
	},
	narrow: map[equipmentType]bool{
		torch:   true,
		neither: true,
	},
}

func canTravelTo(nextRegionType regionType, equippedTool equipmentType) bool {
	return mapRegionsToToolsMap[nextRegionType][equippedTool]
}

func getSwitchableEquipment(n node) equipmentType {
	for i := 0; i < 3; i++ {
		eq := equipmentType(i)
		if eq != n.equipped && mapRegionsToToolsMap[n.regionType][eq] {
			return eq
		}
	}

	panic("must be one piece of equipment to switch to")
}
