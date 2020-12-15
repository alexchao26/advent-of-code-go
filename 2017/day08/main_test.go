package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func Test_calcRegisters(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example-part1", example, 1, 1},
		{"actual-part1", util.ReadFile("input.txt"), 1, 4066},
		{"example-part2", example, 2, 10},
		{"actual-part2", util.ReadFile("input.txt"), 2, 4829},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcRegisters(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
