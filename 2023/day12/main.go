package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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
	stringConditions := parseInput(input)

	ans := 0
	// brute force creating all possible combination per line
	// then check each possibility
	// input is 1000 lines with 10k ?'s total, so approx 10/line
	// 2^10 = 1024 options per line approx. * 1000 = 1_024_000 checks total... seems ok... for part 1...
	for _, sc := range stringConditions {
		possibilities := generatePossibilities(sc.record)

		for _, p := range possibilities {
			if checkIfSpringRecordFitsDamagedGroupCounts(p, sc.damagedGroupCounts) {
				ans++
			}
		}
	}

	return ans
}

func generatePossibilities(record []string) [][]string {
	var recurse func(record []string, index int) [][]string
	recurse = func(record []string, index int) [][]string {
		if index == len(record) {
			cp := make([]string, len(record))
			copy(cp, record)
			return [][]string{cp}
		}

		if record[index] != "?" {
			return recurse(record, index+1)
		}
		possibilities := [][]string{}
		record[index] = "#"
		possibilities = append(possibilities, recurse(record, index+1)...)

		record[index] = "."
		possibilities = append(possibilities, recurse(record, index+1)...)

		record[index] = "?"

		return possibilities
	}

	return recurse(record, 0)
}

func checkIfSpringRecordFitsDamagedGroupCounts(condition []string, damagedGroupCounts []int) bool {
	consecutiveDamagedCount := 0
	foundDamageGroupCounts := []int{}
	for _, cond := range condition {
		if cond == "." {
			if consecutiveDamagedCount != 0 {
				foundDamageGroupCounts = append(foundDamageGroupCounts, consecutiveDamagedCount)
			}
			consecutiveDamagedCount = 0
		} else {
			consecutiveDamagedCount++
		}
	}
	if consecutiveDamagedCount != 0 {
		foundDamageGroupCounts = append(foundDamageGroupCounts, consecutiveDamagedCount)
	}

	if len(damagedGroupCounts) == len(foundDamageGroupCounts) {
		for i := 0; i < len(damagedGroupCounts); i++ {
			if damagedGroupCounts[i] != foundDamageGroupCounts[i] {
				return false
			}
		}
		return true
	}

	return false
}

func part2(input string) int {
	// brute force will not work for part 2 presumably. 2^10 becomes 2^50 which is 1 trillion times larger?

	stringConditions := parseInput(input)

	// hacky hacky way to update string conditions...
	for i, sc := range stringConditions {
		for x := 0; x < 4; x++ {
			stringConditions[i].record = append(stringConditions[i].record, "?")
			stringConditions[i].record = append(stringConditions[i].record, sc.record...)
			stringConditions[i].damagedGroupCounts = append(stringConditions[i].damagedGroupCounts, sc.damagedGroupCounts...)
		}
		// adding a "." at the end helps future logic ensure that the final damaged group will be ended
		stringConditions[i].record = append(stringConditions[i].record, ".")
	}

	ans := 0
	for _, sc := range stringConditions {
		memoOfPossibilities := map[[3]int]int{}
		ans += memo(sc, 0, 0, 0, memoOfPossibilities)
	}

	return ans
}

func memo(sc springCondition, index, doneGroups, currentGroupSize int, memoOfPossibilities map[[3]int]int) int {
	// key of 0, 0, 0 holds final answer of possible results
	key := [3]int{index, doneGroups, currentGroupSize}

	if ans, ok := memoOfPossibilities[key]; ok {
		return ans
	}

	// if the end of the record is reached, and all damaged groups are accounted for
	// do not need to check for currentGroupSize because of the trailing "." that was added
	if index == len(sc.record) && doneGroups == len(sc.damagedGroupCounts) {
		memoOfPossibilities[key] = 1
		return 1
	}

	// any other scenario where we've reached the final index means this possibility is invalid
	if index == len(sc.record) {
		memoOfPossibilities[key] = 0
		return 0
	}

	// damaged spring groups are all accounted for but ran into an additional broken spring,
	// this branch is not valid
	if doneGroups == len(sc.damagedGroupCounts) && sc.record[index] == "#" {
		memoOfPossibilities[key] = 0
		return 0
	}

	// handle ".", "#" or "?"
	possibilities := 0
	if sc.record[index] == "." {
		// end the previous group
		if index == 0 {
			possibilities = memo(sc, index+1, 0, 0, memoOfPossibilities)
		} else if currentGroupSize == 0 {
			possibilities = memo(sc, index+1, doneGroups, 0, memoOfPossibilities)
		} else if currentGroupSize != 0 {
			// we have a non-zero current group size so if all damaged groups are accounted for,
			// there are no possibilities left for this branch
			if doneGroups == len(sc.damagedGroupCounts) {
				possibilities = 0
			} else {
				// not all damaged groups are accounted for
				// if the current group is the right size, recurse; if not, then zero possibilities remain
				if currentGroupSize == sc.damagedGroupCounts[doneGroups] {
					possibilities = memo(sc, index+1, doneGroups+1, 0, memoOfPossibilities)
				} else if currentGroupSize != sc.damagedGroupCounts[doneGroups] {
					// last group is the wrong size, zero possibilities for this branch
					possibilities = 0
				}
			}
		}

	} else if sc.record[index] == "#" {
		// build group
		currentGroupSize++
		// if current group size is too big, this branch has zero possibilities
		if currentGroupSize > sc.damagedGroupCounts[doneGroups] {
			possibilities = 0
		} else {
			possibilities = memo(sc, index+1, doneGroups, currentGroupSize, memoOfPossibilities)
		}

	} else if sc.record[index] == "?" {
		// ?
		// add two possibilities: a damaged spring or OK spring

		// if it is a #
		// do not need to account for if the group is too big here, it'll be handled by a future "#"
		// check or a ".", again part of the reason why a trailing period was added
		possibilities += memo(sc, index+1, doneGroups, currentGroupSize+1, memoOfPossibilities)
		// currentGroupSize--

		// take as .
		// same code as above for if "." block, but possibilities is added to instead of just set
		if index == 0 {
			possibilities += memo(sc, index+1, 0, 0, memoOfPossibilities)
		} else if currentGroupSize == 0 {
			possibilities += memo(sc, index+1, doneGroups, currentGroupSize, memoOfPossibilities)
		} else {
			if doneGroups == len(sc.damagedGroupCounts) {
				possibilities += 0
			} else {
				if currentGroupSize == sc.damagedGroupCounts[doneGroups] {
					possibilities += memo(sc, index+1, doneGroups+1, 0, memoOfPossibilities)
				} else if currentGroupSize != sc.damagedGroupCounts[doneGroups] {
					possibilities += 0
				}
			}
		}

	} else {
		panic("unexpected string condition record character: " + sc.record[index])
	}

	memoOfPossibilities[key] = possibilities
	return possibilities
}

type springCondition struct {
	record             []string
	damagedGroupCounts []int
}

func parseInput(input string) (ans []springCondition) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		sc := springCondition{
			record:             strings.Split(parts[0], ""),
			damagedGroupCounts: []int{},
		}
		for _, str := range strings.Split(parts[1], ",") {
			sc.damagedGroupCounts = append(sc.damagedGroupCounts, cast.ToInt(str))
		}
		ans = append(ans, sc)
	}

	return ans
}
