package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
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

	ans := trenchMap(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func trenchMap(input string, part int) int {
	alg, img := parseInput(input)

	// determine what the infinite space will do, if alg's zero index is "." then it will STAY "."
	infiniteSpaceStaysOff := alg[0] == "."

	totalSteps := 2
	if part == 2 {
		totalSteps = 50
	}
	for steps := 0; steps < totalSteps; steps++ {
		// initial state is off, so index 0 -> false
		infiniteSpaceIsOn := steps%2 == 1
		img = enhanceImg(img, alg, infiniteSpaceStaysOff, infiniteSpaceIsOn)
	}

	var count int
	for _, pix := range img {
		if pix == "#" {
			count++
		}
	}
	return count
}

func enhanceImg(img map[[2]int]string, alg []string, infiniteSpaceStaysOff, infiniteSpaceIsOn bool) map[[2]int]string {
	// get bounds
	var firstRow, lastRow, firstCol, lastCol int
	for coord := range img {
		firstRow = mathy.MinInt(firstRow, coord[0])
		lastRow = mathy.MaxInt(lastRow, coord[0])
		firstCol = mathy.MinInt(firstCol, coord[1])
		lastCol = mathy.MaxInt(lastCol, coord[1])
	}

	// just extend the entire img by 3 spaces around all borders (up, down, left, right)
	// choose which character to extend with based on if the infinite space stays off AND this is a
	// where it would be toggled on
	infChar := "."
	if !infiniteSpaceStaysOff && infiniteSpaceIsOn {
		infChar = "#"
	}
	// fmt.Println("inf Char", infChar)

	for c := firstCol - 3; c <= lastCol+3; c++ {
		img[[2]int{firstRow - 3, c}] = infChar
		img[[2]int{firstRow - 2, c}] = infChar
		img[[2]int{firstRow - 1, c}] = infChar
		img[[2]int{lastRow + 1, c}] = infChar
		img[[2]int{lastRow + 2, c}] = infChar
		img[[2]int{lastRow + 3, c}] = infChar
	}
	for r := firstRow - 3; r <= lastRow+3; r++ {
		img[[2]int{r, firstCol - 3}] = infChar
		img[[2]int{r, firstCol - 2}] = infChar
		img[[2]int{r, firstCol - 1}] = infChar
		img[[2]int{r, lastCol + 1}] = infChar
		img[[2]int{r, lastCol + 2}] = infChar
		img[[2]int{r, lastCol + 3}] = infChar
	}

	// fmt.Println("BEFORE")
	// debugging helper to print an infinite grid for
	// halp.PrintInfiniteGridStrings(img, ".")

	// now only need to check within firstRow - 2 through lastRow + 2 (same for cols)
	// because flickers will kill that third row/col out
	next := map[[2]int]string{}
	for r := firstRow - 2; r <= lastRow+2; r++ {
		for c := firstCol - 2; c <= lastCol+2; c++ {
			ind := getAlgIndex(img, r, c)
			char := alg[ind]
			next[[2]int{r, c}] = char
		}
	}

	// fmt.Println("AFTER")
	// halp.PrintInfiniteGridStrings(next, ".")

	return next
}

func getAlgIndex(img map[[2]int]string, r, c int) int {
	var rawBinary string
	for _, d := range [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	} {
		coord := [2]int{r + d[0], c + d[1]}
		if img[coord] == "#" {
			rawBinary += "1"
		} else {
			// else also captures indexes that are in the infinite space that are not in img map yet
			rawBinary += "0"
		}
	}
	dec, err := strconv.ParseInt(rawBinary, 2, 64)
	if err != nil {
		panic("parsing rawBinary " + err.Error())
	}
	return int(dec)
}

func parseInput(input string) (alg []string, img map[[2]int]string) {
	parts := strings.Split(input, "\n\n")
	alg = strings.Split(parts[0], "")

	img = map[[2]int]string{}
	for r, line := range strings.Split(parts[1], "\n") {
		for c, char := range strings.Split(line, "") {
			img[[2]int{r, c}] = char
		}
	}

	return alg, img
}
