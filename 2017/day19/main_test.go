package main

import (
	"strings"
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

// easier to look at this input in the format, but need to remove leading/trailing newlines
var example = strings.Trim(`
    |         
    |  +--+   
    A  |  C   
F---|----E|--+
    |  |  |  D
    +B-+  +--+
`, "\n")

func Test_movePacket(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		wantVisitedChars string
		wantSteps        int
	}{
		{"example", example, "ABCDEF", 38},
		{"actual", util.ReadFile("input.txt"), "EPYDUXANIT", 17544},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVisitedChars, gotSteps := movePacket(tt.input)
			if gotVisitedChars != tt.wantVisitedChars {
				t.Errorf("movePacket() gotVisitedChars = %v, want %v", gotVisitedChars, tt.wantVisitedChars)
			}
			if gotSteps != tt.wantSteps {
				t.Errorf("movePacket() gotSteps = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}
