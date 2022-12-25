package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

// to six adjacent face...
var diffs = [][3]int{
	{0, 0, 1},
	{0, 1, 0},
	{1, 0, 0},
	{0, 0, -1},
	{0, -1, 0},
	{-1, 0, 0},
}

func part1(input string) int {
	rawCoords := parseInput(input)
	mapCoords := convertRawCoordsToMap(rawCoords)

	totalSurfaceArea := 0
	for _, coord := range rawCoords {
		neighbors := 6
		for _, d := range diffs {
			if mapCoords[[3]int{
				coord[0] - d[0],
				coord[1] - d[1],
				coord[2] - d[2],
			}] {
				neighbors--
			}
		}
		totalSurfaceArea += neighbors
	}

	return totalSurfaceArea
}

func part2(input string) int {
	rawCoords := parseInput(input)
	mapCoords := convertRawCoordsToMap(rawCoords)

	// get bounds
	var limitX, limitY, limitZ int
	for c := range mapCoords {
		limitX = mathy.MaxInt(limitX, c[0])
		limitY = mathy.MaxInt(limitY, c[1])
		limitZ = mathy.MaxInt(limitZ, c[2])
	}

	// bfs to see if an edge can be reached
	// delete if not useful

	totalExternalSurfaceArea := 0

	for coord := range mapCoords {
		totalExternalSurfaceArea += facesThatCanReachEdge(coord, mapCoords,
			limitX, limitY, limitZ)
	}

	// too low: 1036
	return totalExternalSurfaceArea
}

func parseInput(input string) (ans [][3]int) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")
		ans = append(ans, [3]int{
			cast.ToInt(parts[0]),
			cast.ToInt(parts[1]),
			cast.ToInt(parts[2]),
		})
	}
	return ans
}
func convertRawCoordsToMap(rawCoords [][3]int) map[[3]int]bool {
	set := map[[3]int]bool{}
	for _, coord := range rawCoords {
		set[coord] = true
	}
	return set
}

// there would be a big optimization here to keep track of all coords that have
// a known path to an edge, that would eliminate a lot of duplicate work... but
// i think this is a small enough problem space to ignore that...
func facesThatCanReachEdge(coord [3]int, set map[[3]int]bool, limitX, limitY, limitZ int) int {
	ans := 0
	for _, d := range diffs {
		next := [3]int{
			coord[0] + d[0],
			coord[1] + d[1],
			coord[2] + d[2],
		}

		reachResult := canReachEdge(next, set, limitX, limitY, limitZ)
		if reachResult {
			ans++
		}
	}

	return ans
}

func canReachEdge(coord [3]int, set map[[3]int]bool, limitX, limitY, limitZ int,
) bool {
	queue := [][3]int{coord}
	seen := map[[3]int]bool{}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		// seen already or hit some other droplet, skip
		if seen[front] || set[front] {
			continue
		}
		seen[front] = true

		// edge reached
		if front[0] <= 0 || front[0] >= limitX ||
			front[1] <= 0 || front[1] >= limitY ||
			front[2] <= 0 || front[2] >= limitZ {
			return true
		}

		for _, d := range diffs {
			next := [3]int{
				front[0] + d[0],
				front[1] + d[1],
				front[2] + d[2],
			}
			queue = append(queue, next)
		}
	}
	return false
}
