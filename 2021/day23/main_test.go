package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"strings"
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

func Test_amphipodDay23(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{
			name:  "part1 example",
			input: example,
			part:  1,
			want:  12521,
		},
		{
			name: "part1 simple",
			input: `#############
#.A.........#
###.#B#C#D###
  #A#B#C#D#
  #########  `,
			part: 1,
			want: 2,
		},
		{
			name: "part1 simple: A then B",
			input: `#############
#BA.........#
###.#.#C#D###
  #A#B#C#D#
  #########  `,
			part: 1,
			want: 52,
		},
		{
			// NOTE found a bug! A moving from a deep room to another room is calculating energy
			// NOTE as if it is walking through the wall
			name: "part1 reversed B room", // B has to get out of A's way first
			input: `#############
#.B.........#
###.#B#C#D###
  #A#A#C#D#
  #########  `,
			part: 1,
			want: 95,
		},
		{
			name: "part1 some shuffling",
			input: `#############
#...........#
###A#B#C#D###
  #A#C#B#D#
  #########  `,
			part: 1,
			want: 1120,
		},
		{
			name:  "part1 doneInput",
			input: doneInput,
			part:  1,
			want:  0,
		},
		{
			name:  "part1 actual",
			input: input,
			part:  1,
			want:  15299,
		},

		{
			name:  "example part 2",
			input: example,
			part:  2,
			want:  44169,
		},
		{
			name:  "part2 actual",
			input: input,
			part:  2,
			want:  47193,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() && strings.Contains(tt.name, "actual") {
				t.Skip(fmt.Sprintf("skipping %q in -short mode", tt.name))
			}
			if got := amphipodDay23(tt.input, tt.part); got != tt.want {
				t.Errorf("amphipodDay23() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_state_getUnsettledCoords(t *testing.T) {
	tests := []struct {
		name                string
		input               string
		roomCoordToWantChar map[[2]int]string
		want                [][2]int
	}{
		{
			name:                "already done",
			input:               doneInput,
			roomCoordToWantChar: roomCoordToWantCharPart1,
			want:                nil,
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
			roomCoordToWantChar: roomCoordToWantCharPart1,
			want:                [][2]int{{2, 3}, {3, 5}, {2, 5}, {2, 7}, {3, 9}, {2, 9}},
		},
		{
			name: "four unsettled coords",
			input: `#############
#AB.....D..D#
###.#.#C#.###
  #A#B#C#.#  
  #########  `,
			roomCoordToWantChar: roomCoordToWantCharPart1,
			want:                [][2]int{{1, 1}, {1, 2}, {1, 8}, {1, 11}},
		},
		{
			name: "part2 test",
			input: `#############
#AB.....D..D#
###.#.#C#.###
  #A#A#C#.#  
  #B#B#C#D#  
  #A#B#C#D#  
  #########  `,
			roomCoordToWantChar: roomCoordToWantCharPart2,
			want: [][2]int{
				// hallway
				{1, 1}, {1, 2}, {1, 8}, {1, 11},
				{4, 3}, {3, 3}, // A room
				{3, 5}, // B room
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parseInput(tt.input)
			if got := s.getUnsettledCoords(tt.roomCoordToWantChar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("state.getUnsettledCoords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_state_getNextPossibleMoves(t *testing.T) {
	tests := []struct {
		name                string
		input               string
		unsettledCoord      [2]int
		roomCoordToWantChar map[[2]int]string
		want                [][2]int
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
			unsettledCoord:      [2]int{3, 3}, // DEEP room in A slot
			roomCoordToWantChar: roomCoordToWantCharPart1,
			want:                nil,
		},
		{
			name:                "example-frontA",
			input:               example,
			unsettledCoord:      [2]int{2, 3}, // FRONT room in A slot
			roomCoordToWantChar: roomCoordToWantCharPart1,
			want:                [][2]int{{1, 2}, {1, 4}, {1, 1}, {1, 6}, {1, 8}, {1, 10}, {1, 11}},
		},
		{
			name:                "example-deepA-should be stuck",
			input:               example,
			unsettledCoord:      [2]int{3, 3}, // FRONT room in A slot
			roomCoordToWantChar: roomCoordToWantCharPart1,
			want:                nil,
		},
		{
			name: "hallways to rooms ONLY",
			input: `#############
#AB.....D..D#
###.#.#C#.###
  #A#B#C#.#  
  ######### `,
			unsettledCoord:      [2]int{1, 2},
			roomCoordToWantChar: roomCoordToWantCharPart1,
			want:                [][2]int{{2, 5}},
		},
		{
			name: "hallway to DEEP room",
			input: `#############
#AB.....D..D#
###.#.#C#.###
  #A#B#C#.#  
  ######### `,
			unsettledCoord:      [2]int{1, 8},
			roomCoordToWantChar: roomCoordToWantCharPart1,
			want:                [][2]int{{3, 9}},
		},
		{
			name: "part2 simple",
			input: `#############
#B..........#
###A#.#C#D###
  #A#B#C#D#  
  #A#B#C#D#  
  #A#B#C#D#  
  ######### `,
			unsettledCoord:      [2]int{1, 1},
			roomCoordToWantChar: roomCoordToWantCharPart2,
			want:                [][2]int{{2, 5}},
		},
		{
			name: "part2 back of room",
			input: `#############
#B......B.BB#
###A#.#C#D###
  #A#.#C#D#  
  #A#.#C#D#  
  #A#.#C#D#  
  ######### `,
			unsettledCoord:      [2]int{1, 1},
			roomCoordToWantChar: roomCoordToWantCharPart2,
			want:                [][2]int{{5, 5}},
		},
		{
			name: "part2 back of room",
			input: `#############
#B......B..B#
###A#.#C#D###
  #A#.#C#D#  
  #A#.#C#D#  
  #A#B#C#D#  
  ######### `,
			unsettledCoord:      [2]int{1, 8},
			roomCoordToWantChar: roomCoordToWantCharPart2,
			want:                [][2]int{{4, 5}},
		},
		{
			name: "part2 bug moving C from B room to C room",
			input: `#############
#AA.....B.BD#
###B#.#.#.###
  #D#C#.#.#  
  #D#B#C#C#  
  #A#D#C#A#  
  #########  `,
			unsettledCoord:      [2]int{3, 5},
			roomCoordToWantChar: roomCoordToWantCharPart2,
			want:                [][2]int{{1, 4}, {1, 6}, {3, 7}},
		},
		{
			name: "part2 bug moving B out of B room bc of a blocked D",
			input: `#############
#AA.....B.BD#
###B#.#.#.###
  #D#.#C#.#  
  #D#B#C#C#  
  #A#D#C#A#  
  #########  `,
			unsettledCoord:      [2]int{4, 5},
			roomCoordToWantChar: roomCoordToWantCharPart2,
			want:                [][2]int{{1, 4}, {1, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := parseInput(tt.input)
			if got := s.getNextPossibleMoves(tt.unsettledCoord, tt.roomCoordToWantChar); !reflect.DeepEqual(got, tt.want) {
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
		{
			name: "part2 C",
			args: args{
				char:  "C",
				start: [2]int{5, 11},
				end:   [2]int{3, 7},
			},
			want: 1000,
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
