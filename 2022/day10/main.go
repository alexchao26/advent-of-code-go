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
	instructions := parseInput(input)

	X := 1
	sum := 0

	i := 0 // what the current instruction is
	for cycle := 1; cycle <= 220; cycle++ {
		// "during" equates to the start of the cycle...
		if (cycle-20)%40 == 0 {
			sum += X * cycle
		}

		switch instructions[i].name {
		case "addx":
			// decrement cycles on that instruction
			// IF it hits zero add V
			// AND move to next step
			instructions[i].cycles--
			if instructions[i].cycles == 0 {
				X += instructions[i].val
				i++
			}
		case "noop":
			// just increment to next instruction
			i++
		}
	}

	return sum
}

func part2(input string) string {
	instructions := parseInput(input)

	X := 1 // doubles as sprite's center coordinate

	// 6 rows by 40 wide screen, starts all off
	CRT := [6][40]string{}
	for i, rows := range CRT {
		for j := range rows {
			CRT[i][j] = "."
		}
	}

	i := 0 // what the current instruction is
	for cycle := 1; i < len(instructions); cycle++ {
		// if (cycle-20)%40 == 0 {
		// 	sum += X * cycle
		// }

		/*
			X = horizontal position of middle of (3 pixel wide) sprite
			axis draws left to right, top to bottom, 40 wide x 6 high
			1---40
			41---80
			...
			201---240

			draws 1 pixel per cycle
			light up pixels IF the pixel being drawn is the same as one of the sprite's 3 pixels

		*/

		// calculate which pixel is being drawn... ZERO INDEXED
		pixelRow := (cycle - 1) / 40
		pixelCol := (cycle - 1) % 40

		// see if the spite's horizontal location overlaps that pixelCol
		spriteLeft, spriteRight := X-1, X+1
		if spriteLeft <= pixelCol && spriteRight >= pixelCol {
			CRT[pixelRow][pixelCol] = "#"
		}

		switch instructions[i].name {
		case "addx":
			// decrement cycles on that instruction
			// IF it hits zero add V
			// AND move to next step
			instructions[i].cycles--
			if instructions[i].cycles == 0 {
				X += instructions[i].val
				i++
			}
		case "noop":
			// just increment to next instruction
			i++
		}

	}
	log := ""
	for _, rows := range CRT {
		for _, cell := range rows {
			log += cell
		}
		log += "\n"
	}
	fmt.Println(log)
	return log
}

type instruction struct {
	name   string
	val    int
	cycles int
}

func parseInput(input string) (ans []instruction) {
	for _, l := range strings.Split(input, "\n") {
		switch l[:4] {
		case "addx":
			ans = append(ans, instruction{
				name:   "addx",
				val:    cast.ToInt(l[5:]),
				cycles: 2,
			})
		case "noop":
			ans = append(ans, instruction{
				name:   "noop",
				cycles: 1,
			})
		default:
			panic("input line: " + l)
		}
	}
	return ans
}
