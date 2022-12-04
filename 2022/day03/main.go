package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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

var priorities = map[string]int{}

// lol why am i using init like this...
func init() {
	// generate priorities
	for i := 0; i < 26; i++ {
		priorities[cast.ASCIIIntToChar('a'+i)] = i + 1
		priorities[cast.ASCIIIntToChar('A'+i)] = i + 27
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
	prioritiesSum := 0
	for _, l := range strings.Split(input, "\n") {
		rightHalf := l[len(l)/2:]
		if len(rightHalf) != len(l)/2 {
			panic("did not divide in half: " + l)
		}
		for i := 0; i < len(l)/2; i++ {
			leftChar := l[i : i+1]
			// obv unnecessary n^2 but easily fast enough...
			if strings.Contains(rightHalf, leftChar) {
				prioritiesSum += priorities[leftChar]
				break
			}
		}
	}

	return prioritiesSum
}

func part2(input string) int {
	prioritiesSum := 0
	sacks := strings.Split(input, "\n")

	for i := 0; i < len(sacks); i += 3 {
		set2, set3 := stringToSet(sacks[i+1]), stringToSet(sacks[i+2])
		for _, char := range strings.Split(sacks[i], "") {
			if set2[char] && set3[char] {
				prioritiesSum += priorities[char]
				break
			}
		}
	}

	return prioritiesSum
}

func stringToSet(s string) map[string]bool {
	set := map[string]bool{}
	for _, char := range strings.Split(s, "") {
		set[char] = true
	}
	return set
}
