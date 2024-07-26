package mathy

func LeastCommonMultiple(i, j int) int {
	gcd := GreatestCommonMultiple(i, j)

	return (i * j) / gcd
}

func GreatestCommonMultiple(i, j int) int {
	if j > i {
		i, j = j, i
	}

	if j == 0 {
		return i
	}
	return GreatestCommonMultiple(j, i%j)
}
