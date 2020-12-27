package mathy

import (
	"math"
)

func PythagoreanDistance(x1, y1, x2, y2 int) float64 {
	xDiff := float64(x1 - x2)
	yDiff := float64(y1 - y2)

	sumOfSquares := math.Pow(xDiff, 2) + math.Pow(yDiff, 2)
	return math.Sqrt(sumOfSquares)
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	xDiff := x1 - x2
	yDiff := y1 - y2
	if xDiff < 0 {
		xDiff *= -1
	}
	if yDiff < 0 {
		yDiff *= -1
	}
	return xDiff + yDiff
}
