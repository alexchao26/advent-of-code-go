package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/algos"
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
	bricks := parseInput(input)

	dropBricks(bricks)

	// basically have a graph...
	// if a brick supports nothing, it can be removed
	// if a brick supports
	removableBricks := map[int]bool{}
	for _, brick := range bricks {
		if len(brick.supports) == 0 {
			// supports nothing
			removableBricks[brick.index] = true
		} else {
			// check ALL supported bricks, if this is the ONLY support for those
			// then this brick CANNOT be removed
			hasUniqueDependency := false
			for supportedBrickIndex := range brick.supports {
				if len(bricks[supportedBrickIndex].supportedBy) == 1 {
					hasUniqueDependency = true
				}
			}
			if !hasUniqueDependency {
				removableBricks[brick.index] = true
			}
		}
	}

	return len(removableBricks)
}

func dropBricks(bricks []*brick) {
	// process bricks by lowest z values
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].start[2] < bricks[j].start[2]
	})

	// re-index, pesky bug...
	for i, brick := range bricks {
		brick.index = i
	}

	// all bricks are < 10 units of volume, so can store their coords in a map...
	// also all bricks in the input are straight lines, only one dimension will be > 1
	occupiedCells := map[[3]int]int{}
	for _, brick := range bricks {

		isBlocked := false
		for brick.start[2] > 1 && !isBlocked {
			for _, coord := range brick.coords {
				downOne := [3]int{coord[0], coord[1], coord[2] - 1}
				if index, ok := occupiedCells[downOne]; ok {
					isBlocked = true

					brick.supportedBy[index] = true
					bricks[index].supports[brick.index] = true
				}
			}

			if !isBlocked {
				for i := range brick.coords {
					brick.coords[i][2]--
				}
				brick.start[2]--
				brick.end[2]--
			}
		}

		for _, coord := range brick.coords {
			occupiedCells[coord] = brick.index
		}
	}
}

func part2(input string) int {
	bricks := parseInput(input)
	dropBricks(bricks)

	total := 0

	// chain reaction
	for i := range bricks {
		bricksCopy := copyAllBricks(bricks)
		startingBrick := bricksCopy[i]

		queueToRemove := []*brick{}
		for in := range startingBrick.supports {
			delete(bricksCopy[in].supportedBy, startingBrick.index)
			queueToRemove = append(queueToRemove, bricksCopy[in])
		}

		removed := 0
		for len(queueToRemove) > 0 {
			br := queueToRemove[0]
			queueToRemove = queueToRemove[1:]

			if len(br.supportedBy) > 0 {
				continue
			}

			removed++

			// check every brick it supports, remove self from it's supportedBy map
			// then add to queue to be checked
			for supportedBrickIndex := range br.supports {
				delete(bricksCopy[supportedBrickIndex].supportedBy, br.index)
				if len(bricksCopy[supportedBrickIndex].supportedBy) == 0 {
					queueToRemove = append(queueToRemove, bricksCopy[supportedBrickIndex])
				}
			}
		}

		total += removed
	}

	return total
}

func copyAllBricks(bricks []*brick) []*brick {
	copiedBricks := []*brick{}
	for _, b := range bricks {
		newBrick := &brick{
			// start:       []int{},
			// end:         []int{},
			index: b.index,
			// coords:      [][3]int{},
			supportedBy: map[int]bool{},
			supports:    map[int]bool{},
		}
		// need full copies of these otherwise they'll point to the same underlying maps
		for k, v := range b.supportedBy {
			newBrick.supportedBy[k] = v
		}
		for k, v := range b.supports {
			newBrick.supports[k] = v
		}
		copiedBricks = append(copiedBricks, newBrick)
	}
	return copiedBricks
}

type brick struct {
	start, end  []int
	index       int
	coords      [][3]int
	supportedBy map[int]bool
	supports    map[int]bool
}

func parseInput(input string) (ans []*brick) {
	for _, line := range strings.Split(input, "\n") {
		coords := [6]int{}
		for i, part := range algos.SplitStringOn(line, []string{",", "~"}) {
			coords[i] = cast.ToInt(part)
		}

		if coords[0] > coords[3] || coords[1] > coords[4] || coords[2] > coords[5] {
			panic("unordered input")
		}

		allCoords := [][3]int{}
		for x := coords[0]; x <= coords[3]; x++ {
			for y := coords[1]; y <= coords[4]; y++ {
				for z := coords[2]; z <= coords[5]; z++ {
					allCoords = append(allCoords, [3]int{x, y, z})
				}
			}
		}

		ans = append(ans, &brick{
			start:       coords[:3],
			end:         coords[3:],
			index:       len(ans),
			coords:      allCoords,
			supportedBy: map[int]bool{},
			supports:    map[int]bool{},
		})
	}

	return ans
}
