package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", "[1,2,3]", 6},
		{"example", "{\"a\":[-1,1]}", 0},
		{"actual", util.ReadFile("input.txt"), 119433},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"flat example 1", "[1,2,\"red\",5]", 8},
		{"flat example 2", "5", 5},
		{"actual", util.ReadFile("input.txt"), 68466},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
