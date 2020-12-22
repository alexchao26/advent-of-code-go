package cast_test

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/cast"
)

func TestToInt(t *testing.T) {
	if got := cast.ToInt("123"); got != 123 {
		t.Errorf("cast.Int(123) = %v, want 123", got)
	}
	if got := cast.ToInt("9835"); got != 9835 {
		t.Errorf("cast.Int(9835) = %v, want 9835", got)
	}
}

func TestToString(t *testing.T) {
	byteTests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{"byte", byte('a'), "a"},
		{"byte", byte('x'), "x"},
		{"int", 1234, "1234"},
		{"int", 512, "512"},
		{"rune", rune(65), "A"},
		{"rune", rune(97), "a"},
	}
	for _, tt := range byteTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cast.ToString(tt.input); got != tt.want {
				t.Errorf("ToString(byte) = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestToASCIIConstants(t *testing.T) {
	if cast.ASCIICodeCapA != 65 {
		t.Errorf("Expected cast.ASCIICodeCapA to be 65, got %d", cast.ASCIICodeCapA)
	}
	if cast.ASCIICodeLowerA != 97 {
		t.Errorf("Expected cast.ASCIICodeLowerA to be 97, got %d", cast.ASCIICodeLowerA)
	}
}

func TestToASCIICode(t *testing.T) {
	type args struct {
		arg interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example_string", args{"a"}, cast.ASCIICodeLowerA},
		{"example_string", args{"b"}, cast.ASCIICodeLowerA + 1},
		{"example_string", args{"z"}, cast.ASCIICodeLowerA + 25},
		{"example_string", args{"C"}, cast.ASCIICodeCapA + 2},
		{"example_rune", args{rune(97)}, 97},
		{"example_byte", args{'a'}, 97},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cast.ToASCIICode(tt.args.arg); got != tt.want {
				t.Errorf("ToASCIICode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestASCIIIntToChar(t *testing.T) {
	type args struct {
		code int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example", args{97}, "a"},
		{"example", args{98}, "b"},
		{"example", args{65}, "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cast.ASCIIIntToChar(tt.args.code); got != tt.want {
				t.Errorf("ASCIIIntToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
