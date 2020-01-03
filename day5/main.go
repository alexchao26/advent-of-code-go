package main

import "fmt"

// func calcValue(input []int, target int, noun int, verb int) bool {
// 	// update the 1 and 2 values here
// 	input[1] = noun
// 	input[2] = verb

// 	// fmt.Println(input)

// 	// loop through all "instructions"
// 	for i := 0; i < len(input); i += 4 {
// 		// check opertor type (1 2 or 99)
// 		operator := input[i]
// 		// if it's not 99 (don't terminate, and run a calculation)
// 		if operator != 99 {
// 			// grab the two values to be added or multiplied
// 			value1 := input[input[i+1]]
// 			value2 := input[input[i+2]]

// 			// grab the index that will get the new value
// 			indexToUpdate := input[i+3]

// 			// perform calcuation based on operator value
// 			if operator == 1 {
// 				input[indexToUpdate] = value1 + value2
// 			} else if operator == 2 {
// 				input[indexToUpdate] = value1 * value2
// 			}
// 		} else {
// 			// if it is 99, break out of this loop
// 			break
// 		}
// 	}

// 	// if the target value is found, return true
// 	if input[0] == target {
// 		return true
// 	}
// 	return false
// }

func returnCodes(number int) []int {
	ans := make([]int, 4)
	ans[0] = number % 100
	ans[1] = (number%1000 - ans[0]) / 100
	ans[2] = (number%10000 - number%1000) / 1000
	ans[3] = (number%100000 - number%10000) / 10000

	return ans
}

func getValue(puzzleInput []int, positionOrImmediateCode int, val int) int {
	if positionOrImmediateCode == 0 {
		return puzzleInput[val]
	}
	return val
}

func runDiagnostics(puzzleInput []int, inputValue int) string {
	fmt.Println(puzzleInput)

	for i := 0; i < len(puzzleInput); {
		// find op code (last 2 digits of number), a 1, 2, 3, 4, or 99
		slicedInstruction := returnCodes(puzzleInput[i])
		opCode := slicedInstruction[0]
		val1 := slicedInstruction[1]
		val2 := slicedInstruction[2]
		val3 := slicedInstruction[3] // defaults to position so no longer need this variable

		fmt.Println("opCode", opCode, "index", i)
		// fmt.Println("68 slice", puzzleInput[69])

		if opCode == 99 {
			fmt.Println("99 halted")
			return "halted"
		} else if opCode == 1 {
			// additon of next two elements( i + 1 and 2), placed @ or @ location of next element (i + 3)
			firstToAdd, secondToAdd := getValue(puzzleInput, val1, puzzleInput[i+1]), getValue(puzzleInput, val2, puzzleInput[i+2])
			// if val3 == 0 {
			// position mode
			fmt.Println("params are", val1, val2, val3)
			fmt.Println("adding", firstToAdd, secondToAdd, "to position", puzzleInput[i+3])
			puzzleInput[puzzleInput[i+3]] = firstToAdd + secondToAdd
			// } else {
			// immediately to that spot
			// puzzleInput[i+3] = firstToAdd + secondToAdd
			// }
			i += 4
		} else if opCode == 2 {
			// multiply
			firstToMulitply, secondToMulitply := getValue(puzzleInput, val1, puzzleInput[i+1]), getValue(puzzleInput, val2, puzzleInput[i+2])
			// if val3 == 0 {
			// position mode
			fmt.Println("params are", val1, val2, val3)
			fmt.Println("multiplying", firstToMulitply, secondToMulitply, "to position", puzzleInput[i+3])
			puzzleInput[puzzleInput[i+3]] = firstToMulitply + secondToMulitply
			// } else {
			// immediately to that spot
			// puzzleInput[i+3] = firstToMulitply + secondToMulitply
			// }
			i += 4
		} else if opCode == 3 {
			// fmt.Println("A 3 opCode has fired", i)
			// one more input, val1 is position or immediate mode
			indexPlusOne := puzzleInput[i+1]
			// if val1 == 0 {
			// position mode
			fmt.Println("writing input value", inputValue, "to position", indexPlusOne)
			puzzleInput[indexPlusOne] = inputValue
			// } else {
			// immediate mode
			// 	puzzleInput[i+1] = inputValue
			// }
			i += 2
		} else if opCode == 4 {
			// return a value opCode, returns value @ next position, or the next value itself
			indexPlusOne := puzzleInput[i+1]
			fmt.Println("output val is", val1)
			if val1 == 0 {
				// position mode
				fmt.Println("******************output for val1==0: ", puzzleInput[indexPlusOne])
				// if puzzleInput[indexPlusOne] != 0 {
				// 	fmt.Println("non-zero output", opCode, i, indexPlusOne, puzzleInput[indexPlusOne])
				// }
			} else {
				// immediate mode - this will never run?
				fmt.Println("******************output immediate val!=0: ", indexPlusOne)
			}
			i += 2
		}
	}
	return "EOF"
}

func main() {
	// defualt input slice, index 1 and 2 will be replaced
	puzzleInput := []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1, 192, 154, 224, 101, -161, 224, 224, 4, 224, 102, 8, 223, 223, 101, 5, 224, 224, 1, 223, 224, 223, 1001, 157, 48, 224, 1001, 224, -61, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1102, 15, 28, 225, 1002, 162, 75, 224, 1001, 224, -600, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 224, 223, 223, 102, 32, 57, 224, 1001, 224, -480, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 224, 223, 223, 1101, 6, 23, 225, 1102, 15, 70, 224, 1001, 224, -1050, 224, 4, 224, 1002, 223, 8, 223, 101, 5, 224, 224, 1, 224, 223, 223, 101, 53, 196, 224, 1001, 224, -63, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 224, 223, 223, 1101, 64, 94, 225, 1102, 13, 23, 225, 1101, 41, 8, 225, 2, 105, 187, 224, 1001, 224, -60, 224, 4, 224, 1002, 223, 8, 223, 101, 6, 224, 224, 1, 224, 223, 223, 1101, 10, 23, 225, 1101, 16, 67, 225, 1101, 58, 10, 225, 1101, 25, 34, 224, 1001, 224, -59, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 329, 101, 1, 223, 223, 107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 344, 1001, 223, 1, 223, 107, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 359, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 374, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 389, 101, 1, 223, 223, 1007, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 404, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 419, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 1001, 223, 1, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 449, 101, 1, 223, 223, 108, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 464, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 479, 1001, 223, 1, 223, 1007, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 494, 101, 1, 223, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 509, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 539, 1001, 223, 1, 223, 1107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 1001, 223, 1, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 8, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 584, 101, 1, 223, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 101, 1, 223, 223, 1007, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 614, 1001, 223, 1, 223, 8, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 644, 101, 1, 223, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 659, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 674, 1001, 223, 1, 223, 4, 223, 99, 226}

	// fmt.Println(puzzleInput[225])
	// fmt.Println(puzzleInput)

	// 1 as the input?
	runDiagnostics(puzzleInput, 1)

	// fmt.Println(puzzleInput)

	// brute force to try all options for nouns and verbs
	// outer:
	// 	for i := 0; i < 100; i++ {
	// 		for j := 0; j < 100; j++ {
	// crete a copy of the input slice
	// clone := make([]int, 120)
	// copy(clone, input)

	// fmt.Println(clone)
	// if the calcValue function returns true, break out of the loops and print the values to the console
	// if calcValue(clone, 19690720, i, j) == true {
	// print answers to console (manually add to advent of code)
	// fmt.Println("noun is", i, "verb is", j)
	// fmt.Println("actual result value noun * 10 + verb = ", i*100+j)

	// use labeled outer loop to break out of both for loops
	// 			break outer
	// 		}
	// 	}
	// }
}
