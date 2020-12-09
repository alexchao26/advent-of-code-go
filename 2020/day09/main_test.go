package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var exampleInput = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

var tests1 = []struct {
	name           string
	want           int
	input          string
	preambleLength int
}{
	{"example", 127, exampleInput, 5},
	{"actual", 15690279, util.ReadFile("input.txt"), 25},
}

func TestPart1(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.preambleLength); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

var tests2 = []struct {
	name           string
	want           int
	input          string
	preambleLength int
}{
	{"example", 62, exampleInput, 5},
	{"actual", 2174232, util.ReadFile("input.txt"), 25},
}

func TestPart2(t *testing.T) {
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, tt.preambleLength); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
