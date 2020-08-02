package trig

import (
	"math"
)

// TangentAndDistance docz
// startX will "always" be 13
// startY will "always" be 11
// 0 <= angleOffVery < 360
func TangentAndDistance(startX, startY, endX, endY int) (angleOffVert, distance float64) {
	rise, run := float64(endX)-float64(startX), float64(endY)-float64(startY)
	// fmt.Println(rise, run)

	// edge cases for verticals
	if run == 0 && rise < 0 {
		return 0, -1 * rise // endXY is up
	}
	if run == 0 && rise > 0 {
		return 180, rise // endXY is down
	}

	// handle left or right?
	if rise == 0 && run < 0 {
		return 270, -1 * run // left
	}
	if rise == 0 && run > 0 {
		return 90, run // right
	}

	// not verticals
	// calculate return distance
	distance = rise*rise + run*run
	distance = math.Sqrt(distance)

	// calculate arctangent which will be in radians
	// determine quadrent
	if rise < 0 && run > 0 { // top right
		angleOffVert = -1 * math.Atan(run/rise) * 180 / math.Pi
	} else if rise > 0 && run > 0 { // bottom right
		angleOffVert = 90 + math.Atan(rise/run)*180/math.Pi
	} else if rise > 0 && run < 0 { // bottom left
		angleOffVert = 180 + -1*math.Atan(run/rise)*180/math.Pi
	} else if rise < 0 && run < 0 { // top left
		angleOffVert = 270 + math.Atan(rise/run)*180/math.Pi
	}

	return angleOffVert, distance
}

/*
AngleOffVertical takes in two 2D points, it calculates the angle
between the line between then and a vertical line. The angle
returned is to the right of the vertical, i.e. from the vertical
and through the (mathematical) first quadrant
*/
func AngleOffVertical(startX, startY, endX, endY int) float64 {

	return 0
}

// Distance calculates the distance between two sets of 2D coordinates via Pythagorean's theorem
func Distance(startX, startY, endX, endY int) float64 {
	dx := startX - endX
	dy := startY - endY
	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}
