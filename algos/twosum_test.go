package algos

import "testing"

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name      string
		args      args
		wantNum1  int
		wantNum2  int
		wantFound bool
	}{
		{"found 1", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6}, 2, 4, true},
		{"found 2", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 17}, 8, 9, true},
		{"not found 1", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 62}, 0, 0, false},
		{"not found 2", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -12}, 0, 0, false},
		{"found 3", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -13}, -12}, 1, -13, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum1, gotNum2, gotFound := TwoSum(tt.args.nums, tt.args.target)
			if gotNum1 != tt.wantNum1 {
				t.Errorf("TwoSum() gotNum1 = %v, want %v", gotNum1, tt.wantNum1)
			}
			if gotNum2 != tt.wantNum2 {
				t.Errorf("TwoSum() gotNum2 = %v, want %v", gotNum2, tt.wantNum2)
			}
			if gotFound != tt.wantFound {
				t.Errorf("TwoSum() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
		want3 bool
	}{
		{"found 1", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6}, 0, 2, 4, true},
		{"found 2", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 17}, 0, 8, 9, true},
		{"not found 1", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 62}, 0, 0, 0, false},
		{"not found 2", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -12}, 0, 0, 0, false},
		{"found 3", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -13}, -12}, 0, 1, -13, true},
		{"found 3", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 27}, 8, 9, 10, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := ThreeSum(tt.args.nums, tt.args.target)
			if got != tt.want {
				t.Errorf("ThreeSum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ThreeSum() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ThreeSum() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("ThreeSum() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
