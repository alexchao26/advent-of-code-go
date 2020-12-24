package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_dragonChecksum(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  string
	}{
		{"part1", util.ReadFile("input.txt"), 1, "10010110010011110"},
		{"part2", util.ReadFile("input.txt"), 2, "01101011101100011"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dragonChecksum(tt.input, tt.part); got != tt.want {
				t.Errorf("dragonChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
