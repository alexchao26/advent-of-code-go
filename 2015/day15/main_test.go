package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_cookieScience(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 int
		want2 int
	}{
		{"actual both parts", util.ReadFile("input.txt"), 13882464, 11171160},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := cookieScience(tt.input)
			if got1 != tt.want1 {
				t.Errorf("cookieScience() part1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("cookieScience() part2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
