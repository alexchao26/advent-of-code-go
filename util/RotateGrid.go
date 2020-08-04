package util

// RotateGrid returns the inputted grid, rotated counterclockwise
// call it multiple times for 180, & 270 degree rotations
// TODO modify this to take in any type (not just strings...)
func RotateGrid(grid [][]string) [][]string {
	rotated := make([][]string, len(grid[0]))
	for i := range rotated {
		rotated[i] = make([]string, len(grid))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			rotated[len(grid[0])-1-j][i] = grid[i][j]
		}
	}
	return rotated
}
