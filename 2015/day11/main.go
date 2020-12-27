package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := passwordIncrementing(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func passwordIncrementing(input string, part int) string {
	pw := input
	for !isValid(pw) {
		pw = incrementString(pw)
	}

	if part == 1 {
		return pw
	}

	pw = incrementString(pw)
	for !isValid(pw) {
		pw = incrementString(pw)
	}

	return pw
}

func incrementString(in string) string {
	chars := strings.Split(in, "")
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] == "z" {
			// continue loop to carry "carry over the one"
			chars[i] = "a"
		} else {
			asciiCode := cast.ToASCIICode(chars[i])
			chars[i] = cast.ASCIIIntToChar(asciiCode + 1)
			break
		}
	}
	return strings.Join(chars, "")
}

func isValid(in string) bool {
	rule1 := func(in string) bool {
		for i := 2; i < len(in); i++ {
			if in[i-2]+1 == in[i-1] && in[i-1]+1 == in[i] {
				return true
			}
		}
		return false
	}

	rule2 := func(in string) bool {
		return !regexp.MustCompile("[iol]").MatchString(in)
	}

	rule3 := func(in string) bool {
		pairs := map[string]bool{}
		for i := 1; i < len(in); i++ {
			if in[i-1] == in[i] {
				pairs[in[i-1:i+1]] = true
			}
		}
		return len(pairs) >= 2
	}

	return rule1(in) && rule2(in) && rule3(in)
}
