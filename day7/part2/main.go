package main

import (
	"fmt"

	"./intcode"
	"./permutations"
)

func main() {
	// input to first amp = 0
	// output of each amp is input of next amp
	// final output is from amp #5 / E to thrusters

	// input := []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 42, 55, 64, 77, 94, 175, 256, 337, 418, 99999, 3, 9, 102, 4, 9, 9, 1001, 9, 5, 9, 102, 2, 9, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 1001, 9, 3, 9, 1002, 9, 5, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99}

	// should give 18216
	input := []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}

	// create all permutations of 5, 6, 7, 8, 9
	perms := permutations.CreatePermutations(5, 9)
	// fmt.Println(perms)

	highestReturn := 0
	for _, onePerm := range perms {
		// needs to have the right length to copy into
		input1 := make([]int, len(input))
		input2 := make([]int, len(input))
		input3 := make([]int, len(input))
		input4 := make([]int, len(input))
		input5 := make([]int, len(input))

		copy(input1, input)
		copy(input2, input)
		copy(input3, input)
		copy(input4, input)
		copy(input5, input)

		// fmt.Println(input1)
		// fmt.Println(onePerm)

		last5Return := 0

		index1, _, lastOutput1 := intcode.RunDiagnostics(input1, onePerm[0], 0)
		index2, _, lastOutput2 := intcode.RunDiagnostics(input2, onePerm[1], 0)
		index3, _, lastOutput3 := intcode.RunDiagnostics(input3, onePerm[2], 0)
		index4, _, lastOutput4 := intcode.RunDiagnostics(input4, onePerm[3], 0)
		index5, exitCode5, lastOutput5 := intcode.RunDiagnostics(input5, onePerm[4], 0)

		// fmt.Println(lastOutput5)
		for exitCode5 != 99 {
			index1, _, lastOutput1 = intcode.RunDiagnostics(input1, lastOutput5, index1)
			index2, _, lastOutput2 = intcode.RunDiagnostics(input2, lastOutput1, index2)
			index3, _, lastOutput3 = intcode.RunDiagnostics(input3, lastOutput2, index3)
			index4, _, lastOutput4 = intcode.RunDiagnostics(input4, lastOutput3, index4)
			index5, exitCode5, lastOutput5 = intcode.RunDiagnostics(input5, lastOutput4, index5)
			if exitCode5 == 4 {
				last5Return = lastOutput5
			} else if exitCode5 == 99 {
				fmt.Println(lastOutput5)
			}
		}

		if last5Return > highestReturn {
			highestReturn = last5Return
		}
		// fmt.Println(lastOutput5)
	}
	fmt.Println("highestReturn is", highestReturn)
}
