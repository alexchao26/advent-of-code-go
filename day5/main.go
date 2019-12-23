package main

import "fmt"

func calcValue(input []int, target int, noun int, verb int) bool {
	// update the 1 and 2 values here
	input[1] = noun
	input[2] = verb

	// fmt.Println(input)

	// loop through all "instructions"
	for i := 0; i < len(input); i += 4 {
		// check opertor type (1 2 or 99)
		operator := input[i]
		// if it's not 99 (don't terminate, and run a calculation)
		if operator != 99 {
			// grab the two values to be added or multiplied
			value1 := input[input[i+1]]
			value2 := input[input[i+2]]

			// grab the index that will get the new value
			indexToUpdate := input[i+3]

			// perform calcuation based on operator value
			if operator == 1 {
				input[indexToUpdate] = value1 + value2
			} else if operator == 2 {
				input[indexToUpdate] = value1 * value2
			}
		} else {
			// if it is 99, break out of this loop
			break
		}
	}

	// if the target value is found, return true
	if input[0] == target {
		return true
	}
	return false
}

func main() {
	// defualt input slice, index 1 and 2 will be replaced
	input := []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 6, 19, 1, 19, 5, 23, 2, 13, 23, 27, 1, 10, 27, 31, 2, 6, 31, 35, 1, 9, 35, 39, 2, 10, 39, 43, 1, 43, 9, 47, 1, 47, 9, 51, 2, 10, 51, 55, 1, 55, 9, 59, 1, 59, 5, 63, 1, 63, 6, 67, 2, 6, 67, 71, 2, 10, 71, 75, 1, 75, 5, 79, 1, 9, 79, 83, 2, 83, 10, 87, 1, 87, 6, 91, 1, 13, 91, 95, 2, 10, 95, 99, 1, 99, 6, 103, 2, 13, 103, 107, 1, 107, 2, 111, 1, 111, 9, 0, 99, 2, 14, 0, 0}

	// brute force to try all options for nouns and verbs
outer:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			// crete a copy of the input slice
			clone := make([]int, 120)
			copy(clone, input)

			// if the calcValue function returns true, break out of the loops and print the values to the console
			if calcValue(clone, 19690720, i, j) == true {
				// print answers to console (manually add to advent of code)
				fmt.Println("noun is", i, "verb is", j)
				fmt.Println("actual result value noun * 10 + verb = ", i*100+j)

				// use labeled outer loop to break out of both for loops
				break outer
			}
		}
	}
}
