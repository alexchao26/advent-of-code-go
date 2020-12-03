package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

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
	grid, rowBounds := parseInputs(input)

	pour(grid)

	// // Uncomment to print part of the output
	// end := len(grid[0])
	// if end > 550 {
	// 	end = 550
	// }
	// for i, v := range grid {
	// 	fmt.Println(v[450:end])
	// 	if i >= 60 {
	// 		break
	// 	}
	// }

	var ans int
	for r, row := range grid {
		for _, v := range row {
			// NOTE: used x's instead of ~'s because they're easier to see
			if (v == "|" || v == "x") && r <= rowBounds[1] && r >= rowBounds[0] {
				ans++
			}
		}
	}
	return ans
}

func part2(input string) int {
	grid, rowBounds := parseInputs(input)

	pour(grid)

	var ans int
	for r, row := range grid {
		for _, v := range row {
			if v == "x" && r <= rowBounds[1] && r >= rowBounds[0] {
				ans++
			}
		}
	}
	return ans
}

// x, y -> col, row...
func parseInputs(input string) (grid [][]string, rowBounds [2]int) {
	lines := strings.Split(input, "\n")

	verts, horis := [][3]int{}, [][3]int{}
	var largestX, largestY int
	lowestY := math.MaxInt32

	for _, l := range lines {
		var char1, char2 string
		var num1, start, end int
		fmt.Sscanf(l, "%1s=%d, %1s=%d..%d", &char1, &num1, &char2, &start, &end)

		if char1 == "x" { // vert
			verts = append(verts, [3]int{num1, start, end})
			if num1 > largestX {
				largestX = num1
			}
			if end > largestY {
				largestY = end
			}
			if start < lowestY {
				lowestY = start
			}
		} else {
			horis = append(horis, [3]int{num1, start, end})
			if num1 > largestY {
				largestY = num1
			}
			if num1 < lowestY {
				lowestY = num1
			}
			if end > largestX {
				largestX = end
			}
		}
	}

	grid = make([][]string, largestY+1)
	for i := range grid {
		grid[i] = make([]string, largestX+1)
	}

	for _, coords := range verts {
		col, start, end := coords[0], coords[1], coords[2]
		for i := start; i <= end; i++ {
			grid[i][col] = "#"
		}
	}
	for _, coords := range horis {
		row, start, end := coords[0], coords[1], coords[2]
		for i := start; i <= end; i++ {
			grid[row][i] = "#"
		}
	}

	for i, row := range grid {
		for j := range row {
			if grid[i][j] != "#" {
				grid[i][j] = "."
			}
		}
	}

	// add an empty row ot the bottom

	grid = append(grid, make([]string, len(grid[0])))
	for c := range grid[0] {
		grid[len(grid)-1][c] = "."
	}

	return grid, [2]int{lowestY, largestY}
}

func pour(grid [][]string) {
	// stack stores the coordinates that have been poured into
	// the stack is used ot backtrack to previous cells to see if they can
	// pour into additional spaces
	stack := [][2]int{{0, 500}}

	for len(stack) > 0 {
		// take coordinate at top of stack
		top := stack[len(stack)-1]
		currentVal := grid[top[0]][top[1]]

		// if it's a wall, pop it and continue
		if currentVal == "#" {
			stack = stack[:len(stack)-1]
			continue
		}

		down := getNextCoord(top, "down")

		// ensure it's in bounds, if not pop and continue
		if !isInBounds(grid, down) {
			stack = stack[:len(stack)-1]
			continue
		}

		// this will happen on the second visit to a coordinate (it has to be
		// changed into a pipe first @ the bottom of this loop)
		if currentVal == "|" {
			// transform will check if the row below (down variable) is water (pipes)
			// bound by walls (assume that there isn't a sneaky hole in the floor)
			// if it is bound by walls, it will replace all pipes with x's to
			// indicate still water
			transformStillWater(grid, down)

			// pop off stack
			stack = stack[:len(stack)-1]

			// continue with the rest of the loop to add coords on the stack
			// this handles two cases in particular
			//      |
			//      |   pouring over to the right side here
			//   #  |     | #
			//   #  |     v #
			//   #-----###..#
			//   #-----# #..#
			//   #-----###..#
			//   #----------#
			//   ############
			//
			//      |
			//      |
			//   #  |       #
			//   #  |       #
			//   #..|.......#
			//   #..|.......#
			//   #..|.......# <- filling up this row
			//   #----------#
			//   ############
		}

		// if below is a wall, append left and right to the stack
		valDown := getValAt(grid, top, "down")

		// add left and right to stack if they're sand
		if valDown == "#" || valDown == "x" {
			if getValAt(grid, top, "left") == "." {
				stack = append(stack, getNextCoord(top, "left"))
			}
			if getValAt(grid, top, "right") == "." {
				stack = append(stack, getNextCoord(top, "right"))
			}
		}
		// if down is sand, add it to stack
		if valDown == "." {
			stack = append(stack, down)
		}

		// make self a water pipe
		grid[top[0]][top[1]] = "|"
	}
}

// helper functions to make getting the next coordinate or its value or if its in bounds
func isInBounds(grid [][]string, coord [2]int) bool {
	return coord[0] < len(grid) && coord[1] < len(grid[0])
}

func getNextCoord(coord [2]int, direction string) [2]int {
	if !strings.Contains("downleftright", direction) {
		panic("invalid direction passed to getNextCoord")
	}

	switch direction {
	case "down":
		return [2]int{coord[0] + 1, coord[1]}
	case "left":
		return [2]int{coord[0], coord[1] - 1}
	case "right":
		return [2]int{coord[0], coord[1] + 1}
	}

	return [2]int{} // should never be hit...
}

func getValAt(grid [][]string, coord [2]int, direction string) string {
	nextCoord := getNextCoord(coord, direction)

	return grid[nextCoord[0]][nextCoord[1]]
}

// check a particular row to see if it can be transformed into still water
// i.e. is it all water pipes bound by walls on either end
func transformStillWater(grid [][]string, coord [2]int) {
	var left, right int
	isWalled := true
	for col := coord[1] - 1; isInBounds(grid, [2]int{coord[0], col}); col-- {
		if grid[coord[0]][col] == "#" {
			left = col + 1
			break
		}
		if grid[coord[0]][col] == "." {
			isWalled = false
			break
		}
	}
	for col := coord[1] + 1; isInBounds(grid, [2]int{coord[0], col}); col++ {
		if grid[coord[0]][col] == "#" {
			right = col - 1
			break
		}
		if grid[coord[0]][col] == "." {
			isWalled = false
			break
		}
	}

	if isWalled {
		for i := left; i <= right; i++ {
			// only transform waters, not preexisting floors
			if grid[coord[0]][i] == "|" {
				grid[coord[0]][i] = "x"
			}
		}
	}
}
