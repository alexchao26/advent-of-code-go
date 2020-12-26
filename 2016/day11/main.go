package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := rtgHellDay(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func rtgHellDay(input string, part int) int {
	currentState := newInitialState(input)

	if part == 2 {
		currentState.floors[0] = append(currentState.floors[0],
			halves{isChip: false, material: "elerium"},
			halves{isChip: true, material: "elerium"},
			halves{isChip: false, material: "dilithium"},
			halves{isChip: true, material: "dilithium"},
		)
	}

	queue := []state{currentState}
	prevStates := map[string]bool{}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if front.isDone() {
			return front.steps
		}

		// do not visit previous states
		// hashKey method does not differentiate material types because
		// they are effectively the same
		hash := front.hashKey()
		if prevStates[hash] {
			continue
		}
		prevStates[hash] = true

		nextStates := front.getNextStates()
		queue = append(queue, nextStates...)
	}

	return -1
}

// halves are either a chip or generator
type halves struct {
	isChip   bool // false if is generator
	material string
}

// for easier debugging
func (t halves) String() string {
	tType := " generator"
	if t.isChip {
		tType = " microchip"
	}
	return fmt.Sprint(t.material, tType)
}

// state of the puzzle with a bunch of methods for getting next states, checking
// validity of a state, if it represents a finish state...
type state struct {
	floors        [4][]halves
	elevatorLevel int
	steps         int
}

// parsing the input file, this probably would've been easier to do manually...
func newInitialState(input string) state {
	s := state{}

	for lineIndex, line := range strings.Split(input, "\n") {
		// The first floor contains a promethium generator and a promethium-compatible microchip.
		parts := strings.Split(line, " ")
		// trim commas and periods, this input is pretty inconsistent
		for i, v := range parts {
			parts[i] = strings.Trim(v, ",.")
		}

		// iterate through the words and if generator or microchip is found
		// then parse the previous word for the material type
		for i, word := range parts {
			if word == "generator" {
				material := parts[i-1]
				s.floors[lineIndex] = append(s.floors[lineIndex], halves{
					isChip:   false,
					material: material,
				})
			} else if word == "microchip" {
				// also parse off the "-compatible" portion
				material := parts[i-1][:strings.Index(parts[i-1], "-comp")]
				s.floors[lineIndex] = append(s.floors[lineIndex], halves{
					isChip:   true,
					material: material,
				})
			}
		}
	}

	return s
}

// for printability & debugging
func (s state) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Level %d x Steps %d\n", s.elevatorLevel, s.steps)
	for i, f := range s.floors {
		fmt.Fprintf(&sb, "  %d: %v\n", i, f)
	}
	return sb.String()
}

// Generates a hash for this state that is comparable, to not repeat states.
// I spent over an hour figuring out that I left out the elevator level, which
// is a key component of the state hash
func (s state) hashKey() string {
	// get the indices for each generator and chip
	mapGenToIndex := map[string]int{}
	mapChipToIndex := map[string]int{}
	for flIndex, fl := range s.floors {
		for _, half := range fl {
			if half.isChip {
				mapChipToIndex[half.material] = flIndex
			} else {
				mapGenToIndex[half.material] = flIndex
			}
		}
	}

	// then put that into slice form so it ignores the material types
	// this is b/c the types don't really matter... e.g.
	// 0: LithGen, LithChip      0: PluGen, PluChip
	// 1: PluGen             ==  1: LithGen
	// 2: PluChip                2: LithChip
	var genChipPairs [][2]int
	for material := range mapGenToIndex {
		genChipPairs = append(genChipPairs, [2]int{
			mapGenToIndex[material], mapChipToIndex[material],
		})
	}
	// sort it
	sort.Slice(genChipPairs, func(i, j int) bool {
		if genChipPairs[i][0] != genChipPairs[j][0] {
			return genChipPairs[i][0] < genChipPairs[j][0]
		}
		return genChipPairs[i][1] < genChipPairs[j][1]
	})

	// fmt.Sprint is my best friend for making hashes
	return fmt.Sprint(s.elevatorLevel, genChipPairs)
}

func (s state) isValid() bool {
	// check every level, I lost another hour here because I was only checking
	// the active level, but moving some halves off a level could make an old
	// level invalid
	for i := range s.floors {
		// make a hashmap of all the generators on this level
		gensSeen := map[string]bool{}
		for _, half := range s.floors[i] {
			if !half.isChip {
				gensSeen[half.material] = true
			}
		}
		// if there are no gens on this level, it's safe
		if len(gensSeen) == 0 {
			continue
		}

		// there are generators, so if there is any chip that is not protected
		// then it is an invalid level & thus an invalid state
		for _, half := range s.floors[i] {
			if half.isChip && !gensSeen[half.material] {
				return false
			}
		}
	}
	// all chips protected, return true
	return true
}

// is this the final state?
func (s state) isDone() bool {
	var lenSum int
	for _, fl := range s.floors[:3] {
		lenSum += len(fl)
	}
	return lenSum == 0
}

// get perms of INDICES that can be moved off of this level in the next perm
// one or two things can be moved, this generates all perms of one or two
// indices on the elevator's floor
func (s state) getMovablePermIndices() [][]int {
	var permsToMove [][]int

	currentLevel := s.floors[s.elevatorLevel]
	// get pairs first
	for i := 0; i < len(currentLevel); i++ {
		for j := i + 1; j < len(currentLevel); j++ {
			permsToMove = append(permsToMove, []int{i, j})
		}
	}
	// then get singles
	for i := range currentLevel {
		permsToMove = append(permsToMove, []int{i})
	}
	return permsToMove
}

// make a deep clone
func (s state) clone() state {
	cl := state{
		elevatorLevel: s.elevatorLevel,
		steps:         s.steps,
	}
	// slices are effectively reference types in go... they need to be cloned...
	for i, fl := range s.floors {
		cl.floors[i] = append([]halves{}, fl...)
	}
	return cl
}

// get all valid next states that can be reached from this state
func (s state) getNextStates() []state {
	var futureStates []state

	// all combinations of indices that can be moved from this level
	movablePermIndices := s.getMovablePermIndices()

	// get diffs that the elevator can move in, i.e. don't let it move up from
	// the top level, or down from level 0
	var eleDiffs []int
	if s.elevatorLevel < len(s.floors)-1 {
		eleDiffs = append(eleDiffs, 1)
	}
	if s.elevatorLevel > 0 {
		eleDiffs = append(eleDiffs, -1)
	}

	for _, eleDiff := range eleDiffs {
		// for any elevator direction, iterate over moveable perms and make a
		// clone that's modified with those halves moved to the target floor
		for _, permIndices := range movablePermIndices {
			cl := s.clone()
			cl.elevatorLevel += eleDiff
			cl.steps++ // increment steps
			oldLevel := s.elevatorLevel
			newLevel := cl.elevatorLevel

			// move halves to the clone's active level
			for _, index := range permIndices {
				cl.floors[newLevel] = append(cl.floors[newLevel], cl.floors[oldLevel][index])
			}
			// remove halves from the current state's level (in the clone)
			// this code is gross...
			for in := len(permIndices) - 1; in >= 0; in-- {
				cl.floors[oldLevel][permIndices[in]] = cl.floors[oldLevel][len(cl.floors[oldLevel])-1]
				cl.floors[oldLevel] = cl.floors[oldLevel][:len(cl.floors[oldLevel])-1]
			}
			// add to final states if its valid
			if cl.isValid() {
				futureStates = append(futureStates, cl)
			}
		}
	}

	return futureStates
}
