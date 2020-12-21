package set_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/alexchao26/advent-of-code-go/data-structures/set"
)

func TestIntSet(t *testing.T) {
	tests := []struct {
		valToAdd int
		wantHas  []int
	}{
		{5, []int{5}},
		{5, []int{5}},
		{10, []int{5, 10}},
		{20, []int{5, 10, 20}},
		{20, []int{5, 10, 20}},
		{2230, []int{5, 10, 20, 2230}},
		{123, []int{5, 10, 20, 123, 2230}},
	}

	intSet := set.NewIntSet(nil)
	for _, tt := range tests {
		intSet.Add(tt.valToAdd)
		for _, want := range tt.wantHas {
			if !intSet.Has(want) {
				t.Errorf("want IntSet.Has(%d) = true, got false", want)
			}
		}
	}

	got := intSet.Keys()
	sort.Ints(got)
	want := []int{5, 10, 20, 123, 2230}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want intSet.Keys() to be %v, got %v", want, got)
	}

	valsToRemove := []int{5, 10, 20, 123, 2230}
	for _, tt := range valsToRemove {
		intSet.Remove(tt)
		if intSet.Has(tt) {
			t.Errorf("want IntSet.Has(%d) = false, got true", tt)
		}
	}

	got = intSet.Keys()
	if len(got) != 0 {
		t.Errorf("want zero-length slice after removing all keys, got %v", got)
	}
}
