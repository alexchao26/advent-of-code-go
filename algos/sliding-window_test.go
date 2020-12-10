package algos

import "testing"

func TestSlidingWindowSum(t *testing.T) {
	type args struct {
		nums      []int
		targetSum int
	}
	tests := []struct {
		name           string
		args           args
		wantLeftIndex  int
		wantRightIndex int
		wantFound      bool
	}{
		{"found 1", args{[]int{-1, 23, 12, 14, 3, 4}, 21}, 3, 6, true},
		{"found 2", args{[]int{-1, 23, 12, 14, 3, 4, 59}, 21}, 3, 6, true},
		{"not found 1", args{[]int{0, 1, 2, 3, 4, 5, 6}, 45}, 0, 0, false},
		{"not found 2", args{[]int{0, 1, 2, 3, 4, 5, 6}, -34}, 0, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeftIndex, gotRightIndex, gotFound := SlidingWindowSum(tt.args.nums, tt.args.targetSum)
			if gotLeftIndex != tt.wantLeftIndex {
				t.Errorf("SlidingWindowSum() gotLeftIndex = %v, want %v", gotLeftIndex, tt.wantLeftIndex)
			}
			if gotRightIndex != tt.wantRightIndex {
				t.Errorf("SlidingWindowSum() gotRightIndex = %v, want %v", gotRightIndex, tt.wantRightIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("SlidingWindowSum() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
