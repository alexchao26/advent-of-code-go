package main

import (
	_ "embed"
	"fmt"
	"testing"
)

var example = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`

func Test_trenchMap(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{
			name:  "example",
			input: example,
			part:  1,
			want:  35,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  4917,
		},
		{
			name:  "example",
			input: example,
			part:  2,
			want:  3351,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  16389,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.name, tt.part), func(t *testing.T) {
			if got := trenchMap(tt.input, tt.part); got != tt.want {
				t.Errorf("trenchMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAlgIndex(t *testing.T) {
	type args struct {
		img map[[2]int]string
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "around middle of example",
			args: args{
				img: func() map[[2]int]string {
					_, img := parseInput(example)
					return img
				}(),
				r: 2,
				c: 2,
			},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAlgIndex(tt.args.img, tt.args.r, tt.args.c); got != tt.want {
				t.Errorf("getAlgIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
