package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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
	parts := strings.Split(input, "\n\n")

	coords := map[[2]int]bool{}
	// parse coords
	for _, line := range strings.Split(parts[0], "\n") {
		sp := strings.Split(line, ",")
		coords[[2]int{cast.ToInt(sp[0]), cast.ToInt(sp[1])}] = true
	}
	fmt.Println(len(coords))

	// // printing is a pita
	// grid := make([][]int, 15)
	// for i := range grid {
	// 	grid[i] = make([]int, 15)
	// }
	// for c := range coords {
	// 	grid[c[1]][c[0]] = 1
	// }
	// for _, row := range grid {
	// 	fmt.Println(row)
	// }

	for _, fold := range strings.Split(parts[1], "\n") {
		cap := regexp.MustCompile(`fold along (x|y)=(\d+)`).FindStringSubmatch(fold)
		// remove full match
		cap = cap[1:]

		dir := cap[0]
		foldCoord := cast.ToInt(cap[1])

		fmt.Println("folding on", dir, foldCoord)

		// fmt.Println(fold, dir, line)
		// dots will never appear exactly on a fold line
		isFoldOnX := dir == "x"
		nextMap := map[[2]int]bool{}
		if isFoldOnX {
			for c := range coords {
				if c[0] > foldCoord {
					folded := [2]int{
						foldCoord - (c[0] - foldCoord),
						c[1],
					}
					fmt.Println(c, "to", folded)
					nextMap[folded] = true
				} else {
					nextMap[c] = true
				}
			}
		} else {
			// fold on y
			for c := range coords {
				if c[1] > foldCoord {
					folded := [2]int{
						c[0],
						foldCoord - (c[1] - foldCoord),
					}
					fmt.Println(c, "to", folded)
					nextMap[folded] = true
				} else {
					nextMap[c] = true
				}
			}
		}

		coords = nextMap

		// TODO just break for part 1
		// break
	}

	// printing is a pita
	max := 0
	for c := range coords {
		if c[0] > max {
			max = c[0]
		}
		if c[1] > max {
			max = c[1]
		}
	}

	grid := make([][]int, max+1)
	for i := range grid {
		grid[i] = make([]int, max+1)
	}
	for c := range coords {
		grid[c[1]][c[0]] = 1
	}
	for _, row := range grid {
		str := ""
		for _, val := range row {
			if val == 1 {
				str += "#"
			} else {
				str += "."
			}
		}
		fmt.Println(str)
	}

	return len(coords)
}

func part2(input string) int {
	return 0
}
