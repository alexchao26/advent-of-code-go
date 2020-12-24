package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := oneTimePad(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func oneTimePad(input string, part int) int {
	hashCycles := 1
	if part == 2 {
		hashCycles = 2016 + 1
	}

	hashes := []string{}
	for i := 0; i < 1000; i++ {
		hashes = append(hashes, hash(input, i, hashCycles))
	}
	var keys []int
	for index := 0; len(keys) < 64; index++ {
		currentHash := hashes[0]

		// maintain the next 1000 hashes
		hashes = append(hashes, hash(input, index+1000, hashCycles))
		hashes = hashes[1:]

		if char := hasTriple(currentHash); char != "" {
			// this is slow... but it's simple
			pattern := regexp.MustCompile(fmt.Sprintf("[%s]{5}", char))
			if pattern.MatchString(strings.Join(hashes, ",")) {
				keys = append(keys, index)
			}
		}
	}

	return keys[len(keys)-1]
}

func hash(input string, index, cycles int) string {
	hashed := fmt.Sprintf("%s%d", input, index)
	for i := 0; i < cycles; i++ {
		hashed = fmt.Sprintf("%x", md5.Sum([]byte(hashed)))
	}
	return hashed
}

func hasTriple(in string) string {
	for i := 2; i < len(in); i++ {
		if in[i-2] == in[i-1] && in[i-1] == in[i] {
			return string(in[i])
		}
	}
	return ""
}
