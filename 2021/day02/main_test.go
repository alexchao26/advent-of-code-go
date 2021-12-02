package main

import (
	"testing"
)

var example = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func Test_day2(t *testing.T) {
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
			want:  150,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  1813801,
		},
		{
			name:  "example",
			input: example,
			part:  2,
			want:  900,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  1960569556,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day2(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
