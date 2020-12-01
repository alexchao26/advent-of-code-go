package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var tests1 = []struct {
	name  string
	want  int
	input string
}{
	{"example", 325, `initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
`},
	{"actual", 2542, util.ReadFile("input.txt")},
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

var tests2 = []struct {
	name  string
	want  int
	input string
}{
	{"actual", 2550000000883, util.ReadFile("input.txt")},
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
