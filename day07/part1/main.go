package main

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode/day07/part1/intcode"
	"adventofcode/day07/part1/permutations"
	"adventofcode/util"
)

// input to first amp = 0
// output of each amp is input of next amp
// final output is from amp #5 / E to thrusters
func main() {
	readInput := util.ReadFile("../input.txt")
	strSplit := strings.Split(readInput, ",")

	input := make([]int, len(strSplit))
	for i, v := range strSplit {
		input[i], _ = strconv.Atoi(v)
	}

	// create all permutations of 0, 1, 2, 3, 4
	perms := permutations.CreatePermutations()
	// fmt.Println(perms)

	highestReturn := 0
	for _, onePerm := range perms {
		lastOutput := 0
		for _, firstInput := range onePerm {
			lastOutput = intcode.RunDiagnostics(input, firstInput, lastOutput)
		}
		if lastOutput > highestReturn {
			highestReturn = lastOutput
		}
	}
	fmt.Println("highestReturn is", highestReturn)
}
