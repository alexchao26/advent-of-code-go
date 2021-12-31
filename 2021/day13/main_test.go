package main

import (
	"fmt"
	"testing"
)

var example = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

func Test_transparentOrigamiDay13(t *testing.T) {
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
			want:  17,
		},
		{
			name:  "actual",
			part:  1,
			input: input,
			want:  731,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s part%d", tt.name, tt.part), func(t *testing.T) {
			if got := transparentOrigamiDay13(tt.input, tt.part); got != tt.want {
				t.Errorf("transparentOrigamiDay13() = %v, want %v", got, tt.want)
			}
		})
	}
}
