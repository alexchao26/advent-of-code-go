package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	seen := make(map[string]bool)
	counted := make(map[string]bool)
	var overlap int

	for _, line := range lines {
		row := cast.ToInt(line[strings.Index(line, "@")+2 : strings.Index(line, ",")])
		col := cast.ToInt(line[strings.Index(line, ",")+1 : strings.Index(line, ":")])
		width := cast.ToInt(line[strings.Index(line, ":")+2 : strings.Index(line, "x")])
		height := cast.ToInt(line[strings.Index(line, "x")+1:])

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				coords := fmt.Sprintf("%vx%v", row+i, col+j)
				if seen[coords] && !counted[coords] {
					overlap++
					counted[coords] = true
				}
				seen[coords] = true
			}
		}
	}

	return overlap
}

func part2(input string) int {
	lines := strings.Split(input, "\n")

	coords := makeMapOfCoordinates(lines)

	for _, line := range lines {
		uniqueID := hasNoOverlap(coords, line)
		if uniqueID != -1 {
			return uniqueID
		}
	}
	panic("expect return from loop")
}

func makeMapOfCoordinates(lines []string) map[string]int {
	seen := make(map[string]int)
	for _, line := range lines {
		// ID := line[:strings.Index(line, " @")]
		row := cast.ToInt(line[strings.Index(line, "@")+2 : strings.Index(line, ",")])
		col := cast.ToInt(line[strings.Index(line, ",")+1 : strings.Index(line, ":")])
		width := cast.ToInt(line[strings.Index(line, ":")+2 : strings.Index(line, "x")])
		height := cast.ToInt(line[strings.Index(line, "x")+1:])

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				coords := fmt.Sprintf("%vx%v", row+i, col+j)
				seen[coords]++
			}
		}
	}
	return seen
}

// if cut is unique returns the ID, otherwise -1
func hasNoOverlap(seen map[string]int, line string) int {
	ID := cast.ToInt(line[1:strings.Index(line, " @")])
	row := cast.ToInt(line[strings.Index(line, "@")+2 : strings.Index(line, ",")])
	col := cast.ToInt(line[strings.Index(line, ",")+1 : strings.Index(line, ":")])
	width := cast.ToInt(line[strings.Index(line, ":")+2 : strings.Index(line, "x")])
	height := cast.ToInt(line[strings.Index(line, "x")+1:])

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			coords := fmt.Sprintf("%vx%v", row+i, col+j)
			if seen[coords] != 1 {
				return -1
			}
		}
	}
	return ID
}
