package main

import (
	_ "embed"
	"reflect"
	"testing"
)

var example = `#############
#...........#
###B#C#B#D###
  #A#D#C#A#  
  #########  `

var doneInput = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#  
  #########  `

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  12521,
		},
		{
			name: "simple",
			input: `#############
#.A.........#
###.#B#C#D###
  #A#B#C#D#  
  #########  `,
			want: 2,
		},
		{
			name: "simple: A then B",
			input: `#############
#BA.........#
###.#.#C#D###
  #A#B#C#D#  
  #########  `,
			want: 52,
		},
		{
			// NOTE found a bug! A moving from a deep room to another room is calculating energy
			// NOTE as if it is walking through the wall
			name: "reversed B room", // B has to get out of A's way first
			input: `#############
#.B.........#
###.#B#C#D###
  #A#A#C#D#  
  #########  `,
			want: 95,
		},
		{
			name: "some shuffling",
			input: `#############
#...........#
###A#B#C#D###
  #A#C#B#D#  
  #########  `,
			want: 1120,
		},
		{
			name:  "doneInput",
			input: doneInput,
			want:  0,
		},
		{
			name:  "actual",
			input: input,
			want:  15299,
		},
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
		{
			name:  "example",
			input: example,
			want:  44169,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_state_getUnsettledCoords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][2]int
	}{
		{
			name:  "already done",
			input: doneInput,
			want:  nil,
		},
		{
			name:  "example - 2 \"done\" amphipods",
			input: example,
			/*
				#############
				#...........#
				###B#C#B#D###
				  #A#D#C#A#
				  #########
			*/
			want: [][2]int{{2, 3}, {2, 5}, {2, 7}, {2, 9}, {3, 5}, {3, 9}},
		},
		{
			name: "four unsettled coords",
			input: ` #############
#AB.....D..D#
###.#.#C#.###
  #A#B#C#.#  
  ######### `,
			want: [][2]int{{1, 1}, {1, 2}, {1, 8}, {1, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parseInput(tt.input)
			if got := s.getUnsettledCoords(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("state.getUnsettledCoords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_state_getNextPossibleMoves(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		unsettledCoord [2]int
		want           [][2]int
	}{
		{
			name:  "example, deep room is stuck",
			input: example,
			/*
				#############
				#...........#
				###B#C#B#D###
				  #A#D#C#A#
				  #########
			*/
			unsettledCoord: [2]int{3, 3}, // DEEP room in A slot
			want:           nil,
		},
		{
			name:           "example-frontA",
			input:          example,
			unsettledCoord: [2]int{2, 3}, // FRONT room in A slot
			want:           [][2]int{{1, 2}, {1, 4}, {1, 1}, {1, 6}, {1, 8}, {1, 10}, {1, 11}},
		},
		{
			name:           "example-deepA-should be stuck",
			input:          example,
			unsettledCoord: [2]int{3, 3}, // FRONT room in A slot
			want:           nil,
		},
		{
			name: "hallways to rooms ONLY",
			input: ` #############
#AB.....D..D#
###.#.#C#.###
  #A#B#C#.#  
  ######### `,
			unsettledCoord: [2]int{1, 2},
			want:           [][2]int{{2, 5}},
		},
		{
			name: "hallway to DEEP room",
			input: ` #############
#AB.....D..D#
###.#.#C#.###
  #A#B#C#.#  
  ######### `,
			unsettledCoord: [2]int{1, 8},
			want:           [][2]int{{3, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parseInput(tt.input)
			if got := s.getNextPossibleMoves(tt.unsettledCoord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("state.getNextPossibleMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcEnergy(t *testing.T) {
	type args struct {
		char  string
		start [2]int
		end   [2]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A 4 spaces away",
			args: args{
				char:  "A",
				start: [2]int{1, 1},
				end:   [2]int{1, 5},
			},
			want: 4,
		},
		{
			name: "A goes from B's room to A's room",
			args: args{
				char:  "A",
				start: [2]int{3, 5},
				end:   [2]int{2, 3},
			},
			want: 5,
		},
		{
			name: "D 6 spaces away",
			args: args{
				char:  "D",
				start: [2]int{3, 3},
				end:   [2]int{2, 8},
			},
			want: 8000,
		},
		{
			name: "C",
			args: args{
				char:  "C",
				start: [2]int{1, 11},
				end:   [2]int{3, 7},
			},
			want: 600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEnergy(tt.args.char, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("calcEnergy() = %v, want %v", got, tt.want)
			}
		})
	}
}
