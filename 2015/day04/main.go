package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := md5StockingStuffer(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func md5StockingStuffer(input string, part int) int {
	prefixZeroes := 5
	if part == 2 {
		prefixZeroes = 6
	}

	for i := 0; i < math.MaxInt32; i++ {
		toHash := fmt.Sprintf("%s%d", input, i)
		hashed := fmt.Sprintf("%x", md5.Sum([]byte(toHash)))
		if strings.HasPrefix(hashed, strings.Repeat("0", prefixZeroes)) {
			return i
		}
	}

	panic("no hash found")
}
