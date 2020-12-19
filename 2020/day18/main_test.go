package main

import (
	"reflect"
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// {"example1", "2 * 3 + (4 * 5)", 26},
		{"example2", "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
		{"actual", util.ReadFile("input.txt"), 53660285675207},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"actual", util.ReadFile("input.txt"), 141993988282687},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcFlatSlicePart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"example",
			args{input: []string{"1", "+", "3", "*", "4", "+", "5"}},
			"36",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFlatSlicePart2(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcFlatSlicePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splice(t *testing.T) {
	type args struct {
		sli        []string
		startIndex int
		items      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"example1", args{[]string{"a", "b", "c", "d", "e"}, 1, 1}, []string{"a", "c", "d", "e"}},
		{"example1", args{[]string{"a", "b", "c", "d", "e"}, 1, 4}, []string{"a"}},
		{"example1", args{[]string{"a", "b", "c", "d", "e"}, 0, 1}, []string{"b", "c", "d", "e"}},
		{"example1", args{[]string{"a", "b", "c", "d", "e"}, 0, 5}, []string{}},
		{"example1", args{[]string{"a", "b", "c", "d", "e"}, 3, 1}, []string{"a", "b", "c", "e"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splice(tt.args.sli, tt.args.startIndex, tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splice() = %v, want %v", got, tt.want)
			}
		})
	}
}
