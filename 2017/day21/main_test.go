package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`

func Test_fractalArt(t *testing.T) {
	type args struct {
		input  string
		rounds int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{example, 2}, 12},
		{"actual_part1", args{util.ReadFile("input.txt"), 5}, 194},
		{"actual_part2", args{util.ReadFile("input.txt"), 18}, 2536879},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractalArt(tt.args.input, tt.args.rounds); got != tt.want {
				t.Logf("Ruleset: %s", tt.args.input)
				t.Logf("Rounds: %d", tt.args.rounds)
				t.Errorf("fractalArt() = %v, want %v", got, tt.want)
			}
		})
	}
}
