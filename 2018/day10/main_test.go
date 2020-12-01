/*
NOTE This is not an easily testable problem, these tests are for helper functions
*/

package main

import (
	"testing"
)

func TestHasLoneIsland(t *testing.T) {
	tests := []struct {
		name string
		grid [][2]int
		want bool
	}{
		{"simple_all_one_island", [][2]int{
			{1, 1},
			{1, 2},
			{2, 2},
			{2, 1},
		}, false},
		{"simple_has_lone_cell", [][2]int{
			{1, 1},
			{1, 2},
			{2, 2},
			{2, 1},
			{4, 1},
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasLoneIsland(tt.grid); got != tt.want {
				t.Errorf("hasLoneIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintGrid(t *testing.T) {
	tests := []struct {
		name      string
		positions [][2]int
		want      string
	}{
		{"horizontal line", [][2]int{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, -1},
		}, "0000\n"},
		{"horizontal line with gap", [][2]int{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, -1},
			{0, -4},
			{0, -5},
		}, "00  0000\n"},
		{"box", [][2]int{
			{0, 0},
			{0, 1},
			{0, 2},
			{1, 0},
			{2, 0},
			{2, 1},
			{2, 2},
			{1, 2},
		}, "000\n0 0\n000\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printGrid(tt.positions); got != tt.want {
				t.Errorf("printGrid() = %q, want %q", got, tt.want)
			}
		})
	}
}
