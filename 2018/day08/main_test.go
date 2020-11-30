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
	{"example", 138, "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"},
	{"actual", 48155, util.ReadFile("input.txt")},
}

func TestPart1(t *testing.T) {
	for _, test := range tests1 {
		t.Run(test.name, func(*testing.T) {
			got := part1(test.input)
			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

var tests2 = []struct {
	name  string
	want  int
	input string
}{
	{"example", 66, "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"},
	{"actual", 40292, util.ReadFile("input.txt")},
}

func TestPart2(t *testing.T) {
	for _, test := range tests2 {
		t.Run(test.name, func(*testing.T) {
			got := part2(test.input)
			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}
