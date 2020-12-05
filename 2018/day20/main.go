package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	var furthest int

	coordsToRooms := parseInput(input)

	roomsVisited := make(map[[2]int]bool)
	queue := [][3]int{{0, 0, 0}} // coord, coord, distance
	// dijkstra traverse
	for len(queue) != 0 {
		front := queue[0]
		currentCoords := [2]int{front[0], front[1]}
		currentDistance := front[2]
		currentRoom := coordsToRooms[[2]int{front[0], front[1]}]

		// do not visit the same room twice
		if roomsVisited[currentCoords] {
			queue = queue[1:]
			continue
		}

		if furthest < currentDistance {
			furthest = currentDistance
		}
		roomsVisited[currentCoords] = true

		if currentRoom.northDoor {
			queue = append(queue, [3]int{currentCoords[0] - 1, currentCoords[1], currentDistance + 1})
		}
		if currentRoom.southDoor {
			queue = append(queue, [3]int{currentCoords[0] + 1, currentCoords[1], currentDistance + 1})
		}
		if currentRoom.eastDoor {
			queue = append(queue, [3]int{currentCoords[0], currentCoords[1] + 1, currentDistance + 1})
		}
		if currentRoom.westDoor {
			queue = append(queue, [3]int{currentCoords[0], currentCoords[1] - 1, currentDistance + 1})
		}

		queue = queue[1:]
	}

	return furthest
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

type room struct {
	coords                                   [2]int
	northDoor, eastDoor, southDoor, westDoor bool
}

type roomMap map[[2]int]*room

// String method for Stringer interface for easier debugging of a pointer-ed
// struct so it won't print as just an address
func (rm roomMap) String() string {
	ans := "{\n"
	for _, room := range rm {
		ans += fmt.Sprintf("coords: %v, NSEW doors: %v %v %v %v\n",
			room.coords, room.northDoor, room.eastDoor, room.southDoor, room.westDoor)
	}
	return ans + "}"
}

var dirToDiff = map[string][2]int{
	"N": [2]int{-1, 0},
	"S": [2]int{1, 0},
	"E": [2]int{0, 1},
	"W": [2]int{0, -1},
}

func parseInput(input string) map[[2]int]*room {
	coordsToRooms := roomMap{
		{0, 0}: {coords: [2]int{0, 0}},
	}

	paths := flattenRegexpPaths(input[1 : len(input)-1])

	// inputRegexp to validate all generated paths - beauty of input being regexp
	inputRegexp := regexp.MustCompile(input)

	for _, p := range paths {
		if !inputRegexp.MatchString(p) {
			panic("BAD PATH: " + p)
		}

		// step through path and update te coordsToRooms map
		currentCoords := [2]int{0, 0}
		for _, dirRune := range p {
			dir := string(dirRune)

			// if next room is not in map, add it
			nextRoomCoords := [2]int{
				currentCoords[0] + dirToDiff[dir][0],
				currentCoords[1] + dirToDiff[dir][1],
			}
			if _, ok := coordsToRooms[nextRoomCoords]; !ok {
				coordsToRooms[nextRoomCoords] = &room{
					coords: nextRoomCoords,
				}
			}

			switch dir {
			case "N":
				coordsToRooms[currentCoords].northDoor = true
				coordsToRooms[nextRoomCoords].southDoor = true
			case "S":
				coordsToRooms[currentCoords].southDoor = true
				coordsToRooms[nextRoomCoords].northDoor = true
			case "E":
				coordsToRooms[currentCoords].eastDoor = true
				coordsToRooms[nextRoomCoords].westDoor = true
			case "W":
				coordsToRooms[currentCoords].westDoor = true
				coordsToRooms[nextRoomCoords].eastDoor = true
			default:
				panic("INVALID DIRECTION" + string(dir))
			}
			currentCoords = nextRoomCoords
		}
	}

	return coordsToRooms
}

var noNestedStepsRegexp = regexp.MustCompile("^[NS|EW]*$")

// generates a string of all possible paths, inefficient but easier to parse through
func flattenRegexpPaths(input string) []string {
	if noNestedStepsRegexp.MatchString(input) {
		return strings.Split(input, "|")
	}

	var paths []string
	topLevelOptions := breakIntoTopLevelOptions(input)

	for _, step := range topLevelOptions {
		rootStr, branchStr := ingestNextBalancedParen(step)
		if rootStr[0] == '(' {
			rootStr = rootStr[1 : len(rootStr)-1]
		}
		rootPaths := flattenRegexpPaths(rootStr)
		branchPaths := flattenRegexpPaths(branchStr)

		for _, r := range rootPaths {
			for _, b := range branchPaths {
				paths = append(paths, r+b)
			}
		}
	}

	return paths
}

var parensMap = map[string]int{
	"(": 1,
	")": -1,
}

// break into all balanced parens separated by a pipe
func breakIntoTopLevelOptions(input string) []string {
	var topLevel []string

	var left, right, parenBalance int
	for right < len(input) {
		char := string(input[right])
		if char == "|" && parenBalance == 0 {
			topLevel = append(topLevel, input[left:right])
			right++
			left = right
		} else {
			right++
			parenBalance += parensMap[char]
		}

		parenBalance += parensMap[char]
	}

	if parenBalance != 0 {
		fmt.Println("top level", topLevel)
		panic(fmt.Sprintf("paren balance off for %q, %d", input, parenBalance))
	}

	// append on last option
	topLevel = append(topLevel, input[left:])

	return topLevel
}

// splits the string into the first section that's a balanced paren & the rest
// both sections will be balanced parens
//
// if the first character is a paren, it will parse until the string is balanced
// else it splits at the first paren
func ingestNextBalancedParen(input string) (balanced string, remainder string) {
	if len(input) == 0 {
		return "", ""
	}

	// if not opening paren, parse up to first open paren
	if input[0] != '(' {
		firstParenIndex := strings.Index(input, "(")
		if firstParenIndex == -1 {
			return input, ""
		}
		return input[:firstParenIndex], input[firstParenIndex:]
	}

	// otherwise first character is an opening paren
	index := 1
	parenBalance := 1

	// assume all in puts are valid - i.e. a panic for out of range input[index]
	// is a good warning
	for parenBalance != 0 {
		char := string(input[index])
		parenBalance += parensMap[char]

		index++
	}

	return input[:index], input[index:]
}
