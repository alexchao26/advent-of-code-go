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
		ans := blizzardJourney(input, 1)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := blizzardJourney(input, 2)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func blizzardJourney(input string, part int) int {
	start, end, blizzards, totalRows, totalCols := parseInput(input)

	// some form of BFS seems to be the obvious answer for "least steps"
	//
	// if maps are used for all blizzard coords... then the map will have to be
	// constantly copied for new states
	// want something... less involved
	//
	// could represent all blizzards fairly easily via maths
	// < and > stay in same row, col = (starting +/- time) % (# of cols)
	// then it's just math for "is there a blizzard here" but that does require
	// iterating through every blizzard, but that's a known amount so maybe that's fine... aka close enough to constant
	//
	// or just calculate state at time t and store that matrix of "occupied" spaces, and then no need to recalculate
	// per blizzard per time

	stepsForFirstLeg := bfs(blizzards, start, end, totalRows, totalCols, 0)
	if part == 1 {
		return stepsForFirstLeg
	}
	stepsBackToStart := bfs(blizzards, end, start, totalRows, totalCols, stepsForFirstLeg)
	return bfs(blizzards, start, end, totalRows, totalCols, stepsBackToStart)
}

func bfs(blizzards []blizzard, start, end [2]int, totalRows, totalCols, stepsElapsedAlready int) int {
	cacheRoomStates := map[int][][]string{}

	type node struct {
		coords    [2]int
		steps     int
		debugPath string
	}

	queue := []node{}
	queue = append(queue, node{
		coords:    start,
		steps:     stepsElapsedAlready,
		debugPath: fmt.Sprint(0, start),
	})

	seenCoordsSteps := map[[3]int]bool{}
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]

		roomState := calcOrGetRoomState(blizzards, popped.steps+1, totalRows, totalCols, cacheRoomStates)

		for _, diff := range [][2]int{
			{1, 0},
			{0, 1},
			{0, -1},
			{-1, 0},
		} {
			nextCoords := [2]int{
				popped.coords[0] + diff[0],
				popped.coords[1] + diff[1],
			}

			if nextCoords == start {
				continue
			}
			if nextCoords != start && nextCoords != end {
				if nextCoords[0] < 0 || nextCoords[0] >= totalRows ||
					nextCoords[1] < 0 || nextCoords[1] >= totalCols {
					continue
				}
			}

			// no point in processing a coordinate & steps pair that has already been seen
			hash := [3]int{nextCoords[0], nextCoords[1], popped.steps + 1}
			if seenCoordsSteps[hash] {
				continue
			}
			seenCoordsSteps[hash] = true

			// because of how i indexed the room, need to do literal checks to see if we're in start
			// or end coords

			// if blocked, continue
			if nextCoords != start && nextCoords != end &&
				roomState[nextCoords[0]][nextCoords[1]] != "." {
				continue
			}

			// if out of bounds, continue
			if nextCoords != start && nextCoords != end {
				if nextCoords[0] < 0 || nextCoords[0] >= totalRows ||
					nextCoords[1] < 0 || nextCoords[1] >= totalCols {
					continue
				}
			}

			// done
			if nextCoords == end {
				return popped.steps + 1
			}

			queue = append(queue, node{
				coords:    nextCoords,
				steps:     popped.steps + 1,
				debugPath: popped.debugPath + fmt.Sprint(popped.steps+1, nextCoords),
			})
		}
		// if possible to stay still, add "wait" move
		if popped.coords == start ||
			roomState[popped.coords[0]][popped.coords[1]] == "." {
			queue = append(queue, node{
				coords:    popped.coords,
				steps:     popped.steps + 1,
				debugPath: popped.debugPath + fmt.Sprint(popped.steps+1, popped.coords),
			})
		}
	}

	panic("should return from loop")
}

type blizzard struct {
	startRow, startCol   int
	rowSlope, colSlope   int
	totalRows, totalCols int
	char                 string
}

func (b blizzard) calculateCoords(steps int) [2]int {
	row := (b.startRow + b.rowSlope*steps) % b.totalRows
	col := (b.startCol + b.colSlope*steps) % b.totalCols

	row += b.totalRows
	col += b.totalCols
	row %= b.totalRows
	col %= b.totalCols

	return [2]int{
		row, col,
	}
}

// occupied coordinates are easy to calculate based on each blizzard's movement
// and steps/time elapsed, return a matrix that represents occupied cells
// and store the result in a map to reduce future calcs
func calcOrGetRoomState(blizzards []blizzard, steps, totalRows, totalCols int, memo map[int][][]string) [][]string {
	if m, ok := memo[steps]; ok {
		return m
	}

	matrix := make([][]string, totalRows)
	for r := range matrix {
		matrix[r] = make([]string, totalCols)
	}

	for _, b := range blizzards {
		coords := b.calculateCoords(steps)
		matrix[coords[0]][coords[1]] = b.char
	}
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if matrix[r][c] == "" {
				matrix[r][c] = "."
			}
		}
	}

	memo[steps] = matrix

	return matrix
}

func parseInput(input string) ([2]int, [2]int, []blizzard, int, int) {
	var start, end [2]int
	blizzards := []blizzard{}

	lines := strings.Split(input, "\n")

	for c := 0; c < len(lines); c++ {
		if lines[0][c] == '.' {
			start = [2]int{-1, c - 1}
			break
		}
	}

	// 0,0 will be within the BOX we start in
	// start and end will be off the bounds of that box
	totalRows := len(lines) - 2
	totalCols := len(lines[0]) - 2

	for c := 0; c < len(lines[0]); c++ {
		if lines[len(lines)-1][c] == '.' {
			end = [2]int{totalRows, c - 1}
			break
		}
	}

	for l := 1; l < len(lines)-1; l++ {
		chars := strings.Split(lines[l], "")
		for c := 1; c < len(chars)-1; c++ {
			switch chars[c] {
			case ">":
				blizzards = append(blizzards, blizzard{
					startRow:  l - 1,
					startCol:  c - 1,
					rowSlope:  0,
					colSlope:  1,
					totalRows: totalRows,
					totalCols: totalCols,
					char:      ">",
				})
			case "<":
				blizzards = append(blizzards, blizzard{
					startRow:  l - 1,
					startCol:  c - 1,
					rowSlope:  0,
					colSlope:  -1,
					totalRows: totalRows,
					totalCols: totalCols,
					char:      "<",
				})
			case "^":
				blizzards = append(blizzards, blizzard{
					startRow:  l - 1,
					startCol:  c - 1,
					rowSlope:  -1,
					colSlope:  0,
					totalRows: totalRows,
					totalCols: totalCols,
					char:      "^",
				})
			case "v":
				blizzards = append(blizzards, blizzard{
					startRow:  l - 1,
					startCol:  c - 1,
					rowSlope:  1,
					colSlope:  0,
					totalRows: totalRows,
					totalCols: totalCols,
					char:      "v",
				})
			case ".", "#":
			default:
				panic("unhandled char")
			}
		}
	}

	return start, end, blizzards, totalRows, totalCols
}
