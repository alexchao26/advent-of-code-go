package main

import (
	"testing"
)

var example = `Player 1 starting position: 4
Player 2 starting position: 8`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  739785,
		},
		{
			name:  "actual",
			input: input,
			want:  1067724,
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
		want  int64
	}{
		{
			name:  "example",
			input: example,
			want:  444356092776315,
		},
		{
			name:  "actual",
			input: input,
			want:  630947104784464,
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
