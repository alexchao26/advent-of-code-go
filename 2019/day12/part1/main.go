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

func main() {
	input := util.ReadFile("../input.txt")
	stringSlice := strings.Split(input, "\n")

	// need to set x y z, dx, dy, dz of each of the starting moons
	sliceMoons := makeMoonSlice(stringSlice)

	// iterate through for each step
	for i := 0; i < 1000; i++ {
		// iterate through all the moons in the slice
		for start := 0; start < len(sliceMoons); start++ {
			// update velocities dx dy dz
			// requires iterating through all of the moons again...
			for restIndex := start + 1; restIndex < len(sliceMoons); restIndex++ {
				if sliceMoons[start].x < sliceMoons[restIndex].x {
					sliceMoons[start].dx++
					sliceMoons[restIndex].dx--
				} else if sliceMoons[start].x > sliceMoons[restIndex].x {
					sliceMoons[start].dx--
					sliceMoons[restIndex].dx++
				}
				if sliceMoons[start].y < sliceMoons[restIndex].y {
					sliceMoons[start].dy++
					sliceMoons[restIndex].dy--
				} else if sliceMoons[start].y > sliceMoons[restIndex].y {
					sliceMoons[start].dy--
					sliceMoons[restIndex].dy++
				}
				if sliceMoons[start].z < sliceMoons[restIndex].z {
					sliceMoons[start].dz++
					sliceMoons[restIndex].dz--
				} else if sliceMoons[start].z > sliceMoons[restIndex].z {
					sliceMoons[start].dz--
					sliceMoons[restIndex].dz++
				}
			}
		}

		// then update coordinates x y z
		for i2, e := range sliceMoons {
			sliceMoons[i2].x += e.dx
			sliceMoons[i2].y += e.dy
			sliceMoons[i2].z += e.dz
		}
	}

	// get final kinetic energy
	fmt.Println(kinetic(sliceMoons))
}

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

// get total "kinetic energy" of a slice of Moons
func kinetic(moons []Moon) (result int) {
	for _, e := range moons {
		sumXYZ := abs(e.x) + abs(e.y) + abs(e.z)
		velXYZ := abs(e.dx) + abs(e.dy) + abs(e.dz)
		result += (sumXYZ * velXYZ)
	}
	return result
}

func abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}
