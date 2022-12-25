package main

import (
	"testing"
)

var example = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  64,
		},
		{
			name:  "actual",
			input: input,
			want:  4636,
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

var flatDisc = `5,5,5
5,5,6
5,5,7
5,6,5
5,6,6
5,6,7
5,7,5
5,7,6
5,7,7`

// 3x3x1 disc...
func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "dumb simple",
			input: "1,1,1\n2,1,1",
			want:  10,
		},
		{
			name:  "dumber simpleer",
			input: "2,1,1",
			want:  6,
		},
		{
			name:  "example",
			input: example,
			want:  58,
		},
		{
			name:  "flatDisc",
			input: flatDisc,
			// 9 + 9 + 3 * 4 = 30
			want: 30,
		},
		{
			name:  "actual",
			input: input,
			// PAIN, used coord instead of front in the bfs check :/
			want: 2572,
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
