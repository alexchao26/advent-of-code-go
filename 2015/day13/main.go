package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/algos"
	"github.com/alexchao26/advent-of-code-go/mathutil"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := knightsOfTheDinnerTable(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func knightsOfTheDinnerTable(input string, part int) int {
	graph := makeHappinessGraph(input)
	var allPeople []string
	for name := range graph {
		allPeople = append(allPeople, name)
	}

	// for part 2 add "me" to the list of people, and an empty map in the grid
	// for happiness diffs off of me. hashmap misses will be populated with the
	// zero value (0) which is correct for this problem.
	if part == 2 {
		allPeople = append(allPeople, "me")
		graph["me"] = map[string]int{}
	}

	perms := algos.PermuteStringSlice(allPeople)

	maxDiff := math.MinInt32
	for _, p := range perms {
		maxDiff = mathutil.MaxInt(maxDiff, calcHappinessDiff(graph, p))
	}

	return maxDiff
}

type mapGraph map[string]map[string]int

func makeHappinessGraph(input string) mapGraph {
	graph := make(mapGraph)
	for _, line := range strings.Split(input, "\n") {
		var person1, gainLoss, person2 string
		var amount int
		fmt.Sscanf(strings.Trim(line, "."), "%s would %s %d happiness units by sitting next to %s",
			&person1, &gainLoss, &amount, &person2)

		// ensure nested map exists
		if graph[person1] == nil {
			graph[person1] = make(map[string]int)
		}

		graph[person1][person2] = amount
		if gainLoss == "lose" {
			graph[person1][person2] = -amount
		}
	}
	return graph
}

func calcHappinessDiff(graph mapGraph, seating []string) int {
	var diffs int
	for i, person := range seating {
		indexToLeft := (i - 1 + len(seating)) % len(seating)
		indexToRight := (i + 1) % len(seating)
		diffs += graph[person][seating[indexToLeft]] + graph[person][seating[indexToRight]]
	}
	return diffs
}
