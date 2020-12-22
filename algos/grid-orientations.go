package algos

// AllGridOrientations returns the 8 possible orientations of a given grid
// i.e. rotated 4 times and the mirror image of all of those
func AllGridOrientations(grid [][]string) [][][]string {
	orientations := [][][]string{grid}
	// add the 3 other rotations
	for i := 0; i < 3; i++ {
		orientations = append(orientations, RotateStringGrid(orientations[len(orientations)-1]))
	}
	// then add the mirror images of all 4 rotations
	for i := 0; i < 4; i++ {
		orientations = append(orientations, MirrorStringGrid(orientations[i]))
	}

	return orientations
}
