package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_md5Bfs(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  string
	}{
		{"part1_example1", "ihgpwlah", 1, "DDRRRD"},
		{"part1_actual", util.ReadFile("input.txt"), 1, "DDRUDLRRRD"},
		{"part2_example1", "ihgpwlah", 2, "370"},
		{"part2_actual", util.ReadFile("input.txt"), 2, "398"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5Bfs(tt.input, tt.part); got != tt.want {
				t.Errorf("md5Bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
