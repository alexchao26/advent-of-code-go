package main

import (
	"testing"
)

var example = `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`

func Test_blizzardJourney(t *testing.T) {
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
			want:  18,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  240,
		},
		{
			name:  "example",
			input: example,
			part:  2,
			want:  54,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  717,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blizzardJourney(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
