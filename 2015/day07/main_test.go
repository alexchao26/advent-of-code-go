package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var exampleOutputH = `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
h -> a` // added last rule to output to a (same as real question)

// Expect these final registers
// d: 72
// e: 507
// f: 492
// g: 114
// h: 65412
// i: 65079
// x: 123
// y: 456

func Test_someAssemblyRequired(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example h -> a", exampleOutputH, 1, 65412},
		{"part1 actual", util.ReadFile("input.txt"), 1, 16076},
		{"part2 actual", util.ReadFile("input.txt"), 2, 2797},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := someAssemblyRequired(tt.input, tt.part); got != tt.want {
				t.Errorf("someAssemblyRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
