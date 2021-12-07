package main

import (
	_ "embed"
	"testing"
)

var example = `16,1,2,0,4,2,7,1,2,14`

func Test_calcMinFuel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part  int
		want  int
	}{
		{
			name:  "example",
			input: example,
			part:  1,
			want:  37,
		},
		{
			name:  "actual",
			input: input,
			part:  1,
			want:  329389,
		},
		{
			name:  "example",
			input: example,
			part:  2,
			want:  168,
		},
		{
			name:  "actual",
			input: input,
			part:  2,
			want:  86397080,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcMinFuel(tt.input, tt.part); got != tt.want {
				t.Errorf("calcMinFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcSummationFromOneToEnd(t *testing.T) {
	// auto generated these tests because math is rough...
	type args struct {
		end int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{end: 4},
			want: 10,
		},
		{
			args: args{end: 5},
			want: 15,
		},
		{
			args: args{end: 6},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcSummationFromOneToEnd(tt.args.end); got != tt.want {
				t.Errorf("calcSummationFromOneToEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
