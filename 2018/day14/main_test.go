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
	{"example1", 5158916779, "9"},
	{"example2", 124515891, "5"}, // first digit is a zero... my solution doesn't really account for this
	{"example2", 9251071085, "18"},
	{"example2", 5941429882, "2018"},
	{"actual", 3147574107, util.ReadFile("input.txt")},
}

func TestPart1(t *testing.T) {
	for _, test := range tests1 {
		t.Run(test.name, func(*testing.T) {
			got := part1(test.input)
			if got != test.want {
				t.Errorf("after %s recipes, got %v, want %v", test.input, got, test.want)
			}
		})
	}
}

var tests2 = []struct {
	name  string
	want  int
	input string
}{
	{"example1", 9, "51589"},
	{"example2", 18, "92510"},
	{"example2", 2018, "59414"},
	{"actual", 20280190, util.ReadFile("input.txt")},
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
