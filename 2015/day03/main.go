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

var dirs = map[string][2]int{
	"^": [2]int{-1, 0}, // row then col
	"v": [2]int{1, 0},
	"<": [2]int{0, -1},
	">": [2]int{0, 1},
}

func part1(input string) int {
	houseCount := map[[2]int]int{[2]int{}: 1}
	coord := [2]int{0, 0}
	for _, char := range strings.Split(input, "") {
		diff := dirs[string(char)]
		nextCoord := [2]int{
			coord[0] + diff[0],
			coord[1] + diff[1],
		}
		coord = nextCoord
		houseCount[coord]++
	}
	return len(houseCount)
}

func part2(input string) int {
	houseCount := map[[2]int]int{[2]int{}: 2}
	santaCoord := [2]int{0, 0}
	robotCoord := [2]int{0, 0}
	for i, char := range strings.Split(input, "") {
		diff := dirs[string(char)]
		if i%2 == 0 {
			nextSantaCoord := [2]int{
				santaCoord[0] + diff[0],
				santaCoord[1] + diff[1],
			}
			santaCoord = nextSantaCoord
			houseCount[santaCoord]++
		} else {
			nextRobotCoord := [2]int{
				robotCoord[0] + diff[0],
				robotCoord[1] + diff[1],
			}
			robotCoord = nextRobotCoord
			houseCount[robotCoord]++
		}
	}
	return len(houseCount)
}
