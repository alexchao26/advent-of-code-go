package main

import (
	"testing"
)

var example = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func Test_syntaxScoring(t *testing.T) {
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
			want:  26397,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  387363,
		},
		{
			name:  "example",
			input: example,
			part:  2,
			want:  288957,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  4330777059,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := syntaxScoring(tt.input, tt.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
