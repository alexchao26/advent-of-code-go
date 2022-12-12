package main

import (
	"testing"
)

var example = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  13,
		},
		{
			name:  "actual",
			input: input,
			want:  6236,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
		t.Run("reimplementation_"+tt.name, func(t *testing.T) {
			if got := reImplPart1(tt.input); got != tt.want {
				t.Errorf("reImplPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

var largerExample = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  1,
		},
		{
			name:  "larger_example",
			input: largerExample,
			want:  36,
		},
		{
			name:  "actual",
			input: input,
			want:  2449,
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
