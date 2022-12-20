package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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
	matrix := parseInput(input)
	originCol := 0
	for i, c := range matrix[0] {
		if c == "+" {
			originCol = i
		}
	}

	ans := 0
	for !dropSand(matrix, originCol) {
		ans++
	}

	return ans
}

func part2(input string) int {
	matrix := parseInput(input)
	originCol := 0
	for i, c := range matrix[0] {
		if c == "+" {
			originCol = i
		}
		matrix[len(matrix)-1][i] = "#"
	}

	ans := 0
	for !dropSand(matrix, originCol) {
		ans++
		// COULD incorporate this into the for loop conditional but then the
		// ordering is important... must check origin cell BEFORE running
		// dropSand... it's easier to read here...
		if matrix[0][originCol] == "o" {
			break
		}
	}

	return ans
}

func parseInput(input string) (matrix [][]string) {
	coordSets := [][][2]int{}
	lowestCol := math.MaxInt64
	highestRow := 0
	for _, line := range strings.Split(input, "\n") {
		rawCoords := strings.Split(line, " -> ")
		coords := [][2]int{}
		for _, rawCoord := range rawCoords {
			rawNums := strings.Split(rawCoord, ",")
			col, row := cast.ToInt(rawNums[0]), cast.ToInt(rawNums[1])
			coord := [2]int{
				col, row,
			}
			coords = append(coords, coord)

			lowestCol = mathy.MinInt(lowestCol, col)
			highestRow = mathy.MaxInt(highestRow, row)
		}
		coordSets = append(coordSets, coords)
	}

	// lowering this number to 1 makes it easier to print the matrix, which I
	// used for part 1... but then needed to up it for part 2... or just have a
	// massive screen and make the terminal text tiny...
	ExtraLeftSpace := 200

	highestCol := 0
	for s, set := range coordSets {
		for i := range set {
			coordSets[s][i][0] -= lowestCol - ExtraLeftSpace
			highestCol = mathy.MaxInt(highestCol, coordSets[s][i][0])
		}
	}

	matrix = make([][]string, highestRow+3)
	for r := range matrix {
		matrix[r] = make([]string, highestCol+ExtraLeftSpace*2)
	}

	for _, set := range coordSets {
		for i := 1; i < len(set); i++ {
			cols := []int{set[i-1][0], set[i][0]}
			rows := []int{set[i-1][1], set[i][1]}

			sort.Ints(cols)
			sort.Ints(rows)

			if cols[0] == cols[1] {
				for r := rows[0]; r <= rows[1]; r++ {
					matrix[r][cols[0]] = "#"
				}
			} else if rows[0] == rows[1] {
				for c := cols[0]; c <= cols[1]; c++ {
					matrix[rows[0]][c] = "#"
				}
			}
		}
	}

	originCol := 500 - lowestCol + ExtraLeftSpace
	// make it a plus so it's searchable in the next step... or could just
	// return this value too...
	matrix[0][originCol] = "+"

	for i, r := range matrix {
		for j := range r {
			if matrix[i][j] == "" {
				matrix[i][j] = "."
			}
		}
	}

	// printMatrix(matrix)
	return matrix
}

func printMatrix(matrix [][]string) {
	for _, r := range matrix {
		fmt.Println(r)
	}
}

func dropSand(matrix [][]string, originCol int) (fallsIntoAbyss bool) {
	r, c := 0, originCol

	for r < len(matrix)-1 {
		below := matrix[r+1][c]
		diagonallyLeft := matrix[r+1][c-1]
		diagonallyRight := matrix[r+1][c+1]
		if below == "." {
			r++
		} else if diagonallyLeft == "." {
			r++
			c--
		} else if diagonallyRight == "." {
			r++
			c++
		} else {
			matrix[r][c] = "o"
			return false
		}
	}

	return true
}
