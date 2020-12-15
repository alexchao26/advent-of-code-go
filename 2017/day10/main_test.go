package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		listLength int
		want       int
	}{
		{"example", "3,4,1,5", 5, 12},
		{"actual", util.ReadFile("input.txt"), 256, 212},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.listLength); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"example", "1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"actual", util.ReadFile("input.txt"), "96de9657665675b51cd03f0b3528ba26"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
