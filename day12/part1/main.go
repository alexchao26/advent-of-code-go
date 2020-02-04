package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type moon struct {
	x, y, z, dx, dy, dz int
}

func main() {
	stringSlice := readInputFile("../input.txt")
	// fmt.Println(stringSlice)

	// need to set x y z, dx, dy, dz of each of the starting moons
	sliceMoons := makeMoonSlice(stringSlice)
	// fmt.Println(sliceMoons)

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
		// fmt.Println(sliceMoons)
	}

	// get final kinetic energy?
	fmt.Println(kinetic(sliceMoons))
}

func readInputFile(path string) []string {
	// var pixelString string
	pixelSlice := make([]string, 0)
	absPath, _ := filepath.Abs(path)

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// pixelString = line
		pixelSlice = append(pixelSlice, line)
	}

	// return pixelString
	return pixelSlice
}

func makeMoonSlice(stringSlice []string) []moon {
	sliceMoons := make([]moon, 0)
	for _, str := range stringSlice {
		x := str[strings.Index(str, "x=")+2 : strings.Index(str, ",")]
		// this is gross
		y := str[strings.Index(str, "y=")+2 : strings.Index(str, "y=")+strings.Index(str[strings.Index(str, ",")+1:], ",")-1]
		z := str[strings.Index(str, "z=")+2 : len(str)-1]

		intx, _ := strconv.Atoi(x)
		inty, _ := strconv.Atoi(y)
		intz, _ := strconv.Atoi(z)
		sliceMoons = append(sliceMoons, moon{intx, inty, intz, 0, 0, 0})
	}
	return sliceMoons
}

func kinetic(moons []moon) (result int) {
	for _, e := range moons {
		sumXYZ := abs(e.x) + abs(e.y) + abs(e.z)
		velXYZ := abs(e.dx) + abs(e.dy) + abs(e.dz)
		result += (sumXYZ * velXYZ)
	}
	return result
}

func abs(value int) int {
	flt := float64(value)
	return int(math.Abs(flt))
}
