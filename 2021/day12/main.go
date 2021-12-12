package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

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

// 610/1663
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

func part1(input string) int {
	parsed := parseInput(input)

	// part 1
	// cannot visit SMALL caves more than once (lowercase ones)
	// how many paths are there from start -> end

	graph := map[string]map[string]bool{}
	for _, pair := range parsed {
		if graph[pair[0]] == nil {
			graph[pair[0]] = map[string]bool{}
		}
		if graph[pair[1]] == nil {
			graph[pair[1]] = map[string]bool{}
		}
		graph[pair[0]][pair[1]] = true
		graph[pair[1]][pair[0]] = true
	}

	return walk(graph, "start", map[string]bool{"start": true}, []string{"start"})
}

func walk(graph map[string]map[string]bool, current string, visited map[string]bool, path []string) int {
	if current == "end" {
		return 1
	}

	var pathsToEnd int

	for visitable := range graph[current] {
		if visited[visitable] && strings.ToUpper(visitable) != visitable {
			continue
		}
		visited[current] = true
		// path is basically unused (and wasting memory), but useful to debug sometimes
		path = append(path, visitable)

		pathsToEnd += walk(graph, visitable, visited, path)

		// backtrack
		visited[visitable] = false
		path = path[:len(path)-1]
	}

	return pathsToEnd
}

func part2(input string) int {
	parsed := parseInput(input)

	// cannot visit SMALL caves more than once (lowercase ones)
	// how many paths are there from start -> end

	graph := map[string]map[string]bool{}
	for _, pair := range parsed {
		if graph[pair[0]] == nil {
			graph[pair[0]] = map[string]bool{}
		}
		if graph[pair[1]] == nil {
			graph[pair[1]] = map[string]bool{}
		}
		graph[pair[0]][pair[1]] = true
		graph[pair[1]][pair[0]] = true
	}

	return walk2(graph, "start", map[string]int{"start": 5}, []string{"start"}, false)
}

func walk2(graph map[string]map[string]bool, current string, visited map[string]int, path []string, doubleUsed bool) int {
	if current == "end" {
		fmt.Println("path", path)
		return 1
	}

	visited[current]++

	var pathsToEnd int

	for visitable := range graph[current] {
		if visitable == "start" {
			continue
		}

		if strings.ToUpper(visitable) != visitable && visited[visitable] > 0 {
			if doubleUsed {
				continue
			} else {
				doubleUsed = true
			}
		}

		path = append(path, visitable)
		pathsToEnd += walk2(graph, visitable, visited, path, doubleUsed)

		// backtrack
		visited[visitable]--
		path = path[:len(path)-1]
		// backtrack doubleUsed IF this is a smallcave and reducing its visited count still has the
		// cave marked as visited (aka count == 1)
		if strings.ToUpper(visitable) != visitable && visited[visitable] == 1 {
			doubleUsed = false
		}
	}

	return pathsToEnd
}

func parseInput(input string) (ans [][]string) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, strings.Split(line, "-"))
	}
	return ans
}
