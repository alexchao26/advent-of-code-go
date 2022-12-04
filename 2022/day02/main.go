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

const (
	Win  = 6
	Loss = 0
	Draw = 3

	Rock     = 1
	Paper    = 2
	Scissors = 3
)

// this code is heinoussssss
// there's some kind of cheeky map to circular LL nodes to determine if you won or lost
// but there are so few branches to handle that it's faster to just code it manually...

func part1(input string) int {
	lines := parseInput(input)

	// opp choices: ABC rock-paper-scissors
	// my choices:  XYZ rock-paper-scissors
	choices := map[string]int{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	totalScore := 0
	for _, l := range lines {
		if _, ok := choices[l[1]]; !ok {
			panic("choice not in choices map")
		}
		totalScore += choices[l[1]]
		switch l[1] {
		case "X": // i played rock
			switch l[0] {
			case "A":
				totalScore += Draw
			case "B":
				totalScore += Loss
			case "C":
				totalScore += Win
			default:
				panic("unacceptable opp choice " + l[0])
			}
		case "Y": // i played paper
			switch l[0] {
			case "A": // rock
				totalScore += Win
			case "B": // paper
				totalScore += Draw
			case "C": // scissors
				totalScore += Loss
			default:
				panic("unacceptable opp choice " + l[0])
			}
		case "Z": // i played scissors
			switch l[0] {
			case "A": // rock
				totalScore += Loss
			case "B": // paper
				totalScore += Win
			case "C": // scissors
				totalScore += Draw
			default:
				panic("unacceptable opp choice " + l[0])
			}
		}
	}

	return totalScore
}

func part2(input string) int {
	/*
		second column is result, not your choice
		X -> you lose
		Y -> draw
		Z -> win
	*/
	lines := parseInput(input)

	winningScores := map[string]int{
		"X": Loss,
		"Y": Draw,
		"Z": Win,
	}

	totalScore := 0
	for _, l := range lines {
		if _, ok := winningScores[l[1]]; !ok {
			panic("unacceptable result " + l[1])
		}
		totalScore += winningScores[l[1]]
		// switch on opp choice instead
		switch l[0] {
		case "A": // opp: rock
			switch l[1] {
			case "X": // lose
				totalScore += Scissors
			case "Y": // draw
				totalScore += Rock
			case "Z": // win
				totalScore += Paper
			default:
				panic("unacceptable choice " + l[1])
			}
		case "B": // opp: paper
			switch l[1] {
			case "X": // lose
				totalScore += Rock
			case "Y": // draw
				totalScore += Paper
			case "Z": // win
				totalScore += Scissors
			default:
				panic("unacceptable choice " + l[1])
			}
		case "C": // opp: scissors
			switch l[1] {
			case "X": // lose
				totalScore += Paper
			case "Y": // draw
				totalScore += Scissors
			case "Z": // win
				totalScore += Rock
			default:
				panic("unacceptable choice " + l[1])
			}
		}
	}

	return totalScore
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, " "))
	}
	return ans
}
