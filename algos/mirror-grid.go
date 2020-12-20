package algos

// MirrorStringGrid returns the grid mirrored over the y-axis (i.e. left to right)
func MirrorStringGrid(grid [][]string) (flipped [][]string) {
	for i := range grid {
		flipped = append(flipped, []string{})
		for j := len(grid[i]) - 1; j >= 0; j-- {
			flipped[i] = append(flipped[i], grid[i][j])
		}
	}
	return flipped
}
