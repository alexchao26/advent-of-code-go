package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_rpgSimulator(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 int
		want2 int
	}{
		{"actual", util.ReadFile("input.txt"), 121, 201},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := rpgSimulator(tt.input)
			if got1 != tt.want1 {
				t.Errorf("rpgSimulator() = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("rpgSimulator() = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
