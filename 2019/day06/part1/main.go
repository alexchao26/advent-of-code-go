package main

import (
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	input := util.ReadFile("../input.txt")

	// split it into a slice of all strings
	inputSliced := strings.Split(input, ",")

	// send the input to a helper function that will return a map {planetName string: planetItOrbits string}
	mapOfOrbits := makeGraphDependencyList(inputSliced)

	result := 0

	// don't care about outer planet value
	for _, innerPlanet := range mapOfOrbits {
		result += recurseToCOM(mapOfOrbits, innerPlanet)
	}
	fmt.Println(result)
}

func makeGraphDependencyList(inputSlice []string) map[string]string {
	resultMap := make(map[string]string)
	for _, v := range inputSlice {
		slicedPlanets := strings.Split(v, ")")
		inner, outer := slicedPlanets[0], slicedPlanets[1]
		// set the key-value pair on the map/"object"
		resultMap[outer] = inner
	}

	return resultMap
}

func recurseToCOM(mapOfOrbits map[string]string, planet string) int {
	// start the result at one because triggering this function is using the outerPlanet
	// and needs to include that initial orbit from innerPlanet to outerPlanet
	result := 1
	for {
		planetString, ok := mapOfOrbits[planet]
		if ok == true {
			result++
			planet = planetString
		} else {
			break
		}
	}

	return result
}
