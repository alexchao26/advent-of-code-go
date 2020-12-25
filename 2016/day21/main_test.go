package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d`

func Test_part1(t *testing.T) {
	type args struct {
		input, starting, target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"part1 example", args{example, "abcde", ""}, "decab"},
		{"part1 actual", args{util.ReadFile("input.txt"), "abcdefgh", ""}, "bfheacgd"},
		{"part2 actual", args{util.ReadFile("input.txt"), "abcdefgh", "fbgdceah"}, "gcehdbfa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input, tt.args.starting, tt.args.target); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
