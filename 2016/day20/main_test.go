package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `5-8
0-2
4-7`

func Test_firewall(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example part1", example, 1, 3},
		{"actual part1", util.ReadFile("input.txt"), 1, 23923783},
		{"actual part2", util.ReadFile("input.txt"), 2, 125},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firewall(tt.input, tt.part); got != tt.want {
				t.Errorf("firewall() = %v, want %v", got, tt.want)
			}
		})
	}
}
