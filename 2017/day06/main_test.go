package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_memoryReallocation(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{"actual", util.ReadFile("input.txt"), 1, 6681},
		{"actual", util.ReadFile("input.txt"), 2, 2392},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := memoryReallocation(tt.input, tt.part); got != tt.want {
				t.Errorf("memoryReallocation() = %v, want %v", got, tt.want)
			}
		})
	}

}
