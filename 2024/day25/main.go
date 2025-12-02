package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

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

	ans := part1(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	parts := strings.Split(input, "\n\n")

	locks := [][]int{}
	keys := [][]int{}
	maxHeight := 0
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		maxHeight = len(lines) - 1

		// lock detection: first line is all '#'
		fullLock := strings.Repeat("#", len(lines[0]))
		if lines[0] == fullLock {
			lockHeights := []int{}
			// for each column, find first non-dot from bottom
			for col := range len(lines[0]) {
				for row := len(lines) - 1; row >= 0; row-- {
					if lines[row][col:col+1] != "." {
						lockHeights = append(lockHeights, row)
						break
					}
				}
			}
			locks = append(locks, lockHeights)
		} else {
			// key detection: for each column, find first non-dot from top
			keyHeights := []int{}
			for col := range len(lines[0]) {
				for row := range len(lines) {
					if lines[row][col:col+1] != "." {
						keyHeights = append(keyHeights, len(lines)-1-row)
						break
					}
				}
			}
			keys = append(keys, keyHeights)
		}
	}

	count := 0

	for _, lock := range locks {
		for _, key := range keys {
			fits := true
			for col := range len(key) {
				if lock[col]+key[col] >= maxHeight {
					fits = false
					break
				}
			}
			if fits {
				count += 1
			}
		}
	}

	return count
}
