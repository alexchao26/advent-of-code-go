package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/data-structures/heap"
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
		ans := clumsyCart(input, 1, 3)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := clumsyCart(input, 4, 10)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func clumsyCart(input string, minMoves, maxMoves int) int {
	grid := parseInput(input)

	minHeap := heap.NewMinHeap()
	minHeap.Add(bfsNode{
		heatLoss:  0,
		row:       0,
		col:       0,
		lastDir:   right,
		debugPath: "",
	})
	minHeap.Add(bfsNode{
		heatLoss:  0,
		row:       0,
		col:       0,
		lastDir:   down,
		debugPath: "",
	})

	// store lowest heat loss for each coordinate for each direction, slightly suboptimal
	// because it can be divided into vertical and horizontal
	// instead of L, R, U, D. But it's a constant time optimization but it runs fast enough
	cache := map[string]int{}

	for minHeap.Length() > 0 {
		node := minHeap.Remove().(bfsNode)

		key := fmt.Sprintf("%d %d - %v", node.row, node.col, node.lastDir)
		if val, ok := cache[key]; ok {
			if node.heatLoss >= val {
				// exit if the current heatLoss isn't better
				continue
			} else {
				cache[key] = node.heatLoss
			}
		} else {
			cache[key] = node.heatLoss
		}

		if node.row == len(grid)-1 && node.col == len(grid[0])-1 {
			return node.heatLoss
		}

		// just add a node for each vertical direction, then those will move vertically as well
		// which covers all possibilities
		for _, nextDir := range verticalTurns[node.lastDir] {
			summedHeatLoss := 0
			for i := 1; i <= maxMoves; i++ {
				nextRow := node.row + nextDir[0]*i
				nextCol := node.col + nextDir[1]*i

				// skip if out of range
				if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
					continue
				}

				summedHeatLoss += grid[nextRow][nextCol]

				// do not add to heap if the cart has moved less than the minimum required moves (part 2)
				if i < minMoves {
					continue
				}

				minHeap.Add(bfsNode{
					heatLoss:  node.heatLoss + summedHeatLoss,
					row:       nextRow,
					col:       nextCol,
					lastDir:   nextDir,
					debugPath: node.debugPath + fmt.Sprintf("%d,%d ", nextRow, nextCol),
				})
			}
		}
	}

	panic("should return from heap processing")
}

type bfsNode struct {
	heatLoss  int
	row, col  int
	lastDir   direction
	debugPath string
}

func (b bfsNode) Value() int {
	return b.heatLoss
}

type direction [2]int

var up = direction{-1, 0}
var down = direction{1, 0}
var left = direction{0, -1}
var right = direction{0, 1}

var verticalTurns = map[direction][2]direction{
	up:    {left, right},
	down:  {left, right},
	left:  {up, down},
	right: {up, down},
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		row := []int{}
		for _, str := range strings.Split(line, "") {
			row = append(row, cast.ToInt(str))
		}
		ans = append(ans, row)
	}
	return ans
}
