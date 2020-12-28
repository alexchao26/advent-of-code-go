package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_balancingPackages(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"part1 actual", util.ReadFile("input.txt"), 1, 10439961859},
		{"part2 actual", util.ReadFile("input.txt"), 2, 72050269},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancingPackages(tt.input, tt.part); got != tt.want {
				t.Errorf("balancingPackages() = %v, want %v", got, tt.want)
			}
		})
	}
}
