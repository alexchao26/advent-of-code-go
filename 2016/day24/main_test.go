package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

func Test_cleaningRobot(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"part1 example", example, 1, 14},
		{"part1 actual", util.ReadFile("input.txt"), 1, 442},
		{"part2 actual", util.ReadFile("input.txt"), 2, 660},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleaningRobot(tt.input, tt.part); got != tt.want {
				t.Errorf("cleaningRobot() = %v, want %v", got, tt.want)
			}
		})
	}
}
