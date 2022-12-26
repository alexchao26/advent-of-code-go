package main

import (
	_ "embed"
	"testing"
)

var example = `1
2
-3
3
-2
0
4`
var example2 = `1
2
-3
9
-2
0
4`
var example3 = `1
2
-3
3
-8
0
4`

func Test_mixList(t *testing.T) {
	tests := []struct {
		name                 string
		input                string
		decryptionKey, mixes int // for part 2 mostly
		want                 int
	}{
		{
			name:          "example",
			input:         example,
			decryptionKey: 1,
			mixes:         1,
			want:          3,
		},
		{
			name:          "example2",
			input:         example2,
			decryptionKey: 1,
			mixes:         1,
			want:          3,
		},
		{
			name:          "example3",
			input:         example3,
			decryptionKey: 1,
			mixes:         1,
			want:          3,
		},
		{
			name:          "actual",
			input:         input,
			decryptionKey: 1,
			mixes:         1,
			want:          9945,
		},
		{
			name:          "example",
			input:         example,
			decryptionKey: part2DecryptionKey,
			mixes:         10,
			want:          1623178306,
		},
		{
			name:          "actual",
			input:         input,
			decryptionKey: part2DecryptionKey,
			mixes:         10,
			want:          3338877775442,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mixList(tt.input, tt.decryptionKey, tt.mixes); got != tt.want {
				t.Errorf("mixList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_llNode_move(t *testing.T) {
	zeroNode, nodeSlice := parseInput(`0
1
2
3
4
5`)

	zeroNode.move(len(nodeSlice))
	originalString := "0,1,2,3,4,5,"
	// should be the same
	if got := listToString(zeroNode, len(nodeSlice)); got != originalString {
		t.Errorf("moving zero, want no change %q, got %q", originalString, got)
	}

	zeroNode.prev.move(len(nodeSlice))
	if got := listToString(zeroNode, len(nodeSlice)); got != originalString {
		t.Errorf("moving 5, want no change %q, got %q", originalString, got)
	}

	oneNode := zeroNode.next
	oneNode.move(len(nodeSlice))
	want := "0,2,1,3,4,5,"
	if got := listToString(zeroNode, len(nodeSlice)); got != want {
		t.Errorf("moving 1, want %q got %q", want, got)
	}
}
