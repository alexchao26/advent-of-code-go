package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

// leaderboard: 537
func part1(input string) int {
	deck1, deck2 := parseInput(input)

	for !(len(deck1) == 0 || len(deck2) == 0) {
		top1, top2 := deck1[0], deck2[0]
		if top1 > top2 {
			deck1 = append(deck1, top1, top2)
		} else {
			deck2 = append(deck2, top2, top1)
		}
		deck1 = deck1[1:]
		deck2 = deck2[1:]
	}

	winningDeck := append(deck1, deck2...)
	var sumOfProducts int
	multiplier := 1
	for i := len(winningDeck) - 1; i >= 0; i-- {
		sumOfProducts += multiplier * winningDeck[i]
		multiplier++
	}

	return sumOfProducts
}

func part2(input string) int {
	deck1, deck2 := parseInput(input)
	winningScore, _ := recursiveGame(deck1, deck2, true)
	return winningScore
}

// leaderboard: 997
func recursiveGame(deck1, deck2 []int, isMainGame bool) (finalScore int, player1Wins bool) {
	// after the fact optimization from: https://www.reddit.com/r/adventofcode/comments/khyjgv/2020_day_22_solutions/ggpcsnd
	// reduces part2 time from ~30s to <0.5s
	// IF player 1 has the largest present card AND the card's value is greater
	// than the largest number of cards that a subgame could contain (lengths - 2)
	// THEN a subgame will never start when player1's top card is that max card,
	//   and player 1 can never lose that card, so at some point, a pattern will
	//   repeat which leads to player 1 winning
	if !isMainGame {
		max1, max2 := mathy.MaxInt(deck1...), mathy.MaxInt(deck2...)
		if max1 > max2 && max1 >= len(deck1)+len(deck2)-2 {
			return 0, true
		}
	}

	previousHands1 := map[string]bool{}
	previousHands2 := map[string]bool{}

	for !(len(deck1) == 0 || len(deck2) == 0) {
		top1, top2 := deck1[0], deck2[0]

		if previousHands1[fmt.Sprintf("%v", deck1)] || previousHands2[fmt.Sprintf("%v", deck2)] {
			player1Wins = true
		} else {
			previousHands1[fmt.Sprintf("%v", deck1)] = true
			previousHands2[fmt.Sprintf("%v", deck2)] = true

			// if not enough cards in either deck, just compare cards
			if top1 > len(deck1)-1 || top2 > len(deck2)-1 {
				player1Wins = top1 > top2
			} else {
				// otherwise recurse
				_, player1Wins = recursiveGame(append([]int{}, deck1[1:top1+1]...), append([]int{}, deck2[1:top2+1]...), false)
			}
		}

		if player1Wins {
			deck1 = append(deck1, top1, top2)
		} else {
			deck2 = append(deck2, top2, top1)
		}

		deck1 = deck1[1:]
		deck2 = deck2[1:]
	}

	if !isMainGame {
		// player1Wins boolean is equivalent to if their deck does not have zero cards
		return 0, len(deck1) != 0
	}

	winningDeck := append(deck1, deck2...)
	var sumOfProducts int
	multiplier := 1
	for i := len(winningDeck) - 1; i >= 0; i-- {
		sumOfProducts += multiplier * winningDeck[i]
		multiplier++
	}

	return sumOfProducts, false // 997
}

func parseInput(input string) ([]int, []int) {
	players := strings.Split(input, "\n\n")
	var deck1, deck2 []int
	for _, l := range strings.Split(players[0], "\n")[1:] {
		deck1 = append(deck1, cast.ToInt(l))
	}
	for _, l := range strings.Split(players[1], "\n")[1:] {
		deck2 = append(deck2, cast.ToInt(l))
	}
	return deck1, deck2
}
