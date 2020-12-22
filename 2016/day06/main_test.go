package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

var example = `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

func Test_signalsAndNoise(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  string
	}{
		{"example", example, 1, "easter"},
		{"actual", util.ReadFile("input.txt"), 1, "afwlyyyq"},
		{"example", example, 2, "advent"},
		{"actual", util.ReadFile("input.txt"), 2, "bhkzekao"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := signalsAndNoise(tt.input, tt.part); got != tt.want {
				t.Errorf("signalsAndNoise() = %v, want %v", got, tt.want)
			}
		})
	}
}
