package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	pairs := parseInput(input)
	// sum all the indexes that are in the right order
	// ONE INDEXED NOT ZERO
	goodIndexSum := 0
	for i, pair := range pairs {
		left, right := pair[0], pair[1]
		if isInOrder(left, right) {
			goodIndexSum += i + 1
		}
	}

	return goodIndexSum
}

func part2(input string) int {
	pairs := parseInput(input)
	allPackets := [][]interface{}{
		// good reminder that json.Unmarshal will convert numbers to float64...
		// so need this to match for the way to isInOrder() function works..
		{[]interface{}{float64(2)}},
		{[]interface{}{float64(6)}},
	}
	for _, pair := range pairs {
		allPackets = append(allPackets, pair[0])
		allPackets = append(allPackets, pair[1])
	}

	sort.Slice(allPackets, func(i, j int) bool {
		left, right := allPackets[i], allPackets[j]
		return isInOrder(left, right)
	})

	ans := 1
	for i, p := range allPackets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			ans *= i + 1
		}
	}

	return ans
}

func parseInput(input string) (ans [][2][]interface{}) {
	for _, packetPairs := range strings.Split(input, "\n\n") {
		pairs := strings.Split(packetPairs, "\n")
		ans = append(ans, [2][]interface{}{
			parseRawString(pairs[0]),
			parseRawString(pairs[1]),
		})
	}
	return ans
}

// will parse as JSON with elements as either int or []int...
func parseRawString(raw string) []interface{} {
	ans := []interface{}{}
	json.Unmarshal([]byte(raw), &ans)
	return ans
}

func isInOrder(left, right []interface{}) bool {
	for l := 0; l < len(left); l++ {
		if l > len(right)-1 {
			return false
		}

		// attempt to convert both to ints...
		leftNum, isLeftNum := left[l].(float64)
		rightNum, isRightNum := right[l].(float64)

		leftList, isLeftList := left[l].([]interface{})
		rightList, isRightList := right[l].([]interface{})
		if isLeftNum && isRightNum {
			if leftNum != rightNum {
				return leftNum < rightNum
			}
		} else if isLeftNum || isRightNum {
			if isLeftNum {
				leftList = []interface{}{leftNum}
			} else if isRightNum {
				rightList = []interface{}{rightNum}
			} else {
				panic(fmt.Sprintf("expected one num %T:%v, %T:%v", left[l],
					left[l], right[l], right[l]))
			}
			return isInOrder(leftList, rightList)
		} else {
			// both lists
			if !isLeftList || !isRightList {
				panic(fmt.Sprintf("expected two lists %T:%v, %T:%v", left[l],
					left[l], right[l], right[l]))
			}
			return isInOrder(leftList, rightList)
		}
	}
	return true
}
