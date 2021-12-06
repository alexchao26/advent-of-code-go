package main

import (
	"testing"
)

var example = `3,4,3,1,2`

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
			want:  5934, // after 80 days
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  362666,
		},
		{
			name:  "example",
			input: example,
			part:  2,
			want:  26984457539,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  1640526601595,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.part == 0 {
				t.Error("part value cannot be zero")
			}
			if got := step(tt.input, tt.part); got != tt.want {
				t.Errorf("step() = %v, want %v", got, tt.want)
			}
		})
	}
}
