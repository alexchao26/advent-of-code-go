package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/alexchao26/advent-of-code-go/cast"
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
	graph, starting := parseInput(input)

	possibles := map[string]bool{}

	for i, mol := range starting {
		if products, ok := graph[mol]; ok {
			for _, p := range products {
				starting[i] = p
				possibles[strings.Join(starting, "")] = true
			}
		}
		// reset
		starting[i] = mol
	}

	return len(possibles)
}

func parseInput(input string) (graph map[string][]string, startingMaterial []string) {
	blocks := strings.Split(input, "\n\n")
	startingMaterial = splitMolecules(blocks[1])

	graph = map[string][]string{}

	for _, l := range strings.Split(blocks[0], "\n") {
		parts := strings.Split(l, " => ")
		graph[parts[0]] = append(graph[parts[0]], parts[1])
	}

	return graph, startingMaterial
}

func splitMolecules(input string) []string {
	var molecules []string
	for _, char := range input {
		code := cast.ToASCIICode(char)
		if code >= cast.ASCIICodeCapA && code <= cast.ASCIICodeCapZ {
			molecules = append(molecules, string(char))
		} else {
			molecules[len(molecules)-1] += string(char)
		}
	}
	return molecules
}

// This makes some very large assumptions about the answer, but it all revolves
// around the fact that there is only one solution for the given input.
// It also assumes that some products need to be replaced by their reactants
//   in a particular order, and when those replacements are made, ALL instances
//   of that product can be replaced by its reactant
//
// I should learn CYK...
// Other things I tried initially were A* starting from 'e', A* from the final
// molecule, but the space was huge.
func part2(input string) int {
	reverseGraph, startingMols := parseInput(input)

	// reverse the graph so it's products to reactants
	productToReactant := map[string]string{}
	for react, products := range reverseGraph {
		for _, p := range products {
			if _, ok := productToReactant[p]; ok {
				panic("dup found")
			}
			productToReactant[p] = react
		}
	}

	// slice of all products to have an order of which products to replace
	var allProducts []string
	for prod := range productToReactant {
		allProducts = append(allProducts, prod)
	}

	start := strings.Join(startingMols, "")
	mol := start

	var steps int
	for mol != "e" {
		var changeMade bool
		for _, prod := range allProducts {
			count := strings.Count(mol, prod)
			if count <= 0 {
				continue
			}
			changeMade = true
			steps += count
			mol = strings.ReplaceAll(mol, prod, productToReactant[prod])
			// break out to restart from the beginning of allProducts slice
			break
		}
		// if no change was made, then this ordering of allProducts will not
		// resolve in an electron, shuffle and reset mol and steps
		if !changeMade {
			allProducts = shuffleSlice(allProducts)
			mol = start
			steps = 0
		}
	}

	return steps
}

var rn = rand.New(rand.NewSource(time.Now().UnixNano()))

func shuffleSlice(in []string) []string {
	// shuffle slice, lazy sorting method
	sort.Slice(in, func(i, j int) bool {
		return rn.Intn(2) == 1
	})
	return in
}
