package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/alexchao26/advent-of-code-go/util"
)

var tests1 = []struct {
	name  string
	want  int
	input string
}{
	{"example", 114, "depth: 510\ntarget: 10,10"},
	{"actual", 11972, util.ReadFile("input.txt")},
}

func TestPart1(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
			t.Logf("Run time: %v", time.Since(startTime))
		})
	}
}

var tests2 = []struct {
	name  string
	want  int
	input string
}{
	{"example", 45, "depth: 510\ntarget: 10,10"},
	{"actual", 1092, util.ReadFile("input.txt")},
}

func TestPart2(t *testing.T) {
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
			t.Logf("Run time: %v", time.Since(startTime))
		})
	}
}

func Test_memoRegionTypeCalculator(t *testing.T) {
	depth, targetX, targetY := 510, 10, 10
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want regionType
	}{
		{"0,0", args{0, 0}, rocky},
		{"1,0", args{1, 0}, wet},
		{"0,1", args{0, 1}, rocky},
		{"1,1", args{1, 1}, narrow},
		{"10,10", args{10, 10}, rocky},
		{"11,10", args{11, 10}, wet},
		{"0,6", args{0, 6}, narrow},
	}
	for _, tt := range tests {
		memoFunc := memoRegionTypeCalculator(depth, targetX, targetY)
		t.Run(tt.name, func(t *testing.T) {
			if got := memoFunc(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoRegionTypeCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
