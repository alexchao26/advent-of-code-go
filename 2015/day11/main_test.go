package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_passwordIncrementing(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  string
	}{
		{"part1 actual", util.ReadFile("input.txt"), 1, "hxbxxyzz"},
		{"part2 actual", util.ReadFile("input.txt"), 2, "hxcaabcc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := passwordIncrementing(tt.input, tt.part); got != tt.want {
				t.Errorf("passwordIncrementing() = %v, want %v", got, tt.want)
			}
		})
	}
}
