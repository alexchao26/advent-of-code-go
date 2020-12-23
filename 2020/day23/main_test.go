package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_crabCups(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example_part1", "389125467", 1, 67384529},
		{"actual_part1", util.ReadFile("input.txt"), 1, 47382659},
		{"example_part2", "389125467", 2, 149245887792},
		{"actual_part2", util.ReadFile("input.txt"), 2, 42271866720},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crabCups(tt.input, tt.part); got != tt.want {
				t.Errorf("crabCups() = %v, want %v", got, tt.want)
			}
		})
	}
}
