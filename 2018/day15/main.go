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

func part1(input string) int {
	g := newGame(input)
	return g.runFullGame()
}

func part2(input string) int {
	var outcome int
	for elfPower := 4; ; elfPower++ {
		g := newGame(input)
		elvesBefore := g.countElves()

		// update all elves to the new attack power
		for _, c := range g.coordsToChars {
			if c.charType == "E" {
				c.attackPower = elfPower
			}
		}

		// run the game until it ends... not optimized: could abort when an elf dies
		// but it's good enough
		outcome = g.runFullGame()

		// check if all elves are still alive
		if elvesBefore == g.countElves() {
			break
		}
	}

	return outcome
}

type game struct {
	grid          [][]string
	coordsToChars map[[2]int]*character
	rounds        int
}

func (g game) String() string {
	ans := fmt.Sprintf("Rounds: %d\n", g.rounds)
	for rowNum, row := range g.grid {
		ans += fmt.Sprintf("\n%02d: ", rowNum)
		for _, v := range row {
			ans += v
		}
	}
	ans += "\nAlive characters:"
	for coord, char := range g.coordsToChars {
		ans += fmt.Sprintf("\n%v: Char: %v", coord, char)
	}
	return ans
}

func (g *game) countElves() int {
	var elves int
	for _, c := range g.coordsToChars {
		if c.charType == "E" {
			elves++
		}
	}
	return elves
}

type character struct {
	coord       [2]int
	hp          int
	charType    string // "E" or "G"
	attackPower int
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
					coord:       coord,
					hp:          200,
					charType:    val,
					attackPower: 3, // default to 3
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

func (g *game) runFullGame() int {
	var gameover bool
	for !gameover {
		gameover = g.runTurn()
	}

	var totalHp int
	for _, c := range g.coordsToChars {
		totalHp += c.hp
	}

	return g.rounds * totalHp
}

func (g *game) runTurn() (gameover bool) {
	charsInOrder := g.getTurnOrder()

	for _, char := range charsInOrder {
		// if char is already dead, just continue on
		if char.hp <= 0 {
			continue
		}

		// check if there are enemies in entire game
		enemyType := getEnemyType(char.charType)
		var enemiesFound bool
		for _, c := range g.coordsToChars {
			if c.charType == enemyType {
				enemiesFound = true
				break
			}
		}
		if !enemiesFound {
			return true
		}

		// check if the character has a unit next to ir right now
		enemy := g.pickTarget(char.coord)
		if enemy != nil {
			// attack & move on
			g.attack(g.coordsToChars[char.coord], enemy)
		} else {
			// else try to move, then try to pick an enemy again
			inRangeCoordsMap := g.getInRangeOfEnemies(char.charType)
			// if no in range coords, that does not mean all enemies are dead
			// it just means there is no open floor around enemies
			if len(inRangeCoordsMap) == 0 {
				continue
			}
			// get next move
			nextCoord, willMove := g.determineNextMove(char.coord, inRangeCoordsMap)
			if willMove {
				// update grid for this character
				g.grid[nextCoord[0]][nextCoord[1]] = char.charType
				g.grid[char.coord[0]][char.coord[1]] = "."

				// coords of this char have changed, update variables
				delete(g.coordsToChars, char.coord) // delete old entry using char's outdated coords
				g.coordsToChars[nextCoord] = char   // add new entry
				char.coord = nextCoord              // update char's coords too

				// pick a target and attack it
				enemy := g.pickTarget(nextCoord)
				if enemy != nil {
					g.attack(g.coordsToChars[nextCoord], enemy)
				}
			}
		}
	}

	g.rounds++
	return false
}

// returns a slice of coordinates where there are characters, in turn order
func (g *game) getTurnOrder() []*character {
	var charsInOrder []*character
	for i, row := range g.grid {
		for j, tile := range row {
			if tile == "E" || tile == "G" {
				charsInOrder = append(charsInOrder, g.coordsToChars[[2]int{i, j}])
			}
		}
	}
	return charsInOrder
}

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
	target.hp -= attacker.attackPower
	if target.hp <= 0 {
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
		{coord: startingCoord, dist: 0, initialMove: [2]int{}}, // some zero values are redundant, but readable
	}
	visitedCoords := map[[2]int]bool{[2]int{0, 0}: true}

	// store the closest in range nodes to tie break
	var closestInRange []bfsNode

	// run while the closet in range slice is still empty and queue is not empty
	for checkDist := 0; len(closestInRange) == 0 && len(queue) > 0; checkDist++ {
		// process front of queue while its distance is equal to the check distance
		for len(queue) > 0 && queue[0].dist == checkDist {
			front := queue[0]
			queue = queue[1:]

			// if front is in range of an enemy, add to closest in range slice
			if inRangeCoordsMap[front.coord] {
				closestInRange = append(closestInRange, front)
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
	}

	if len(closestInRange) == 0 {
		return [2]int{}, false
	}

	// sort destination nodes via reading order of coords, break ties on initialMove
	sort.Slice(closestInRange, func(i, j int) bool {
		nodeI, nodeJ := closestInRange[i], closestInRange[j]
		if nodeI.coord != nodeJ.coord {
			return readingOrderSortFunc(nodeI.coord, nodeJ.coord)
		}
		return readingOrderSortFunc(nodeI.initialMove, nodeJ.initialMove)
	})

	// return the initial move of the winning bfs node, will be used to move
	// the character
	return closestInRange[0].initialMove, true
}

// returns a slice of coordinates that are next to enemies and tile is floor
// to be run when a character wants to figure out where to move
// if the returned map is empty (len 0), that indicates no one should move
func (g *game) getInRangeOfEnemies(attackingType string) map[[2]int]bool {
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

// should i go before j in a slice where we're sorting by reading order
func readingOrderSortFunc(i, j [2]int) (iBeforeJ bool) {
	// compare via first indices if not equal
	if i[0] != j[0] {
		return i[0] < j[0]
	}
	// otherwise tie break via second indices
	return i[1] < j[1]
}
