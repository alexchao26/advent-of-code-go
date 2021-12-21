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
	positions := parseInput(input)

	// die, two pawns, circular track w/ 10 spaces (1-10)

	// players take turns
	// roll die 3 times, take sum
	// move forward that number of spaces
	// increase score by space they stop on
	// game ends when someone has 1000 points

	// deterministic die, 100 sided, rolls 1, then 2, 3, 4, 5, 6, etc
	die := 0 // starts at zero, first time it rolls will produce a 1
	isPlayer1sTurn := true
	var scores [2]int
	var rolls int
	for scores[0] < 1000 && scores[1] < 1000 {
		var threeRolls int
		for i := 0; i < 3; i++ {
			rolls++
			die++
			if die == 101 {
				die = 1
			}
			threeRolls += die
		}

		if isPlayer1sTurn {
			positions[0] += threeRolls
			for positions[0] > 10 {
				positions[0] -= 10
			}
			scores[0] += positions[0]
		} else {
			positions[1] += threeRolls
			for positions[1] > 10 {
				positions[1] -= 10
			}
			scores[1] += positions[1]
		}
		// switch turns
		isPlayer1sTurn = !isPlayer1sTurn
	}

	loser := scores[0]
	if scores[1] < 1000 {
		loser = scores[1]
	}

	// final answer is the losing player's score TIMES number of times the die was rolled
	return loser * rolls
}

func part2(input string) int64 {
	// who wins in more universes
	// dirac die, 3 sided, splits into 3 copies, one for each result, on EVERY roll
	positions := parseInput(input)
	w1, w2 := play([2]int{positions[0], positions[1]}, [2]int{}, 3, true, map[string][2]int64{})

	if w1 > w2 {
		return w1
	}
	return w2
}

func play(positions, scores [2]int, rollsLeftInTurn int, isPlayer1sTurn bool, memo map[string][2]int64) (wins1, wins2 int64) {
	key := fmt.Sprint(positions, scores, rollsLeftInTurn, isPlayer1sTurn)
	if res, ok := memo[key]; ok {
		return res[0], res[1]
	}

	// NOTE 0-indexed array, so player 2 is at index 1...
	playerIndex := 1
	if isPlayer1sTurn {
		playerIndex = 0
	}

	scoresCopy := [2]int{scores[0], scores[1]}
	if rollsLeftInTurn == 0 {
		scoresCopy[playerIndex] += positions[playerIndex]
		// TERMINATION CASE
		if scoresCopy[playerIndex] >= 21 {
			if playerIndex == 0 {
				return 1, 0
			}
			return 0, 1
		}

		isPlayer1sTurn = !isPlayer1sTurn
		rollsLeftInTurn = 3
		// update playerIndex because they're switching now! (not getting that hour of my life back...)
		playerIndex++
		playerIndex %= 2
	}

	for roll := 1; roll <= 3; roll++ {
		// recurse with a given roll
		// copy positions so each recurse gets its own copy
		positionsCopy := [2]int{positions[0], positions[1]}
		positionsCopy[playerIndex] += roll
		if positionsCopy[playerIndex] > 10 {
			positionsCopy[playerIndex] -= 10
		}
		r1, r2 := play(positionsCopy, scoresCopy, rollsLeftInTurn-1, isPlayer1sTurn, memo)
		wins1 += r1
		wins2 += r2
	}

	memo[key] = [2]int64{wins1, wins2}
	return wins1, wins2
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, "\n") {
		// Player 1 starting position: 5
		var player, startingPosition int
		fmt.Sscanf(line, "Player %d starting position: %d", &player, &startingPosition)
		ans = append(ans, startingPosition)
	}
	return ans
}
