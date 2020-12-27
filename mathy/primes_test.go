package mathy

import (
	"fmt"
	"testing"
)

func TestGeneratePrimes(t *testing.T) {
	tests := []struct {
		previousPrimes []int
		n, expected    int
	}{
		{[]int{2, 3}, 1, 2},
		{[]int{2, 3}, 2, 3},
		{[]int{2, 3}, 3, 5},
		{[]int{2, 3}, 4, 7},
		{[]int{2, 3, 5, 7}, 4, 7},
		{[]int{2, 3, 5, 7}, 5, 11},
		{[]int{2, 3, 5, 7}, 6, 13},
		{[]int{2, 3}, 7, 17},
		{[]int{2, 3}, 8, 19},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v-th prime number", test.n),
			func(t *testing.T) {
				ans := GeneratePrimes(test.previousPrimes, test.n)
				if ans != test.expected {
					t.Errorf("Expected %v-th prime to be %v, got %v", test.n, test.expected, ans)
				}
			},
		)
	}
}

// run go test -bench=. from within this util folder
// Benchmark
func benchGenPrimes(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePrimes([]int{2, 3}, n)
	}
}

// Benchmark generating different magnitudes of primes...
func BenchmarkGeneratePrimes10(b *testing.B)      { benchGenPrimes(10, b) }
func BenchmarkGeneratePrimes100(b *testing.B)     { benchGenPrimes(100, b) }
func BenchmarkGeneratePrimes1000(b *testing.B)    { benchGenPrimes(1000, b) }
func BenchmarkGeneratePrimes10000(b *testing.B)   { benchGenPrimes(10000, b) }
func BenchmarkGeneratePrimes100000(b *testing.B)  { benchGenPrimes(100000, b) }
func BenchmarkGeneratePrimes1000000(b *testing.B) { benchGenPrimes(1000000, b) }

// takes ~2 minutes on my mac
// func BenchmarkGeneratePrimes10000000(b *testing.B) { benchGenPrimes(10000000, b) }
