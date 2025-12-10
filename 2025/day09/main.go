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
	coords := parseInput(input)

	largest := 0
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {

			leftCol := min(c1[0], c2[0])
			rightCol := max(c1[0], c2[0])
			topRow := min(c1[1], c2[1])
			bottomRow := max(c1[1], c2[1])
			size := (bottomRow - topRow + 1) * (rightCol - leftCol + 1)

			largest = max(largest, size)
		}
	}

	return largest
}

func part2(input string) int {
	coords := parseInput(input)

	// store both coords that make up an edge
	horizontalEdges := [][][2]int{}
	verticalEdges := [][][2]int{}

	for i := range len(coords) {
		c1, c2 := coords[i], coords[(i+1)%len(coords)]
		if c1[0] == c2[0] {
			// cols are equal, have a vertical edge
			verticalEdges = append(verticalEdges, [][2]int{
				c1, c2,
			})
		} else {
			// horizontal edge
			horizontalEdges = append(horizontalEdges, [][2]int{
				c1, c2,
			})
		}
	}

	largest := 0
	for i, c1 := range coords {
		for _, c2 := range coords[i+1:] {

			// check every potential pair of coords

			leftCol := min(c1[0], c2[0])
			rightCol := max(c1[0], c2[0])
			topRow := min(c1[1], c2[1])
			bottomRow := max(c1[1], c2[1])
			size := (bottomRow - topRow + 1) * (rightCol - leftCol + 1)

			// optimization to skip pairs that are smaller than a previous result
			if size <= largest {
				continue
			}

			// check if any edges break into this rectangle, if so do not update largest
			isContinuous := true

			for _, edge := range horizontalEdges {
				row := edge[0][1]
				// check horizontal edge is within the rows of the current rectangle
				if row <= topRow || row >= bottomRow {
					continue
				}

				edgeLeft := min(edge[0][0], edge[1][0])
				edgeRight := max(edge[0][0], edge[1][0])

				// final bug fix was removing = from these conditionals
				// i'm not entirely sure why that works but something about the corners of two edges
				if edgeLeft <= leftCol && edgeRight > leftCol {
					isContinuous = false
					break
				} else if edgeLeft < rightCol && edgeRight >= rightCol {
					isContinuous = false
					break
				}
			}
			if !isContinuous {
				continue
			}

			// same thing for verticalEdges
			for _, edge := range verticalEdges {
				col := edge[0][0]
				// check horizontal edge is within the rows of the current rectangle
				if col <= leftCol || col >= rightCol {
					continue
				}

				edgeTop := min(edge[0][1], edge[1][1])
				edgeBottom := max(edge[0][1], edge[1][1])

				if edgeTop <= topRow && edgeBottom > topRow {
					isContinuous = false
					break
				} else if edgeTop < bottomRow && edgeBottom >= bottomRow {
					isContinuous = false
					break
				}
			}
			if !isContinuous {
				continue
			}

			largest = size
		}
	}

	return largest
}

func parseInput(input string) (ans [][2]int) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")
		ans = append(ans, [2]int{
			cast.ToInt(parts[0]),
			cast.ToInt(parts[1]),
		})
	}
	return ans
}
