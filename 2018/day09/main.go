package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	players, lastPoints := parseInput(input)

	currentMarble := &CircularLinkedListNode{nil, nil, 0}
	currentMarble.left = currentMarble
	currentMarble.right = currentMarble

	playerTurn := 0
	playerScores := make([]int, players)

	for points := 1; points <= lastPoints; points++ {
		if points%23 == 0 {
			for i := 0; i < 7; i++ {
				currentMarble = currentMarble.left
			}

			playerScores[playerTurn] += points
			playerScores[playerTurn] += currentMarble.val
			currentMarble = currentMarble.RemoveSelf()
		} else {
			currentMarble = currentMarble.right
			currentMarble.AddToRight(points)
			currentMarble = currentMarble.right
		}

		playerTurn++
		playerTurn %= players
	}

	return util.MaxInts(playerScores...)
}

func part2(input string) int {
	// lazily modify input...
	split := strings.Split(input, " ")
	steps := util.StrToInt(split[6]) * 100
	split[6] = strconv.Itoa(steps)

	return part1(strings.Join(split, " "))
}

func parseInput(input string) (players int, lastPoints int) {
	split := strings.Split(input, " ")
	return util.StrToInt(split[0]), util.StrToInt(split[6])
}

type CircularLinkedListNode struct {
	left, right *CircularLinkedListNode
	val         int
}

func (c *CircularLinkedListNode) AddToRight(val int) {
	newNode := CircularLinkedListNode{
		left:  c,
		right: c.right,
		val:   val,
	}
	c.right.left = &newNode
	c.right = &newNode
}

func (c *CircularLinkedListNode) RemoveSelf() *CircularLinkedListNode {
	// not handling edge case of < 2 nodes because it's not going to get hit in this case
	c.left.right = c.right
	c.right.left = c.left
	return c.right
}
