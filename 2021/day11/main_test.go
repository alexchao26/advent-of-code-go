package main

import (
	"testing"
)

var example = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func Test_flashingOctopiLol(t *testing.T) {
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
			want:  1656,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  1723,
		},
		{
			name:  "example",
			input: example,
			part:  2,
			want:  195,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  327,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flashingOctopiLol(tt.input, tt.part); got != tt.want {
				t.Errorf("flashingOctopiLol() = %v, want %v", got, tt.want)
			}
		})
	}
}
