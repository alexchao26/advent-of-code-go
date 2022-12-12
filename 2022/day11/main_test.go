package main

import (
	"testing"
)

var example = ``

func Test_part1(t *testing.T) {
	tests := []struct {
		name         string
		useRealInput bool
		want         int
	}{
		{
			name:         "example",
			useRealInput: false,
			want:         10605,
		},
		{
			name:         "actual",
			useRealInput: true,
			want:         151312,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.useRealInput); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name         string
		useRealInput bool
		want         int
	}{
		{
			name:         "example",
			useRealInput: false,
			want:         2713310158,
		},
		{
			name:         "actual",
			useRealInput: true,
			want:         51382025916,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.useRealInput); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
