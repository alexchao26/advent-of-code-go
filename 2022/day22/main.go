package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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
	matrix, path := parseInput(input)
	var row, col int
	for c := 0; c < len(matrix[0]); c++ {
		if matrix[0][c] == "." {
			col = c
			break
		}
	}

	diffIndex := 0
	diffs := [][2]int{
		{0, 1},  // start facing right
		{1, 0},  // turning right will point you down
		{0, -1}, // left
		{-1, 0}, // turning left from index 0 makes you face up
	}

	// walking over the edge wraps you around within the same direction...
	// unless if it is a wall, then you're stuck

	p := 0
	for p < len(path) {
		// get next direction...
		indexOfLorR := p
		for indexOfLorR < len(path) &&
			path[indexOfLorR] != 'L' && path[indexOfLorR] != 'R' {
			indexOfLorR++
		}
		steps := cast.ToInt(path[p:indexOfLorR])
		// try to move that many steps
		for s := 0; s < steps; s++ {
			diff := diffs[diffIndex]
			nextRow, nextCol := row+diff[0], col+diff[1]
			// mod them so they wrap if necessary
			nextRow += len(matrix)
			nextCol += len(matrix[0])

			nextRow %= len(matrix)
			nextCol %= len(matrix[0])

			// if it's an empty space you need to keep looping around...
			for matrix[nextRow][nextCol] == " " || matrix[nextRow][nextCol] == "" {
				nextRow += diff[0]
				nextCol += diff[1]

				// wrapping math...
				nextRow += len(matrix)
				nextCol += len(matrix[0])

				nextRow %= len(matrix)
				nextCol %= len(matrix[0])
			}

			// wall: break
			if matrix[nextRow][nextCol] == "#" {
				break
			}
			row = nextRow
			col = nextCol
		}

		if indexOfLorR == len(path) {
			break
		}
		// handle turn if indexOfLorR is still in bounds
		switch path[indexOfLorR] {
		case 'L':
			diffIndex--
		case 'R':
			diffIndex++
		}
		diffIndex += 4
		diffIndex %= 4
		p = indexOfLorR + 1
	}

	// final row, col, facing
	// row & col are 1 indexed
	// facing is indexed same as diffs slice
	// 1000 * row + 4 * col + facing_index
	return 1000*(row+1) + 4*(col+1) + diffIndex
}

func parseInput(input string) ([][]string, string) {
	parts := strings.Split(input, "\n\n")

	matrix := [][]string{}
	topRowLen := len(strings.Split(parts[0], "\n")[0])

	for _, line := range strings.Split(parts[0], "\n") {
		matrix = append(matrix, make([]string, topRowLen))
		split := strings.Split(line, "")
		copy(matrix[len(matrix)-1], split)
	}

	return matrix, parts[1]
}

func part2(input string) int {
	matrix, path := parseInput(input)
	var row, col int
	for c := 0; c < len(matrix[0]); c++ {
		if matrix[0][c] == "." {
			col = c
			break
		}
	}

	diffIndex := 0
	diffs := [][2]int{
		{0, 1},  // start facing right
		{1, 0},  // turning right will point you DOWN
		{0, -1}, // left
		{-1, 0}, // turning left from index 0 makes you face up
	}

	// shape of example
	//	 #
	// ###
	//   ##
	//
	//
	// shape of my input
	//   ##
	//   #
	//  ##
	//  #

	// a lot of (hah) edge case handling to determine where to "teleport" to
	// pen and paper math... might not be worth doing the example input
	//   because i'm going to make a literal calculation for my input shape...

	p := 0
	for p < len(path) {
		// get next direction...
		indexOfLorR := p
		for indexOfLorR < len(path) &&
			path[indexOfLorR] != 'L' && path[indexOfLorR] != 'R' {
			indexOfLorR++
		}
		steps := cast.ToInt(path[p:indexOfLorR])

		// try to move that many steps
		for s := 0; s < steps; s++ {
			diff := diffs[diffIndex]
			nextRow, nextCol := row+diff[0], col+diff[1]

			// DO NOT UPDATE diffIndex here because if it's a wall we DON'T want
			// to change directions
			nextRow, nextCol, nextDiffIndex := handleWrap(row, col, nextRow, nextCol, diffIndex)

			// we'll never see empty spaces now because we're handling wrapping above
			// wall: break
			if matrix[nextRow][nextCol] == "#" {
				break
			}
			// only update if we didn't hit a wall
			row = nextRow
			col = nextCol
			diffIndex = nextDiffIndex
		}

		if indexOfLorR == len(path) {
			break
		}
		// handle turn if indexOfLorR is still in bounds
		switch path[indexOfLorR] {
		case 'L':
			diffIndex--
		case 'R':
			diffIndex++
		}
		diffIndex += 4
		diffIndex %= 4
		p = indexOfLorR + 1
	}

	// final answer calculated from flattened map coords
	// 1000 * row + 4 * col + facing_index
	// too low: 111043
	return 1000*(row+1) + 4*(col+1) + diffIndex
}

// handles edge cases ;)
// how i'll number my boxes...
//     21
//     3
//    54
//    6

const (
	RightIndex = 0
	DownIndex  = 1
	LeftIndex  = 2
	UpIndex    = 3
)

// handleWrap checks if the movement from r,c to nextRow, nextCol is off the
// edge of the matrix, if so it does the maths and direction change to wrap
// around the edge of the cube, this is very manual and based upon a drawing i
// made of my input (i'll upload it when i remember...)
//
// got a little carried away with assertions in here trying to debug...
func handleWrap(r, c, nextRow, nextCol, diffIndex int) (newRow, newCol, newDiffIndex int) {
	// there will be roughly 14 checks in here... this is gonna get ugly, esp
	// b/c i'm too lazy to dry this up

	// 2 -> 5 conversion
	if getBoxNumber(r, c) == 2 &&
		0 <= nextRow && nextRow < 50 && nextCol == 49 {
		if diffIndex != LeftIndex {
			panic(fmt.Sprintf("expected LeftIndex, got %d", diffIndex))
		}

		newCol = 0
		newRow = 149 - nextRow
		if getBoxNumber(newRow, newCol) != 5 {
			panic(fmt.Sprintf("expected to move to box 5, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, RightIndex
	}
	// 5 -> 2
	if getBoxNumber(r, c) == 5 &&
		nextCol == -1 && 100 <= nextRow && nextRow < 150 {
		if diffIndex != LeftIndex {
			panic(fmt.Sprintf("expected LeftIndex got %d", diffIndex))
		}
		newCol = 50
		newRow = 149 - nextRow
		if getBoxNumber(newRow, newCol) != 2 {
			panic(fmt.Sprintf("expected to move to box 2, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, RightIndex
	}

	// 3 -> 5
	if getBoxNumber(r, c) == 3 &&
		nextCol == 49 && 50 <= nextRow && nextRow < 100 {
		if diffIndex != LeftIndex {
			panic(fmt.Sprintf("expected LeftIndex, got %d", diffIndex))
		}
		newRow = 100
		newCol = nextRow - 50
		if getBoxNumber(newRow, newCol) != 5 {
			panic(fmt.Sprintf("expected to move to box 5, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, DownIndex
	}
	// 5 -> 3
	if getBoxNumber(r, c) == 5 &&
		nextRow == 99 && 0 <= nextCol && nextCol < 50 {
		if diffIndex != UpIndex {
			panic(fmt.Sprintf("expected UpIndex, got %d", diffIndex))
		}
		newRow = nextCol + 50
		newCol = 50
		if getBoxNumber(newRow, newCol) != 3 {
			panic(fmt.Sprintf("expected to move to box 3, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, RightIndex
	}

	// 2 -> 6
	if getBoxNumber(r, c) == 2 &&
		nextRow == -1 && 50 <= nextCol && nextCol < 100 {
		if diffIndex != UpIndex {
			panic(fmt.Sprintf("expected UpIndex, got %d", diffIndex))
		}
		newRow = nextCol + 100
		newCol = 0
		if getBoxNumber(newRow, newCol) != 6 {
			panic(fmt.Sprintf("expected to move to box 6, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, RightIndex
	}
	// 6 -> 2
	if getBoxNumber(r, c) == 6 &&
		nextCol == -1 && 150 <= nextRow && nextRow < 200 {
		if diffIndex != LeftIndex {
			panic(fmt.Sprintf("expected LeftIndex, got %d", diffIndex))
		}
		newRow = 0
		newCol = nextRow - 100
		if getBoxNumber(newRow, newCol) != 2 {
			panic(fmt.Sprintf("expected to move to box 2, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, DownIndex
	}

	// 1 -> 6
	if getBoxNumber(r, c) == 1 &&
		nextRow == -1 && 100 <= nextCol && nextCol < 150 {
		if diffIndex != UpIndex {
			panic(fmt.Sprintf("expected UpIndex, got %d", diffIndex))
		}
		newRow = 199
		newCol = nextCol - 100
		if getBoxNumber(newRow, newCol) != 6 {
			panic(fmt.Sprintf("expected to move to box 6, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, UpIndex
	}
	// 6 -> 1
	if getBoxNumber(r, c) == 6 &&
		nextRow == 200 && 0 <= nextCol && nextCol < 50 {
		if diffIndex != DownIndex {
			panic(fmt.Sprintf("expected DownIndex, got %d", diffIndex))
		}
		newRow = 0
		newCol = nextCol + 100
		if getBoxNumber(newRow, newCol) != 1 {
			panic(fmt.Sprintf("expected to move to box 1, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, DownIndex
	}

	// 4 -> 6
	if getBoxNumber(r, c) == 4 &&
		nextRow == 150 && 50 <= nextCol && nextCol < 100 {
		if diffIndex != DownIndex {
			panic(fmt.Sprintf("expected DownIndex, got %d", diffIndex))
		}
		newRow = nextCol + 100
		newCol = 49
		if getBoxNumber(newRow, newCol) != 6 {
			panic(fmt.Sprintf("expected to move to box 6, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, LeftIndex
	}
	// 6 -> 4
	if getBoxNumber(r, c) == 6 &&
		nextCol == 50 && 150 <= nextRow && nextRow < 200 {
		if diffIndex != RightIndex {
			panic(fmt.Sprintf("expected RightIndex, got %d", diffIndex))
		}
		newRow = 149
		newCol = nextRow - 100
		if getBoxNumber(newRow, newCol) != 4 {
			panic(fmt.Sprintf("expected to move to box 4, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, UpIndex
	}

	// 4 -> 1
	if getBoxNumber(r, c) == 4 &&
		nextCol == 100 && 100 <= nextRow && nextRow < 150 {
		if diffIndex != RightIndex {
			panic(fmt.Sprintf("expected RightIndex, got %d", diffIndex))
		}
		newRow = 149 - nextRow
		newCol = 149
		if getBoxNumber(newRow, newCol) != 1 {
			panic(fmt.Sprintf("expected to move to box 1, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, LeftIndex
	}
	// 1 -> 4
	if getBoxNumber(r, c) == 1 &&
		nextCol == 150 && 0 <= nextRow && nextRow < 50 {
		if diffIndex != RightIndex {
			panic(fmt.Sprintf("expected RightIndex, got %d", diffIndex))
		}
		newRow = 149 - nextRow
		newCol = 99
		if getBoxNumber(newRow, newCol) != 4 {
			panic(fmt.Sprintf("expected to move to box 4, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, LeftIndex
	}

	// 3 -> 1
	if getBoxNumber(r, c) == 3 &&
		nextCol == 100 && 50 <= nextRow && nextRow < 100 {
		if diffIndex != RightIndex {
			panic(fmt.Sprintf("expected RightIndex, got %d", diffIndex))
		}
		newRow = 49
		newCol = nextRow + 50
		if getBoxNumber(newRow, newCol) != 1 {
			panic(fmt.Sprintf("expected to move to box 1, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, UpIndex
	}
	// 1 -> 3
	if getBoxNumber(r, c) == 1 &&
		nextRow == 50 && 100 <= nextCol && nextCol < 150 {
		if diffIndex != DownIndex {
			panic(fmt.Sprintf("expected DownIndex, got %d", diffIndex))
		}
		newRow = nextCol - 50
		newCol = 99
		if getBoxNumber(newRow, newCol) != 3 {
			panic(fmt.Sprintf("expected to move to box 3, got %d", getBoxNumber(newRow, newCol)))
		}
		return newRow, newCol, LeftIndex
	}

	// no edge conversion required, just pass through
	return nextRow, nextCol, diffIndex
}

func getBoxNumber(r, c int) int {
	if 0 <= r && r < 50 && 100 <= c && c < 150 {
		return 1
	}
	if 0 <= r && r < 50 && 50 <= c && c < 100 {
		return 2
	}
	if 50 <= r && r < 100 && 50 <= c && c < 100 {
		return 3
	}
	if 100 <= r && r < 150 && 50 <= c && c < 100 {
		return 4
	}
	if 100 <= r && r < 150 && 0 <= c && c < 50 {
		return 5
	}
	if 150 <= r && r < 200 && 0 <= c && c < 50 {
		return 6
	}

	panic(fmt.Sprintf("bad row %d and col %d", r, c))
}
