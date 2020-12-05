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
	{"example", 7, `#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`},
	{"actual", 1350, util.ReadFile("input.txt")},
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
	// add extra args if needed
}{
	{"actual", 15844608, util.ReadFile("input.txt")},
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

func TestPart2Cheeky(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{"actual", util.ReadFile("input.txt"), 15844608},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2Cheeky(tt.args); got != tt.want {
				t.Errorf("part2Cheeky() = %v, want %v", got, tt.want)
			}
		})
	}
}
