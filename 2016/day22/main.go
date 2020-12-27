package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	nodes := parseInput(input)

	var viable int
	for i1, n1 := range nodes {
		for i2, n2 := range nodes {
			if i1 == i2 || n1.used == 0 {
				continue
			}
			if n2.avail >= n1.used {
				viable++
			}
		}
	}

	return viable
}

// NOTE: this is not a generalized solution, this was done after solving it somewhat
// manually by printing the entire grid and getting to the top right corner
// there was a blocking row that had really large memory usage, these had to be routed around
// then swapping the top right tile each step to the left, required 5 steps
//    4 to get in front of it and 1 to do the actual swap
func part2(input string) int {
	nodes := parseInput(input)

	var maxX, maxY int
	var x, y int
	for c, n := range nodes {
		maxX = mathy.MaxInt(c[0], maxX)
		maxY = mathy.MaxInt(c[1], maxY)
		// getting the starting node, i.e. has zero used space
		if n.used == 0 {
			x = n.coord[1]
			y = n.coord[0]
		}
	}

	// // uncomment to print a useable grid
	// grid := make([][]*node, maxY+1)
	// for i := range grid {
	// 	grid[i] = make([]*node, maxX+1)
	// }
	// for c, n := range nodes {
	// 	grid[c[1]][c[0]] = n
	// }
	// for _, line := range grid {
	// 	for _, n := range line {
	// 		fmt.Print(n)
	// 	}
	// 	fmt.Println()
	// }

	var stepsTaken int
	// get to the cell that's needed
	for !(x == maxX && y == 0) {
		if y > 0 {
			if nodes[[2]int{x, y - 1}].used < 100 { // realizing that x/y are "flipped"..
				// grid[y-1][x].used < 100 {
				y--
			} else {
				// go left to get around the "blocking" chips who's used size
				// is so large that it cannot be copied into the zero chip
				x--
			}
		} else if y == 0 {
			x++
		}
		stepsTaken++
	}
	// decrement x because we shifted into it already
	x--

	// then you need five steps to move the target cell to the left by one cell
	for x != 0 {
		stepsTaken += 5
		x--
	}

	return stepsTaken
}

type node struct {
	coord [2]int
	size  int
	used  int
	avail int
}

func (n node) String() string {
	// str := fmt.Sprintf("%v: %d used of %d, %d avail", n.coord, n.used, n.size, n.avail)
	str := fmt.Sprintf("| %d/%d ", n.used, n.size)
	for len(str) < 10 {
		str += " "
	}
	return str
}

func parseInput(input string) map[[2]int]*node {
	allNodes := map[[2]int]*node{}

	spaces := regexp.MustCompile("[ ]{2,}")
	for _, line := range strings.Split(input, "\n")[2:] {
		str := spaces.ReplaceAllString(line, " ")
		var percentage int
		n := node{}
		fmt.Sscanf(str, "/dev/grid/node-x%d-y%d %dT %dT %dT %d%",
			&n.coord[0], &n.coord[1], &n.size, &n.used, &n.avail, &percentage)
		allNodes[n.coord] = &n
	}

	return allNodes
}
