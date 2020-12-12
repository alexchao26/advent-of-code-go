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

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		// ans := part2(util.ReadFile("./input.txt"))
		// fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	g := newGame(input)
	fmt.Println("game: ", g)

	var gameover bool
	for !gameover {
		gameover = g.runTurn()
		fmt.Println("AFTER IN PART 1", g)
	}

	var totalHp int
	for _, c := range g.coordsToChars {
		totalHp += c.hp
	}

	// NOT 182715, too low
	return g.rounds * totalHp
}

type game struct {
	grid          [][]string
	coordsToChars map[[2]int]*character
	rounds        int
}

func (g game) String() string {
	ans := fmt.Sprintf("Rounds: %d\n", g.rounds)
	for _, row := range g.grid {
		for _, v := range row {
			ans += v
		}
		ans += "\n"
	}

	for coord, char := range g.coordsToChars {
		ans += fmt.Sprintf("%v: Char: %v\n", coord, char)
	}
	return ans
}

type character struct {
	coord    [2]int
	hp       int
	charType string // "E" or "G"
}

func (c character) String() string {
	return fmt.Sprintf("%v %v HP:%v", c.charType, c.coord, c.hp)
}

func newGame(input string) *game {
	lines := strings.Split(input, "\n")
	var grid [][]string
	coordsToChars := map[[2]int]*character{}

	for row, line := range lines {
		grid = append(grid, make([]string, len(line)))
		for col, val := range strings.Split(line, "") {
			grid[row][col] = val
			switch val {
			case "E", "G":
				coord := [2]int{row, col}
				coordsToChars[coord] = &character{
					coord:    coord,
					hp:       200,
					charType: val,
				}
			}
		}
	}

	return &game{
		grid:          grid,
		coordsToChars: coordsToChars,
		rounds:        0,
	}
}

func (g *game) runTurn() (gameover bool) {
	turnOrder := g.getTurnOrder()

	for _, charCoords := range turnOrder {
		// ensure this character is still alive
		currentChar, ok := g.coordsToChars[charCoords]
		if !ok {
			fmt.Println("character is dead, skipping turn; ", charCoords)
			continue
		}
		fmt.Println("turn's coordinates:", charCoords)

		// check if there are enemies still
		enemyType := getEnemyType(g.grid[charCoords[0]][charCoords[1]])
		var enemiesFound bool
		for _, char := range g.coordsToChars {
			if char.charType == enemyType {
				enemiesFound = true
				break
			}
		}
		if !enemiesFound {
			return true
		}

		// check if the character has a unit next to ir right now
		enemy := g.pickTarget(charCoords)
		if enemy != nil {
			// attack & move on
			g.attack(g.coordsToChars[charCoords], enemy)
			fmt.Println("  immediately attacking:", enemy)
		} else {
			// try to move, then try to pick an enemy again
			inRangeCoordsMap := g.calcInRangeCoordsMap(g.grid[charCoords[0]][charCoords[1]])
			// fmt.Println("for starting", startingCoord, "\n    IN RANGE MAP", inRangeCoordsMap)
			// if no enemies are in the map, they're all dead
			if len(inRangeCoordsMap) == 0 {
				fmt.Println("--nowhere to move, continuing...")
				continue
			}

			nextCoord, willMove := g.determineNextMove(charCoords, inRangeCoordsMap)
			if willMove {
				fmt.Println("moving to", nextCoord)
				// update grid and coordinates for this character
				g.grid[nextCoord[0]][nextCoord[1]] = currentChar.charType
				g.grid[charCoords[0]][charCoords[1]] = "."
				g.coordsToChars[nextCoord] = g.coordsToChars[charCoords]
				currentChar.coord = nextCoord
				delete(g.coordsToChars, charCoords)

				fmt.Println("  searching for enemy after moving")
				enemy := g.pickTarget(nextCoord)
				if enemy != nil {
					// attack & move on
					g.attack(g.coordsToChars[nextCoord], enemy)
					fmt.Println("    after attack:", enemy)
				}
			} else {
				fmt.Println("NO IN RANGE TARGETS TO MOVE TO")
			}
		}
	}

	g.rounds++
	return false
}

// returns a slice of coordinates where there are characters, in turn order
func (g *game) getTurnOrder() [][2]int {
	var charCoords [][2]int
	for i, row := range g.grid {
		for j, tile := range row {
			if tile == "E" || tile == "G" {
				charCoords = append(charCoords, [2]int{i, j})
			}
		}
	}
	return charCoords
}

// order diffs in such a way that the shortest paths found will be in reading
// list order
var diffs = [][2]int{
	{-1, 0}, // up
	{0, -1}, // left
	{0, 1},  // right
	{1, 0},  // down
}

// checks the four directions around the given coordinate
func (g *game) pickTarget(currentCoords [2]int) *character {
	enemyType := getEnemyType(g.grid[currentCoords[0]][currentCoords[1]])

	var chosenEnemy *character
	for _, d := range diffs {
		nextRow := currentCoords[0] + d[0]
		nextCol := currentCoords[1] + d[1]
		next := [2]int{nextRow, nextCol}
		enemy, ok := g.coordsToChars[next]
		// fmt.Printf("  picking target, checking %v, enemy? %v\n", next, enemy)
		if ok && enemy.charType == enemyType {
			if chosenEnemy == nil || chosenEnemy.hp > enemy.hp {
				chosenEnemy = enemy
			}
			// due to the ordering of diffs slice, the reading-order enemy will
			// be chosen first if there is an HP tie... I think...
		}
	}

	return chosenEnemy
}

func (g *game) attack(attacker, target *character) {
	target.hp -= 3
	if target.hp <= 0 {
		fmt.Println("  KILLED:", target)
		// remove target from map and update grid
		targetCoords := target.coord
		delete(g.coordsToChars, target.coord)
		g.grid[targetCoords[0]][targetCoords[1]] = "."
	}
}

type bfsNode struct {
	coord       [2]int
	dist        int
	initialMove [2]int
}

func (g *game) determineNextMove(startingCoord [2]int, inRangeCoordsMap map[[2]int]bool) (nextCoord [2]int, willMove bool) {
	queue := []bfsNode{
		{coord: startingCoord, dist: 0, initialMove: [2]int{}},
	}
	visitedCoords := map[[2]int]bool{[2]int{0, 0}: true}

	for len(queue) > 0 {
		// get front of queue
		front := queue[0]
		queue = queue[1:]

		// if front is in range of an enemy, return the initial move
		if inRangeCoordsMap[front.coord] {
			return front.initialMove, true
		}

		// if it has not been visited before, then check its four directions
		if !visitedCoords[front.coord] {
			for _, d := range diffs {
				nextCoord := [2]int{d[0] + front.coord[0], d[1] + front.coord[1]}
				// only proceed if next coordinate is walkable
				if g.grid[nextCoord[0]][nextCoord[1]] == "." {
					// add next coord to queue
					node := bfsNode{
						coord:       nextCoord,
						dist:        front.dist + 1,
						initialMove: front.initialMove,
					}
					if front.dist == 0 {
						node.initialMove = nextCoord
					}
					queue = append(queue, node)
				}
			}
		}
		visitedCoords[front.coord] = true
	}

	fmt.Println("WILL NOT MOVE FROM ", startingCoord, "\nON GAME\n", g)
	return [2]int{}, false
}

// returns a slice of coordinates that are next to enemies and tile is floor
// to be run when a character wants to figure out where to move
// if the returned map is empty (len 0), that indicates no one should move
func (g *game) calcInRangeCoordsMap(attackingType string) map[[2]int]bool {
	enemyType := getEnemyType(attackingType)

	inRangeCoords := map[[2]int]bool{}
	for row := 1; row < len(g.grid)-1; row++ {
		for col := 1; col < len(g.grid[0])-1; col++ {
			// if search type is found, check four neighbors for a ground
			if g.grid[row][col] == enemyType {
				for _, d := range diffs {
					nextRow := row + d[0]
					nextCol := col + d[1]
					if g.grid[nextRow][nextCol] == "." {
						inRangeCoords[[2]int{nextRow, nextCol}] = true
					}
				}
			}
		}
	}
	return inRangeCoords
}

func getEnemyType(attacker string) string {
	if attacker == "G" {
		return "E"
	}
	return "G"
}

// Path finding...
//   Ties broken in READING order
//   top to bottom, left to right
//
// # wall
// . open
// G goblin
// E Elf
//
// Rounds:
// each unit takes a turn & completes all its actions until the next unit goes
// IN order of reading position (record at start of round)
// 1. IF not in range of energy, tries to move towrds one
//    1.1. identify all possible targets, if no targets, end combat
//    1.2. get open squares (.) in range of all targets
//         1.2.1 if no open squares found, end turn
//    1.3. determine closest open square, tie break via reading order
//    1.4. takes single step towards chosen target, along SHORTEST path, ties broken via reading order
// 2. IF in range, attack
//    2.1. if no units next to it, move on
//    2.2. select neighbor with fewest hitpoints, tie break via reading order
//    2.3. do damage equal to attack power (starts w/ 200HP & 3 attack power)
//    2.4. if unit dies, make it a (.)

// Part 1:
// - find number of FULL rounds (i.e. do not include last one)
// - find sum of remaining hit points on board
// multiply them together
