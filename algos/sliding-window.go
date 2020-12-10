package algos

// SlidingWindowSum returns the left and right indices of a window within the
// nums slice, where all numbers in the slice [left:right] sum up to targetSum
// It also returns a boolean indicating if a valid window is found
func SlidingWindowSum(nums []int, targetSum int) (leftIndex, rightIndex int, found bool) {
	var left, right, sum int
	for right < len(nums) {
		switch {
		case left == right:
			sum += nums[right]
			right++
		case sum > targetSum:
			sum -= nums[left]
			left++
		case sum < targetSum:
			sum += nums[right]
			right++
		}
		if sum == targetSum {
			return left, right, true
		}
	}
	return 0, 0, false
}
