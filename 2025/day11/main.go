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
	graph := parseInput(input)
	// big question is if there are any cycles in the actual input...
	// the answer is no, and that also did not help me solve part 2 necessarily :')

	return dfs(graph, "you", map[string]bool{})
}

func dfs(graph map[string][]string, current string, visited map[string]bool) int {
	if current == "out" {
		return 1
	}
	visited[current] = true
	ans := 0
	for _, next := range graph[current] {
		if !visited[next] {
			ans += dfs(graph, next, visited)
		}
	}
	visited[current] = false
	return ans
}

func part2(input string) int {
	graph := parseInput(input)

	return memoDfs(graph, "svr", map[string]bool{}, map[string]int{})
}

func memoDfs(graph map[string][]string, current string, visited map[string]bool, memo map[string]int) int {
	// memoize "have we visited dac or fft yet, and if so how many paths are there to out"
	key := fmt.Sprint(current+"_", visited["dac"], visited["fft"])
	if val, ok := memo[key]; ok {
		return val
	}

	if current == "out" {
		if visited["dac"] && visited["fft"] {
			memo[key] = 1
			return 1
		}
		memo[key] = 0
		return 0
	}

	visited[current] = true
	ans := 0
	for _, next := range graph[current] {
		if !visited[next] {
			ans += memoDfs(graph, next, visited, memo)
		}
	}
	// backtrack
	visited[current] = false
	memo[key] = ans
	return ans
}

func parseInput(input string) map[string][]string {
	graph := map[string][]string{}
	for _, line := range strings.Split(input, "\n") {
		line := strings.ReplaceAll(line, ":", "")
		parts := strings.Split(line, " ")
		graph[parts[0]] = parts[1:]
	}
	return graph
}
