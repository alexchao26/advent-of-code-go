package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_simpleAssemblyComputer(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"actual", util.ReadFile("input.txt"), 1, 307},
		{"actual", util.ReadFile("input.txt"), 2, 160},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simpleAssemblyComputer(tt.input, tt.part); got != tt.want {
				t.Errorf("simpleAssemblyComputer() = %v, want %v", got, tt.want)
			}
		})
	}
}
