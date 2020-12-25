package main

import (
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	ans := part1(util.ReadFile("./input.txt"))
	fmt.Println("Output:", ans)
}

// ~1900..
func part1(input string) int {
	var publicKeys []int
	for _, l := range strings.Split(input, "\n") {
		publicKeys = append(publicKeys, cast.ToInt(l))
	}

	var loopSizes [2]int
	val := 1
	// calculate both loop sizes in one pass to avoid recalculating over and over...
	for loops := 1; loopSizes[0] == 0 || loopSizes[1] == 0; loops++ {
		val *= 7
		val %= 20201227
		if val == publicKeys[0] {
			loopSizes[0] = loops
		}
		if val == publicKeys[1] {
			loopSizes[1] = loops
		}
	}

	// ensure both loop sizes are correct by reaching the same encryption key
	key1 := runTranformations(publicKeys[0], loopSizes[1])
	key2 := runTranformations(publicKeys[1], loopSizes[0])
	if key1 != key2 {
		panic(fmt.Sprintf("encryption keys should be the same, got %d and %d", key1, key2))
	}

	return key1
}

func runTranformations(subject, loops int) int {
	val := 1
	for i := 0; i < loops; i++ {
		val *= subject
		val %= 20201227
	}
	return val
}
