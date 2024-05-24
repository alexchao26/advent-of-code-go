package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

func part1(input string) int {
	seeds, listOfMappingRanges := parseInput(input)

	lowestFinalNum := math.MaxInt64

	for _, seed := range seeds {
		finalMapping := getFinalMapping(seed, listOfMappingRanges)
		lowestFinalNum = mathy.MinInt(lowestFinalNum, finalMapping)
	}

	return lowestFinalNum
}

func getFinalMapping(seed int, listOfMappingRanges [][]mappingRange) int {
	currentVal := seed

	// jeez these are horrible variable names...
	for _, mappingRanges := range listOfMappingRanges {
		// if currentVal is between the sourceStart and sourceStart plus count, a mappingRange is found
		// otherwise it maps to the same value and move on
		// [sourceStart, sourceStart + count) <- not inclusive of sourceStart + count
		for _, m := range mappingRanges {
			if m.sourceStart <= currentVal && currentVal < m.sourceStart+m.count {
				currentVal = m.destinationStart + (currentVal - m.sourceStart)
				// then break so number is only mapped once per round of mappings
				break
			}
		}
	}

	return currentVal
}

type mappingRange struct {
	// seed to soil map: 50 98 2
	// means 2 mappings total: seed 98 maps to water 50 and seed 99 maps to soil 51
	// if a mappingRange does not exist, then it maps to the same number
	sourceStart, destinationStart, count int
}

func parseInput(input string) (seeds []int, listOfMappingRanges [][]mappingRange) {
	parts := strings.Split(input, "\n\n")

	seedParts := strings.Split(parts[0], " ")
	// skip first part which is "seeds: "
	for i := 1; i < len(seedParts); i++ {
		seeds = append(seeds, cast.ToInt(seedParts[i]))
	}

	// parse mappings
	for p := 1; p < len(parts); p++ {
		// get separate lines and ignore throw away first line of text
		mappingLines := strings.Split(parts[p], "\n")[1:]
		var mappings []mappingRange
		for _, l := range mappingLines {
			lineParts := strings.Split(l, " ")
			mappings = append(mappings, mappingRange{
				destinationStart: cast.ToInt(lineParts[0]),
				sourceStart:      cast.ToInt(lineParts[1]),
				count:            cast.ToInt(lineParts[2]),
			})
		}

		listOfMappingRanges = append(listOfMappingRanges, mappings)
	}

	return seeds, listOfMappingRanges
}

func part2(input string) int {
	seedRanges, listOfMappingRanges := parseInput(input)

	lowestFinalNum := math.MaxInt64

	// store final mappings to save on duplicate calcs?
	// not sure if that helped. brute force solution worked in a minute or two
	// finalMappings := map[int]int{}

	for i := 0; i < len(seedRanges); i += 2 {
		for count := 0; count < seedRanges[i+1]; count++ {
			seed := seedRanges[i] + count
			finalMapping := getFinalMapping(seed, listOfMappingRanges)
			lowestFinalNum = mathy.MinInt(lowestFinalNum, finalMapping)
		}

		// progress check...
		fmt.Println(i, len(seedRanges))
	}

	return lowestFinalNum
}

// ////////////
// naive solutions that are too slow with actual input
//

func naivePart1(input string) int {
	seeds, allMaps := naiveParseInput(input)
	lowestFinalNum := math.MaxInt64

	for _, s := range seeds {
		finalLocation := naiveMapSeedToFinalLocation(s, allMaps)
		lowestFinalNum = mathy.MinInt(lowestFinalNum, finalLocation)
	}

	return lowestFinalNum
}

func naiveMapSeedToFinalLocation(seed int, allMaps []map[int]int) int {
	mappedNum := seed

	for _, m := range allMaps {
		num, ok := m[mappedNum]
		if ok {
			mappedNum = num
		}
	}

	return mappedNum
}

// assume that the mappings are in a logical order in the input
// seeds -> soil -> fertilizer -> water -> etc and not out of order...
// so a slice works fine for storing the mappings
func naiveParseInput(input string) (seeds []int, allMaps []map[int]int) {
	// seed to soil map: 50 98 2
	// means 2 mappings total: seed 98 maps to water 50 and seed 99 maps to soil 51
	// if a mappingRange does not exist, then it maps to the same number

	parts := strings.Split(input, "\n\n")

	seedParts := strings.Split(parts[0], " ")
	// skip first part which is "seeds: "
	for i := 1; i < len(seedParts); i++ {
		seeds = append(seeds, cast.ToInt(seedParts[i]))
	}

	// mappings
	for p := 1; p < len(parts); p++ {
		// get separate lines and ignore throw away first line of text
		mappingLines := strings.Split(parts[p], "\n")[1:]
		m := map[int]int{}
		for _, line := range mappingLines {
			nums := strings.Split(line, " ")
			destinationStart, sourceStart, count := cast.ToInt(nums[0]), cast.ToInt(nums[1]), cast.ToInt(nums[2])
			for c := 0; c < count; c++ {
				m[sourceStart+c] = destinationStart + c
			}
		}
		allMaps = append(allMaps, m)
	}

	return seeds, allMaps
}
