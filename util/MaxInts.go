package util

// MaxInts takes a variable number of integers and returns the largest one
func MaxInts(nums ...int) int {
	maxNum := nums[0]
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}
