package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

/**
This implementation is too tangled to cover part 1 and 2 because the sort
logic is encapsulated in the interface which cannot (nicely) be made aware
of which part is being run...
There are two main differences highlighted by "// NOTE: diff to part 1 here"
*/

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

	ans := part2(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func part2(input string) int {
	allPlayers := parseInput(input, 1)

	sort.Sort(allPlayers)

	var ans int
	for i, player := range allPlayers {
		ans += int(player.bid) * int(i+1)
	}
	return ans
}

type player struct {
	hand string
	bid  int

	handTypeScore int
}

func parseInput(input string, part int) (ans SortablePlayers) {
	for _, line := range strings.Split(input, "\n") {
		handAndBid := strings.Split(line, " ")
		pl := player{
			hand: handAndBid[0],
			bid:  cast.ToInt(handAndBid[1]),
		}
		pl.handTypeScore = scoreHandType(pl, part)
		ans = append(ans, pl)
	}
	return ans
}

/*	point assignments:
	five of a kind -> 7
	four of a kind -> 6
	full house -> 5
	three of a kind -> 4
	two pair -> 3
	one pair -> 2
	high card -> 1
*/

func scoreHandType(p player, part int) int {
	counts := map[string]int{}
	for _, card := range strings.Split(p.hand, "") {
		counts[card]++
	}

	// NOTE: diff to part 1 here
	if counts["J"] != 0 {
		// add the J counts to the highest card's count as that will always result in the best hand?
		jCount := counts["J"]
		// remove "J" from map in case if it's the highest count
		delete(counts, "J")

		var highestCard string
		var highestCount int
		for card, ct := range counts {
			if ct > highestCount {
				highestCount = ct
				highestCard = card
			}
		}

		counts[highestCard] += jCount
	}

	// high card
	if len(counts) == 5 {
		return 1
	}
	// one pair
	if len(counts) == 4 {
		return 2
	}

	if len(counts) == 3 {
		// either two pair or three of a kind
		for _, ct := range counts {
			if ct == 3 {
				return 4 // 3 of a kind, 3 1 1
			}
		}
		return 3 // two pair, 2 2 1
	}

	if len(counts) == 2 {
		// full house (3 2) or four of a kind (4 1)
		for _, ct := range counts {
			if ct == 3 {
				return 5
			}
		}
		return 6
	}

	if len(counts) == 1 {
		return 7
	}

	panic(fmt.Sprintf("error scoring hand: %+v", p.hand))
}

type SortablePlayers []player

func (ps SortablePlayers) Len() int { return len(ps) }
func (ps SortablePlayers) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func (ps SortablePlayers) Less(i, j int) bool {
	iTypeScore := ps[i].handTypeScore
	jTypeScore := ps[j].handTypeScore

	if iTypeScore == jTypeScore {
		// higher score goes to end of ps slice
		return !doesPlayer1WinTiebreak(ps[i], ps[j])
	}
	return iTypeScore < jTypeScore
}

var cardValuesMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	// can just leave the same values though, the heiarchy is what matters
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	// NOTE: diff to part 1 here: J is the worst card now
	"J": 0,
}

// returns true if player1 wins the tie break
func doesPlayer1WinTiebreak(p1, p2 player) bool {
	if p1.handTypeScore != p2.handTypeScore {
		panic("p1 and p2 scores do not have the same level hand")
	}
	p1Cards := strings.Split(p1.hand, "")
	p2Cards := strings.Split(p2.hand, "")
	for i := 0; i < len(p1Cards); i++ {
		if cardValuesMap[p1Cards[i]] > cardValuesMap[p2Cards[i]] {
			return true
		} else if cardValuesMap[p1Cards[i]] < cardValuesMap[p2Cards[i]] {
			return false
		}
	}
	panic("not expecting to have two matching hands")
}
