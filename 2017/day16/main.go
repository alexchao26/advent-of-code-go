package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := permPromenade(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func permPromenade(input string, part int) string {
	// electing to be laxy and parse values out of each step when I get there
	// instead of doing it all upfront
	steps := strings.Split(input, ",")

	// init programs slice
	var programs []string
	for i := 0; i < 16; i++ {
		programs = append(programs, string(rune('a'+i)))
	}

	rounds := 1
	if part == 2 {
		rounds = 1000000000
	}

	// for part 2 to determine if a state has been seen before, and leveraging
	// that frequency to minimize operations of the steps between them
	seenStateToIndex := map[string]int{}
	for i := 0; i < rounds; i++ {
		for _, step := range steps {
			switch step[0] {
			case 's':
				countToSpin := mathutil.StrToInt(step[1:])
				fromEnd := programs[len(programs)-countToSpin:]
				fromFront := programs[:len(programs)-countToSpin]
				programs = append(fromEnd, fromFront...)
			case 'x':
				var index1, index2 int
				_, err := fmt.Sscanf(step, "x%d/%d", &index1, &index2)
				if err != nil {
					panic("error parsing an 'x' step " + err.Error())
				}
				programs[index1], programs[index2] = programs[index2], programs[index1]
			case 'p':
				var char1, char2 string
				_, err := fmt.Sscanf(step, "p%1s/%1s", &char1, &char2)
				if err != nil {
					panic("error parsing a 'p' step " + err.Error())
				}
				// find index then swap. Programs is 16 elements so this isn't THAT slow
				index1, index2 := -1, -1
				for i, v := range programs {
					if v == char1 {
						index1 = i
					} else if v == char2 {
						index2 = i
					}
				}
				programs[index1], programs[index2] = programs[index2], programs[index1]
			default:
				panic("unfound step type " + string(step[0]))
			}
		}
		// stringify so they're comparable and can be used as a map key
		state := stringify(programs)
		if lastSeenIndex, ok := seenStateToIndex[state]; ok {
			diff := i - lastSeenIndex
			for i+diff < rounds {
				i += diff
			}
		}
		seenStateToIndex[state] = i
	}

	return stringify(programs)
}

func stringify(programs []string) string {
	var state string
	for _, v := range programs {
		state += v
	}
	return state
}
