package algos

import (
	"math/rand"
	"time"
)

var rn = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomInt returns a random int from zero through upper - 1
func RandomInt(upper int) int {
	return rn.Intn(upper)
}
