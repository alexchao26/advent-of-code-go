package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `Generator A starts with 65
Generator B starts with 8921`

func Test_duelingGenerators(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example_part1", example, 1, 588},
		{"actual_part1", util.ReadFile("input.txt"), 1, 592},
		{"example_part2", example, 2, 309},
		{"actual_part2", util.ReadFile("input.txt"), 2, 320},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := duelingGenerators(tt.input, tt.part); got != tt.want {
				t.Errorf("duelingGenerators() = %v, want %v", got, tt.want)
			}
		})
	}
}
