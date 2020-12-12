package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_part1And2(t *testing.T) {
	type args struct {
		input string
		part  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"part1 example 1", args{"^WNE$", 1}, 3},
		{"part1 example 2", args{"^ENWWW(NEEE|SSE(EE|N))$", 1}, 10},
		{"part1 example 3", args{"^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$", 1}, 18},
		{"part1 example 4", args{"^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$", 1}, 23},
		{"part1 example 5", args{"^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$", 1}, 31},
		{"part1 actual", args{util.ReadFile("input.txt"), 1}, 4121},
		{"part2 actual", args{util.ReadFile("input.txt"), 2}, 8636},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1And2(tt.args.input, tt.args.part); got != tt.want {
				t.Errorf("part1And2() = %v, want %v", got, tt.want)
			}
		})
	}
}
