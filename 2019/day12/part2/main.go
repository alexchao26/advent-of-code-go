package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

// Moon stores coordinates and velocities of a moon
type Moon struct {
	x, y, z, dx, dy, dz int
}
type oneDim struct {
	m, dm int
}

func main() {
	input := util.ReadFile("../input.txt")
	stringSlice := strings.Split(input, "\n")

	// need to set x y z, dx, dy, dz of each of the starting moons
	sliceMoons := makeMoonSlice(stringSlice)
	// fmt.Println(sliceMoons)

	// manually make the three dimension slices, to be passed into stringify & iterate helper functions
	xDim := make([]oneDim, 0)
	yDim := make([]oneDim, 0)
	zDim := make([]oneDim, 0)

	for _, m := range sliceMoons {
		xDim = append(xDim, oneDim{m.x, m.dx})
		yDim = append(yDim, oneDim{m.y, m.dy})
		zDim = append(zDim, oneDim{m.z, m.dz})
	}

	// fmt.Println(xDim, yDim, zDim)

	// "stringify" them so they are easy to compare later
	initialX, initialY, initialZ := stringifyOneDim(xDim), stringifyOneDim(yDim), stringifyOneDim(zDim)

	// find the number of steps for each dimension to reach it's initial position & velocity
	xSteps, ySteps, zSteps := iterate(xDim, initialX), iterate(yDim, initialY), iterate(zDim, initialZ)
	// fmt.Println(xSteps, ySteps, zSteps)

	// print the final least common multiple of the three number of steps
	fmt.Println(lcm(xSteps, ySteps, zSteps))
	fmt.Println(lcm2([]int{xSteps, ySteps, zSteps}))
}

// helper function to take the slice of strings and return a slice of Moon structs
func makeMoonSlice(stringSlice []string) []Moon {
	sliceMoons := make([]Moon, 0)
	for _, str := range stringSlice {
		xStart := strings.Index(str, "x=") + 2
		xEnd := strings.Index(str, ",")
		yStart := xEnd + 4
		zStart := strings.Index(str, "z=") + 2
		yEnd := zStart - 4
		zEnd := len(str) - 1

		x := str[xStart:xEnd]
		y := str[yStart:yEnd]
		z := str[zStart:zEnd]

		intX, _ := strconv.Atoi(x)
		intY, _ := strconv.Atoi(y)
		intZ, _ := strconv.Atoi(z)
		sliceMoons = append(sliceMoons, Moon{intX, intY, intZ, 0, 0, 0})
	}
	return sliceMoons
}

// helper function that updates the velocity then coordinate of a slice of oneDim structs
func updateVelThenCoords(sliceOneDim []oneDim) {
	for start := 0; start < len(sliceOneDim); start++ {
		// update velocity
		// requires iterating through all of the moons again...
		for restIndex := start + 1; restIndex < len(sliceOneDim); restIndex++ {
			if sliceOneDim[start].m < sliceOneDim[restIndex].m {
				sliceOneDim[start].dm++
				sliceOneDim[restIndex].dm--
			} else if sliceOneDim[start].m > sliceOneDim[restIndex].m {
				sliceOneDim[start].dm--
				sliceOneDim[restIndex].dm++
			}
		}
	}

	// then update coordinates x y z
	for i2, e := range sliceOneDim {
		sliceOneDim[i2].m += e.dm
	}
}

// helper function that will stringify a slice of oneDims to compare its values to another
func stringifyOneDim(sliceOneDim []oneDim) (result string) {
	for _, m := range sliceOneDim {
		result += strconv.Itoa(m.m) + ","
		result += strconv.Itoa(m.dm) + ","
	}
	return result
}

// helper function that will return the number of steps until the initial state is reached
// uses string comparison and the stringifyOneDim helper function
func iterate(dims []oneDim, initialString string) int {
	for i := 0; ; i++ {
		updateVelThenCoords(dims)
		if stringifyOneDim(dims) == initialString {
			return i + 1
		}
	}
}

// helper function that returns the least common multiple of three integers
func lcm(x, y, z int) int {
	pFactX, pFactY, pFactZ := primeFactorization(x), primeFactorization(y), primeFactorization(z)
	fmt.Println(pFactX, pFactY, pFactZ)

	ans := 1

	// multiple by every value in every slice, but do not count duplicates
	for i, j, k := 0, 0, 0; i < len(pFactX) || j < len(pFactY) || k < len(pFactZ); {
		if i < len(pFactX) {
			ans *= pFactX[i]
			if pFactX[i] == pFactY[j] {
				j++
			}
			if pFactX[i] == pFactZ[k] {
				k++
			}
			i++
		} else if j < len(pFactY) {
			ans *= pFactY[j]
			if pFactY[j] == pFactZ[k] {
				k++
			}
			j++
		} else if k < len(pFactZ) {
			ans *= pFactZ[k]
			k++
		}
	}

	return ans
}

func primeFactorization(num int) []int {
	ans := make([]int, 0)
	for i := 2; num > 1; {
		if num%i == 0 {
			ans = append(ans, i)
			num /= i
		} else {
			i++
		}
	}
	return ans
}

// Better solution to finding a least common multiple
func lcm2(sNum []int) int {
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
