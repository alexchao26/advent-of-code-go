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
	parsed := parseInput(input)
	ans := 0
	for _, hist := range parsed {
		ans += findNextValue(hist)
	}

	return ans
}

func findNextValue(history []int) int {
	matrix := [][]int{
		history,
	}
	nonZeroFound := true
	for nonZeroFound {
		nonZeroFound = false
		next := []int{}

		matrixHistoryRow := matrix[len(matrix)-1]
		for i := 1; i < len(matrixHistoryRow); i++ {
			prev := matrixHistoryRow[i-1]
			curr := matrixHistoryRow[i]
			next = append(next, curr-prev)
			if next[len(next)-1] != 0 {
				nonZeroFound = true
			}
		}
		matrix = append(matrix, next)
	}

	ans := 0
	for _, row := range matrix {
		ans += row[len(row)-1]
	}

	return ans
}

func part2(input string) int {
	parsed := parseInput(input)
	ans := 0
	for _, hist := range parsed {
		ans += findPrevValue(hist)
	}

	return ans
}

func findPrevValue(history []int) int {
	matrix := [][]int{
		history,
	}
	nonZeroFound := true
	for nonZeroFound {
		nonZeroFound = false
		next := []int{}

		matrixHistoryRow := matrix[len(matrix)-1]
		for i := 1; i < len(matrixHistoryRow); i++ {
			prev := matrixHistoryRow[i-1]
			curr := matrixHistoryRow[i]
			next = append(next, curr-prev)
			if next[len(next)-1] != 0 {
				nonZeroFound = true
			}
		}
		matrix = append(matrix, next)
	}

	ans := 0
	for r := len(matrix) - 1; r >= 0; r-- {
		ans = matrix[r][0] - ans
	}

	return ans
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		nums := []int{}
		for _, str := range strings.Split(line, " ") {
			nums = append(nums, cast.ToInt(str))
		}
		ans = append(ans, nums)
	}
	return ans
}
