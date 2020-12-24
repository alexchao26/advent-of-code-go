package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"regexp"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := md5Bfs(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

var openPattern = regexp.MustCompile("[b-f]")

type node struct {
	coords   [2]int
	path     string
	distance int
}

func md5Bfs(input string, part int) string {
	queue := []node{{path: input}}
	var longestPath string
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if front.coords == [2]int{3, 3} {
			validPath := front.path[len(input):]
			if part == 1 {
				return validPath
			}

			if len(longestPath) < len(validPath) {
				longestPath = validPath
			}
			// cannot pass through the end point
			continue
		}

		hash := fmt.Sprintf("%x", md5.Sum([]byte(front.path)))

		for i, direction := range []struct {
			char    string
			rowDiff int
			colDiff int
		}{
			{"U", -1, 0},
			{"D", 1, 0},
			{"L", 0, -1},
			{"R", 0, 1},
		} {
			nextRow := front.coords[0] + direction.rowDiff
			nextCol := front.coords[1] + direction.colDiff
			if nextRow >= 0 && nextRow < 4 && nextCol >= 0 && nextCol < 4 {
				if openPattern.MatchString(hash[i : i+1]) {
					queue = append(queue, node{
						coords:   [2]int{nextRow, nextCol},
						path:     front.path + direction.char,
						distance: front.distance + 1,
					})
				}
			}
		}
	}

	// part 2, stringified number...
	return cast.ToString(len(longestPath))
}
