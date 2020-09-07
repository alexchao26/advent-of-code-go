package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func main() {
	input := util.ReadFile("../input.txt")
	lines := strings.Split(input, "\n")

	var twos, threes int
	for _, boxID := range lines {
		if hasTwo(boxID) {
			twos++
		}
		if hasThree(boxID) {
			threes++
		}
	}

	fmt.Println("checksum", twos*threes)
}

func hasTwo(box string) bool {
	chars := make(map[rune]int)
	for _, c := range box {
		chars[c]++
	}
	for _, v := range chars {
		if v == 2 {
			return true
		}
	}
	return false
}

func hasThree(box string) bool {
	chars := make(map[rune]int)
	for _, c := range box {
		chars[c]++
	}
	for _, v := range chars {
		if v == 3 {
			return true
		}
	}
	return false
}
