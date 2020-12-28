package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := letItSnow(util.ReadFile("./input.txt"))
	fmt.Println("Output:", ans)
}

func letItSnow(input string) int {
	var row, col int
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, ".", "")
	for _, part := range strings.Split(input, " ") {
		if regexp.MustCompile("[0-9]").MatchString(part) {
			if row == 0 { // jeez i am getting lazy
				row = cast.ToInt(part)
			} else {
				col = cast.ToInt(part)
			}
		}
	}

	// the number of iterations to run can be calculated by:
	// - finding the number of cells that in the triangle that is formed by each
	//   diagonal prior to the incomplete one that the target cell is on
	//     This triangle is generated from adding 1+2+3+4...+(row+col) 0-indexed
	// - then the number of iterations for the incomplete diagonal, is equal to
	//   the current column number (1-indexed)
	var triangleBefore int
	for i := 1; i <= row+col-2; i++ {
		triangleBefore += i
	}

	numberOnThisDiagonal := col

	// subtract one for starting cell
	iterations := triangleBefore + numberOnThisDiagonal - 1

	// and thankfully this runs quickly
	code := 20151125
	for i := 0; i < iterations; i++ {
		code *= 252533
		code %= 33554393
	}

	return code
}
