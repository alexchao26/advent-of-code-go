package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_permPromenade(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  string
	}{
		{"actual", util.ReadFile("input.txt"), 1, "ebjpfdgmihonackl"},
		{"actual", util.ReadFile("input.txt"), 2, "abocefghijklmndp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permPromenade(tt.input, tt.part); got != tt.want {
				t.Errorf("permPromenade() = %v, want %v", got, tt.want)
			}
		})
	}
}
