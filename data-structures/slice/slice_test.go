package slice_test

import (
	"reflect"
	"testing"

	"github.com/alexchao26/advent-of-code-go/data-structures/slice"
)

func TestDedupeStrings(t *testing.T) {
	tests := []struct {
		arg  []string
		want []string
	}{
		{[]string{"abc", "def", "ghi", "jkl", "abc"}, []string{"abc", "def", "ghi", "jkl"}},
		{[]string{"A", "a"}, []string{"A", "a"}},
		{[]string{"A", "B", "B", "B"}, []string{"A", "B"}},
	}
	for _, tt := range tests {
		if got := slice.DedupeStrings(tt.arg); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("DedupeStrings() = %v, want %v", got, tt.want)
		}
	}
}
func TestIntersection(t *testing.T) {
	tests := []struct {
		arg1, arg2 []string
		want       []string
	}{
		{[]string{"abc", "def", "ghi"}, []string{"jkl", "abc"}, []string{"abc"}},
		{[]string{"A", "a"}, []string{"A", "a"}, []string{"A", "a"}},
		{[]string{"A", "B", "C", "X"}, []string{"X", "p"}, []string{"X"}},
	}
	for _, tt := range tests {
		if got := slice.IntersectionStrings(tt.arg1, tt.arg2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("DedupeStrings() = %v, want %v", got, tt.want)
		}
	}
}

func TestSpliceString(t *testing.T) {
	type args struct {
		sli   []string
		index int
		items int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"example 1", args{[]string{"a", "b", "c", "d", "e"}, 4, 1}, []string{"a", "b", "c", "d"}},
		{"example 2", args{[]string{"a", "b", "c", "d", "e"}, 4, 2}, []string{"a", "b", "c", "d"}},
		{"example 3", args{[]string{"a", "b", "c", "d", "e"}, 3, 2}, []string{"a", "b", "c"}},
		{"example 4", args{[]string{"a", "b", "c", "d", "e"}, 0, 2}, []string{"c", "d", "e"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slice.SpliceStrings(tt.args.sli, tt.args.index, tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpliceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
