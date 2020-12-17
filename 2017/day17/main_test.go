package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_spinlock(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example_part1", "3", 1, 638},
		{"actual_part1", util.ReadFile("input.txt"), 1, 1547},
		{"actual_part2", util.ReadFile("input.txt"), 2, 31154878},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running %s", tt.name)
			if tt.name == "actual_part2" {
				t.Log("This one could take a while, it has 50 million steps to run")
			}
			if got := spinlock(tt.input, tt.part); got != tt.want {
				t.Errorf("spinlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
