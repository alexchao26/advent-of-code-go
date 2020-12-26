package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_lookAndSay(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"actual", util.ReadFile("input.txt"), 1, 252594},
		{"actual", util.ReadFile("input.txt"), 2, 3579328},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lookAndSay(tt.input, tt.part); got != tt.want {
				t.Errorf("lookAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
