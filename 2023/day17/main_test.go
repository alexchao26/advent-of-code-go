package main

import (
	"testing"
)

var example = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func Test_clumsyCart(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		minMoves int
		maxMoves int
		want     int
	}{
		{
			name:     "example",
			input:    example,
			minMoves: 1,
			maxMoves: 3,
			want:     102,
		},
		{
			name:     "actual",
			input:    input,
			minMoves: 1,
			maxMoves: 3,
			want:     1001,
		},
		{
			name:     "example_part2",
			input:    example,
			minMoves: 4,
			maxMoves: 10,
			want:     94,
		},
		{
			name:     "actual",
			input:    input,
			minMoves: 1,
			maxMoves: 3,
			want:     1197,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clumsyCart(tt.input, tt.minMoves, tt.maxMoves); got != tt.want {
				t.Errorf("clumsyCart() = %v, want %v", got, tt.want)
			}
		})
	}
}
