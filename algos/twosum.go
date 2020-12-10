package algos

// TwoSum returns the two numbers found within a slice that add up to
// the input target, and a boolean to check if the sum is found or not
func TwoSum(nums []int, target int) (num1 int, num2 int, found bool) {
	seen := make(map[int]bool, len(nums))

	for _, v := range nums {
		if seen[target-v] {
			return target - v, v, true
		}
		seen[v] = true
	}

	return 0, 0, false
}

// ThreeSum returns the three values within a slice that sum to the
// target input and a boolean stating if a match was found
func ThreeSum(nums []int, target int) (int, int, int, bool) {
	for i, v := range nums {
		if num1, num2, found := TwoSum(nums[i+1:], target-v); found {
			return v, num1, num2, true
		}
	}
	return 0, 0, 0, false
}
