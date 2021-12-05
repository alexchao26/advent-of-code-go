package main

import (
	"testing"
)

var example = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{
			name:  "example",
			input: example,
			part:  1,
			want:  5,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  4993,
		},
		{
			name:  "example",
			input: example,
			part:  2,
			want:  12,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  21101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.part == 0 {
				t.Error("part number cannot be zero")
			}
			if got := countIntersections(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
