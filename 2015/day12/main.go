package main

import (
	"encoding/json"
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

	var ans int
	if part == 1 {
		ans = part1(util.ReadFile("./input.txt"))
	} else {
		ans = part2(util.ReadFile("./input.txt"))
	}
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	var totalSum int
	var runningNum strings.Builder
	for _, char := range input {
		// got lucky on the input here
		if regexp.MustCompile("[-0-9]").MatchString(string(char)) {
			runningNum.WriteRune(char)
		} else if runningNum.Len() != 0 {
			totalSum += cast.ToInt(runningNum.String())
			runningNum.Reset()
		}
	}

	// this is for part 2 to handle ends of strings
	if runningNum.Len() != 0 {
		totalSum += cast.ToInt(runningNum.String())
	}

	return totalSum
}

// This solution leverages the error or nil that is returned from json.Marshal
//
// A full Go solution would requiring writing my own JSON parser
// A much easier way would use javascript's JSON.parse and just do a dfs on it
func part2(input string) int {
	// if input does not have object braces or an instance of "red", just pass it through part1
	if !regexp.MustCompile("[{}]").MatchString(input) ||
		!regexp.MustCompile("red").MatchString(input) {
		return part1(input)
	}

	// try to parse into an object if that's
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(input), &obj)
	// not a json object, assume it's an array
	if err != nil {
		// parse into an array
		var arr []interface{}
		err := json.Unmarshal([]byte(input), &arr)
		if err != nil {
			panic(err)
		}

		var arrayTotal int
		for _, v := range arr {
			// marshal each array element into a string, then pass it back into part2
			str, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			arrayTotal += part2(string(str))
		}

		return arrayTotal
	}

	// if any value in the object is "red" this object & its children RETURN ZERO
	for _, v := range obj {
		// have to convert interface into a string first
		str, ok := v.(string)
		if ok && str == "red" {
			return 0
		}
	}

	var total int
	for _, v := range obj {
		str, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		total += part2(string(str))
	}

	return total
}
