package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
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
		// drg,gvw,jbp,jgc,qjb,z15,z22,z35
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	calc := parseInput(input)
	res := 0

	cache := map[string]int{}

	for k, v := range calc.wires {
		cache[k] = v
	}

	for k := range calc.outputToSource {
		if k[0] != 'z' {
			continue
		}

		val, err := getWireVal(k, calc, cache, 0)

		if err != nil {
			log.Fatalf("Did not expect error\n")
		}

		if val == 0 {
			continue
		}

		i := cast.ToInt(k[1:])

		res += 1 << int(i)
	}

	return res
}

func part2(input string) string {
	calc := parseInput(input)

	swaps := map[string]bool{}
	ans := map[string]bool{}

	fmt.Println("this will take a while to run")
	backtrack(0, calc, swaps, ans)

	wires := []string{}
	for wire := range ans {
		wires = append(wires, wire)
	}

	sort.Strings(wires)

	return strings.Join(wires, ",")
}

func getWireVal(wire string, calc Calculator, cache map[string]int, depth int) (int, error) {
	if val, ok := cache[wire]; ok {
		return val, nil
	}

	if depth > 200 {
		return -1, fmt.Errorf("likely a loop, exit")
	}

	op, ok := calc.outputToSource[wire]

	if !ok {
		return -1, fmt.Errorf("cannot get value for wire %v", wire)
	}

	v1, _ := getWireVal(op.inputs[0], calc, cache, depth+1)
	v2, _ := getWireVal(op.inputs[1], calc, cache, depth+1)

	val := 0
	if op.op == "AND" {
		val = v1 & v2
	} else if op.op == "OR" {
		val = v1 | v2
	} else if op.op == "XOR" {
		val = v1 ^ v2
	}

	cache[wire] = val

	return val, nil
}

func getFirstWrong(calc Calculator) (int, map[string]int, error) {
	cRes := map[string]int{}
	bRes := 100

	for range 10 {
		cache := map[string]int{}

		carry := 0
		for b := range 45 {
			xStr := fmt.Sprintf("x%02d", b)
			yStr := fmt.Sprintf("y%02d", b)
			zStr := fmt.Sprintf("z%02d", b)

			cache[xStr] = rand.Intn(2)
			cache[yStr] = rand.Intn(2)

			x := cache[xStr]
			y := cache[yStr]

			zCalc, err := getWireVal(zStr, calc, cache, 0)

			if err != nil {
				return -1, nil, err
			}

			zExp := (x + y + carry) % 2
			carry = (x + y + carry) / 2

			if zCalc != zExp && b < bRes {
				bRes = b
				cRes = cache
				break
			}
		}
	}

	if bRes == 100 {
		bRes = -1
	}

	return bRes, cRes, nil
}

func checkCalc(calc Calculator) bool {
	bErr, _, err := getFirstWrong(calc)

	return bErr == -1 && err == nil
}

func backtrack(b int, calc Calculator, swaps map[string]bool, ans map[string]bool) {
	if len(ans) > 0 {
		return
	}
	// fmt.Printf("Checking %v with swaps\n%v\n", b, swaps)
	if len(swaps) == 8 {
		if checkCalc(calc) {
			for k := range swaps {
				ans[k] = true
			}
		}

		return
	}

	bErr, conns, err := getFirstWrong(calc)

	if bErr < b || err != nil {
		return
	}

	conns2 := map[string]Connections{}

	for k, v := range calc.outputToSource {
		conns2[k] = v
	}

	for c1 := range conns {
		if swaps[c1] || c1[0] == 'x' || c1[0] == 'y' {
			continue
		}

		for c2 := range conns2 {
			if c2 == c1 || swaps[c2] || c2[0] == 'x' || c2[0] == 'y' {
				continue
			}

			calc.outputToSource[c1], calc.outputToSource[c2] = calc.outputToSource[c2], calc.outputToSource[c1]

			swaps[c1] = true
			swaps[c2] = true

			backtrack(bErr+1, calc, swaps, ans)

			delete(swaps, c2)
			delete(swaps, c1)

			calc.outputToSource[c1], calc.outputToSource[c2] = calc.outputToSource[c2], calc.outputToSource[c1]
		}
	}
}

type Connections struct {
	inputs []string
	op     string
}

type Calculator struct {
	outputToSource map[string]Connections
	wires          map[string]int
}

func parseInput(input string) Calculator {
	parts := strings.Split(input, "\n\n")

	wires := map[string]int{}
	for _, line := range strings.Split(parts[0], "\n") {
		wires[line[:3]] = cast.ToInt(line[5:])
	}

	outputToSource := map[string]Connections{}
	for _, line := range strings.Split(parts[1], "\n") {
		lineParts := strings.Split(line, " ")
		in1, op, in2, out := lineParts[0], lineParts[1], lineParts[2], lineParts[4]
		outputToSource[out] = Connections{
			inputs: []string{in1, in2},
			op:     op,
		}
	}

	return Calculator{outputToSource, wires}
}
