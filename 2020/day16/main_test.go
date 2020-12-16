package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example1 = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func Test_ticketTranslation(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"example_part1", example1, 1, 71},
		{"actual", util.ReadFile("input.txt"), 1, 32835},
		{"actual", util.ReadFile("input.txt"), 2, 514662805187},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ticketTranslation(tt.input, tt.part); got != tt.want {
				t.Errorf("ticketTranslation() = %v, want %v", got, tt.want)
			}
		})
	}
}
