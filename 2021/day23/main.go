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
	start := parseInput(input)

	minHeap := heap.NewMinHeap()

	minHeap.Add(start)
	seenGrids := map[string]bool{}
	for minHeap.Length() > 0 {
		front := minHeap.Remove().(*state)

		key := fmt.Sprint(front.grid)
		if seenGrids[key] {
			continue
		}
		fmt.Println(minHeap.Length(), "\n", front)
		seenGrids[key] = true

		if front.allDone() {
			return front.energyUsed
		}

		unsettledCoords := front.getUnsettledCoords()
		for _, unsettledCoord := range unsettledCoords {
			// // do not try to move the last one that was moved, otherwise it'll infinite loop
			// if front.coordOfLastMoved == unsettledCoord {
			// 	continue
			// }

			ur, uc := unsettledCoord[0], unsettledCoord[1]
			nextMoves := front.getNextPossibleMoves(unsettledCoord)
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
				cp.coordOfLastMoved = nextCoord

				// add it to the min heap
				minHeap.Add(cp)
			}
		}
	}

	// 10901 TOO LOW

	panic("should return from loop")
}

func part2(input string) int {
	return 0
}

type state struct {
	grid             [][]string
	coordOfLastMoved [2]int // store so you don't try to move the same one twice in a row
	energyUsed       int
	path             string
}

// Value is to implement the heap.Node interface so I can dump states into a Min Heap
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

	sb.WriteString(fmt.Sprintf("nrg: %d, last_moved: %v, path: %s\n", s.energyUsed, s.coordOfLastMoved, s.path))

	return sb.String()
}

// copy method to generate copies to make future heap nodes
func (s *state) copy() *state {
	cp := state{
		grid:             make([][]string, len(s.grid)),
		coordOfLastMoved: s.coordOfLastMoved,
		energyUsed:       s.energyUsed,
		path:             s.path,
	}

	// need to directly copy grid or else underlying arrays will be the same & interfere
	for i := range cp.grid {
		cp.grid[i] = make([]string, len(s.grid[i]))
		copy(cp.grid[i], s.grid[i])
	}

	return &cp
}

var roomCoordToWantChar = map[[2]int]string{
	{2, 3}: "A", {3, 3}: "A",
	{2, 5}: "B", {3, 5}: "B",
	{2, 7}: "C", {3, 7}: "C",
	{2, 9}: "D", {3, 9}: "D",
}

func (s *state) allDone() bool {
	for coord, want := range roomCoordToWantChar {
		if s.grid[coord[0]][coord[1]] != want {
			return false
		}
	}
	return true
}

func (s *state) getUnsettledCoords() [][2]int {
	var unsettled [][2]int
	for r, row := range s.grid {
		for c, v := range row {
			// iterate through the entire grid, for every letter "/[A-D]/"
			if strings.Contains("ABCD", v) {
				// add it to the unsettled list
				coord := [2]int{r, c}
				// IF not in coords map
				if want, ok := roomCoordToWantChar[coord]; !ok {
					unsettled = append(unsettled, coord)
					continue // these are all probably unnecessary but helpful for my brain
				} else {
					// IF in coords map but not matching the wantChar
					if want != v {
						unsettled = append(unsettled, coord)
						continue
					} else {
						// IF it matches wantChar but the cell below is
						// ALSO in coords->want map AND it is the wrong want unsettledChar
						// this means that it is in the right place but needs to get out
						// of the way for a cell below
						below := [2]int{r + 1, c}
						wantBelow, ok := roomCoordToWantChar[below]
						// ok means that it is a "room" cell, not wall
						if ok && wantBelow != s.grid[r+1][c] {
							unsettled = append(unsettled, coord)
							continue
						}
					}
				}

			}
		}
	}
	return unsettled
}

// cannot stop in front of a room
var coordsInFrontOfRooms = map[[2]int]bool{
	{1, 3}: true,
	{1, 5}: true,
	{1, 7}: true,
	{1, 9}: true,
}

func isInHallway(coord [2]int) bool {
	return coord[0] == 1
}

func (s *state) getNextPossibleMoves(unsettledCoord [2]int) [][2]int {
	// get all the eligible locations for this coord to go to
	unsettledChar := s.grid[unsettledCoord[0]][unsettledCoord[1]]

	if !strings.Contains("ABCD", unsettledChar) {
		panic("unexpected character to get next moves for " + unsettledChar)
	}

	startedInHallway := isInHallway(unsettledCoord)
	// fmt.Println(unsettledChar, unsettledCoord, "in hallway", startedInHallway)
	var possible [][2]int

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
					maybeLower := [2]int{front[0] + 1, front[1]}
					if _, ok := roomCoordToWantChar[maybeLower]; ok {
						lowerChar := s.grid[maybeLower[0]][maybeLower[1]]
						if lowerChar == "." {
							possible = append(possible, maybeLower)
							// can only go deeper into the room so just kill the traverse here
							continue
						}
						// if lower char is the same, then can move into the front of the room
						if lowerChar == unsettledChar {
							possible = append(possible, front)
							// no where else to go, so just continue and end this iteration
							continue
						}
					} else {
						// otherwise already deep part of the room, append it
						// ?probably unreachable code
						fmt.Println("unreachable code in else block?")
						possible = append(possible, front)
						continue
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

func parseInput(input string) *state {
	grid := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	// // uncomment to check if coordToWantChars are correct...
	// for c, unsettledChar := range roomCoordToWantChar {
	// 	grid[c[0]][c[1]] = unsettledChar
	// }
	// st := state{grid: grid}
	// fmt.Println(st.String())
	// if !st.allDone() {
	// 	panic("state.allDone() should be true ")
	// }

	return &state{
		grid: grid,
	}
}
