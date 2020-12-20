package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/algos"

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

func part1(input string) int {
	tiles := parseTilesFromInput(input)
	edgeSize := int(math.Sqrt(float64(len(tiles))))

	assembledTiles := backtrackAssemble(tiles, nil, map[int]bool{})

	// return product of corners
	product := assembledTiles[0][0].id
	product *= assembledTiles[0][edgeSize-1].id
	product *= assembledTiles[edgeSize-1][0].id
	product *= assembledTiles[edgeSize-1][edgeSize-1].id
	return product
}

func part2(input string) int {
	tiles := parseTilesFromInput(input)
	edgeSize := int(math.Sqrt(float64(len(tiles))))

	assembledTiles := backtrackAssemble(tiles, nil, map[int]bool{})

	// remove ALL borders from all tiles
	for r := range assembledTiles {
		for c, cell := range assembledTiles[r] {
			assembledTiles[r][c].contents = removeBordersFromGrid(cell.contents)
		}
	}

	// generate image from assembledTiles...
	// assembledTiles is effectively a 2D grid where each cell is a 2D grid representing a tile
	// there has to be an easier way to "flatten" a 4D grid into a 2D grid...
	var image [][]string
	for bigRow := 0; bigRow < edgeSize; bigRow++ {
		for subRow := 0; subRow < len(assembledTiles[0][0].contents); subRow++ {
			image = append(image, []string{})
			for bigCol := 0; bigCol < edgeSize; bigCol++ {
				subLine := assembledTiles[bigRow][bigCol].contents[subRow]
				image[len(image)-1] = append(image[len(image)-1], subLine...)
			}
		}
	}

	// get the coordinates of all monsters by iterating over all possible
	// orientations of the image
	var monsterCoords [][2]int
	for _, opt := range generateGridOrientations(image) {
		monsterCoords = findMonsterCoords(opt)
		// assuming there's only one orientation of image with valid monsters
		if len(monsterCoords) > 0 {
			image = opt
			break
		}
	}

	// modify all monster coordinates to "O" characters (anything but "#")
	for _, coord := range monsterCoords {
		image[coord[0]][coord[1]] = "O"
	}

	// count up remaining "#" cells
	var roughWatersCount int
	for _, row := range image {
		for _, cell := range row {
			if cell == "#" {
				roughWatersCount++
			}
		}
	}

	return roughWatersCount
}

type tile struct {
	contents [][]string
	id       int
}

func parseTilesFromInput(input string) []*tile {
	ans := []*tile{}
	for _, block := range strings.Split(input, "\n\n") {
		split := strings.Split(block, "\n")
		var tileID int
		_, err := fmt.Sscanf(split[0], "Tile %d:", &tileID)
		if err != nil {
			panic(err)
		}

		var contents [][]string
		for _, line := range split[1:] {
			contents = append(contents, strings.Split(line, ""))
		}
		ans = append(ans, &tile{id: tileID, contents: contents})
	}
	return ans
}

func generateGridOrientations(grid [][]string) [][][]string {
	var options [][][]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			options = append(options, grid)
			grid = algos.RotateStringGrid(grid)
		}
		grid = algos.MirrorStringGrid(grid)
	}
	// note: there will likely be duplicates in there... but that's fine...
	return options
}

func backtrackAssemble(tiles []*tile, assembledTiles [][]*tile, usedIndices map[int]bool) [][]*tile {
	// pray it's a square...
	edgeSize := int(math.Sqrt(float64(len(tiles))))
	if assembledTiles == nil {
		assembledTiles = make([][]*tile, edgeSize)
		for i := 0; i < edgeSize; i++ {
			assembledTiles[i] = make([]*tile, edgeSize)
		}
	}

	// iterate through all cells, skipping cells that have already been set
	for row := 0; row < edgeSize; row++ {
		for col := 0; col < edgeSize; col++ {
			// skip cells that have already been assigned
			if assembledTiles[row][col] == nil {
				// iterate over all available tiles (skip ones that are tagged "used")
				for i, t := range tiles {
					if !usedIndices[i] {
						// iterate over the OPTIONS for a particular tile, i.e. all 8 images of it...
						for _, opt := range generateGridOrientations(t.contents) {
							// check if setting this tile is okay with (if applicable) tiles above
							// and to the left
							if row != 0 { // check above
								currentTopRow := getRow(opt, true)
								bottomOfAbove := getRow(assembledTiles[row-1][col].contents, false)
								// if they don't match, continue onto next option b/c this one doesn't match
								if currentTopRow != bottomOfAbove {
									continue
								}
							}
							if col != 0 { // check left, same logic checking above
								currentLeftCol := getCol(opt, true)
								rightColOfLeft := getCol(assembledTiles[row][col-1].contents, false)
								if currentLeftCol != rightColOfLeft {
									continue
								}
							}
							// set tile, mark tile as used, recurse
							t.contents = opt // side effects apply b/c t is a pointer
							assembledTiles[row][col] = t
							// if non-nil response, relay the return value up the call stack
							usedIndices[i] = true
							recurseResult := backtrackAssemble(tiles, assembledTiles, usedIndices)
							if recurseResult != nil {
								return recurseResult
							}
							// backtrack if nil response from recursing
							assembledTiles[row][col] = nil
							usedIndices[i] = false
						}
					}
				}
				// Note: this is a key part of backtracking, to escape if there are
				// no valid options for to set for this row/col
				if assembledTiles[row][col] == nil {
					return nil
				}
			}
		}
	}

	// if entire loop finishes, that means every cell has been assigned
	// return the assembled tiles to collapse the call stack
	return assembledTiles
}

// helper functions to get a string (easily comparable) of a single side
func getCol(grid [][]string, firstCol bool) string {
	var str string
	for i := range grid {
		if firstCol {
			str += grid[i][0]
		} else {
			str += grid[i][len(grid[0])-1]
		}
	}
	return str
}

func getRow(grid [][]string, firstRow bool) string {
	var str string
	for i := range grid[0] {
		if firstRow {
			str += grid[0][i]
		} else {
			str += grid[len(grid)-1][i]
		}
	}
	return str
}

func removeBordersFromGrid(grid [][]string) [][]string {
	var result [][]string

	for i := 1; i < len(grid)-1; i++ {
		result = append(result, []string{})
		for j := 1; j < len(grid[0])-1; j++ {
			result[i-1] = append(result[i-1], grid[i][j])
		}
	}

	return result
}

// returns all coordinates that make up any valid monsters...
// valid monsters are the '#' like so... dots are for visual effect
//                 ..#.
// #....##....##....###
// .#..#..#..#..#..#...
var monster = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

func findMonsterCoords(image [][]string) [][2]int {
	var monsterOffsets [][2]int
	var monsterHeight, monsterLength int
	for r, line := range strings.Split(monster, "\n") {
		for c, char := range line {
			if char == '#' {
				monsterOffsets = append(monsterOffsets, [2]int{r, c})
			}
			monsterLength = c + 1
		}
		monsterHeight++
	}

	// determine the top left corners of a found monster
	var monsterStartingCoords [][2]int
	for r := 0; r < len(image)-monsterHeight+1; r++ {
		for c := 0; c < len(image[0])-monsterLength+1; c++ {
			monsterFound := true
			for _, diff := range monsterOffsets {
				rowToCheck := r + diff[0]
				colToCheck := c + diff[1]
				if image[rowToCheck][colToCheck] != "#" {
					monsterFound = false
				}
			}
			if monsterFound {
				monsterStartingCoords = append(monsterStartingCoords, [2]int{r, c})
			}
		}
	}

	// generate a list of all the coordinates that are monsters
	var monsterCoords [][2]int
	for _, startingCoord := range monsterStartingCoords {
		for _, diff := range monsterOffsets {
			monsterCoords = append(monsterCoords, [2]int{startingCoord[0] + diff[0], startingCoord[1] + diff[1]})
		}
	}

	return monsterCoords
}
