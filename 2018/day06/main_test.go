package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var arg1 = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

func TestPart1(t *testing.T) {
	// Examples
	want := 17
	got := part1(arg1)
	if got != want {
		t.Errorf("arg1: wanted %d, got %d", want, got)
	}

	// Run actual problem input
	want = 5333
	got = part1(util.ReadFile("input.txt"))
	if got != want {
		t.Errorf("actual AOC input, wanted %d, got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	// Examples
	distArg := 32
	want := 16
	got := part2(arg1, distArg)
	if got != want {
		t.Errorf("part2(arg1, %v): want %v, got %v", distArg, want, got)
	}

	// Run actual problem input
	// part2(util.ReadFile("input.txt"))
}
