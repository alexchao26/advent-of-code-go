package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	var nice int

	disallowPattern := regexp.MustCompile("(ab|cd|pq|xy)")
	for _, line := range strings.Split(input, "\n") {
		var vowels int
		for _, char := range line {
			if strings.ContainsRune("aeiou", char) {
				vowels++
			}
		}
		var hasDouble bool
		for i := 0; i < len(line)-1; i++ {
			if line[i] == line[i+1] {
				hasDouble = true
				break
			}
		}
		if vowels >= 3 && !disallowPattern.MatchString(line) && hasDouble {
			nice++
		}
	}

	return nice
}

func part2(input string) int {
	var nice int

	// put a double for loop check inside of a separate function b/c it makes
	// returning out of both loops possible, and avoids using a label which
	// makes me sad
	passesRule1 := func(line string) bool {
		for i := 0; i < len(line)-2; i++ {
			toMatch := line[i : i+2]
			for j := i + 2; j < len(line)-1; j++ {
				if line[j:j+2] == toMatch {
					return true
				}
			}
		}
		return false
	}

	for _, line := range strings.Split(input, "\n") {
		rule1 := passesRule1(line)

		var rule2 bool
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				rule2 = true
				break
			}
		}
		if rule1 && rule2 {
			nice++
		}
	}

	return nice
}
