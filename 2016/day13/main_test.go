package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_bfs(t *testing.T) {
	type args struct {
		input       string
		destination [2]int
		part        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{"10", [2]int{7, 4}, 1}, 11},
		{"actual_part1", args{util.ReadFile("input.txt"), [2]int{31, 39}, 1}, 86},
		{"actual_part2", args{util.ReadFile("input.txt"), [2]int{31, 39}, 2}, 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bfs(tt.args.input, tt.args.destination, tt.args.part); got != tt.want {
				t.Errorf("bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
