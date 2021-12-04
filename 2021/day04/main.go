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
	nums, boards := parseInput(input)

	for _, n := range nums {
		for _, b := range boards {
			didWin := b.PickNum(n)
			if didWin {
				// multiply score of winning board by number that was just called
				return b.Score() * n
			}
		}
	}

	panic("a board should've won and returned from the loop")
}

func part2(input string) int {
	nums, boards := parseInput(input)

	lastWinningScore := -1
	alreadyWon := map[int]bool{}
	for _, n := range nums {
		for bi, b := range boards {
			if alreadyWon[bi] {
				continue
			}
			didWin := b.PickNum(n)
			if didWin {
				// WHICH BOARD WINS LAST
				lastWinningScore = b.Score() * n

				// mark board as already won
				alreadyWon[bi] = true
			}
		}
	}

	return lastWinningScore

}

// BoardState maintains a parsed board and a boolean matrix of cells that have
// been picked/marked
type BoardState struct {
	board  [][]int
	picked [][]bool
}

func NewBoardState(board [][]int) BoardState {
	picked := make([][]bool, len(board))
	for i := range picked {
		picked[i] = make([]bool, len(board[0]))
	}
	return BoardState{
		board:  board,
		picked: picked,
	}
}

func (b *BoardState) PickNum(num int) bool {
	for r, rows := range b.board {
		for c, v := range rows {
			if v == num {
				b.picked[r][c] = true
			}
		}
	}

	// is this fast enough to do on every "cycle"?
	// guess so. probably a constant time way to do this but oh well
	for i := 0; i < len(b.board); i++ {
		isFullRow, isFullCol := true, true
		// board is square so this works fine, otherwise would need another pair of nested loops
		for j := 0; j < len(b.board); j++ {
			// check row at index i
			if !b.picked[i][j] {
				isFullRow = false
			}
			// check col at index j
			if !b.picked[j][i] {
				isFullCol = false
			}
		}
		if isFullRow || isFullCol {
			// returns true if is winning board
			return true
		}
	}

	// false for incomplete board
	return false
}

func (b *BoardState) Score() int {
	var score int

	for r, rows := range b.board {
		for c, v := range rows {
			// adds up all the non-picked/marked cells
			if !b.picked[r][c] {
				score += v
			}
		}
	}

	return score
}

func parseInput(input string) (nums []int, boards []BoardState) {
	lines := strings.Split(input, "\n\n")

	for _, v := range strings.Split(lines[0], ",") {
		nums = append(nums, cast.ToInt(v))
	}

	for _, grid := range lines[1:] {
		b := [][]int{}
		for _, line := range strings.Split(grid, "\n") {
			line = strings.ReplaceAll(line, "  ", " ")
			for line[0] == ' ' {
				line = line[1:]
			}
			parts := strings.Split(line, " ")

			row := []int{}
			for _, p := range parts {
				row = append(row, cast.ToInt(p))
			}
			b = append(b, row)
		}

		boards = append(boards, NewBoardState(b))
	}
	return nums, boards
}
