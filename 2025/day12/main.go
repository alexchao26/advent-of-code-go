package main

import (
	_ "embed"
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
	ans := part1(input)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func part1(input string) int {
	shapeSizes, regions := parseInput(input)

	count := 0
	for _, r := range regions {
		size := 0
		for j, cnt := range r.counts {
			size += cnt * shapeSizes[j]
		}

		if size <= r.width*r.height {
			count++
		}
	}

	return count
}

type region struct {
	width  int
	height int
	counts []int
}

func parseInput(input string) (map[int]int, []region) {
	lines := strings.Split(input, "\n")
	shapeSizes := make(map[int]int)
	var regions []region
	i := 0

	for i < len(lines) {
		text := strings.TrimSpace(lines[i])
		if text == "" {
			i++
			continue
		}

		if strings.Contains(text, "x") {
			break
		}

		idx := cast.ToInt(text[:len(text)-1])
		size := 0
		i++

		for i < len(lines) {
			text := strings.TrimSpace(lines[i])
			if text == "" {
				break
			}
			size += len(text)
			i++
		}

		shapeSizes[idx] = size
	}

	for i < len(lines) {
		text := strings.TrimSpace(lines[i])
		if text == "" {
			i++
			continue
		}

		split := strings.Split(text, ": ")
		sizes := strings.Split(split[0], "x")
		width := cast.ToInt(sizes[0])
		height := cast.ToInt(sizes[1])

		var counts []int
		for _, countStr := range strings.Split(split[1], " ") {
			counts = append(counts, cast.ToInt(countStr))
		}

		regions = append(regions, region{
			width:  width,
			height: height,
			counts: counts,
		})

		i++
	}

	return shapeSizes, regions
}
