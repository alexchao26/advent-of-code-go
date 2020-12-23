package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_decompressLength(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example part1_0", "ADVENT", 1, len("ADVENT")},
		{"example part1_1", "A(1x5)BC", 1, len("ABBBBBC")},
		{"example part1_2", "(6x1)(1x3)A", 1, len("(1x3)A")},
		{"actual part1", util.ReadFile("input.txt"), 1, 107035},
		{"example part2_1", "X(8x2)(3x3)ABCY", 2, len("XABCABCABCABCABCABCY")},
		{"example part2_2", "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 2, 445},
		{"actual part2", util.ReadFile("input.txt"), 2, 11451628995},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompressLength(tt.input, tt.part); got != tt.want {
				t.Errorf("decompressLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
