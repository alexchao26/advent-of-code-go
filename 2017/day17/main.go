package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := spinlock(util.ReadFile("./input.txt"), part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

type llNode struct {
	value int
	next  *llNode
}

func spinlock(input string, part int) int {
	steps := cast.ToInt(input)

	lastNumToAdd := 2017
	if part == 2 {
		lastNumToAdd = 50000000
	}

	current := &llNode{value: 0}
	current.next = current
	for i := 1; i <= lastNumToAdd; i++ {
		for j := 0; j < steps; j++ {
			current = current.next
		}

		saveNext := current.next
		current.next = &llNode{
			value: i,
			next:  saveNext,
		}
		current = current.next

		// progress log for part 2 brute force... this is SLOW
		if i%1000000 == 0 {
			log.Println(i, "steps done")
		}
	}

	// iterate to the node to find, then return the value of the next node
	valueToFind := 2017
	if part == 2 {
		valueToFind = 0
	}
	for current.value != valueToFind {
		current = current.next
	}

	return current.next.value
}
