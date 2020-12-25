package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example", example, 1, 3},
		{"actual", util.ReadFile("input.txt"), 1, 11893},
		{"actual", util.ReadFile("input.txt"), 2, 479008453},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assemblyComputer(tt.input, tt.part); got != tt.want {
				t.Errorf("assemblyComputer() = %v, want %v", got, tt.want)
			}
		})
	}
}
