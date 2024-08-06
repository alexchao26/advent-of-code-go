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
	parsed := parseInput(input)
	ans := 0
	for _, step := range parsed {
		ans += hash(step)
	}

	return ans
}

func hash(step string) int {
	ans := 0
	for _, char := range strings.Split(step, "") {
		asciiVal := cast.ToASCIICode(char)
		ans += asciiVal
		ans *= 17
		ans %= 256
	}
	return ans
}

func part2(input string) int {
	boxes := make([]box, 256)
	// optimization to keep a linked list within in box, but likely not necessary...
	labelToBoxIndex := map[string]int{}

	steps := parseInput(input)

	for _, step := range steps {
		if strings.Contains(step, "=") {
			parts := strings.Split(step, "=")
			label := parts[0]
			focalLength := cast.ToInt(parts[1])

			boxIndex := hash(label)

			if oldBoxIndex, ok := labelToBoxIndex[label]; ok {
				if oldBoxIndex != boxIndex {
					panic("hashes should be the same...")
				}
				// iterate and update focalLength of found box
				for i := range len(boxes[boxIndex]) {
					if boxes[boxIndex][i].label == label {
						boxes[boxIndex][i].focalLength = focalLength
					}
				}
			} else {
				boxes[boxIndex] = append(boxes[boxIndex], lense{
					label:       label,
					focalLength: focalLength,
				})
				labelToBoxIndex[label] = boxIndex
			}

		} else if strings.Contains(step, "-") {
			label := step[:len(step)-1]
			if boxIndex, ok := labelToBoxIndex[label]; ok {
				// switch it all the way to the end
				for i := range len(boxes[boxIndex]) - 1 {
					if boxes[boxIndex][i].label == label {
						boxes[boxIndex][i], boxes[boxIndex][i+1] = boxes[boxIndex][i+1], boxes[boxIndex][i]
					}
				}
				// cut off end, remove from map
				boxes[boxIndex] = boxes[boxIndex][:len(boxes[boxIndex])-1]
				delete(labelToBoxIndex, label)
			}
		} else {
			panic("unexpected step format: " + step)
		}
	}

	ans := 0

	for boxIndex, box := range boxes {
		for lenseIndex, lense := range box {
			ans += (boxIndex + 1) * (lenseIndex + 1) * lense.focalLength
		}
	}

	return ans
}

type lense struct {
	label       string
	focalLength int
}
type box []lense

func parseInput(input string) (ans []string) {
	return strings.Split(input, ",")
}
