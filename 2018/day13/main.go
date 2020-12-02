// Accidentally named all the carts miners for some reason...
package main

import (
	"flag"
	"fmt"
	"sort"
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
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) string {
	grid, miners := parseInputs(input)

	var collisionCoords map[[2]int]bool
	for len(collisionCoords) == 0 {
		collisionCoords = tick(grid, miners)
		// fmt.Println(printGridWithMiners(grid, miners))
	}

	var coordString string
	// should only be one entry, but range over the map to make it easier
	// it's a map only because of how I wanted to handle part 2
	for c := range collisionCoords {
		coordString = fmt.Sprintf("%d,%d", c[1], c[0])
	}
	return coordString // X,Y i.e. COL, ROW
}

func part2(input string) string {
	grid, miners := parseInputs(input)

	for len(miners) > 1 {
		collisionCoords := tick(grid, miners)

		// remove all miners that have collided
		if len(collisionCoords) > 0 {
			// wholesale replacement of miners slice
			remainingMiners := []*Miner{}
			for _, m := range miners {
				hasCollided := collisionCoords[m.coords]
				if !hasCollided {
					remainingMiners = append(remainingMiners, m)
				}
			}

			miners = remainingMiners
		}

		// fmt.Println(printGridWithMiners(grid, miners))
	}

	finalAnsCoords := miners[0].coords
	return fmt.Sprintf("%d,%d", finalAnsCoords[1], finalAnsCoords[0]) // X,Y i.e. COL, ROW
}

type Miner struct {
	coords    [2]int
	direction string
	lastTurn  int
}

func (m *Miner) Move() {
	switch m.direction {
	case "up":
		m.coords[0]--
	case "down":
		m.coords[0]++
	case "left":
		m.coords[1]--
	case "right":
		m.coords[1]++
	}
}

// HandleIntersection is for "+" cells
func (m *Miner) HandleIntersection() {
	turns := []string{"left", "straight", "right"}
	m.lastTurn++
	m.lastTurn %= 3

	turnDirection := turns[m.lastTurn]
	if turnDirection == "straight" {
		// do nothing
		return
	}

	// get new direction
	directions := []string{"left", "up", "right", "down"}
	var index int
	for i, d := range directions {
		if d == m.direction {
			index = i
			break
		}
	}
	if turnDirection == "left" {
		index--
	} else if turnDirection == "right" {
		index++
	}

	index %= 4
	if index < 0 {
		index = 3
	}

	m.direction = directions[index]

	return
}

func parseInputs(input string) (grid [][]string, miners []*Miner) {
	// overwrite miners with their corresponding grid value
	minerVals := map[rune]string{
		'^': "|",
		'v': "|",
		'>': "-",
		'<': "-",
	}
	minerDirections := map[rune]string{
		'^': "up",
		'v': "down",
		'>': "right",
		'<': "left",
	}
	lines := strings.Split(input, "\n")
	for r, l := range lines {
		if l == "" {
			continue
		}
		row := []string{}
		// each line represents a row
		for col, v := range l {
			if minerVals[v] != "" {
				row = append(row, minerVals[v])
				miners = append(miners, &Miner{
					coords:    [2]int{r, col},
					direction: minerDirections[v],
					lastTurn:  2,
				})
			} else {
				row = append(row, string(v))
			}
		}

		grid = append(grid, row)
	}

	return grid, miners
}

// collision will not be "" if there is a collision
func tick(grid [][]string, miners []*Miner) (collisionCoords map[[2]int]bool) {
	// sort miners into order of who will move first
	sort.Slice(miners, func(i, j int) bool {
		iCoords := miners[i].coords
		jCoords := miners[j].coords
		if iCoords[0] != jCoords[0] {
			return iCoords[0] < jCoords[0]
		}
		return iCoords[1] < jCoords[1]
	})

	// map to detect collisions
	minerCoords := map[[2]int]bool{}
	// populate all coords to start because collisions don't necessarily happen
	// after one or the other has moved
	for _, m := range miners {
		minerCoords[m.coords] = true
	}

	collisionCoords = make(map[[2]int]bool)
	for _, m := range miners {
		coords := m.coords

		// if a collision has already occured here, do not move this cart
		if collisionCoords[coords] {
			continue
		}

		switch grid[coords[0]][coords[1]] {
		case "/":
			switch m.direction {
			case "left":
				m.direction = "down"
			case "right":
				m.direction = "up"
			case "up":
				m.direction = "right"
			case "down":
				m.direction = "left"
			}
		case "\\":
			switch m.direction {
			case "left":
				m.direction = "up"
			case "right":
				m.direction = "down"
			case "up":
				m.direction = "left"
			case "down":
				m.direction = "right"
			}
		case "+":
			m.HandleIntersection()
		}

		// remove current coordinates (as moving off this cell)
		minerCoords[m.coords] = false

		m.Move()

		// if a Miner is at the new coordinates, add a collision entry
		if minerCoords[m.coords] {
			collisionCoords[m.coords] = true
		}
		// update coordinates in map
		minerCoords[m.coords] = true
	}

	return collisionCoords
}

// helper func to watch the miners move
func printGridWithMiners(grid [][]string, miners []*Miner) string {
	minerDirections := map[string]string{
		"up":    "^",
		"down":  "v",
		"right": ">",
		"left":  "<",
	}
	str := ""

	mapMinerCoords := map[[2]int]*Miner{}
	for _, m := range miners {
		mapMinerCoords[m.coords] = m
	}

	for r, row := range grid {
		for c, val := range row {
			if m, ok := mapMinerCoords[[2]int{r, c}]; ok {
				str += minerDirections[m.direction]
			} else {
				str += val
			}
		}
		str += "\n"
	}
	return str
}
