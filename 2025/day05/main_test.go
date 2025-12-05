package main

import (
	"testing"
)

var example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  3,
		},
		{
			name:  "actual",
			input: input,
			want:  739,
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
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  14,
		},
		{
			name:  "test1",
			input: "10-20\n\n1",
			want:  11,
		},
		{
			name:  "test1",
			input: "10-20\n20-30\n\n1",
			want:  21,
		},
		{
			name:  "test1",
			input: "10-20\n20-30\n1-5\n\n1",
			want:  26,
		},
		{
			name:  "test1",
			input: "10-20\n20-30\n1-5\n29-33\n\n1",
			want:  29,
		},
		{
			name:  "actual",
			input: input,
			want:  344486348901788,
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
