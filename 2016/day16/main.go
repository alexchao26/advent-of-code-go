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

	ans := dragonChecksum(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func dragonChecksum(input string, part int) string {
	disk := input
	diskLength := 272
	if part == 2 {
		diskLength = 35651584
	}
	for len(disk) < diskLength {
		var sb strings.Builder
		sb.WriteString(disk)
		sb.WriteByte('0')
		for i := len(disk) - 1; i >= 0; i-- {
			if disk[i] == '1' {
				sb.WriteByte('0')
			} else {
				sb.WriteByte('1')
			}
		}
		disk = sb.String()
	}

	disk = disk[0:diskLength]
	for len(disk)%2 == 0 {
		var sb strings.Builder
		for i := 0; i < len(disk); i += 2 {
			if disk[i] == disk[i+1] {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
		disk = sb.String()
	}

	return disk
}
