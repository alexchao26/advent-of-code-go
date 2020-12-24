package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example_part1", "abc", 1, 22728},
		{"actual_part1", util.ReadFile("input.txt"), 1, 23769},
		{"example_part2", "abc", 2, 22551},
		{"actual_part2", util.ReadFile("input.txt"), 2, 20606},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneTimePad(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
