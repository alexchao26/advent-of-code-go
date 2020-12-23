package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`

func Test_balanceBots(t *testing.T) {
	type args struct {
		input            string
		part1CompareVals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example_part1", args{example, []int{2, 5}}, 2},
		{"actual_part1", args{util.ReadFile("input.txt"), []int{17, 61}}, 181},
		{"actual_part2", args{util.ReadFile("input.txt"), nil}, 12567},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balanceBots(tt.args.input, tt.args.part1CompareVals); got != tt.want {
				t.Errorf("balanceBots() = %v, want %v", got, tt.want)
			}
		})
	}
}
