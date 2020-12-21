package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/algos"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	rooms := parseInput(input)

	rooms = getValidRooms(rooms)
	var sum int
	for _, rm := range rooms {
		sum += rm.sectorID
	}

	return sum
}

func part2(input string) int {
	rooms := getValidRooms(parseInput(input))

	for _, rm := range rooms {
		for i := 0; i < rm.sectorID; i++ {
			// rotate each character forward
			for i, part := range rm.nameParts {
				rm.nameParts[i] = algos.CaesarShift(part, rm.sectorID)
			}
			// printed all new name parts and searched for "north" to find what the goal text was
			if rm.nameParts[0] == "northpole" &&
				rm.nameParts[1] == "object" &&
				rm.nameParts[2] == "storage" {
				return rm.sectorID
			}
		}
	}

	panic("loop should have returned sectorID")
}

type room struct {
	nameParts []string
	sectorID  int
	checksum  string
}

func parseInput(input string) (ans []room) {
	for _, line := range strings.Split(input, "\n") {
		r := room{}
		parts := strings.Split(line, "-")
		r.nameParts = parts[:len(parts)-1]
		fmt.Sscanf(parts[len(parts)-1], "%d[%5s]", &r.sectorID, &r.checksum)

		ans = append(ans, r)
	}

	return ans
}

func getValidRooms(rooms []room) []room {
	var validRooms []room
	for _, rm := range rooms {
		countChars := map[string]int{}
		for _, part := range rm.nameParts {
			for _, char := range part {
				countChars[string(char)]++
			}
		}
		var allCounts []int
		for _, v := range countChars {
			allCounts = append(allCounts, v)
		}
		// sort in reverse order so five highest are at the front
		sort.Sort(sort.Reverse((sort.IntSlice(allCounts))))

		isValid := true

		var counts []int
		for i, char := range rm.checksum {
			counts = append(counts, countChars[string(char)])
			// compare to five highest
			if counts[i] != allCounts[i] {
				isValid = false
			}

			// tie break equal counts
			if i != 0 {
				if counts[i-1] < counts[i] ||
					(counts[i-1] == counts[i] && string(rm.checksum[i-1]) > string(char)) {
					isValid = false
				}
			}
		}

		if isValid {
			validRooms = append(validRooms, rm)
		}
	}
	return validRooms
}
