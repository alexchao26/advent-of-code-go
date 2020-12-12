package main

import (
	"flag"
	"fmt"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := part1And2(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

// The code is nearly identical for both, so part # is passed in as an arg
func part1And2(input string, part int) int {
	coordsToRooms := generateRoomMap(input)

	// dijkstra traverse
	var furthest int             // part 1
	var countFarRooms int        // part 2
	queue := [][3]int{{0, 0, 0}} // queue node is [3]int{row, col, distance}
	roomsVisited := make(map[[2]int]bool, len(coordsToRooms))
	for len(queue) != 0 {
		front := queue[0]
		currentCoords := [2]int{front[0], front[1]}
		currentDistance := front[2]
		currentRoom := coordsToRooms[currentCoords]

		// do not visit the same room twice
		if roomsVisited[currentCoords] {
			queue = queue[1:]
			continue
		}

		// part 1 check for furthest room
		if furthest < currentDistance {
			furthest = currentDistance
		}
		// part 2 check for rooms at least 1000 doors away
		if currentDistance >= 1000 {
			countFarRooms++
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

	if part == 1 {
		return furthest
	}
	return countFarRooms
}

type room struct {
	coords                                   [2]int
	northDoor, eastDoor, southDoor, westDoor bool
}

// String receiver method to satisfy Stringer interface, for easier debugging
func (r room) String() string {
	return fmt.Sprintf("%v: N %v S %v E %v W %v", r.coords, r.northDoor, r.southDoor, r.eastDoor, r.westDoor)
}

// returns slice of rooms that represent the ends of child paths, these need
// to be extended upon
func generateRoomMap(input string) map[[2]int]*room {
	coordsToRooms := map[[2]int]*room{
		[2]int{0, 0}: &room{}, // starting room, all zero values are applicable
	}
	iter := coordsToRooms[[2]int{0, 0}]
	var stack []*room

	for _, r := range input[1 : len(input)-1] {
		switch dir := string(r); dir {
		case "N":
			nextCoords := [2]int{iter.coords[0] - 1, iter.coords[1]}
			// add room to map if it's no in there already
			if _, ok := coordsToRooms[nextCoords]; !ok {
				coordsToRooms[nextCoords] = &room{
					coords: nextCoords,
				}
			}
			// update valid doors, next's south, current's north
			nextRoom := coordsToRooms[nextCoords]
			nextRoom.southDoor = true
			iter.northDoor = true
			// move to next room
			iter = nextRoom
		case "S":
			nextCoords := [2]int{iter.coords[0] + 1, iter.coords[1]}
			if _, ok := coordsToRooms[nextCoords]; !ok {
				coordsToRooms[nextCoords] = &room{
					coords: nextCoords,
				}
			}
			nextRoom := coordsToRooms[nextCoords]
			nextRoom.northDoor = true
			iter.southDoor = true
			iter = nextRoom
		case "E":
			nextCoords := [2]int{iter.coords[0], iter.coords[1] + 1}
			if _, ok := coordsToRooms[nextCoords]; !ok {
				coordsToRooms[nextCoords] = &room{
					coords: nextCoords,
				}
			}
			nextRoom := coordsToRooms[nextCoords]
			nextRoom.westDoor = true
			iter.eastDoor = true
			iter = nextRoom
		case "W":
			nextCoords := [2]int{iter.coords[0], iter.coords[1] - 1}
			if _, ok := coordsToRooms[nextCoords]; !ok {
				coordsToRooms[nextCoords] = &room{
					coords: nextCoords,
				}
			}
			nextRoom := coordsToRooms[nextCoords]
			nextRoom.eastDoor = true
			iter.westDoor = true
			iter = nextRoom
		case "(":
			// push onto stack
			stack = append(stack, iter)
		case "|":
			// reset to top of stack
			iter = stack[len(stack)-1]
		case ")":
			// backtrack and pop off stack
			iter = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			panic("unhandled character: " + string(r))
		}
	}

	return coordsToRooms
}
