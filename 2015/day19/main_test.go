package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `H => HO
H => OH
O => HH

HOH`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", example, 4},
		{"example", example + "OHO", 7},
		{"actual", util.ReadFile("input.txt"), 576},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

var part2Example = `e => H
e => O
H => HO
H => OH
O => HH

HOH`

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", part2Example, 3},
		{"example", part2Example + "OHO", 6},
		{"actual", util.ReadFile("input.txt"), 207},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
