package main

import (
	"testing"
	"time"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_rambunctiousRecitation(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		desiredStep int
		want        int
	}{
		{"actual", util.ReadFile("input.txt"), 2020, 1259},
		{"actual", util.ReadFile("input.txt"), 30000000, 689},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			if got := rambunctiousRecitation(tt.input, tt.desiredStep); got != tt.want {
				t.Errorf("rambunctiousRecitation() = %v, want %v", got, tt.want)
			}
			t.Log("Op time: ", time.Since(startTime))
		})
	}
}
