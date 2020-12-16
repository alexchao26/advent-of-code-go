package main

import (
	"testing"
	"time"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `0: 3
1: 2
4: 4
6: 4`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", example, 24},
		{"actual", util.ReadFile("input.txt"), 2508},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", example, 10},
		{"actual", util.ReadFile("input.txt"), 3913186},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
			t.Logf("Runtime for %s: %v", tt.name, time.Since(startTime))
		})
	}
}
