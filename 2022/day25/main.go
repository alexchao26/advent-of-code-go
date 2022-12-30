package main

import (
	_ "embed"
	"flag"
	"fmt"
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

func part1(input string) string {
	snafuNums := strings.Split(input, "\n")
	// sum of the fuel requirements
	// power of 5...
	// 2  is 2
	// 1  is 1
	// 0  is 0
	// -1 is -
	// -2 is =
	sum := ""
	for _, n := range snafuNums {
		sum = addSnafu(sum, n)
	}

	return sum
}

func addSnafu(one, two string) string {
	// reversed...
	split1, split2 := strings.Split(one, ""), strings.Split(two, "")
	var reversed1, reversed2 []string
	for i := len(split1) - 1; i >= 0; i-- {
		reversed1 = append(reversed1, split1[i])
	}
	for i := len(split2) - 1; i >= 0; i-- {
		reversed2 = append(reversed2, split2[i])
	}

	longer, shorter := reversed1, reversed2
	if len(longer) < len(shorter) {
		longer, shorter = shorter, longer
	}

	charToVal := map[string]int{
		"=": -2,
		"-": -1,
		"0": 0,
		"1": 1,
		"2": 2,
	}
	valToChar := map[int]string{
		-2: "=",
		-1: "-",
		0:  "0",
		1:  "1",
		2:  "2",
	}

	ans := make([]int, len(longer)+1)
	for i := 0; i < len(longer); i++ {
		sum := charToVal[longer[i]]
		if i < len(shorter) {
			sum += charToVal[shorter[i]]
		}
		ans[i] += sum
		if ans[i] > 2 {
			ans[i] -= 5
			ans[i+1]++
		} else if ans[i] < -2 {
			ans[i] += 5
			ans[i+1]--
		}
	}

	for ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}

	snafu := ""
	for _, a := range ans {
		snafu = valToChar[a] + snafu
	}
	return snafu
}

func part2(input string) string {
	return ":)"
}
