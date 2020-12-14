package main

import (
	"flag"
	"fmt"
	"strings"

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

func part1(input string) int {
	groups := parseGroups(input)
	memory := map[int]int{}
	for _, group := range groups {
		for _, instruction := range group.instructions {
			address, value := instruction[0], instruction[1]
			// overwrite val w/ bit shifts and shit
			for index, overwrite := range group.mask {
				if overwrite == "1" {
					// run bitwise or
					value |= 1 << index
				} else if overwrite == "0" {
					// OR w/ a 1 -> 1; XOR w/ a 1 -> 0
					value |= 1 << index
					value ^= 1 << index
				}
			}

			// set the overwritten value into memory
			memory[address] = value
		}
	}

	return getTotal(memory)
}

func part2(input string) int {
	groups := parseGroups(input)
	memory := map[int]int{}
	for _, group := range groups {
		for _, instruction := range group.instructions {
			address, value := instruction[0], instruction[1]

			var floatingIndices []int
			for index, overwrite := range group.mask {
				if overwrite == "1" {
					address |= 1 << index
				} else if overwrite == "X" {
					floatingIndices = append(floatingIndices, index)
				}
			}

			// ugly way to get all the permutations (with one extra) of all possible addresses
			perms := []int{address}
			for _, index := range floatingIndices {
				for _, perm := range perms {
					// for each existing permutation, get the permutation with this floating index
					// flipped to a 1, and to a zero
					with1 := perm | 1<<index  // OR with a shifted 1 -> 1
					with0 := with1 ^ 1<<index // with1 XOR'ed with a shifted 1 -> 0
					perms = append(perms, with1, with0)
				}
			}
			// make writes to memory for each perm
			for _, index := range perms {
				memory[index] = value
			}
		}
	}

	return getTotal(memory)
}

// each grouping contains a mask at the start, then a pairs of memory addressed to values
type group struct {
	mask         map[int]string
	instructions [][2]int // memory address, overwrite value
}

func parseGroups(input string) []*group {
	var groups []*group

	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "mask") {
			// start a new group
			newGroup := &group{mask: map[int]string{}}
			groups = append(groups, newGroup)

			var maskValue string
			fmt.Sscanf(line, "mask = %s", &maskValue)

			// iterate backwards so that keys in the mask map will correspond to powers of two
			for i := len(maskValue) - 1; i >= 0; i-- {
				newGroup.mask[len(maskValue)-i-1] = string(maskValue[i])
			}
		} else {
			currentGroup := groups[len(groups)-1]
			var index, val int
			fmt.Sscanf(line, "mem[%d] = %d", &index, &val)
			currentGroup.instructions = append(currentGroup.instructions, [2]int{index, val})
		}
	}
	return groups
}

func getTotal(memory map[int]int) int {
	var sum int
	for _, v := range memory {
		sum += v
	}
	return sum
}
