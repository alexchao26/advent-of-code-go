package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/data-structures/heap"
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

	ans := amphipodDay23(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

var roomCoordToWantCharPart1 = map[[2]int]string{
	{2, 3}: "A", {3, 3}: "A",
	{2, 5}: "B", {3, 5}: "B",
	{2, 7}: "C", {3, 7}: "C",
	{2, 9}: "D", {3, 9}: "D",
}
var roomCoordToWantCharPart2 = map[[2]int]string{
	{2, 3}: "A", {3, 3}: "A", {4, 3}: "A", {5, 3}: "A",
	{2, 5}: "B", {3, 5}: "B", {4, 5}: "B", {5, 5}: "B",
	{2, 7}: "C", {3, 7}: "C", {4, 7}: "C", {5, 7}: "C",
	{2, 9}: "D", {3, 9}: "D", {4, 9}: "D", {5, 9}: "D",
}

func amphipodDay23(input string, part int) int {
	start := parseInput(input)

	roomCoordToWantChar := roomCoordToWantCharPart1
	if part == 2 {
		roomCoordToWantChar = roomCoordToWantCharPart2

		// update the grid with the 2 new rows, move old ones down
		start.grid = append(start.grid, nil, nil)
		start.grid[6] = start.grid[4]
		start.grid[5] = start.grid[3]

		start.grid[3] = strings.Split("  #D#C#B#A#  ", "")
		start.grid[4] = strings.Split("  #D#B#A#C#  ", "")
	}

	minHeap := heap.NewMinHeap()

	minHeap.Add(start)
	seenGrids := map[string]bool{}
	for minHeap.Length() > 0 {
		front := minHeap.Remove().(*state)

		key := fmt.Sprint(front.grid)
		if seenGrids[key] {
			continue
		}
		seenGrids[key] = true

		if front.allDone(roomCoordToWantChar) {
			return front.energyUsed
		}

		unsettledCoords := front.getUnsettledCoords(roomCoordToWantChar)
		for _, unsettledCoord := range unsettledCoords {
			ur, uc := unsettledCoord[0], unsettledCoord[1]
			nextMoves := front.getNextPossibleMoves(unsettledCoord, roomCoordToWantChar)
			for _, nextCoord := range nextMoves {
				nr, nc := nextCoord[0], nextCoord[1]
				if front.grid[nr][nc] != "." {
					panic(fmt.Sprintf("should only be moving to walkable spaces, got %q at %d,%d", front.grid[nr][nc], nr, nc))
				}

				cp := front.copy()
				// add the energy that will be used, swap the two coords
				cp.energyUsed += calcEnergy(cp.grid[ur][uc], unsettledCoord, nextCoord)
				cp.path += fmt.Sprintf("%s%v->%v{%d},", front.grid[ur][uc], unsettledCoord, nextCoord, cp.energyUsed)
				cp.grid[nr][nc], cp.grid[ur][uc] = cp.grid[ur][uc], cp.grid[nr][nc]

				// add it to the min heap
				minHeap.Add(cp)
			}
		}
	}

	panic("should return from loop")
}

type state struct {
	grid       [][]string
	energyUsed int
	path       string // for debugging
}

func parseInput(input string) *state {
	grid := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	return &state{
		grid: grid,
	}
}

// Value is to implement the heap.heapNode interface so I can dump states into a Min Heap
func (s *state) Value() int {
	return s.energyUsed
}

func (s *state) String() string {
	var sb strings.Builder
	for _, row := range s.grid {
		for _, unsettledChar := range row {
			sb.WriteString(unsettledChar)
		}
		sb.WriteRune('\n')
	}

	sb.WriteString(fmt.Sprintf("nrg: %d, ,path: %s\n", s.energyUsed, s.path))

	return sb.String()
}

// copy method to generate copies to make future heap nodes
func (s *state) copy() *state {
	cp := state{
		grid:       make([][]string, len(s.grid)),
		energyUsed: s.energyUsed,
		path:       s.path,
	}

	// need to directly copy grid or else underlying arrays will be the same & interfere
	for i := range cp.grid {
		cp.grid[i] = make([]string, len(s.grid[i]))
		copy(cp.grid[i], s.grid[i])
	}

	return &cp
}

func (s *state) allDone(roomCoordToWantChar map[[2]int]string) bool {
	for coord, want := range roomCoordToWantChar {
		if s.grid[coord[0]][coord[1]] != want {
			return false
		}
	}
	return true
}

func (s *state) getUnsettledCoords(roomCoordToWantChar map[[2]int]string) [][2]int {
	var unsettled [][2]int
	// check entire hallway
	for col := 1; col < len(s.grid[0]); col++ {
		if strings.Contains("ABCD", s.grid[1][col]) {
			unsettled = append(unsettled, [2]int{1, col})
		}
	}

	for _, col := range []int{3, 5, 7, 9} {
		roomFullFromBack := true
		for row := len(s.grid) - 2; row >= 2; row-- {
			coord := [2]int{row, col}
			wantChar := roomCoordToWantChar[coord]
			gotChar := s.grid[row][col]
			if gotChar != "." {
				if gotChar != wantChar {
					roomFullFromBack = false
					unsettled = append(unsettled, coord)
				} else if gotChar == wantChar && !roomFullFromBack {
					// need to get out of the way of someone in the wrong room
					unsettled = append(unsettled, coord)
				}
			}
		}
	}
	return unsettled
}

// cannot stop in front of a room, still applicable for part2
var coordsInFrontOfRooms = map[[2]int]bool{
	{1, 3}: true,
	{1, 5}: true,
	{1, 7}: true,
	{1, 9}: true,
}

func isInHallway(coord [2]int) bool {
	return coord[0] == 1
}

func (s *state) getNextPossibleMoves(unsettledCoord [2]int, roomCoordToWantChar map[[2]int]string) [][2]int {
	// get all the eligible locations for this coord to go to
	unsettledChar := s.grid[unsettledCoord[0]][unsettledCoord[1]]

	if !strings.Contains("ABCD", unsettledChar) {
		panic("unexpected character to get next moves for " + unsettledChar)
	}

	var possible [][2]int

	startedInHallway := isInHallway(unsettledCoord)

	queue := [][2]int{unsettledCoord}
	seen := map[[2]int]bool{}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if seen[front] {
			continue
		}
		seen[front] = true

		if front != unsettledCoord {
			// is not a coord in front of a room
			if !coordsInFrontOfRooms[front] {
				wantChar, isRoomCoord := roomCoordToWantChar[front]
				// if NOT in a room, append it
				if !isRoomCoord {
					// ONLY add a hallway if it started in a room bc of rule 3
					if !startedInHallway {
						possible = append(possible, front)
					}
				} else if wantChar == unsettledChar {
					// found the correct room
					// check if there is a deeper part of the room (aka lower)

					// if there is a "stuck" amphipod deeper in the room, cannot stop here
					// if not deepest empty coord, cannot stop here
					// in both cases walking further is handles all cases, whether that's
					//   to walk further in or out of the room
					isStuckAmphipod := false
					roomHasDeeperOpenSpaces := false
					for r := front[0] + 1; r < len(s.grid)-1; r++ {
						char := s.grid[r][front[1]]
						if char == "." {
							roomHasDeeperOpenSpaces = true
						}
						if char != "." && char != unsettledChar {
							isStuckAmphipod = true
							break
						}
					}

					if !roomHasDeeperOpenSpaces && !isStuckAmphipod {
						possible = append(possible, front)
					}
				}
			}
		}

		for _, d := range [][2]int{
			// up down left right
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			// do not need to check in range because the entire walkable area is surrounded by walls
			next := [2]int{front[0] + d[0], front[1] + d[1]}
			if s.grid[next[0]][next[1]] == "." {
				// add to queue to keep walking regardless of whether or not it gets added to the possible slice
				queue = append(queue, next)
			}
		}
	}

	return possible
}

func calcEnergy(char string, start, end [2]int) int {
	// start with cols distance
	dist := mathy.AbsInt(end[1] - start[1])
	// add distance to hallway for start and end?
	dist += start[0] - 1
	dist += end[0] - 1

	energyPerType := map[string]int{
		"A": 1,
		"B": 10,
		"C": 100,
		"D": 1000,
	}

	if _, ok := energyPerType[char]; !ok {
		panic(char + " should not call calcEnergy()")
	}
	return energyPerType[char] * dist
}
