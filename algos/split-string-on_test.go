package algos_test

import (
	"reflect"
	"testing"

	"github.com/alexchao26/advent-of-code-go/algos"
)

func TestSplitStringOn(t *testing.T) {
	type args struct {
		in     string
		cutset []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"simple_cutset_len_1", args{"bot 1 gets 4", []string{" gets "}}, []string{"bot 1", "4"}},
		{"simple", args{"hello x world y potato", []string{" x ", " y "}}, []string{"hello", "world", "potato"}},
		{"longer example", args{"onextwoxthree-one-two-threexfour", []string{"x", "-"}}, []string{"one", "two", "three", "one", "two", "three", "four"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := algos.SplitStringOn(tt.args.in, tt.args.cutset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitStringOn() = %v, want %v", got, tt.want)
			}
		})
	}
}
