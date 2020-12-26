package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.`

func Test_rtgHellDay(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"part1 example", example, 1, 11},
		{"part1 actual", util.ReadFile("input.txt"), 1, 33},
		{"part2 actual", util.ReadFile("input.txt"), 2, 57},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rtgHellDay(tt.input, tt.part); got != tt.want {
				t.Errorf("rtgHellDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
