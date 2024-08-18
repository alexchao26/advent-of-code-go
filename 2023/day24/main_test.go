package main

import (
	"testing"
)

var example = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

func Test_part1(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		testRange [2]float64
		want      int
	}{
		{
			name:      "example",
			input:     example,
			testRange: [2]float64{7, 27},
			want:      2,
		},
		{
			name:      "actual",
			input:     input,
			testRange: [2]float64{200000000000000, 400000000000000},
			want:      31921,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.testRange); got != tt.want {
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
		// example input is not big enough for this logic to work
		// {
		// 	name:  "example",
		// 	input: example,
		// 	want:  47,
		// },
		{
			name:  "actual",
			input: input,
			want:  761691907059631,
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
