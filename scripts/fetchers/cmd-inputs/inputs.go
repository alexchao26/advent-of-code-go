package main

import (
	"fmt"
	"path/filepath"

	"github.com/alexchao26/advent-of-code-go/scripts/fetchers"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	// determine day to fetch
	day, year, cookie := fetchers.ParseFlags()
	fmt.Println("fetching for day", day)

	// make the request
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	body := fetchers.GetWithAOCCookie(url, cookie)

	// write to file
	filename := filepath.Join(util.Dirname(), "../../..", fmt.Sprintf("%d/day%02d-input.txt", year, day))
	fetchers.WriteToFile(filename, body)

	fmt.Println("Wrote to file: ", filename)

	fmt.Println("Done!")
}
