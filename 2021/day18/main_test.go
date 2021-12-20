package main

import (
	_ "embed"
	"testing"
)

var shortExample = `[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`

/*
[[[[[1,1],[2,2]],[3,3]],[4,4]],[5,5]]
[[[[ 0,   [3,2]],[3,3]],[4,4]],[5,5]]
[[[[ 3,    0   ], [5,3]],[4,4]],[5,5]]

// reduces to [[[[3,0],[5,3]],[4,4]],[5,5]]

*/

var bigExample = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

var example1 = `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`

/*

[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]]
+ [[[[4,2],2],6],[8,7]]
[[[[[7,7],[7,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]],[[[[4,2],2],6],[8,7]]]
[[[[0,[14,7]],[[8,7],[8,7]]],[[[7,0],[7,7]],9]],[[[[4,2],2],6],[8,7]]]
[[[[14,0],[[15,7],[8,7]]],[[[7,0],[7,7]],9]],[[[[4,2],2],6],[8,7]]]
[[[[14,15],[0,[15,7]]],[[[7,0],[7,7]],9]],[[[[4,2],2],6],[8,7]]]
[[[[14,15],[15,0]],[[[14,0],[7,7]],9]],[[[[4,2],2],6],[8,7]]]
[[[[14,15],[15,14]],[[0,[7,7]],9]],[[[[4,2],2],6],[8,7]]]
[[ [[14,15],[15,14]] , [[7,0],16]],[[[[4,2],2],6],[8,7]]]
                               ^ this 16 was not getting added to
							   indicating that the prev pointers were broken
							   the bug was in the addNodes function to set n2.prev to lastN1 😭
[[ [[14,15],[15,14]] , [[7,0],20] ],[[[0,4],6],[8,7]]]


*/

var splitExample = `[[[[4,3],4],4],[7,[[8,4],9]]]
[1,1]`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: "[[1,2],[[3,4],5]]",
			want:  143,
		},
		{
			name:  "short",
			input: shortExample,
			// reduces to [[[[3,0],[5,3]],[4,4]],[5,5]]
			want: 791,
		},
		{
			name:  "short plus 6s",
			input: shortExample + "\n[6,6]",
			// reduces to [[[[5,0],[7,4]],[5,5]],[6,6]]
			want: 1137,
		},
		{
			name:  "split example",
			input: splitExample,
			// reduces to [[[[0,7],4],[[7,8],[6,0]]],[8,1]]
			want: 1384,
		},

		{
			name:  "example1",
			input: example1,
			want:  3488,
		},

		{
			name:  "bigExample",
			input: bigExample,
			want:  4140,
		},

		{
			name:  "actual",
			input: input,
			want:  4202,
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
			input: bigExample,
			want:  3993,
		},
		{
			name:  "actual",
			input: input,
			want:  4779,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_reduce(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantString string
	}{
		{
			name:       "simple",
			input:      "[[[[[9,8],1],2],3],4]",
			wantString: "v: 0, depth: 4 -> v: 9, depth: 4 -> v: 2, depth: 3 -> v: 3, depth: 2 -> v: 4, depth: 1 -> ",
			// wantString: "[[[[0,9],2],3],4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := parseInput(tt.input)
			head := nums[0]
			// fmt.Println(head)
			gotHead := head.reduce()
			if gotHead.String() != tt.wantString {
				t.Errorf("node.reduce() = %q, want %q", gotHead.String(), tt.wantString)
			}
		})
	}
}

func Test_node_magnitude(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{input: `[[1,2],[[3,4],5]]`, want: 143},
		{input: `[[[[0,7],4],[[7,8],[6,0]]],[8,1]]`, want: 1384},
		{input: `[[[[1,1],[2,2]],[3,3]],[4,4]]`, want: 445},
		{input: `[[[[3,0],[5,3]],[4,4]],[5,5]]`, want: 791},
		{input: `[[[[5,0],[7,4]],[5,5]],[6,6]]`, want: 1137},
		{input: `[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]`, want: 3488},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := parseInput(tt.input)[0]
			if got := n.magnitude(); got != tt.want {
				t.Errorf("node.magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
