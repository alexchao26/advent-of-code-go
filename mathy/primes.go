package mathy

import (
	"math"
)

// GeneratePrimes returns the n-th prime number
// its param primes []int is intended to contain previously found prime numbers
// to reduce duplicated work, but still be testable
func GeneratePrimes(primes []int, n int) int {
	if len(primes) < 2 {
		primes = []int{2, 3}
	}
	if len(primes) >= n {
		return primes[n-1]
	}
	for i := primes[len(primes)-1] + 2; len(primes) <= n; i += 2 {
		// check if i is a prime number by checking if it is divisible by any of the previous values of primes
		// stop at the square root of i
		for _, v := range primes {
			// not a prime, stop this loop
			if i%v == 0 {
				break
			}
			if math.Sqrt(float64(i)) < float64(v) {
				// add to primes
				primes = append(primes, i)
				break
			}
		}
	}

	return primes[n-1]
}
