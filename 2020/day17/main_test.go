package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `.#.
..#
###`

func Test_conwayCubes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example_part1", example, 1, 112},
		{"actual_part1", util.ReadFile("input.txt"), 1, 388},
		{"example_part2", example, 2, 848},
		{"actual_part2", util.ReadFile("input.txt"), 2, 2280},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conwayCubes(tt.input, tt.part); got != tt.want {
				t.Errorf("conwayCubes() = %v, want %v", got, tt.want)
			}
		})
	}
}
