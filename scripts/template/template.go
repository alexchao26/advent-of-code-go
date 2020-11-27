package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/alexchao26/advent-of-code-go/scripts/fetchers"
	"github.com/alexchao26/advent-of-code-go/util"
)

type TemplateData struct {
	Year int
	Day  string // a string to include the prefixing zero
}

var testTemplateString = `package aoc{{.Year}}

import (
	"testing"
	
	"github.com/alexchao26/advent-of-code-go/util"
)

func TestDay{{.Day}}Part1(t *testing.T) {
	// Run actual problem input
	day{{.Day}}Part1(util.ReadFile("./day{{.Day}}-input.txt"))
}

func TestDay{{.Day}}Part2(t *testing.T) {
	// Run actual problem input
	day{{.Day}}Part2(util.ReadFile("./day{{.Day}}-input.txt"))
}
`

var solutionTemplateString = `package aoc{{.Year}}

func day{{.Day}}Part1(input string) int {
	return 0
}

func day{{.Day}}Part2(input string) int {
	return 0
}
`

func main() {
	day, year, _ := fetchers.ParseFlags()
	data := TemplateData{
		Year: year,
		Day:  fmt.Sprintf("%02d", day),
	}

	testTemp, err := template.New("test-template").Parse(testTemplateString)
	if err != nil {
		panic(err)
	}
	solutionTemp, err := template.New("solution-template").Parse(solutionTemplateString)
	if err != nil {
		panic(err)
	}

	solutionFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d.go", year, day))
	testFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d_test.go", year, day))

	EnsureNotOverwriting(solutionFilename)
	EnsureNotOverwriting(testFilename)

	solutionWriter, err := os.Create(solutionFilename)
	if err != nil {
		panic(err)
	}
	testWriter, err := os.Create(testFilename)
	if err != nil {
		panic(err)
	}

	solutionTemp.Execute(solutionWriter, data)
	testTemp.Execute(testWriter, data)
	fmt.Println("templates made")
}

func EnsureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		panic(fmt.Sprintf("File already exists: %s", filename))
	}
}
