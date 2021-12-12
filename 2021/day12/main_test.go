package main

import (
	"testing"
)

var example = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var example2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var example3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  10,
		},
		{
			name:  "example2",
			input: example2,
			want:  19,
		},
		{
			name:  "example3",
			input: example3,
			want:  226,
		},
		{
			name:  "actual",
			input: input,
			want:  3421,
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
			want:  36,
		},
		{
			name:  "example2",
			input: example2,
			want:  103,
		},
		{
			name:  "example3",
			input: example3,
			want:  3509,
		},
		{
			name:  "actual",
			input: input,
			want:  84870,
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
