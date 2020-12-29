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
		{"actual", util.ReadFile("input.txt"), 1, 38813},
		{"actual", util.ReadFile("input.txt"), 2, 141071},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
