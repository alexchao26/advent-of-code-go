package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := part1(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func part1(input string, part int) string {
	directions := parseInput(input)

	var pressedButtons string

	// set starting row, col and phone pad based on part number
	// pad the borders of phonePad with spaces to make border collision logic
	// the same in both parts
	// used a space instead of empty string so it's easier to look at...
	row, col := 2, 2
	phonePad := [][]string{
		{" ", " ", " ", " ", " "},
		{" ", "1", "2", "3", " "},
		{" ", "4", "5", "6", " "},
		{" ", "7", "8", "9", " "},
		{" ", " ", " ", " ", " "},
	}
	if part == 2 {
		phonePad = [][]string{
			{" ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", "1", " ", " ", " "},
			{" ", " ", "2", "3", "4", " ", " "},
			{" ", "5", "6", "7", "8", "9", " "},
			{" ", " ", "A", "B", "C", " ", " "},
			{" ", " ", " ", "D", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " "},
		}
		row, col = 3, 1
	}
	for _, list := range directions {
		for _, direction := range list {
			switch direction {
			case "U":
				if phonePad[row-1][col] != " " {
					row--
				}
			case "D":
				if phonePad[row+1][col] != " " {
					row++
				}
			case "L":
				if phonePad[row][col-1] != " " {
					col--
				}
			case "R":
				if phonePad[row][col+1] != " " {
					col++
				}
			default:
				panic("unhandled direction: " + direction)
			}
		}
		pressedButtons += phonePad[row][col]
	}

	return pressedButtons
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, ""))
	}
	return ans
}
