package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_md5StockingStuffer(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"part1 actual", util.ReadFile("input.txt"), 1, 254575},
		{"part2 actual", util.ReadFile("input.txt"), 2, 1038736},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5StockingStuffer(tt.input, tt.part); got != tt.want {
				t.Errorf("md5StockingStuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
