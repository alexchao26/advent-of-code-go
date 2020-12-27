package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	input := util.ReadFile("../input.txt")
	sli := strings.Split(input, "\n")

	var sum int

	for _, instruction := range sli {
		sign := instruction[:1]
		num, _ := strconv.Atoi(instruction[1:])
		if sign == "+" {
			sum += num
		} else {
			sum -= num
		}
	}

	fmt.Println("Final sum", sum)
}
