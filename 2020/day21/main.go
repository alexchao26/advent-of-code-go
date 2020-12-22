package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/data-structures/slice"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	part1Ans, part2Ans := allergenAssessment(util.ReadFile("./input.txt"))
	if part == 1 {
		fmt.Println("Output:", part1Ans)
	} else {
		fmt.Println("Output:", part2Ans)
	}
}

// leaderboard: 225/127, closest yet!
func allergenAssessment(input string) (part1Ans int, part2Ans string) {
	allergensToPossibleIngredients := map[string][]string{}
	ingredientCounts := map[string]int{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " (contains ")
		ingredients := strings.Split(parts[0], " ")
		allergens := strings.Split(strings.Trim(parts[1], ")"), ", ")

		// count up the appearances for each ingredient
		for _, ingred := range ingredients {
			ingredientCounts[ingred]++
		}

		// generate all possible ingredients that could be a particular allergen
		for _, a := range allergens {
			// if no ingredients are there, set this as the initial list
			if allergensToPossibleIngredients[a] == nil {
				allergensToPossibleIngredients[a] = ingredients
			} else {
				// otherwise take the inner join/overlap to eliminate ingredients
				allergensToPossibleIngredients[a] = slice.IntersectionStrings(allergensToPossibleIngredients[a], ingredients)
			}
		}
	}

	// iterate through the allergens to possible map and if a slice of length 1
	// is found, remove that ingredient from all other value slices
	// do this until every slice has only one possible ingredient
	for {
		allSingle := true
		for allergen, possible := range allergensToPossibleIngredients {
			if len(possible) != 1 {
				allSingle = false
			} else {
				// remove this name from all lists
				for otherAllergen, otherIngredients := range allergensToPossibleIngredients {
					if otherAllergen != allergen {
						allergensToPossibleIngredients[otherAllergen] = slice.RemoveAllStrings(otherIngredients, possible[0])
					}
				}
			}
		}
		if allSingle {
			break
		}
	}

	// remove the allergens from the ingredientsCount map
	for _, hashedName := range allergensToPossibleIngredients {
		delete(ingredientCounts, hashedName[0])
	}

	// for part 1: count up the total occurrences of non-allergen ingredients
	var count int
	for _, ct := range ingredientCounts {
		count += ct
	}

	// for part 2: get a list of all allergens, sort them, then in order, add the
	// hashed name to a canonical dangerous list
	var names []string
	for k := range allergensToPossibleIngredients {
		names = append(names, k)
	}
	sort.Strings(names)

	var canonical []string
	for _, n := range names {
		canonical = append(canonical, allergensToPossibleIngredients[n][0])
	}

	return count, strings.Join(canonical, ",")
}
