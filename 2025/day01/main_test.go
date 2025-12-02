package main

import (
	"testing"
)

var example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  3,
		},
		{
			name:  "actual",
			input: input,
			want:  1084,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  6,
		},
		{
			name:  "L50",
			input: "L50",
			want:  1,
		},
		{
			name:  "R50",
			input: "R50",
			want:  1,
		},
		{
			name:  "L50R200",
			input: "L50\nR200",
			want:  3,
		},
		{
			// ugh off by ones in the "L" branch
			name:  "L50L200",
			input: "L50\nL200",
			want:  3,
		},
		{
			name:  "actual",
			input: input,
			want:  6475,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
