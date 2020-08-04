package trig

import "math"

/*
AngleOffVertical takes in two 2D points, it calculates the angle
between the line and a vertical line (straight up from origin)
NOTE: "up"/"top" and "down" are lexically flipped b/c of drawing a grid
where 0, 0 is the top left corner and higher numbers physically go DOWN
but lexically increase/go UP 🤦‍♂️
*/
func AngleOffVertical(startX, startY, endX, endY int) float64 {
	rise := float64(endX) - float64(startX)
	run := float64(endY) - float64(startY)

	var angle float64
	// basically a big if/elseif/else block
	switch {
	case run == 0 && rise < 0: // up
		angle = 0
	case run == 0 && rise > 0: // down
		angle = 180
	case rise == 0 && run < 0: // left
		angle = 270
	case rise == 0 && run > 0: // right
		angle = 90
	case rise < 0 && run > 0: // top right
		angle = -1 * math.Atan(run/rise) * 180 / math.Pi
	case rise > 0 && run > 0: // bottom right
		angle = 90 + math.Atan(rise/run)*180/math.Pi
	case rise > 0 && run < 0: // bottom left
		angle = 180 + -1*math.Atan(run/rise)*180/math.Pi
	case rise < 0 && run < 0: // top left
		angle = 270 + math.Atan(rise/run)*180/math.Pi
	}

	return angle
}

// Distance calculates the distance between two sets of 2D coordinates via Pythagorean's theorem
func Distance(startX, startY, endX, endY int) float64 {
	dx := startX - endX
	dy := startY - endY
	return math.Sqrt(float64(dx*dx) + float64(dy*dy))
}
