package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_allergenAssessment(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantPart1 int
		wantPart2 string
	}{
		{"actual", util.ReadFile("input.txt"), 1815, "kllgt,jrnqx,ljvx,zxstb,gnbxs,mhtc,hfdxb,hbfnkq"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPart1, gotPart2 := allergenAssessment(tt.input)
			if gotPart1 != tt.wantPart1 {
				t.Errorf("allergenAssessment() = %v, want %v", gotPart1, tt.wantPart1)
			}
			if gotPart2 != tt.wantPart2 {
				t.Errorf("allergenAssessment() = %v, want %v", gotPart2, tt.wantPart2)
			}
		})
	}
}
