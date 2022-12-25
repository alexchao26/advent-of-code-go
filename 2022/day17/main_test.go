package main

import (
	"testing"
)

var example = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  3068,
		},
		{
			name:  "actual",
			input: input,
			want:  3219,
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
		name        string
		input       string
		wantedRocks int
		want        int
	}{
		// part 1 values should work as well, so using them for extra testing
		{
			name:        "example",
			input:       example,
			wantedRocks: 2022,
			want:        3068,
		},
		{
			name:        "actual",
			input:       input,
			wantedRocks: 2022,
			want:        3219,
		},
		{
			name:        "example",
			input:       example,
			wantedRocks: 1000000000000,
			want:        1514285714288,
		},
		{
			name:        "actual",
			input:       input,
			wantedRocks: 1000000000000,
			want:        1582758620701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, tt.wantedRocks); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
