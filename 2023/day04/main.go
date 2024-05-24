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

	var ans int

	for _, card := range parsed {
		ans += scoreCard(card)
	}

	return ans
}

func scoreCard(c card) int {
	var ans int

	for _, num := range c.hand {
		if c.winning[num] {
			if ans == 0 {
				ans = 1
			} else {
				ans *= 2
			}
		}
	}

	return ans
}

func part2(input string) int {
	cards := parseInput(input)

	// tracks the number of cards won, starts with 1 of each card
	numCards := make([]int, len(cards))
	for i := range numCards {
		numCards[i] = 1
	}

	for index, c := range cards {
		cardsWon := countWinningNumbers(c)
		for i := 1; i <= cardsWon; i++ {
			// add number of current card to account for previous wins
			numCards[index+i] += numCards[index]
		}
	}

	// add up total number of cards
	var cardCount int
	for _, n := range numCards {
		cardCount += n
	}
	return cardCount
}

func countWinningNumbers(c card) int {
	var ans int

	for _, num := range c.hand {
		if c.winning[num] {
			ans++
		}
	}

	return ans
}

type card struct {
	// index   int // unused
	winning map[int]bool
	hand    []int
}

func parseInput(input string) (ans []card) {
	for _, line := range strings.Split(input, "\n") {
		c := card{
			// index:   len(ans) + 1,
			winning: map[int]bool{},
		}

		half := strings.Split(line, ": ")
		numParts := strings.Split(half[1], " | ")
		for _, winningNum := range strings.Split(numParts[0], " ") {
			// handles single digits that have an extra empty string between nums
			if winningNum == "" {
				continue
			}
			c.winning[cast.ToInt(winningNum)] = true
		}

		for _, handNum := range strings.Split(numParts[1], " ") {
			if handNum == "" {
				continue
			}
			c.hand = append(c.hand, cast.ToInt(handNum))
		}
		ans = append(ans, c)
	}
	return ans
}
