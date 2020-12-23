package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := crabCups(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

// LLNode is a doubly linked list node
type LLNode struct {
	val         int
	left, right *LLNode
}

// ~1250/ 199
func crabCups(input string, part int) int {
	cups := strings.Split(input, "")

	currentCup := &LLNode{val: cast.ToInt(cups[0])}
	// LRU caches are a hashmap of integers pointing to a pointer of a (doubly) LL node
	// so finding a given node is constant time through the hash map
	lruCacheMap := map[int]*LLNode{currentCup.val: currentCup}

	// setup initial cups
	iter := currentCup
	for _, v := range cups[1:] {
		iter.right = &LLNode{
			val:  cast.ToInt(v),
			left: iter,
		}
		iter = iter.right
		lruCacheMap[iter.val] = iter
	}
	rounds := 100 // part 1 default

	// for part 2 make rounds higher and add a bunch of extra nodes
	if part == 2 {
		rounds = 10000000
		for i := 10; i <= 1000000; i++ {
			iter.right = &LLNode{
				val:  i,
				left: iter,
			}
			iter = iter.right
			lruCacheMap[iter.val] = iter
		}
	}

	iter.right = currentCup
	currentCup.left = iter

	for i := 0; i < rounds; i++ {
		valToFind := currentCup.val - 1

		ignoreVals := map[int]bool{}
		for it := currentCup.right; len(ignoreVals) < 3; it = it.right {
			ignoreVals[it.val] = true
		}

		// if the val to find is in ignore vals or is less than smallest possible
		// value, then loop and decrement until a valid value is found
		for ignoreVals[valToFind] || valToFind <= 0 {
			valToFind--
			// loop back to highest numbered cup if value goes below zero
			if valToFind <= 0 {
				valToFind = len(lruCacheMap)
			}
		}

		// next 3
		startOfThree := currentCup.right
		endOfThree := currentCup.right.right.right

		// remove the three off the end of the current cup
		currentCup.right = endOfThree.right
		currentCup.right.left = currentCup

		// find the LL node to attach the 3 cups to (directly from LRU cache)
		attachAfterMe := lruCacheMap[valToFind]
		// mash the 3 cups into the doubly LL
		endOfThree.right = attachAfterMe.right
		attachAfterMe.right.left = endOfThree
		attachAfterMe.right = startOfThree
		startOfThree.left = attachAfterMe

		// move currentCup pointer
		currentCup = currentCup.right
	}

	oneCup := lruCacheMap[1]

	if part == 2 {
		return oneCup.right.val * oneCup.right.right.val
	}

	var ans int
	for i := 0; i < 8; i++ {
		oneCup = oneCup.right
		ans *= 10
		ans += oneCup.val
	}

	return ans
}
