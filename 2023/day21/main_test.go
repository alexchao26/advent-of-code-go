package main

import (
	"testing"
)

var example = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		steps int
		want  int
	}{
		{
			name:  "example",
			input: example,
			steps: 6,
			want:  16,
		},
		{
			name:  "actual",
			input: input,
			steps: 64,
			want:  3743,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.steps); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		steps int
		want  int
	}{
		// {
		// 	name:  "example-10",
		// 	input: example,
		// 	steps: 10,
		// 	want:  50,
		// },
		// {
		// 	name:  "example-50",
		// 	input: example,
		// 	steps: 50,
		// 	want:  1594,
		// },
		// {
		// 	name:  "example-100",
		// 	input: example,
		// 	steps: 100,
		// 	want:  6536,
		// },
		// {
		// 	name:  "example-5k",
		// 	input: example,
		// 	steps: 5000,
		// 	want:  16733044,
		// },
		{
			name:  "actual",
			input: input,
			steps: 26501365,
			want:  618261433219147,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, tt.steps); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
