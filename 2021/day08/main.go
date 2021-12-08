package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"regexp"
	"sort"
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

	ans := jumbledSevenSegment(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func jumbledSevenSegment(input string, part int) int {
	// four digit seven segment displays
	// seven segments are letters
	// need to unjumble the ten mappings (pre-bar) to outputs (last 4 post-bar)
	var parsedInput [][]string
	for i, line := range strings.Split(input, "\n") {
		// regexp capture group to just get characters
		parts := regexp.MustCompile(`([a-g]+)`).FindAllString(line, -1)

		if len(parts) != 14 {
			log.Fatalf("should be 14 parts in each input line, got %d for line %d", len(parts), i)
		}

		// return them as just all 14 mappings together, alphabetize them for easier comparison later
		var fourteen []string
		for _, v := range parts {
			fourteen = append(fourteen, alphabetizeString(v))
		}
		parsedInput = append(parsedInput, fourteen)
	}

	// PART 1 just count the outputs that are "easy" to determine
	// aka only one displayed number uses that number of segments
	if part == 1 {
		var ans int
		for _, set := range parsedInput {
			// check the back 4: output
			for _, o := range set[10:] {
				switch len(o) {
				// 1, 4, 7, 8 have a unique number of segments lit
				// 2, 4, 4, 7 segments respectively
				case 2, 4, 3, 7:
					ans++
				}
			}
		}
		return ans
	}

	// PART 2
	var ans int
	indexToCharacters := make([]string, 10)
	for _, set := range parsedInput {

		workingSet := set[:10]
		var killIndices []int // store the indices that will need to be removed in batches

		// 1, 4, 7 and 8 are all easy to find because they have a unique number of segments
		for i, mapping := range workingSet {
			switch len(mapping) {
			case 2:
				// these two make up the 1
				indexToCharacters[1] = mapping
				killIndices = append(killIndices, i)
			case 4: // the 4
				indexToCharacters[4] = mapping
				killIndices = append(killIndices, i)
			case 3: // the 7
				indexToCharacters[7] = mapping
				killIndices = append(killIndices, i)
			case 7: // the 8
				indexToCharacters[8] = mapping
				killIndices = append(killIndices, i)
			}
		}

		// remove them from the workingSet
		workingSet = removeSliceIndices(workingSet, killIndices...)

		// only 0 3 and 9 will fully overlap the 1 characters
		var zeroThreeOrNine []string
		killIndices = []int{} // reset these...
		for i, mapping := range workingSet {
			if checkStringOverlap(mapping, indexToCharacters[1]) {
				zeroThreeOrNine = append(zeroThreeOrNine, mapping)
				killIndices = append(killIndices, i)
			}
		}
		if len(zeroThreeOrNine) != 3 {
			log.Fatalf("one three or nine does not have three matches: got %d", len(zeroThreeOrNine))
		}

		// find the 3 based on segment length
		for i, maybe039 := range zeroThreeOrNine {
			if len(maybe039) == 5 {
				// must be the 3
				indexToCharacters[3] = maybe039
				zeroThreeOrNine = removeSliceIndices(zeroThreeOrNine, i)
				break
			}
		}

		// the 9 will have a full overlap with 4
		for i, maybe09 := range zeroThreeOrNine {
			if checkStringOverlap(maybe09, indexToCharacters[4]) {
				indexToCharacters[9] = maybe09
				zeroThreeOrNine = removeSliceIndices(zeroThreeOrNine, i)
			}
		}

		// 0 is only one left of the triplet
		indexToCharacters[0] = zeroThreeOrNine[0]

		// trim down working set again
		workingSet = removeSliceIndices(workingSet, killIndices...)
		if len(workingSet) != 3 {
			log.Fatalf("expected length of 3 at this stage, got %d", len(workingSet))
		}

		// 6 is the last one with a length of 6
		for i, mapping := range workingSet {
			if len(mapping) == 6 {
				indexToCharacters[6] = mapping
				workingSet = removeSliceIndices(workingSet, i)
			}
		}

		// 5 will be fully contained within the 9
		for i, mapping := range workingSet {
			if checkStringOverlap(indexToCharacters[9], mapping) {
				indexToCharacters[5] = mapping
				workingSet = removeSliceIndices(workingSet, i)
			}
		}

		if len(workingSet) != 1 {
			log.Fatalf("expected length of 1 at this stage, got %d", len(workingSet))
		}

		// 2 is the last remaining mapping
		indexToCharacters[2] = workingSet[0]

		// finally, collect the four digits in the output & add it to the answer
		var num int
		for _, out := range set[10:] {
			for i, mapping := range indexToCharacters {
				// because they were all alphabetized, we can just do a direct comparison
				if out == mapping {
					// shift all digits to the left and add the new digit to the end
					num *= 10
					num += i
				}
			}
		}
		ans += num
	}

	return ans
}

func removeSliceIndices(sli []string, indices ...int) []string {
	m := map[int]bool{}
	for _, v := range indices {
		m[v] = true
	}

	var ans []string
	for i, v := range sli {
		if !m[i] {
			ans = append(ans, v)
		}
	}
	return ans
}

func checkStringOverlap(larger, smaller string) bool {
	// safeguard, a smaller string can never contain a larger string
	// i think every use of this function already passes args in the correct order, but this makes
	// the checkStringOverlap algo more robust in general
	if len(larger) < len(smaller) {
		larger, smaller = smaller, larger
	}

	largeMap := map[rune]bool{}
	for _, r := range larger {
		largeMap[r] = true
	}

	// check all runes in the smaller string, return false if not in largeMap
	for _, r := range smaller {
		if !largeMap[r] {
			return false
		}
	}
	return true
}

// alphabetize the strings to make them easier to compare later
func alphabetizeString(input string) string {
	chars := strings.Split(input, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
