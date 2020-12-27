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

	sanOrbitsToCom, youOrbitsToCom := recurseToCOM(mapOfOrbits, "SAN"), recurseToCOM(mapOfOrbits, "YOU")

	// need to find where the two run over the same planet...
	// these two variables will store the indexes within the slices, that corresponds to the same planet
	var indexOfYouPlanet, indexOfSanPlanet int

	// iterate through the SAN orbits, nest a loop to search for the same planet in the YOU orbits
	// a map would have better time complexity... due to lookup time
	for indexSan, sanOrbitPlanetString := range sanOrbitsToCom {
		// a boolean to track if the planet was found in the you orbit
		var foundYouPlanet bool
		// iterate through YOU orbits and if the same planet is found, update the indexYou and foundBool
		for indexYou, youOrbitPlanetString := range youOrbitsToCom {
			if youOrbitPlanetString == sanOrbitPlanetString {
				indexOfYouPlanet, foundYouPlanet = indexYou, true
				break
			}
		}

		// if it was found, set the indexOfSanPlanet and stop the looping
		if foundYouPlanet {
			indexOfSanPlanet = indexSan
			break
		}
	}

	// subtract two for the SAN and YOU planets (they're pointers to a planet, not extra orbits to traverse)
	fmt.Println("result is: ", indexOfYouPlanet+indexOfSanPlanet-2)
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

func recurseToCOM(mapOfOrbits map[string]string, planet string) []string {
	// added to handle part2
	planetsToComSlice := make([]string, 0)

	// start the result at one because triggering this function is using the outerPlanet
	// and needs to include that initial orbit from innerPlanet to outerPlanet
	for {
		planetString, ok := mapOfOrbits[planet]
		if ok {
			// if the planet is found in the map, place the planet on the slice and continue to loop
			planetsToComSlice = append(planetsToComSlice, planet)
			planet = planetString
		} else {
			break
		}
	}

	return planetsToComSlice
}
