package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
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

	ans := unstableDiffusion(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

var diffsSlice = [][][2]int{
	// N
	{
		{-1, -1},
		{-1, 0},
		{-1, 1},
	},
	// S
	{
		{1, -1},
		{1, 0},
		{1, 1},
	},
	// W
	{
		{-1, -1},
		{0, -1},
		{1, -1},
	},
	// E
	{
		{-1, 1},
		{0, 1},
		{1, 1},
	},
}
var targetDiff = [][2]int{
	{-1, 0}, // N
	{1, 0},  // S
	{0, -1}, // W
	{0, 1},  // E
}

func unstableDiffusion(input string, part int) int {
	elfCoords := parseInput(input)
	diffStartIndex := 0

	// part 2
	var lastState string

	round := 1

	for {
		if part == 1 && round == 11 {
			break
		}

		elfPlannedMoves := [][][2]int{}
		elvesTargetingCoord := map[[2]int]int{}

		for coords, val := range elfCoords {
			if val == "#" {
				nonZeroNeighbors := 0
				for _, diffSlice := range diffsSlice {
					for _, d := range diffSlice {
						if elfCoords[[2]int{
							coords[0] + d[0],
							coords[1] + d[1],
						}] == "#" {
							nonZeroNeighbors++
						}
					}
				}

				if nonZeroNeighbors == 0 {
					elfPlannedMoves = append(elfPlannedMoves, [][2]int{
						coords, coords,
					})
					elvesTargetingCoord[coords]++
				} else {
					foundAMove := false
					for i := 0; i < 4; i++ {
						diffSliceIndex := i + diffStartIndex
						diffSlice := diffsSlice[diffSliceIndex%4]
						neighbors := 0
						for _, d := range diffSlice {
							if elfCoords[[2]int{
								coords[0] + d[0],
								coords[1] + d[1],
							}] == "#" {
								neighbors++
							}
						}
						if neighbors == 0 {
							nextCoords := coords
							nextCoords[0] += targetDiff[diffSliceIndex%4][0]
							nextCoords[1] += targetDiff[diffSliceIndex%4][1]

							elfPlannedMoves = append(elfPlannedMoves, [][2]int{
								coords, nextCoords,
							})
							elvesTargetingCoord[nextCoords]++

							foundAMove = true
							break
						}
					}
					if !foundAMove {
						elfPlannedMoves = append(elfPlannedMoves, [][2]int{
							coords, coords,
						})
						elvesTargetingCoord[coords]++
					}
				}

			}
		}

		// reset coords, but only if elves are not blocked...
		elfCoords = map[[2]int]string{}

		for _, plannedMove := range elfPlannedMoves {
			if elvesTargetingCoord[plannedMove[1]] > 1 {
				// stay
				elfCoords[plannedMove[0]] = "#"
			} else {
				// move
				elfCoords[plannedMove[1]] = "#"
			}
		}

		// rotate directions that are checked
		diffStartIndex++

		if part == 2 { // hash the state
			allCoords := []string{}
			for c := range elfCoords {
				allCoords = append(allCoords, fmt.Sprint(c))
			}
			sort.Strings(allCoords)
			thisState := fmt.Sprint(allCoords)
			if lastState == thisState {
				return round
			}
			lastState = thisState
		}

		round++
	}

	lowRow, highRow, lowCol, highCol := math.MaxInt16, math.MinInt16, math.MaxInt16, math.MinInt16
	for coords := range elfCoords {
		lowRow = mathy.MinInt(lowRow, coords[0])
		highRow = mathy.MaxInt(highRow, coords[0])
		lowCol = mathy.MinInt(lowCol, coords[1])
		highCol = mathy.MaxInt(highCol, coords[1])
	}

	ans := 0
	for r := lowRow; r <= highRow; r++ {
		for c := lowCol; c <= highCol; c++ {
			if elfCoords[[2]int{r, c}] != "#" {
				ans++
			}
		}
	}

	return ans
}

func parseInput(input string) map[[2]int]string {
	ans := map[[2]int]string{}
	for r, line := range strings.Split(input, "\n") {
		for c, v := range strings.Split(line, "") {
			if v == "#" {
				ans[[2]int{r, c}] = "#"
			}
		}
	}
	return ans
}
