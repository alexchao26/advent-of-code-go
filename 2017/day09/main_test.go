package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_streamProcessing(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"part1_example 1", "{}", 1, 1},
		{"part1_example 2", "{{}}", 1, 3},
		{"part1_example 1", "{{}, {}}", 1, 5},
		{"part1_example 1", "{{},<aksdljfsd!<> {}}", 1, 5},
		{"part1_actual", util.ReadFile("input.txt"), 1, 16021},
		{"part2_actual", util.ReadFile("input.txt"), 2, 7685},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := streamProcessing(tt.input, tt.part); got != tt.want {
				t.Errorf("streamProcessing() = %v, want %v", got, tt.want)
			}
		})
	}
}
