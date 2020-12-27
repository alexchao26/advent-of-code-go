package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	input := util.ReadFile("../input.txt")
	stoichiometryGraph := makeDependencyGraph(strings.Split(input, "\n"))

	neededChemicals := map[string]int{"FUEL": 1}

	// fmt.Println(stoichiometryGraph, neededChemicals)

	// while the neededChemicals map has positive values for keys besides "ORE"
	for !isOnlyOres(neededChemicals) {
		// iterate through neededChemicals
		for neededChemical, quantityNeeded := range neededChemicals {
			// if a positive value is found for a neededChemical besides "ORE"
			if quantityNeeded > 0 && neededChemical != "ORE" {
				// find its reaction in the stoichiometryGraph
				reactionStoichiometry := stoichiometryGraph[neededChemical]

				// determine the number of times the reactionStoichiometry must be run
				timesToRun := quantityNeeded / reactionStoichiometry[neededChemical]
				if quantityNeeded%reactionStoichiometry[neededChemical] > 0 {
					timesToRun++
				}

				// decrement all of the values in neededChemicals map with this reaction's details (* timesToRun)
				for reactionChemical, chemicalStoich := range reactionStoichiometry {
					neededChemicals[reactionChemical] -= chemicalStoich * timesToRun
				}
			}
		}
	}

	// Print final output
	fmt.Println("ORE needed", neededChemicals["ORE"])
}

// Creates a graph that maps the products name to its full reaction stoichiometry
func makeDependencyGraph(reactions []string) map[string]map[string]int {
	graph := make(map[string]map[string]int)

	for _, reaction := range reactions {
		product, reactionStoichiometry := parseReaction(reaction)
		graph[product] = reactionStoichiometry
	}

	return graph
}

// parseReaction takes in a line of the input i.e. a reaction in a string
// parses all its details and returns the generated product as a string
// and a map of all products to the quantity used/produced by the reaction
//   for produced chemicals, map value will be > 0, inputs will be < 0
func parseReaction(reaction string) (string, map[string]int) {
	reactionStoichiometry := make(map[string]int)

	splitByArrow := strings.Split(reaction, " => ")
	productStr := splitByArrow[1]
	reactantsStr := splitByArrow[0]

	// handle product
	productQty, productName := parseQtyAndName(productStr)

	reactionStoichiometry[productName] = productQty

	// split reactants via comma first
	reactantsSli := strings.Split(reactantsStr, ", ")
	for _, str := range reactantsSli {
		reactantQuantity, reactantName := parseQtyAndName(str)
		reactionStoichiometry[reactantName] = -1 * reactantQuantity
	}

	return productName, reactionStoichiometry
}

// parse an inputted string of the for "<number> <string>" and return the int & string
func parseQtyAndName(input string) (int, string) {
	split := strings.Split(input, " ")

	quantity, _ := strconv.Atoi(split[0])
	name := split[1]

	return quantity, name
}

// helper function to determine if the neededChemicals graph is "complete"
// it is complete if there only positive value is for the chemical "ORE"
func isOnlyOres(neededChemicals map[string]int) bool {
	for key, val := range neededChemicals {
		if key != "ORE" && val > 0 {
			return false
		}
	}
	return true
}
