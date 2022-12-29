package main

import (
	"testing"
)

var example = `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

func Test_unstableDiffusion(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{
			name: "small example",
			input: `.....
..##.
..#..
.....
..##.
.....`,
			part: 1,
			want: 30 - 5,
		},
		{
			name:  "example",
			input: example,
			part:  1,
			want:  110,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  4116,
		},

		//
		{
			name:  "example",
			input: example,
			part:  2,
			want:  20,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  984,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unstableDiffusion(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
