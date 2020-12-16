package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_hexEd(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"actual-part1", util.ReadFile("input.txt"), 1, 794},
		{"actual-part2", util.ReadFile("input.txt"), 2, 1524},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hexEd(tt.input, tt.part); got != tt.want {
				t.Errorf("hexEd() = %v, want %v", got, tt.want)
			}
		})
	}
}
