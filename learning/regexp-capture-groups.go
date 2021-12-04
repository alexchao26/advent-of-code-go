package learning

import (
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
)

func CaptureBingoBoard(board string) [][]int {
	var nums [][]int
	// parens create indexed capture groups
	// when used with (*regexp).FindStringSubmatch a string slice is returned
	pattern := regexp.MustCompile(`\s?(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)

	for _, row := range strings.Split(board, "\n") {
		matches := pattern.FindStringSubmatch(row)
		if matches == nil {
			panic("row does not match pattern: " + row)
		}

		// submatch[0] is the entire string
		// submatch[1:] are the captured groups that i'm interested in
		var rowNums []int
		for _, v := range matches[1:] {
			rowNums = append(rowNums, cast.ToInt(v))
		}
		nums = append(nums, rowNums)
	}

	return nums
}
