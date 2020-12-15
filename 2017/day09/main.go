package main

import (
	"flag"
	"fmt"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := streamProcessing(util.ReadFile("./input.txt"), part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func streamProcessing(input string, part int) int {
	var totalScore, garbageCount int

	var inGarbage bool
	var openCurlies int // equivalent to the current group's score

	for i := 0; i < len(input); i++ {
		char := string(input[i])
		if inGarbage {
			switch char {
			case "!":
				i++
			case ">":
				inGarbage = false
			default:
				garbageCount++ // part 2
			}
		} else {
			switch char {
			case "{":
				openCurlies++
			case "}":
				totalScore += openCurlies // part 1
				openCurlies--
			case "<":
				inGarbage = true
			}
		}
	}

	if part == 1 {
		return totalScore
	}
	return garbageCount
}
