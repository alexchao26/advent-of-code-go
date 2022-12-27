package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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
	parsed := parseInput(input)

	// what will monkey 'root' yell?
	v, _ := bfs("root", parsed, map[string]int{})
	return v
}

var numRegexp = regexp.MustCompile("^[0-9]+$")

func bfs(key string, raw map[string]string, solved map[string]int) (int, error) {
	if v, ok := solved[key]; ok {
		return v, nil
	}

	if numRegexp.MatchString(raw[key]) {
		solved[key] = cast.ToInt(raw[key])
		return solved[key], nil
	}

	equation := raw[key]
	parts := strings.Split(equation, " ")

	if len(parts) != 3 {
		return 0, fmt.Errorf("expected 3 parts for %q, got %q", key, equation)
	}

	left, err := bfs(parts[0], raw, solved)
	if err != nil {
		return 0, err
	}
	right, err := bfs(parts[2], raw, solved)
	if err != nil {
		return 0, err
	}

	switch parts[1] {
	case "+":
		solved[key] = left + right
	case "-":
		solved[key] = left - right
	case "*":
		solved[key] = left * right
	case "/":
		solved[key] = left / right
	default:
		panic("error with key: " + key + " string: " + equation)
	}
	return solved[key], nil
}

func part2(input string) int {
	raw := parseInput(input)
	if len(strings.Split(raw["root"], " ")) != 3 {
		panic(fmt.Sprintf("expected 3 parts to %q", raw["root"]))
	}

	// change humn to something that will error in bfs so we know which branch
	// of the equations is fully solvable
	raw["humn"] = "humn_will_error_in_bfs"

	// basically making the root equation leftSymbol / rightSymbol = 1 in the
	// inverted graph
	invertedGraph := map[string]string{"root": "1"}
	rootParts := strings.Split(raw["root"], " ")
	rootParts[1] = "/"
	raw["root"] = strings.Join(rootParts, " ")

	keyToInvert := "root"
	solvedMap := map[string]int{}

	for keyToInvert != "humn" {
		// find the equation, determine which side is easily solvable, and which
		// is not, reverse the equation for the unsolvable variable (aka the one
		// that needs to know what value humn shouts)
		// end at humn
		eq := raw[keyToInvert]
		parts := strings.Split(eq, " ")

		leftRaw, rightRaw := parts[0], parts[2]

		leftVal, errLeft := bfs(leftRaw, raw, solvedMap)
		if errLeft == nil {
			invertedGraph[leftRaw] = cast.ToString(leftVal)
		}
		rightVal, errRight := bfs(rightRaw, raw, solvedMap)
		if errRight == nil {
			invertedGraph[rightRaw] = cast.ToString(rightVal)
		}

		switch parts[1] {
		case "+":
			if errLeft != nil {
				invertedGraph[leftRaw] = fmt.Sprintf("%s - %s", keyToInvert, rightRaw)
				keyToInvert = leftRaw
			} else if errRight != nil {
				invertedGraph[rightRaw] = fmt.Sprintf("%s - %s", keyToInvert, leftRaw)
				keyToInvert = rightRaw
			} else {
				panic(fmt.Sprintf("both vals did not error '+' %q: %q", keyToInvert, eq))
			}
		case "-":
			if errLeft != nil {
				invertedGraph[leftRaw] = fmt.Sprintf("%s + %s", keyToInvert, rightRaw)
				keyToInvert = leftRaw
			} else if errRight != nil {
				invertedGraph[rightRaw] = fmt.Sprintf("%s - %s", leftRaw, keyToInvert)
				keyToInvert = rightRaw
			} else {
				panic(fmt.Sprintf("both vals did not error '-' %q: %q", keyToInvert, eq))
			}
		case "*":
			if errLeft != nil {
				invertedGraph[leftRaw] = fmt.Sprintf("%s / %s", keyToInvert, rightRaw)
				keyToInvert = leftRaw
			} else if errRight != nil {
				invertedGraph[rightRaw] = fmt.Sprintf("%s / %s", keyToInvert, leftRaw)
				keyToInvert = rightRaw
			} else {
				panic(fmt.Sprintf("both vals did not error '/' %q: %q", keyToInvert, eq))
			}
		case "/":
			if errLeft != nil {
				invertedGraph[leftRaw] = fmt.Sprintf("%s * %s", keyToInvert, rightRaw)
				keyToInvert = leftRaw
			} else if errRight != nil {
				invertedGraph[rightRaw] = fmt.Sprintf("%s / %s", leftRaw, keyToInvert)
				keyToInvert = rightRaw
			} else {
				panic(fmt.Sprintf("both vals did not error '*' %q: %q", keyToInvert, eq))
			}

		default:
			panic(fmt.Sprintf("inverting graph: key: %q, eq: %q", keyToInvert, eq))
		}
	}

	v, _ := bfs("humn", invertedGraph, map[string]int{})
	return v
}

func parseInput(input string) map[string]string {
	ans := map[string]string{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		ans[parts[0]] = parts[1]
	}
	return ans
}
