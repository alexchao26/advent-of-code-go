package main

import (
	"testing"
)

var example = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func Test_part1(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expansionFactor int
		want            int
	}{
		{
			name:            "example",
			input:           example,
			expansionFactor: 2,
			want:            374,
		},
		{
			name:            "actual",
			input:           input,
			expansionFactor: 2,
			want:            9734203,
		},

		// part 2
		{
			name:            "example",
			input:           example,
			expansionFactor: 10,
			want:            1030,
		},
		{
			name:            "example",
			input:           example,
			expansionFactor: 100,
			want:            8410,
		},
		{
			name:            "actual",
			input:           input,
			expansionFactor: 1000000,
			want:            568914596391,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.expansionFactor); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
