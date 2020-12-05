package main

import (
	"reflect"
	"testing"
)

var tests1 = []struct {
	name  string
	want  int
	input string
	// add extra args if needed
}{
	{"example", 3, "^WNE$"},
	{"example", 10, "^ENWWW(NEEE|SSE(EE|N))$"},
	{"example", 18, "^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"},
	{"example", 23, "^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"},
	{"example", 31, "^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$"},
	// {"actual", ACTUAL_ANSWER, util.ReadFile("input.txt")},
}

func TestPart1(t *testing.T) {
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

var tests2 = []struct {
	name  string
	want  int
	input string
	// add extra args if needed
}{
	// {"actual", ACTUAL_ANSWER, util.ReadFile("input.txt")},
}

func TestPart2(t *testing.T) {
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flattenRegexpPaths(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{"no children", "NSEW", []string{"NSEW"}},
		{"no children", "NSEW|SS|EWN", []string{"NSEW", "SS", "EWN"}},
		{"one child", "NSEW(EE|WW|SS)", []string{"NSEWEE", "NSEWWW", "NSEWSS"}},
		{"nested", "NS(EE(WW|SS))WW", []string{"NSEEWWWW", "NSEESSWW"}},
		{"optional path", "NS(EE(WW|SS|))WW", []string{"NSEEWWWW", "NSEESSWW", "NSEEWW"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flattenRegexpPaths(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenRegexpPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ingestNextBalancedParen(t *testing.T) {
	tests := []struct {
		name          string
		args          string
		wantBalanced  string
		wantRemainder string
	}{
		{"nonparen start", "NEWS(NEWS)", "NEWS", "(NEWS)"},
		{"paren start", "(NEWS|NEWS)NEWS(NEWS|)", "(NEWS|NEWS)", "NEWS(NEWS|)"},
		{"nested", "(NEWS(NEW|NEW))NEWS(NEWS|)", "(NEWS(NEW|NEW))", "NEWS(NEWS|)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBalanced, gotRemainder := ingestNextBalancedParen(tt.args)
			if gotBalanced != tt.wantBalanced {
				t.Errorf("ingestNextBalancedParen() gotBalanced = %v, want %v", gotBalanced, tt.wantBalanced)
			}
			if gotRemainder != tt.wantRemainder {
				t.Errorf("ingestNextBalancedParen() gotRemainder = %v, want %v", gotRemainder, tt.wantRemainder)
			}
		})
	}
}

func Test_breakIntoTopLevelOptions(t *testing.T) {
	tests := []struct {
		name string
		args string
		want []string
	}{
		{"flat w/o or", "NEWS", []string{"NEWS"}},
		{"flat w/ or", "NEWS|NWW", []string{"NEWS", "NWW"}},
		{"parens w/ or", "NEWS(NE|NW)|NWW", []string{"NEWS(NE|NW)", "NWW"}},
		{"from example", "ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))", []string{
			"ESSSSW(NWSW|SSEN)", "WSWWN(E|WWS(E|SS))",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakIntoTopLevelOptions(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("breakIntoTopLevelOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
