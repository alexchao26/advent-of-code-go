package main

import (
	_ "embed"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "example1",
			input: "8A004A801A8002F478",
			want:  16,
		},
		{
			name:  "example2",
			input: "620080001611562C8802118E34",
			want:  12,
		},
		{
			name:  "example3",
			input: "C0015000016115A2E0802F182340",
			want:  23,
		},
		{
			name:  "example4",
			input: "A0016C880162017C3686B18A3D4780",
			want:  31,
		},
		{
			name:  "actual",
			input: input,
			want:  953,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: "C200B40A82",
			want:  3,
		},
		{
			name:  "example",
			input: "04005AC33890",
			want:  54,
		},
		{
			name:  "example",
			input: "880086C3E88112",
			want:  7,
		},
		{
			name:  "example",
			input: "CE00C43D881120",
			want:  9,
		},
		{
			name:  "example",
			input: "D8005AC2A8F0",
			want:  1,
		},
		{
			name:  "example",
			input: "F600BC2D8F",
			want:  0,
		},
		{
			name:  "example",
			input: "9C005AC2F8F0",
			want:  0,
		},
		{
			name:  "example",
			input: "9C0141080250320F1802104A08",
			want:  1,
		},

		{
			name:  "actual",
			input: input,
			want:  246225449979,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_handlePacket(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		pack string
// 		want int
// 	}{
// 		{
// 			pack: "110100101111111000101000",
// 			want: 2021,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := handlePacket(tt.pack); got != tt.want {
// 				t.Errorf("handlePacket() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
