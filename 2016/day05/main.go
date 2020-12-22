package main

import (
	"crypto/md5"
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

	ans := md5Chess(util.ReadFile("./input.txt"), part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func md5Chess(input string, part int) string {
	passwordParts := map[int]string{}
	var part1Index int

	for i := 0; len(passwordParts) < 8; i++ {
		in := fmt.Sprintf("%s%d", input, i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(in)))

		if strings.HasPrefix(hash, "00000") {
			if part == 1 {
				passwordParts[part1Index] = hash[5:6]
				part1Index++
			} else {
				if regexp.MustCompile("[0-7]").MatchString(hash[5:6]) {
					index := cast.ToInt(hash[5:6])
					if _, ok := passwordParts[index]; !ok {
						value := hash[6:7]
						passwordParts[index] = value
					}
				}
			}
		}
	}

	var password string
	for i := 0; i < 8; i++ {
		password += passwordParts[i]
	}

	return password
}
