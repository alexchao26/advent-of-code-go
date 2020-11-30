package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var arg1 = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

var tests1 = []struct {
	name  string
	want  int
	input string
}{
	{"example1", 17, arg1},
	{"actual", 5333, util.ReadFile("input.txt")},
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
	name    string
	want    int
	input   string
	distArg int
}{
	{"example1", 16, arg1, 32},
	{"actual", 35334, util.ReadFile("input.txt"), 10000},
}

func TestPart2(t *testing.T) {
	for _, test := range tests2 {
		t.Run(test.name, func(*testing.T) {
			got := part2(test.input, test.distArg)
			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}
