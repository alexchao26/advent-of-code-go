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

	return calcEnergizedTiles(grid, beamHead{
		velocity: right,
		row:      0,
		col:      0,
	})
}

func part2(input string) int {
	grid := parseInput(input)
	best := 0
	for r := range len(grid) {
		// starting from left of grid, headed right
		startBeam := beamHead{
			velocity: right,
			row:      r,
			col:      0,
		}
		// max is a go 1.21 addition: https://pkg.go.dev/builtin#max
		best = max(best, calcEnergizedTiles(grid, startBeam))

		// starting from right of grid, headed left
		startBeam = beamHead{
			velocity: left,
			row:      r,
			col:      len(grid[0]) - 1,
		}
		best = max(best, calcEnergizedTiles(grid, startBeam))
	}

	for c := range len(grid[0]) {
		// starting from top of grid, headed down
		startBeam := beamHead{
			velocity: down,
			row:      0,
			col:      c,
		}
		best = max(best, calcEnergizedTiles(grid, startBeam))

		// starting from bottom of grid, headed up
		startBeam = beamHead{
			velocity: up,
			row:      len(grid) - 1,
			col:      c,
		}
		best = max(best, calcEnergizedTiles(grid, startBeam))
	}

	return best
}

func calcEnergizedTiles(grid [][]string, beam beamHead) int {
	// need to track multiple beams because they can split and generate multiple
	beams := []beamHead{beam}

	// [4]bool represents being hit from left, right, up and down respectively
	// need to track direction so that we can terminate beams that are cyclical
	hitGrid := [][][4]bool{}
	for range grid {
		hitGrid = append(hitGrid, make([][4]bool, len(grid[0])))
	}

	for len(beams) > 0 {
		b := beams[0]
		beams = beams[1:]

		skip := false
		for vel, hitGridIndex := range velToHitGridIndex {
			if b.velocity == vel && hitGrid[b.row][b.col][hitGridIndex] {
				skip = true
			}
		}
		if skip {
			continue
		}

		// record direction hit in hit grid
		hitGrid[b.row][b.col][velToHitGridIndex[b.velocity]] = true

		cell := grid[b.row][b.col]
		nextVelocities, ok := mirrorToNextVelocities[cell][b.velocity]
		if !ok {
			panic("no nextVelocities found for cell type: " + cell)
		}
		for _, nextVelocity := range nextVelocities {
			nextRow := b.row + nextVelocity[0]
			nextCol := b.col + nextVelocity[1]

			if nextRow < 0 || nextRow >= len(grid) ||
				nextCol < 0 || nextCol >= len(grid[0]) {
				continue
			}

			beams = append(beams, beamHead{
				velocity: nextVelocity,
				row:      nextRow,
				col:      nextCol,
			})
		}
	}

	energizedTiles := 0
	for r := range len(hitGrid) {
		for c := range len(hitGrid[0]) {
			for _, dir := range hitGrid[r][c] {
				if dir {
					energizedTiles++
					break
				}
			}
		}
	}
	return energizedTiles
}

var left = [2]int{0, -1}
var right = [2]int{0, 1}
var up = [2]int{-1, 0}
var down = [2]int{1, 0}

var velToHitGridIndex = map[[2]int]int{
	left:  0,
	right: 1,
	up:    2,
	down:  3,
}

var mirrorToNextVelocities = map[string]map[[2]int][][2]int{
	".": {
		left:  {left},
		right: {right},
		up:    {up},
		down:  {down},
	},
	"/": {
		left:  {down},
		down:  {left},
		up:    {right},
		right: {up},
	},
	"\\": {
		left:  {up},
		up:    {left},
		down:  {right},
		right: {down},
	},
	"|": {
		left:  {up, down},
		right: {up, down},
		up:    {up},
		down:  {down},
	},
	"-": {
		left:  {left},
		right: {right},
		up:    {left, right},
		down:  {left, right},
	},
}

type beamHead struct {
	velocity [2]int
	row, col int
}

func (b beamHead) String() string {
	return fmt.Sprintf("vel: %v, coords: %d, %d", b.velocity, b.row, b.col)
}

// for debugging
func condenseHitGrid(grid [][][4]bool) [][]int {
	ans := [][]int{}
	for range grid {
		ans = append(ans, make([]int, len(grid[0])))
	}

	for r, row := range grid {
		for c, sli := range row {
			for _, val := range sli {
				if val {
					ans[r][c]++
				}
			}
		}
	}

	return ans
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
