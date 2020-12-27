package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `20
15
10
5
5`

func Test_eggnogCombinations(t *testing.T) {
	type args struct {
		input  string
		target int
		part   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part1 example", args{example, 25, 1}, 4},
		{"part1 actual", args{util.ReadFile("input.txt"), 150, 1}, 1304},
		{"part2 example", args{example, 25, 2}, 3},
		{"part2 actual", args{util.ReadFile("input.txt"), 150, 2}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eggnogCombinations(tt.args.input, tt.args.target, tt.args.part); got != tt.want {
				t.Errorf("eggnogCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
