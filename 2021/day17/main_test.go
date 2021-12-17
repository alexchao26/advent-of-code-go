package main

import (
	_ "embed"
	"testing"
)

var example = `target area: x=20..30, y=-10..-5`

func Test_trickShot(t *testing.T) {
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
			want:  45,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  12246,
		},
		{
			name:  "part2 example",
			input: example,
			part:  2,
			want:  112,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  3528,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trickShot(tt.input, tt.part); got != tt.want {
				t.Errorf("trickShot() = %v, want %v", got, tt.want)
			}
		})
	}
}
