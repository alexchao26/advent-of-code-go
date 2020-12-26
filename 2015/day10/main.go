package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := lookAndSay(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func lookAndSay(input string, part int) int {
	lastSaid := input

	rounds := 40
	if part == 2 {
		rounds = 50
	}

	for i := 0; i < rounds; i++ {
		var runningCount int
		var said strings.Builder

		for i := 0; i < len(lastSaid); i++ {
			if i == len(lastSaid)-1 || lastSaid[i] != lastSaid[i+1] {
				// add to seen
				said.WriteString(fmt.Sprintf("%d%s", runningCount+1, string(lastSaid[i])))
				runningCount = 0
			} else {
				// build up running count
				runningCount++
			}

		}
		lastSaid = said.String()
	}

	return len(lastSaid)
}
