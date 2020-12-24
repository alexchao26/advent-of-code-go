package main

import (
	"flag"
	"fmt"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := elephant(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

// LLNode represents an elf
type LLNode struct {
	elfNum   int
	presents int
	next     *LLNode
}

func elephant(input string, part int) int {
	startingElves := cast.ToInt(input)
	root := &LLNode{
		elfNum:   1,
		presents: 1,
	}
	iter := root
	for i := 2; i <= startingElves; i++ {
		iter.next = &LLNode{
			elfNum:   i,
			presents: 1,
		}
		iter = iter.next
	}
	iter.next = root

	if part == 1 {
		for root.next != root {
			root.presents += root.next.presents
			root.next = root.next.next
			root = root.next
		}
		return root.elfNum
	}

	// initialize a pointer to the node before the node across from the start
	// need the node before b/c removing a node is like reassigning beforeNode.next
	// if there are an odd number of starting elves, this points to the first one
	// which is also the one of "the left."
	isOddLength := startingElves%2 == 1
	beforeAcross := root
	for i := 0; i < startingElves/2-1; i++ {
		beforeAcross = beforeAcross.next
	}

	for root.next != root {
		root.presents += beforeAcross.next.presents
		// remove beforeAcross node
		beforeAcross.next = beforeAcross.next.next
		// if odd number of total nodes, beforeAcross node skips the node
		// that was previously the "right" side of the across pair
		if isOddLength {
			beforeAcross = beforeAcross.next
		}
		isOddLength = !isOddLength
		root = root.next
	}

	return root.elfNum
}
