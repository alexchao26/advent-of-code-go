package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_likeARogue(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		numRows int
		want    int
	}{
		{"example", ".^^.^.^^^^", 10, 38},
		{"part1_actual", util.ReadFile("input.txt"), 40, 2005},
		{"part2_actual", util.ReadFile("input.txt"), 400000, 20008491},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := likeARogue(tt.input, tt.numRows); got != tt.want {
				t.Errorf("likeARogue() = %v, want %v", got, tt.want)
			}
		})
	}
}
