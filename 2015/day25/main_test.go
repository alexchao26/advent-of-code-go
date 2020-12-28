package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_letItSnow(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"actual", util.ReadFile("input.txt"), 19980801},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letItSnow(tt.input); got != tt.want {
				t.Errorf("letItSnow() = %v, want %v", got, tt.want)
			}
		})
	}
}
