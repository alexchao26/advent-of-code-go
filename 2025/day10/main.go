package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"sort"
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
	machines := parseInput(input)

	// bfs it
	type node struct {
		state []bool
		steps int
	}

	ans := 0

	for _, m := range machines {

		queue := []node{
			node{
				state: make([]bool, len(m.lights)),
				steps: 0,
			},
		}
		seenStates := map[string]bool{}

		for len(queue) > 0 {
			front := queue[0]
			queue = queue[1:]

			if fmt.Sprint(front.state) == fmt.Sprint(m.lights) {
				ans += front.steps
				break
			}

			for _, buttonList := range m.buttons {
				stateCopy := make([]bool, len(front.state))
				copy(stateCopy, front.state)
				for _, button := range buttonList {
					stateCopy[button] = !stateCopy[button]
				}
				key := fmt.Sprint(stateCopy)
				if seenStates[key] {
					continue
				}
				seenStates[key] = true

				queue = append(queue, node{
					state: stateCopy,
					steps: front.steps + 1,
				})
			}
		}

	}
	return ans
}

func part2(input string) int {
	machines := parseInput(input)

	ans := 0

	for _, m := range machines {
		combinations := makeAllButtonCombinations(m.buttons, len(m.joltage))
		n, _ := State(m.joltage).solve2(combinations)
		ans += n
	}
	return ans
}

type State []int

type ButtonCombination struct {
	counter          State
	nbPressedButtons int
}

func makeAllButtonCombinations(buttons [][]int, m int) []ButtonCombination {
	var res []ButtonCombination
	var recurse func(index int, current State, pressed int)
	recurse = func(index int, current State, pressed int) {
		if index == len(buttons) {
			res = append(res, ButtonCombination{append([]int(nil), current...), pressed})
			return
		}
		// not press
		recurse(index+1, current, pressed)
		// press
		newCurrent := make([]int, len(current))
		copy(newCurrent, current)
		for _, idx := range buttons[index] {
			newCurrent[idx]++
		}
		recurse(index+1, newCurrent, pressed+1)
	}
	recurse(0, make([]int, m), 0)
	return res
}

func newState(size int) State {
	return make([]int, size)
}

func (s State) isSmallerOrEqual(b State) bool {
	for i := range len(s) {
		if s[i] > b[i] {
			return false
		}
	}
	return true
}

func (s State) equalsModulo2(b State) bool {
	for i := range len(s) {
		if s[i]%2 != b[i]%2 {
			return false
		}
	}
	return true
}

func (s State) isZero() bool {
	for i := range len(s) {
		if s[i] != 0 {
			return false
		}
	}
	return true
}

func (counter State) solve2(combinaisons []ButtonCombination) (int, bool) {
	if counter.isZero() {
		return 0, true
	}

	res := math.MaxInt
	for _, comb := range combinaisons {
		if !comb.counter.isSmallerOrEqual(counter) {
			continue
		}
		if !comb.counter.equalsModulo2(counter) {
			continue
		}

		nextState := newState(len(counter))
		for i := range len(counter) {
			nextState[i] = (counter[i] - comb.counter[i]) / 2
		}
		rec, ok := nextState.solve2(combinaisons)
		if !ok {
			continue
		}

		if n := 2*rec + comb.nbPressedButtons; n < res {
			res = n
		}
	}
	if res < math.MaxInt {
		return res, true
	}
	return 0, false
}

type machine struct {
	lights  []bool
	buttons [][]int
	joltage []int
}

func parseInput(input string) []machine {
	var ans []machine
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		partZero := parts[0]
		var lights []bool
		for _, l := range strings.Split(partZero[1:len(partZero)-1], "") {
			lights = append(lights, l == "#")
		}

		var buttons [][]int
		for _, part := range parts[1 : len(parts)-1] {
			buttons = append(buttons, []int{})
			for _, str := range strings.Split(part[1:len(part)-1], ",") {
				buttons[len(buttons)-1] = append(buttons[len(buttons)-1], cast.ToInt(str))
			}
		}

		sort.Slice(buttons, func(i, j int) bool {
			return len(buttons[i]) > len(buttons[j])
		})

		lastPart := parts[len(parts)-1]
		var joltage []int
		for _, str := range strings.Split(lastPart[1:len(lastPart)-1], ",") {
			joltage = append(joltage, cast.ToInt(str))
		}

		ans = append(ans, machine{
			lights:  lights,
			buttons: buttons,
			joltage: joltage,
		})
	}
	return ans
}
