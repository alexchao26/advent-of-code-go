package main

import (
	"testing"
)

var example = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

var example2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

func Test_pulsePropagation(t *testing.T) {
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
			want:  32000000,
		},
		{
			name:  "example2",
			input: example2,
			part:  1,
			want:  11687500,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  817896682,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  250924073918341,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pulsePropagation(tt.input, tt.part); got != tt.want {
				t.Errorf("pulsePropagation() = %v, want %v", got, tt.want)
			}
		})
	}
}
