package main

import (
	"testing"
)

var example = `.....
.S-7.
.|.|.
.L-J.
.....`

var complexExample = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

var examplePart2 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`
var examplePart2_2 = `..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........`
var examplePart2_large = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`
var examplePart2_larger = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func Test_pipeMaze(t *testing.T) {
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
			want:  4,
		},
		{
			name:  "complexExample",
			input: complexExample,
			part:  1,
			want:  8,
		},
		{
			name:  "actual part 1",
			input: input,
			part:  1,
			want:  6773,
		},

		// part 2
		{
			name:  "examplePart2",
			input: examplePart2,
			part:  2,
			want:  4,
		},
		{
			name:  "examplePart2_2",
			input: examplePart2_2,
			part:  2,
			want:  4,
		},
		{
			name:  "examplePart2_large",
			input: examplePart2_large,
			part:  2,
			want:  8,
		},
		{
			name:  "examplePart2_larger",
			input: examplePart2_larger,
			part:  2,
			want:  10,
		},
		// {
		// 	name:  "actual part 2",
		// 	input: input,
		// part: 2,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pipeMaze(tt.input, tt.part); got != tt.want {
				t.Errorf("pipeMaze() = %v, want %v", got, tt.want)
			}
		})
	}
}
