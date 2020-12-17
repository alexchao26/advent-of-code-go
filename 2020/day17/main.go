package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

// finished 933/~700 - jeez...
func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	nodes := parseInput3D(input)

	for cycles := 0; cycles < 6; cycles++ {
		toCheck := map[[3]int]bool{}
		for _, node := range nodes {
			for _, dir := range directions {
				x, y, z := node.x+dir[0], node.y+dir[1], node.z+dir[2]
				toCheck[[3]int{x, y, z}] = true
			}
		}

		nextState := map[[3]int]*node3D{}
		for coord := range toCheck {
			// check all neighbors around this coord
			var countNeighbors int
			for _, d := range directions {
				x, y, z := coord[0]+d[0], coord[1]+d[1], coord[2]+d[2]
				neighCoord := [3]int{x, y, z}
				if neigh, ok := nodes[neighCoord]; ok {
					if neigh.active {
						countNeighbors++
					}
				}
			}

			stateInNext := node3D{
				x:      coord[0],
				y:      coord[1],
				z:      coord[2],
				active: false,
			}
			if n, ok := nodes[coord]; ok && n.active {
				if countNeighbors == 2 || countNeighbors == 3 {
					stateInNext.active = true
				}
			} else {
				// inactive originally
				if countNeighbors == 3 {
					stateInNext.active = true
				}
			}
			nextState[coord] = &stateInNext
		}
		nodes = nextState

	}

	var count int
	for _, node := range nodes {
		if node.active {
			count++
		}
	}
	// cubes after 6 cycles
	return count
}

func generate3DDirections() [][3]int {
	directions := [][3]int{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				if !(i == 0 && j == 0 && k == 0) {
					directions = append(directions, [3]int{i, j, k})
				}
			}
		}
	}
	return directions
}

var directions = generate3DDirections()

type node3D struct {
	x, y, z int
	active  bool
}

func parseInput3D(input string) map[[3]int]*node3D {
	nodes := map[[3]int]*node3D{}
	lines := strings.Split(input, "\n")
	for i, l := range lines {
		for j, cell := range strings.Split(l, "") {
			n := &node3D{
				x: i, y: j, z: 0, active: false,
			}
			if cell == "#" {
				n.active = true
			}
			nodes[[3]int{i, j, 0}] = n
		}
	}
	return nodes
}

func part2(input string) int {
	nodes := parseInput4D(input)

	for cycles := 0; cycles < 6; cycles++ {
		toCheck := map[[4]int]bool{}
		for _, node := range nodes {
			for _, dir := range directions4D {
				x, y, z, w := node.x+dir[0], node.y+dir[1], node.z+dir[2], node.w+dir[3]
				toCheck[[4]int{x, y, z, w}] = true
			}
		}

		nextState := map[[4]int]*node4D{}
		for coord := range toCheck {
			// check all neighbors around this coord
			var countNeighbors int
			for _, d := range directions4D {
				x, y, z, w := coord[0]+d[0], coord[1]+d[1], coord[2]+d[2], coord[3]+d[3]
				neighCoord := [4]int{x, y, z, w}
				if neigh, ok := nodes[neighCoord]; ok {
					if neigh.active {
						countNeighbors++
					}
				}
			}

			stateInNext := node4D{
				x:      coord[0],
				y:      coord[1],
				z:      coord[2],
				w:      coord[3],
				active: false,
			}
			if n, ok := nodes[coord]; ok && n.active {
				if countNeighbors == 2 || countNeighbors == 3 {
					stateInNext.active = true
				}
			} else {
				// inactive originally
				if countNeighbors == 3 {
					stateInNext.active = true
				}
			}
			nextState[coord] = &stateInNext

		}
		nodes = nextState

	}

	var count int
	for _, node := range nodes {
		if node.active {
			count++
		}
	}
	// cubes after 6 cycles
	return count
}

type node4D struct {
	x, y, z, w int
	active     bool
}

func parseInput4D(input string) map[[4]int]*node4D {
	nodes := map[[4]int]*node4D{}
	lines := strings.Split(input, "\n")
	for i, l := range lines {
		for j, cell := range strings.Split(l, "") {
			n := &node4D{
				x: i, y: j, z: 0, w: 0, active: false,
			}
			if cell == "#" {
				n.active = true
			}
			nodes[[4]int{i, j, 0, 0}] = n
		}
	}
	return nodes
}

func generate4DDirections() [][4]int {
	directions := [][4]int{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			for k := -1; k < 2; k++ {
				for w := -1; w < 2; w++ {
					if !(i == 0 && j == 0 && k == 0 && w == 0) {
						directions = append(directions, [4]int{i, j, k, w})
					}
				}
			}
		}
	}
	return directions
}

var directions4D = generate4DDirections()

func makeDirections(length int) [][]int {
	perms := [][]int{
		make([]int, length),
	}

	for i := 0; i < length; i++ {
		for _, p := range perms {
			copy1, copy2 := make([]int, length), make([]int, length)
			copy(copy1, p)
			copy(copy2, p)
			copy1[i] = -1
			copy2[i] = 1
			perms = append(perms, copy1, copy2)
		}
	}

	return perms[1:]
}

func getStringKey(slice []int) string {
	var key string
	for _, v := range slice {
		key += fmt.Sprintf("%d-", v)
	}
	return key
}
