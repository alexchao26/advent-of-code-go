package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_reindeerOlympics(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"part1 actual", util.ReadFile("input.txt"), 1, 2660},
		{"part2 actual", util.ReadFile("input.txt"), 2, 1256},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reindeerOlympics(tt.input, tt.part); got != tt.want {
				t.Errorf("reindeerOlympics() = %v, want %v", got, tt.want)
			}
		})
	}
}
