package mathutil

import (
	"fmt"
	"strconv"
)

func MaxInt(nums ...int) int {
	maxNum := nums[0]
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}
func MinInt(nums ...int) int {
	minNum := nums[0]
	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}

func StrToInt(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		panic(fmt.Sprintf("converting string to number: %s", err))
	}
	return num
}

func IntToStr(in int) string {
	return strconv.Itoa(in)
}
