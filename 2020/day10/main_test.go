package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var exampleInputs = []string{`16
10
15
5
1
11
7
19
6
12
4`,
	`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`,
}

var tests1 = []struct {
	name  string
	want  int
	input string
}{
	{"example 1", 35, exampleInputs[0]},
	{"example 2", 220, exampleInputs[1]},
	{"actual", 2176, util.ReadFile("input.txt")},
}

func TestPart1(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

var tests2 = []struct {
	name  string
	want  int
	input string
}{
	{"example 1", 8, exampleInputs[0]},
	{"example 2", 19208, exampleInputs[1]},
	{"actual", 18512297918464, util.ReadFile("input.txt")},
}

func TestPart2(t *testing.T) {
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dynamicProgramming(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		input string
	}{
		{"example 1", 8, exampleInputs[0]},
		{"example 2", 19208, exampleInputs[1]},
		{"actual", 18512297918464, util.ReadFile("input.txt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dynamicProgramming(tt.input); got != tt.want {
				t.Errorf("dynamicProgramming() = %v, want %v", got, tt.want)
			}
		})
	}
}
