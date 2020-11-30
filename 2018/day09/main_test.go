package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var tests1 = []struct {
	name  string
	want  int
	input string
	// add extra args if needed
}{
	{"example1", 8317, "10 players; last marble is worth 1618 points"},
	{"example2", 146373, "13 players; last marble is worth 7999 points"},
	{"example3", 2764, "17 players; last marble is worth 1104 points"},
	{"example4", 54718, "21 players; last marble is worth 6111 points"},
	{"example5", 37305, "30 players; last marble is worth 5807 points"},
	{"actual", 394486, util.ReadFile("input.txt")},
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
	// add extra args if needed
}{
	{"actual", 3276488008, util.ReadFile("input.txt")},
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
