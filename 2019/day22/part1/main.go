package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

type instruction struct {
	name   string // name of instruction
	number int    // number
}

func main() {
	input := util.ReadFile("../input.txt")
	lines := strings.Split(input, "\n")

	// make slice of all instructions
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		splitLine := strings.Split(line, " ")
		if splitLine[len(splitLine)-1] == "stack" {
			instructions[i] = instruction{name: "deal into new stack"}
		} else {
			name := strings.Join(splitLine[0:len(splitLine)-1], " ")
			num, _ := strconv.Atoi(splitLine[len(splitLine)-1])

			instructions[i] = instruction{
				name,
				num,
			}
		}
	}

	// make the deck
	deck := make([]int, 10007)
	for i := range deck {
		deck[i] = i
	}

	// iterate through instructions and apply the correct function to the deck
	for _, inst := range instructions {
		switch inst.name {
		case "deal into new stack":
			deck = dealIntoNewStack(deck)
		case "deal with increment":
			deck = dealWithIncrement(deck, inst.number)
		case "cut":
			deck = cut(deck, inst.number)
		}
	}

	// Find the 2019 card and print its index
	for i, v := range deck {
		if v == 2019 {
			fmt.Println("position of card 2019 is", i)
		}
	}
}

// side effects to handle "deal into new stack" command
func dealIntoNewStack(deck []int) []int {
	// effectively just reverses the entire deck..
	for i := 0; i < len(deck)/2; i++ {
		deck[i], deck[len(deck)-1-i] = deck[len(deck)-1-i], deck[i]
	}

	// technically changed by side effects but will keep syntax similar
	return deck
}

// index could be negative
func cut(deck []int, index int) []int {
	newDeck := make([]int, len(deck))

	if index < 0 {
		index = len(deck) + index
	}

	for i := 0; i < len(deck); i++ {
		newDeck[i] = deck[(i+index)%len(deck)]
	}
	return newDeck
}

func dealWithIncrement(deck []int, jump int) []int {
	newDeck := make([]int, len(deck))

	for i := 0; i < len(deck); i++ {
		newDeckIndex := jump * i % len(deck)
		newDeck[newDeckIndex] = deck[i]
	}

	return newDeck
}
