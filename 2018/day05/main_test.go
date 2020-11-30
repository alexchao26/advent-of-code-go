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
	{"ex1", 10, "dabAcCaCBAcCcaDA"},
	{"actual", 10180, util.ReadFile("input.txt")},
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
	{"ex1", 4, "dabAcCaCBAcCcaDA"},
	{"actual", 5668, util.ReadFile("input.txt")},
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
