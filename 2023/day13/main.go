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
	patterns := parseInput(input)

	ans := 0
	for _, pattern := range patterns {
		maybeMirrorRow := findMirrorRow(pattern, -1)
		if maybeMirrorRow != -1 {
			ans += 100 * maybeMirrorRow
		} else {
			maybeMirrorCol := findMirrorCol(pattern, -1)
			if maybeMirrorCol == -1 {
				panic("did not find mirror row or col")
			}
			ans += maybeMirrorCol
		}
	}

	return ans
}

// returns the zero-index of the line, if the line is between index 3 and 4, it returns 4
// ignoreRow is for part2 where we want to ignore the original mirrored row because it may still be
// valid in the un-smudged pattern
func findMirrorRow(pattern [][]string, ignoreRow int) int {
	// combine the string slices into a string so they're easier to compare
	combinedRows := []string{}
	for _, row := range pattern {
		combinedRows = append(combinedRows, strings.Join(row, ""))
	}

	for i := 1; i < len(combinedRows); i++ {
		mismatchFound := false
		for offset := 1; i-offset >= 0 && i+offset-1 < len(combinedRows); offset++ {
			// fmt.Println("combined row indexes", i-offset, i+offset-1)
			// fmt.Println(combinedRows[i-offset], "\n", combinedRows[i+offset-1])
			if combinedRows[i-offset] != combinedRows[i+offset-1] {
				mismatchFound = true
				// fmt.Println("mismatch found")
				break
			}
		}

		if !mismatchFound {
			if i != ignoreRow {
				return i
			}
		}
	}
	// none found
	return -1
}

func findMirrorCol(pattern [][]string, ignoreCol int) int {
	// rotate the grid, maintaining the indices for easier maths later
	// then just pass it into the findMirrorRow func
	rotatedGrid := [][]string{}
	for c := 0; c < len(pattern[0]); c++ {
		newRow := []string{}
		for r := 0; r < len(pattern); r++ {
			newRow = append(newRow, pattern[r][c])
		}
		rotatedGrid = append(rotatedGrid, newRow)
	}

	return findMirrorRow(rotatedGrid, ignoreCol)
}

func part2(input string) int {
	patterns := parseInput(input)

	ans := 0

	for _, pattern := range patterns {

		// store the original row and col so they can be ignored in the find mirror row func
		originalMirrorRow := findMirrorRow(pattern, -1)
		originalMirrorCol := findMirrorCol(pattern, -1)

	traverse:
		// labels suck but without the breaks this all has to go into a separate function which is
		// arguably less readable. and the break is necessary to not double count reflections
		for r, row := range pattern {
			for c, val := range row {
				if val == "." {
					pattern[r][c] = "#"
					if maybeMirrorRow := findMirrorRow(pattern, originalMirrorRow); maybeMirrorRow != -1 {
						ans += 100 * maybeMirrorRow
						break traverse
					}
					if maybeMirrorCol := findMirrorCol(pattern, originalMirrorCol); maybeMirrorCol != -1 {
						ans += maybeMirrorCol
						break traverse
					}
					pattern[r][c] = "."
				} else if val == "#" {
					pattern[r][c] = "."
					if maybeMirrorRow := findMirrorRow(pattern, originalMirrorRow); maybeMirrorRow != -1 {
						ans += 100 * maybeMirrorRow
						break traverse
					}
					if maybeMirrorCol := findMirrorCol(pattern, originalMirrorCol); maybeMirrorCol != -1 {
						ans += maybeMirrorCol
						break traverse
					}
					pattern[r][c] = "#"

				} else {
					panic("expected input: " + val)
				}
			}
		}
	}

	return ans
}

func parseInput(input string) (ans [][][]string) {
	for _, section := range strings.Split(input, "\n\n") {
		grid := [][]string{}
		for _, line := range strings.Split(section, "\n") {
			grid = append(grid, strings.Split(line, ""))
		}
		ans = append(ans, grid)
	}
	return ans
}
