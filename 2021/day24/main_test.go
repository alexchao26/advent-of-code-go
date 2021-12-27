package main

import (
	"testing"
)

func Test_aluDay24(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int64
	}{
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  59996912981939,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  17241911811915,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aluDay24(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
