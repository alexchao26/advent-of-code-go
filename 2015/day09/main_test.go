package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_travelingSalesman(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantPart1 int
		wantPart2 int
	}{
		{"actual", util.ReadFile("input.txt"), 117, 909},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := travelingSalesman(tt.input)
			if got1 != tt.wantPart1 {
				t.Errorf("travelingSalesman() part1 = %v, want %v", got1, tt.wantPart1)
			}
			if got2 != tt.wantPart2 {
				t.Errorf("travelingSalesman() part2 = %v, want %v", got2, tt.wantPart2)
			}
		})
	}
}
