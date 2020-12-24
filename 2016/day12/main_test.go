package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_assemblyComputer(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"actual_part1", util.ReadFile("input.txt"), 1, 318077},
		{"actual_part2", util.ReadFile("input.txt"), 2, 9227731},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assemblyComputer(tt.input, tt.part); got != tt.want {
				t.Errorf("assemblyComputer() = %v, want %v", got, tt.want)
			}
		})
	}
}
