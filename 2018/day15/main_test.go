package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var exampleInput1 = `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`

// this example is in part 1 but not part 2 -shrug-
var exampleInput2 = `#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`
var exampleInput3 = `#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`
var exampleInput4 = `#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`
var exampleInput5 = `#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`
var exampleInput6 = `#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`

var tests1 = []struct {
	name  string
	input string
	want  int
}{
	{"example1", exampleInput1, 27730},
	{"example2", exampleInput2, 36334},
	{"example3", exampleInput3, 39514},
	{"example4", exampleInput4, 27755},
	{"example5", exampleInput5, 28944},
	{"example6", exampleInput6, 18740},
	{"actual", util.ReadFile("input.txt"), 183300},
}

func TestPart1(t *testing.T) {
	for _, test := range tests1 {
		t.Run(test.name, func(*testing.T) {
			got := part1(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

var reddit1 = `#######
#######
#.E..G#
#.#####
#G#####
#######
#######`

// used a sample test from a reddit thread where I was failing this specific issue
// https://www.reddit.com/r/adventofcode/comments/a6r6kg/2018_day_15_part_1_what_am_i_missing/ebxjjuo?utm_source=share&utm_medium=web2x&context=3
func TestMovement(t *testing.T) {
	t.Log("Expect Elf's first move to go RIGHT")
	part1(reddit1)
}

var tests2 = []struct {
	name  string
	input string
	want  int
}{
	{"example1", exampleInput1, 4988},
	{"example3", exampleInput3, 31284},
	{"example4", exampleInput4, 3478},
	{"example5", exampleInput5, 6474},
	{"example6", exampleInput6, 1140},
	{"actual", util.ReadFile("input.txt"), 40625},
}

func TestPart2(t *testing.T) {
	for _, test := range tests2 {
		t.Run(test.name, func(*testing.T) {
			got := part2(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
