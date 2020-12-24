package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_elephant(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"part1 actual", util.ReadFile("input.txt"), 1, 1834471},
		{"part2 example", "5", 2, 2},
		{"part2 actual", util.ReadFile("input.txt"), 2, 1420064},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := elephant(tt.input, tt.part); got != tt.want {
				t.Errorf("elephant() = %v, want %v", got, tt.want)
			}
		})
	}
}
