package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`

func Test_magneticMoat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example_part1", example, 1, 31},
		{"actual_part1", util.ReadFile("input.txt"), 1, 1868},
		{"example_part2", example, 2, 19},
		{"actual_part2", util.ReadFile("input.txt"), 2, 1841},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magneticMoat(tt.input, tt.part); got != tt.want {
				t.Errorf("magneticMoat() = %v, want %v", got, tt.want)
			}
		})
	}
}
