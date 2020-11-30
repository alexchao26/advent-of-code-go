package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var exampleInput = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
`

var tests1 = []struct {
	name  string
	want  string
	input string
}{
	{"example", "CABDFE", exampleInput},
	{"actual", "JMQZELVYXTIGPHFNSOADKWBRUC", util.ReadFile("input.txt")},
}

func TestPart1(t *testing.T) {
	for _, test := range tests1 {
		t.Run(test.name, func(*testing.T) {
			got := part1(test.input)
			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

var tests2 = []struct {
	name      string
	want      int
	input     string
	workers   int
	fudgeTime int
}{
	{"example", 15, exampleInput, 2, 0},
	{"actual", 1133, util.ReadFile("input.txt"), 5, 60},
}

func TestPart2(t *testing.T) {
	for _, test := range tests2 {
		t.Run(test.name, func(*testing.T) {
			got := part2(test.input, test.workers, test.fudgeTime)
			if got != test.want {
				t.Errorf("got %v; want %v", got, test.want)
			}
		})
	}
}

func TestTimeForStep(t *testing.T) {
	for _, test := range []struct {
		input       string
		want        int
		fudgeFactor int
	}{
		{"A", 1, 0},
		{"F", 6, 0},
		{"J", 10, 0},
		{"Z", 26, 0},
		{"A", 61, 60},
		{"F", 66, 60},
		{"J", 70, 60},
		{"Z", 86, 60},
	} {
		if got := timeForStep(test.input, test.fudgeFactor); got != test.want {
			t.Errorf("timeForStep(%s, %d) = %d, want %d", test.input, test.fudgeFactor, got, test.want)
		}
	}
}
