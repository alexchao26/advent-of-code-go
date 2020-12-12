package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var exampleInput = `F10
N3
F7
R90
F11`

var tests1 = []struct {
	name  string
	want  int
	input string
}{
	{"example", 25, exampleInput},
	{"actual", 820, util.ReadFile("input.txt")},
}

func TestPart1(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

var tests2 = []struct {
	name  string
	want  int
	input string
}{
	{"example", 286, exampleInput},
	{"actual", 66614, util.ReadFile("input.txt")},
}

func TestPart2(t *testing.T) {
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
