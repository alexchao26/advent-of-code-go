package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

	ans := aluDay24(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func aluDay24(input string, part int) int64 {
	instLines := strings.Split(input, "\n")

	/*
		"Compiled" the assembly to understand that it was 14 separate "programs"
		one for each input character.
		z is the only stateful variable, doing some algebra shows that this is
		basically a stack pairing a=1 and a=26

		determining the relationships between the input characters reduces the
		search size, so it can be brute forced
	*/

	/*
		NOTE: this does not need to be looped to find the largest and smallest
		because the differences are all known so it's just a matter of finding
		the largest or smallest digits from the start
		w0 + 4 = w13 -> max w0 is 5, so w13 = 9
	*/

	// w10 + 8 = w11
	// w5 - 8 = w6
	// do not need to loop these bc they are 8 apart so only one possibility
	w11, w10 := 9, 1
	w5, w6 := 9, 1

	var largest, smallest int64 = 0, 99999999999999
	// w0 + 4 = w13
	for w13, w0 := 9, 5; w0 >= 1; w13, w0 = w13-1, w0-1 {
		// w1 - 6 = w12
		for w12, w1 := 3, 9; w12 >= 1; w12, w1 = w12-1, w1-1 {
			// w3 - 3 = w4
			for w3, w4 := 9, 6; w4 >= 1; w3, w4 = w3-1, w4-1 {
				// w2 - 1 = w9
				for w2, w9 := 9, 8; w9 >= 1; w2, w9 = w2-1, w9-1 {
					// w7 + 7 = w8
					for w7, w8 := 2, 9; w7 >= 1; w7, w8 = w7-1, w8-1 {
						str := fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d",
							w0, w1, w2, w3, w4, w5, w6,
							w7, w8, w9, w10, w11, w12, w13,
						)
						num, err := strconv.ParseInt(str, 10, 64)
						if err != nil {
							panic("parseint" + err.Error())
						}

						// check against hardcoded alu and actual alu implementation
						lastZ := hardcodedALU(num)
						lastZ2 := runALU(num, instLines)

						if lastZ == 0 && lastZ2 == 0 {
							if num > largest {
								largest = num
							}
							if num < smallest {
								smallest = num
							}
						}
					}
				}
			}
		}
	}

	if part == 1 {
		return largest
	}
	return smallest
}

// a literal implementation of the ALU, useful for checking the final answer?
func runALU(model int64, instLines []string) (z int64) {
	registers := map[string]int64{}
	// ALU will only have w,x,y,z keys
	// sub model numbers to test are 14 digit numbers (1-9 ONLY)
	// uses 14 separate inp instructions, each a single digit of the model number (starting from left)
	// 1357... has inputs 1, 3, 5, 7...
	// model number is valid IF variable z == 0
	modelNumberIndex := 0
	modelString := fmt.Sprint(model)
	if len(modelString) != 14 {
		panic("model string should be 14 characters: " + modelString)
	}

	for _, line := range instLines {
		parts := strings.Split(line, " ")

		var bVal int64
		if len(parts) == 3 {
			maybeInt64, err := strconv.Atoi(parts[2]) // could use a regexp too...
			if err != nil {
				// parts[2] is not a number, look up its register
				bVal = registers[parts[2]]
			} else {
				bVal = int64(maybeInt64)
			}
		}
		switch parts[0] {
		case "inp":
			val := cast.ToInt(modelString[modelNumberIndex : modelNumberIndex+1])
			modelNumberIndex++
			registers[parts[1]] = int64(val)
		case "add":
			a := parts[1]
			registers[a] += bVal
		case "mul":
			a := parts[1]
			registers[a] *= bVal
		case "div":
			a := parts[1]
			registers[a] /= bVal
		case "mod":
			a := parts[1]
			registers[a] %= bVal
		case "eql":
			a := parts[1]
			if registers[a] == bVal {
				registers[a] = 1
			} else {
				registers[a] = 0
			}
		default:
			panic("unexpected command type " + parts[0])
		}
	}

	return registers["z"]
}

func hardcodedALU(model int64) (z int64) {
	registers := map[string]int64{}
	modelChars := strings.Split(fmt.Sprint(model), "")
	if len(modelChars) != 14 {
		panic(fmt.Sprintf("expected 14 digit number, got %d", len(modelChars)))
	}

	// all 14 steps are the same except for 3 values
	/*
		inp w
		mul x 0
		add x z
		mod x 26
		div z 1  <--
		add x 14 <--
		eql x w
		eql x 0
		mul y 0
		add y 25
		mul y x
		add y 1
		mul z y
		mul y 0
		add y w
		add y 16 <--
		mul y x
		add z y
	*/
	// div z, add x, add y
	differentVals := [][3]int{
		{1, 14, 16},   // 0
		{1, 11, 3},    // 1
		{1, 12, 2},    // 2
		{1, 11, 7},    // 3
		{26, -10, 13}, // 4
		{1, 15, 6},    // 5
		{26, -14, 10}, // 6
		{1, 10, 11},   // 7
		{26, -4, 6},   // 8
		{26, -3, 5},   // 9
		{1, 13, 11},   // 10
		{26, -3, 4},   // 11
		{26, -9, 4},   // 12
		{26, -12, 6},  // 13
	}
	for i, char := range modelChars {
		inp := int64(cast.ToInt(char))

		x := registers["z"]%26 + int64(differentVals[i][1])

		// always divided by 1 or 26... 1 is a no-op
		registers["z"] /= int64(differentVals[i][0])

		if x != inp {
			registers["z"] = 26*registers["z"] + (inp + int64(differentVals[i][2]))
		}
	}

	return registers["z"]
}
