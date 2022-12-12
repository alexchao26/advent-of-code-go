package main

import (
	_ "embed"
	"flag"
	"fmt"
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

	if part == 1 {
		ans := part1(true)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(true)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(useRealInput bool) int {
	monkeys := initInput()
	if !useRealInput {
		monkeys = initExample()
	}

	inspectedCounts := make([]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				newItemVal := monkey.operation(item) / 3

				if newItemVal%monkey.testDivisibleBy == 0 {
					monkeys[monkey.trueMonkey].items = append(
						monkeys[monkey.trueMonkey].items, newItemVal)
				} else {
					monkeys[monkey.falseMonkey].items = append(
						monkeys[monkey.falseMonkey].items, newItemVal)
				}

			}
			inspectedCounts[i] += len(monkey.items)

			// empty out this monkey's items
			monkeys[i].items = []int{}
		}
	}

	sort.Ints(inspectedCounts)
	return inspectedCounts[len(inspectedCounts)-1] * inspectedCounts[len(inspectedCounts)-2]
}

// oh my god i figured out a math-y remainder theorem-y thing myself!
func part2(useRealInput bool) int {
	monkeys := initInput()
	if !useRealInput {
		monkeys = initExample()
	}

	// the worry levels will always increase now that they're not being divided
	// by 3, and we care about remainders because that's what all the tests are
	// BUT we can't just mod by any monkey's testBy number, because they're all
	// throwing the items around,
	// so find a shared common denominator that can be used to keep the numbers
	// under overflow
	bigMod := 1
	for _, m := range monkeys {
		bigMod *= m.testDivisibleBy
	}

	inspectedCounts := make([]int, len(monkeys))
	for round := 0; round < 10000; round++ {

		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				newItemVal := monkey.operation(item)
				newItemVal %= bigMod

				if newItemVal%monkey.testDivisibleBy == 0 {
					monkeys[monkey.trueMonkey].items = append(
						monkeys[monkey.trueMonkey].items, newItemVal)
				} else {
					monkeys[monkey.falseMonkey].items = append(
						monkeys[monkey.falseMonkey].items, newItemVal)
				}

			}
			inspectedCounts[i] += len(monkey.items)

			// empty out this monkey's items
			monkeys[i].items = []int{}
		}
	}

	sort.Ints(inspectedCounts)
	return inspectedCounts[len(inspectedCounts)-1] * inspectedCounts[len(inspectedCounts)-2]
}

type monkey struct {
	items                   []int
	operation               func(int) int
	testDivisibleBy         int
	trueMonkey, falseMonkey int // indices
}

// faster to manually type this than write a parser (and potentially debug)
func initInput() []monkey {
	return []monkey{
		{
			items: []int{50, 70, 89, 75, 66, 66},
			operation: func(old int) int {
				return old * 5
			},
			testDivisibleBy: 2,
			trueMonkey:      2,
			falseMonkey:     1,
		},
		{
			items: []int{85},
			operation: func(old int) int {
				return old * old
			},
			testDivisibleBy: 7,
			trueMonkey:      3,
			falseMonkey:     6,
		},
		{
			items: []int{66, 51, 71, 76, 58, 55, 58, 60},
			operation: func(old int) int {
				return old + 1
			},
			testDivisibleBy: 13,
			trueMonkey:      1,
			falseMonkey:     3,
		},
		{
			items: []int{79, 52, 55, 51},
			operation: func(old int) int {
				return old + 6
			},
			testDivisibleBy: 3,
			trueMonkey:      6,
			falseMonkey:     4,
		},
		{
			items: []int{69, 92},
			operation: func(old int) int {
				return old * 17
			},
			testDivisibleBy: 19,
			trueMonkey:      7,
			falseMonkey:     5,
		},
		{
			items: []int{71, 76, 73, 98, 67, 79, 99},
			operation: func(old int) int {
				return old + 8
			},
			testDivisibleBy: 5,
			trueMonkey:      0,
			falseMonkey:     2,
		},
		{
			items: []int{82, 76, 69, 69, 57},
			operation: func(old int) int {
				return old + 7
			},
			testDivisibleBy: 11,
			trueMonkey:      7,
			falseMonkey:     4,
		},
		{
			items: []int{65, 79, 86},
			operation: func(old int) int {
				return old + 5
			},
			testDivisibleBy: 17,
			trueMonkey:      5,
			falseMonkey:     0,
		},
	}
}

func initExample() []monkey {
	return []monkey{
		{
			items: []int{79, 98},
			operation: func(num int) int {
				return num * 19
			},
			testDivisibleBy: 23,
			trueMonkey:      2,
			falseMonkey:     3,
		},
		{
			items: []int{54, 65, 75, 74},
			operation: func(num int) int {
				return num + 6
			},
			testDivisibleBy: 19,
			trueMonkey:      2,
			falseMonkey:     0,
		},
		{
			items: []int{79, 60, 97},
			operation: func(num int) int {
				return num * num
			},
			testDivisibleBy: 13,
			trueMonkey:      1,
			falseMonkey:     3,
		},
		{
			items: []int{74},
			operation: func(num int) int {
				return num + 3
			},
			testDivisibleBy: 17,
			trueMonkey:      0,
			falseMonkey:     1,
		},
	}
}
