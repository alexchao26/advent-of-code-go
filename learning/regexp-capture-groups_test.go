package learning

import (
	"reflect"
	"testing"
)

var (
	example1 = `22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19`

	example2 = ` 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6`

	example3 = `14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`
)

func TestCaptureBingoBoard(t *testing.T) {
	type args struct {
		board string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				board: example1,
			},
			want: [][]int{
				{22, 13, 17, 11, 0},
				{8, 2, 23, 4, 24},
				{21, 9, 14, 16, 7},
				{6, 10, 3, 18, 5},
				{1, 12, 20, 15, 19},
			},
		},
		{
			name: "example 2",
			args: args{
				board: example2,
			},
			want: [][]int{
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6},
			},
		},
		{
			name: "example 3",
			args: args{
				board: example3,
			},
			want: [][]int{
				{14, 21, 17, 24, 4},
				{10, 16, 15, 9, 19},
				{18, 8, 23, 26, 20},
				{22, 11, 13, 6, 5},
				{2, 0, 12, 3, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaptureBingoBoard(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CaptureBingoBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
