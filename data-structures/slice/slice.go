package slice

// DedupeStrings returns a new slice with duplicates removed, maintains original order
func DedupeStrings(sli []string) []string {
	var result []string
	seen := map[string]bool{}
	for _, v := range sli {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

// DedupeInts returns a new slice with duplicates removed, maintains original order
func DedupeInts(sli []int) []int {
	var result []int
	seen := map[int]bool{}
	for _, v := range sli {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

// IntersectionStrings returns a slice of values in both argument slices, deduped
func IntersectionStrings(sli1, sli2 []string) []string {
	var result []string
	seen := map[string]bool{}
	for _, v := range sli1 {
		seen[v] = true
	}
	for _, v := range sli2 {
		if seen[v] {
			result = append(result, v)
			delete(seen, v)
		}
	}
	return result
}

// RemoveAllStrings returns a new slice with all instances of a given string removed
func RemoveAllStrings(sli []string, val string) []string {
	var result []string
	for _, v := range sli {
		if v != val {
			result = append(result, v)
		}
	}
	return result
}

// RemoveAllInts returns a new slice with all instances of a given int removed
func RemoveAllInts(sli []int, val int) []int {
	var result []int
	for _, v := range sli {
		if v != val {
			result = append(result, v)
		}
	}
	return result
}

// SpliceStrings removes a given number of elements starting at a given index
// if index + items >= len(sli) it does not throw an error
func SpliceStrings(sli []string, index int, items int) []string {
	if items < 0 {
		panic("cannot splice negative number of items")
	}
	if index+items >= len(sli) {
		return sli[:index]
	}
	copy(sli[index:], sli[index+items:])
	sli = sli[:len(sli)-items]
	return sli
}

// SpliceInts removes a given number of elements starting at a given index
// if index + items >= len(sli) it does not throw an error
func SpliceInts(sli []int, index int, items int) []int {
	if items < 0 {
		panic("cannot splice negative number of items")
	}
	if index+items >= len(sli) {
		return sli[:index]
	}
	copy(sli[index:], sli[index+items:])
	sli = sli[:len(sli)-items]
	return sli
}
