package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `Disc #1 has 5 positions; at time=0, it is at position 4.
Disc #2 has 2 positions; at time=0, it is at position 1.`

func Test_timingIsEverything(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example", example, 1, 5},
		{"actual", util.ReadFile("input.txt"), 1, 317371},
		{"actual", util.ReadFile("input.txt"), 2, 2080951},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timingIsEverything(tt.input, tt.part); got != tt.want {
				t.Errorf("timingIsEverything() = %v, want %v", got, tt.want)
			}
		})
	}
}
