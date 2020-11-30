package util

import (
	"fmt"
	"strconv"
)

func StrToInt(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		panic(fmt.Sprintf("converting string to number: %s", err))
	}
	return num
}
