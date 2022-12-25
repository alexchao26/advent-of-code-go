package main

import (
	_ "embed"
	"flag"
	"fmt"
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
		ans := part2(input, 1000000000000)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	// ####

	// .#.
	// ###
	// .#.

	// ..#
	// ..#
	// ###

	// #
	// #
	// #
	// #

	// ##
	// ##

	// chamber is 7 units wide
	// rocks fall w/ 2 spaces to left wall, 3 spaces to next rock or floor
	// rocks pushed by wind, then fall down 1 space
	// WHEN rock touches something below (ground or another settled rock), rock is "done" and next rock appears
	// 		this also means that rocks will be pushed once L/R before fully settling (unless blocked L/R)

	s := newState(input)
	for i := 0; i < 2022; i++ {
		s.dropRock()
	}

	// height after 2022 rocks fall
	return s.highestSettledRow + 1 // 1 indexed
}

func part2(input string, wantedRocks int) int {
	s := newState(input)
	// obviously can't do the calculations literally anymore...
	// need some kind of "last time this state was seen" check to skip steps
	// 	need to hash states
	//    - the top 10ish rows should be plenty
	//    - rock index
	//    - value = steps since last seen AND highestRow to calc skips

	// [2]int{steps when last seen, rocks dropped, highestRow to calc diff}
	pastStates := map[string][3]int{}

	// will keep track of rows that are mathematically skipped otherwise it'll
	// mess with state and the settled matrix
	dupeRows := 0

	rocksDropped := 0
	for rocksDropped < wantedRocks {
		s.dropRock()
		rocksDropped++

		h := s.hash(20)
		if past, ok := pastStates[h]; ok {
			pastSteps, pastRocksDropped, pastHighRow := past[0], past[1], past[2]

			stepsToSkip := s.stepIndex - pastSteps
			rocksToSkip := rocksDropped - pastRocksDropped
			rowsToAdd := s.highestSettledRow - pastHighRow

			iterationsToSkip := (wantedRocks - rocksDropped) / rocksToSkip
			dupeRows += rowsToAdd * iterationsToSkip
			s.stepIndex += stepsToSkip * iterationsToSkip
			rocksDropped += rocksToSkip * iterationsToSkip
		} else {
			pastStates[h] = [3]int{
				s.stepIndex, rocksDropped, s.highestSettledRow,
			}
		}
	}

	// height after 2022 rocks fall
	return s.highestSettledRow + 1 + dupeRows // 1 indexed
}

type state struct {
	settled           [][]string
	highestSettledRow int
	fallingCoords     [][2]int
	nextRockIndex     int
	steps             []string
	stepIndex         int
}

func newState(input string) state {
	s := state{
		settled:           [][]string{},
		highestSettledRow: -1,
		fallingCoords:     nil,
		nextRockIndex:     0,
		steps:             strings.Split(input, ""),
		stepIndex:         0,
	}

	return s
}

// knew I'd need this for debugging...
func (s state) printState() {
	copySettled := [][]string{}
	for _, row := range s.settled {
		copyRow := make([]string, len(row))
		copy(copyRow, row)
		copySettled = append(copySettled, copyRow)
	}

	for _, coord := range s.fallingCoords {
		copySettled[coord[0]][coord[1]] = "@"
	}

	var sb strings.Builder
	for r := len(copySettled) - 1; r >= 0; r-- {
		sb.WriteString(strings.Join(copySettled[r], "") + cast.ToString(r) + "\n")
	}
	fmt.Println(sb.String())
}

func (s *state) dropRock() {
	s.populateNextBaseCoords()

	highestRow := 0
	for _, c := range s.fallingCoords {
		highestRow = mathy.MaxInt(highestRow, c[0])
	}
	for len(s.settled) <= highestRow {
		s.settled = append(s.settled, newEmptyRow())
	}

	// will be set back to nil when settled
	for s.fallingCoords != nil {
		switch s.steps[s.stepIndex%len(s.steps)] {
		case ">":
			// check if can move right
			canMoveRight := true
			for _, c := range s.fallingCoords {
				if c[1] == 6 || s.settled[c[0]][c[1]+1] != "." {
					canMoveRight = false
				}
			}
			if canMoveRight {
				for i := range s.fallingCoords {
					s.fallingCoords[i][1]++
				}
			}
		case "<":
			// check if can move left
			canMoveLeft := true
			for _, c := range s.fallingCoords {
				if c[1] == 0 || s.settled[c[0]][c[1]-1] != "." {
					canMoveLeft = false
				}
			}
			if canMoveLeft {
				for i := range s.fallingCoords {
					s.fallingCoords[i][1]--
				}
			}
		default:
			panic(s.steps[s.stepIndex])
		}
		s.stepIndex++

		// move down
		canMoveDown := true
		for _, c := range s.fallingCoords {
			if c[0] == 0 || s.settled[c[0]-1][c[1]] != "." {
				canMoveDown = false
			}
		}
		// is blocked, draw onto settled then make nil
		if !canMoveDown {
			for _, c := range s.fallingCoords {
				s.settled[c[0]][c[1]] = "#"
			}
			s.fallingCoords = nil

			for r := len(s.settled) - 1; r >= 0; r-- {
				if strings.Join(s.settled[r], "") != "......." {
					s.highestSettledRow = r
					break
				}
			}
		} else {
			for i := range s.fallingCoords {
				s.fallingCoords[i][0]--
			}
		}
	}
}

func newEmptyRow() []string {
	row := make([]string, 7)
	for i := range row {
		row[i] = "."
	}
	return row
}

var baseCoords = [][][2]int{
	{
		// line ####
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	}, {
		// plus
		{0, 1},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 1},
	}, {
		// flipped L
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 2},
		{2, 2},
	}, {
		// vert line
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	}, {
		// square
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	},
}

func init() {
	// add 2 cols to all baseCoords because they fall 2 off of left wall
	for i := range baseCoords {
		for j := range baseCoords[i] {
			baseCoords[i][j][1] += 2
		}
	}
}

func (s *state) populateNextBaseCoords() {
	copyCoords := make([][2]int, len(baseCoords[s.nextRockIndex]))
	copy(copyCoords, baseCoords[s.nextRockIndex])
	s.nextRockIndex++
	s.nextRockIndex %= 5

	// lowest row of baseCoords...

	for i := range copyCoords {
		copyCoords[i][0] += s.highestSettledRow + 1 + 3
	}
	s.fallingCoords = copyCoords
}

// for part 2 to find return states
// NOTE: had to play with the number of rows to be hashed... 20 seems to
// work on the example input
func (s *state) hash(topRowsToHash int) string {
	var sb strings.Builder
	sb.WriteString(cast.ToString(s.nextRockIndex))
	for r := s.highestSettledRow; r >= 0 && r > s.highestSettledRow-topRowsToHash; r-- {
		sb.WriteString("\n" + strings.Join(s.settled[r], ""))
	}
	return sb.String()
}
