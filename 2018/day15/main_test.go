package main

import "testing"

var exampleInput1 = `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`
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
	// {"actual", ACTUAL_ANSWER, util.ReadFile("input.txt")},
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

// var tests2 = []struct {
// 	name  string
// 	want  int
// 	input string
// 	// add extra args if needed
// }{
// 	// {"actual", ACTUAL_ANSWER, util.ReadFile("input.txt")},
// }

// func TestPart2(t *testing.T) {
// 	for _, test := range tests2 {
// 		t.Run(test.name, func(*testing.T) {
// 			got := part2(test.input)
// 			if got != test.want {
// 				t.Errorf("got %v, want %v", got, test.want)
// 			}
// 		})
// 	}
// }
