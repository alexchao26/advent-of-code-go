package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_md5Chess(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  string
	}{
		{"example_part1", "abc", 1, "18f47a30"},
		{"actual_part1", util.ReadFile("input.txt"), 1, "801b56a7"},
		{"example_part2", "abc", 2, "05ace8e3"},
		{"actual_part2", util.ReadFile("input.txt"), 2, "424a0197"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5Chess(tt.input, tt.part); got != tt.want {
				t.Errorf("md5Chess() = %v, want %v", got, tt.want)
			}
		})
	}
}
