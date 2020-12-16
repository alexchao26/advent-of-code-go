package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_diskDefrag(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example_part1", "flqrgnkx", 1, 8108},
		{"actual_part1", util.ReadFile("input.txt"), 1, 8204},
		{"example_part2", "flqrgnkx", 2, 1242},
		// {"actual_part2", util.ReadFile("input.txt"), 2, 1089},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diskDefrag(tt.input, tt.part); got != tt.want {
				t.Errorf("diskDefrag() = %v, want %v", got, tt.want)
			}
		})
	}
}
