package aoc

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func GetInput(day, year int, cookie string) {
	fmt.Printf("fetching for day %d, year %d\n", day, year)

	// make the request
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	body := GetWithAOCCookie(url, cookie)

	if strings.HasPrefix(string(body), "Puzzle inputs differ by user") {
		panic("'Puzzle inputs differ by user' response")
	}

	// write to file
	filename := filepath.Join(util.Dirname(), "../..", fmt.Sprintf("%d/day%02d/input.txt", year, day))
	WriteToFile(filename, body)

	fmt.Println("Wrote to file: ", filename)

	fmt.Println("Done!")
}
