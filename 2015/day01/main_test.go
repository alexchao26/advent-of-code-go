package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_notQuiteLisp(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"actual", util.ReadFile("input.txt"), 1, 232},
		{"actual", util.ReadFile("input.txt"), 2, 1783},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := notQuiteLisp(tt.input, tt.part); got != tt.want {
				t.Errorf("notQuiteLisp() = %v, want %v", got, tt.want)
			}
		})
	}
}
