package main

import (
	"fmt"
	"strconv"
	"strings"
)

func breakString(str string) []string {
	var result []string
	split := strings.Split(str, ",")
	result = append(result, split...)
	return result
}

/*
// function that will return an array of coordinates that live in the form
func returnCoordsSlice(directionsSlice []string) [][]int {
	coords := [][]int{{0, 0}}
	for _, v := range directionsSlice {
		// fmt.Println(v)
		runeSlice := []rune(v)
		// fmt.Println(runeSlice)
		num, _ := strconv.Atoi(string(runeSlice[1:len(runeSlice)]))
		// fmt.Println(num)
		prevCoords := coords[len(coords)-1]
		// fmt.Println(prevCoords)

		if runeSlice[0] == 'R' {
			// fmt.Println(runeSlice[1:len(runeSlice)])

			// previous coordinates
			// is a direction going right
			coords = append(coords, []int{prevCoords[0], prevCoords[1] + num})
			// fmt.Println(coords)
		} else if runeSlice[0] == 'L' {
			coords = append(coords, []int{prevCoords[0], prevCoords[1] - num})
		} else if runeSlice[0] == 'U' {
			coords = append(coords, []int{prevCoords[0] - num, prevCoords[1]})
		} else if runeSlice[0] == 'D' {
			coords = append(coords, []int{prevCoords[0] + num, prevCoords[1]})
		}
	}
	// fmt.Println(coords)
	return coords
}
*/

// return a map that has keys of strings and a value of an int of steps to reach the coordinate
func returnMapOfCoords(directionsSlice []string) map[string]int {
	gridOfCoordinates := map[string]int{}
	prevX := 0
	prevY := 0
	runningLength := 0

	for i := 0; i < len(directionsSlice); i++ {
		// grab the current element
		v := directionsSlice[i]

		// split this element into a slice of runes...
		runeSlice := []rune(v)

		// stores number parsed off of this element
		num, _ := strconv.Atoi(string(runeSlice[1:len(runeSlice)]))
		// fmt.Println(num)

		// loop from 0 to num and add to the map/gridOfCoordinates
		for num > 0 {
			// on each loop increment the runningLength, decrement num
			runningLength++
			num--
			if runeSlice[0] == 'R' {
				// if going right, increment prevX
				prevX++
			} else if runeSlice[0] == 'L' {
				prevX--
			} else if runeSlice[0] == 'U' {
				prevY++
			} else if runeSlice[0] == 'D' {
				prevY--
			}

			// set `${prevX}x${prevY}` to the map with runningLength as the value
			newCoord := strconv.Itoa(prevX) + "x" + strconv.Itoa(prevY)

			_, ok := gridOfCoordinates[newCoord]
			if ok == false {
				gridOfCoordinates[newCoord] = runningLength
			}
		}
	}
	// fmt.Println(gridOfCoordinates)
	return gridOfCoordinates
}

// will return the manhattan distance
/*
func bruteForceCheckFourCoords(coord1 []int, coord2 []int, coord3 []int, coord4 []int) int {
	// coord1 and coord2 are from the first list (and need to be adjacent points to descibe a line)
	// same for coord3 and 4 but they're from the second list
	// check if there is an intersection between the four points
	if coord1[0] == coord2[0] && coord3[0] == coord4[0] {
		// return a -1 if the coordinates do not cross (signal to not update the min Manhattan Distance)
		// fmt.Println(coord1, coord2, coord3, coord4)
		return -1
	} else if coord1[1] == coord2[1] && coord3[1] == coord4[1] {
		// fmt.Println(coord1, coord2, coord3, coord4)
		return -1
	} else {
		// fmt.Println(coord1, coord2, coord3, coord4)
		// check if the lines are intersecting
		if coord1[0] == coord2[0] {
			// x's are equal, so check if the range of y's in the first 2 coords
			// contains y of either 3 or 4 coordinate
			yRange12 := []int{coord1[1], coord2[1]}
			sort.Ints(yRange12)

			xRange34 := []int{coord3[0], coord4[0]}
			sort.Ints(xRange34)

			if coord3[1] > yRange12[0] && coord3[1] < yRange12[1] && coord1[0] < xRange34[1] && coord1[0] > xRange34[1] {
				// return the intersection (which is the equal x's and equal y's, absolute values, summed)
				fmt.Println("returned a non -1 (upper)")
				return int(math.Abs(float64(coord1[0])) + math.Abs(float64(coord3[1])))
			}
			return -1
		}
		// else y's are equal in coord1 and 2
		// calculate
		xRange12 := []int{coord1[0], coord2[0]}
		sort.Ints(xRange12)

		yRange34 := []int{coord3[1], coord4[1]}
		sort.Ints(yRange34)

		if coord3[0] > xRange12[0] && coord3[0] < xRange12[1] && coord1[1] < yRange34[1] && coord1[1] > yRange34[0] {
			// return the intersection (which is the equal x's and equal y's, absolute values, summed)
			fmt.Println("returned a non -1 (lower)")
			return int(math.Abs(float64(coord1[1])) + math.Abs(float64(coord3[0])))
		}
		return -1
	}
	// return 25120398
}
*/

func main() {
	// these should return 610
	// input1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	// input2 := "U62,R66,U55,R34,D71,R55,D58,R83"

	// these should return 410
	// input1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	// input2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	// actual inputs
	input1 := "R1008,U428,L339,U16,R910,U221,R53,D546,L805,U376,L19,U708,R493,D489,L443,D567,R390,D771,R270,U737,R926,U181,L306,D456,L668,D79,L922,U433,L701,U472,R914,U903,L120,U199,L273,D206,L967,U711,R976,U976,R603,U8,L882,U980,R561,D697,L224,D620,L483,U193,R317,D588,L932,D990,R658,U998,L136,U759,R463,U890,L297,U648,R163,U250,R852,U699,R236,D254,L173,U720,L259,U632,L635,U426,R235,D699,R411,U650,L258,D997,L781,D209,L697,D306,L793,U657,L936,U317,R549,D798,L951,D80,R591,D480,R835,U292,L722,U987,L775,U173,R353,U404,L250,U652,L527,D282,L365,D657,L141,D301,R128,D590,L666,U478,L85,D822,L716,U68,R253,D186,R81,U741,L316,D615,L570,U407,L734,D345,L652,U362,L360,D791,R358,U190,L823,U274,L279,D998,L16,D644,L201,D469,R213,D487,L251,D395,R130,U902,L398,U201,L56,D850,R541,D661,R921,U647,R309,D550,L307,D68,R894,U653,L510,D375,R20,U86,R357,D120,L978,D200,L45,D247,R906,U334,L242,D466,R418,U548,R698,D158,R582,U469,L968,U736,R196,U437,R87,D722,R811,U625,L425,D675,L904,D331,R693,D491,R559,U540,L120,D975,R180,U224,R610,U260,L769,D486,R93,D300,L230,U181,L60,U910,L60,D554,L527,U37,R69,D649,R768,D835,L581,U932,L746,U170,L733,U40,L497,D957,R12,U686,R85,D461,R796,D142,R664,U787,R636,D621,R824,D421,R902,D686,L202,D839,R567,D129,L77,D917,L200,D106,R3,D546,L101,D762,R780,U334,L410,D190,R431,D828,L816,D529,R648,D449,L845,U49,R750,U864,L133,D822,R46,U475,L229,U929,L676,D793,R379,U71,L243,U640,L122,U183,R528,U22,R375,D928,R292,U796,R259,U325,L921,U489,L246,D153,L384,D684,L243,U65,L342,U662,R707"
	input2 := "L1008,D243,L602,D497,L395,U81,R589,U94,R640,D965,L397,D781,R464,U642,L130,D740,R938,D260,L106,D323,L626,U869,L495,U450,R640,D675,R602,D449,L542,U917,L244,U702,L644,D809,R902,U163,R118,U982,L867,D795,R546,U194,R397,D877,L354,D255,L477,U45,L707,D624,R806,U642,L926,D233,L800,U691,L990,D979,L431,U999,L423,D794,L238,U25,R986,U595,L167,U480,L555,U831,R496,U799,L897,D895,L993,D11,R486,U176,L90,U842,R499,U792,R787,U859,L100,U169,R170,D89,R297,D944,R362,D460,R147,U831,L45,U972,R582,D90,L934,U284,R555,U235,L138,U837,R965,U915,R928,U982,R157,D566,L953,U653,L629,U460,L335,D912,R355,D683,L710,D562,R392,D44,R707,D979,L749,D223,L776,D723,R735,D356,R49,U567,L563,D220,L868,U223,R448,D505,L411,U662,L188,D536,R55,U718,L108,D289,R435,U98,R775,U933,L127,D84,R253,D523,L2,D905,R266,U675,R758,D331,R122,U988,R215,D500,R89,U306,R833,U763,R570,D818,L985,U127,L87,D210,R355,D532,R870,U196,R695,U633,R170,D540,R506,U708,L663,U566,L633,U306,L452,U180,R463,D21,L220,D268,R608,U986,L493,D598,L957,D116,L454,D584,L405,U651,R352,U681,R807,U767,L988,U692,R474,U710,R607,U313,R711,U12,R371,D561,R72,U522,R270,U904,L49,U525,R562,U895,L232,D319,R902,D236,L601,D816,R836,U419,R610,U644,L733,U355,L836,U228,L895,D39,L44,D848,L965,U475,R56,U62,L458,U99,R236,D763,R912,U295,R515,U179,R20,D777,R511,D906,R903,U855,L507,D512,L63,D630,R442,U595,L701,U654,R238,D35,L31,D469,R6,D222,R132,D837,R921,U838,R986,D441,L950,D530,L397,U41,L81,D60,L245,D75,R620,D455,L937,D180,R215,D684,R724,U561,R479,D353,L501"

	split1 := breakString(input1)
	split2 := breakString(input2)

	// coordsSlice1 := returnCoordsSlice(split1)
	// coordsSlice2 := returnCoordsSlice(split2)

	coordsMap1 := returnMapOfCoords(split1)
	coordsMap2 := returnMapOfCoords(split2)

	// fmt.Println(coordsMap1)
	// fmt.Println(coordsMap2)

	lowestSumOfDistances := 9999999

	// iterate over all keys & values in coordsMap1
	for key, value1 := range coordsMap1 {
		// check if the same key is in coordsMap2
		value2, ok := coordsMap2[key]
		if ok == true {
			// fmt.Println(key, value1, value2) // I was getting an error from the 0x0 coordinates
			// check if value2 + value1 is less than lowestSumOfDistances, if so replace it
			if lowestSumOfDistances > value1+value2 {
				lowestSumOfDistances = value1 + value2
			}
		}
	}

	// for i := 0; i < len(coordsSlice1)-1; i++ {
	// 	for j := 0; j < len(coordsSlice2)-1; j++ {
	// // fmt.Println(j, j+1)
	// val := bruteForceCheckFourCoords(coordsSlice1[i], coordsSlice1[i+1], coordsSlice2[j], coordsSlice2[j+1])
	// // fmt.Println(coordsSlice1[i], coordsSlice1[i+1], coordsSlice2[j], coordsSlice2[j+1])
	// if val != -1 && min > val {
	// 	// fmt.Println(min, val)
	// 	min = val
	// }
	// 	}
	// }

	fmt.Println(lowestSumOfDistances)

	// fmt.Println(split1[2], split2[3])
	// fmt.Println(coordsSlice1[0])
	// fmt.Println(split1[2], split2[3])
	// fmt.Println(coordsSlice1)
	// fmt.Println(coordsSlice2)
}
