package main

import "fmt"

func main() {
	fmt.Println(lcm([]int{2, 4, 8}), " -> expect 8")
	fmt.Println(lcm([]int{25, 15, 18}), " -> expect 450")
}

// Better solution to finding a least common multiple
func lcm(sNum []int) int {
	ans := 1

	// start iterating from the number 2
	for i := 2; ; {
		// booleans to track if a number has been added to the ans (LCM)
		// and if all of the values in the slice are 1, in which case we are done
		changeMade, allOnes := false, true
		// iterate over all elements in the sNum slice
		for index, num := range sNum {
			if num%i == 0 && !changeMade {
				changeMade = true // update this boolean flag
				ans *= i          // increment answer
				sNum[index] /= i  // need to use this notation because num is pass by value
				// do not increment i
			} else if num%i == 0 {
				// need to divide the element, but not increment ans (same prime factor as a previous element)
				sNum[index] /= i
			}

			// if any of the values in the slice are not one, flip the allOnes flag to false
			// to indicate that we're not done with our outer loop
			if num != 1 {
				allOnes = false
			}
		}

		// if all values in sNum are 1, we're done & can return the answer
		if allOnes {
			return ans
		} else if !changeMade {
			// if a change was not made & not all values are one, increment i
			i++
		}
	}
}
