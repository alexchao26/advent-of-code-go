package main

import (
	"testing"
)

var example = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func Test_laboratories(t *testing.T) {
	tests := []struct {
		name  string
		part  int
		input string
		want  int
	}{
		{
			name:  "example",
			part:  1,
			input: example,
			want:  21,
		},
		{
			name:  "actual",
			part:  1,
			input: input,
			want:  1579,
		},
		{
			name:  "example",
			part:  2,
			input: example,
			want:  40,
		},
		{
			name:  "actual",
			part:  2,
			input: input,
			want:  13418215871354,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := laboratories(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
