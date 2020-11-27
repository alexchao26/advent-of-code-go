package main

import (
	"github.com/alexchao26/advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadFile("../input.txt")
	lines := strings.Split(input, "\n")

	seen := make(map[string]bool)
	counted := make(map[string]bool)
	var overlap int

	for _, line := range lines {
		// ID := line[:strings.Index(line, " @")]
		row, _ := strconv.Atoi(line[strings.Index(line, "@")+2 : strings.Index(line, ",")])
		col, _ := strconv.Atoi(line[strings.Index(line, ",")+1 : strings.Index(line, ":")])
		width, _ := strconv.Atoi(line[strings.Index(line, ":")+2 : strings.Index(line, "x")])
		height, _ := strconv.Atoi(line[strings.Index(line, "x")+1:])

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

	fmt.Println("Overlapping cells:", overlap)
}
