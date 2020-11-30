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

var testTemplateString = `package main

import "testing"

var tests1 = []struct {
	name  string
	want  int
	input string
	// add extra args if needed
}{
	// {"actual", ACTUAL_ANSWER, util.ReadFile("input.txt")},
}

func TestPart1(t *testing.T) {
	for _, test := range tests1 {
		t.Run(test.name, func(*testing.T) {
			got := part1(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

var tests2 = []struct {
	name  string
	want  int
	input string
	// add extra args if needed
}{
	// {"actual", ACTUAL_ANSWER, util.ReadFile("input.txt")},
}

func TestPart2(t *testing.T) {
	for _, test := range tests2 {
		t.Run(test.name, func(*testing.T) {
			got := part2(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
`

var solutionTemplateString = `package main

import (
	"flag"
	"fmt"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
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

	solutionFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d/main.go", year, day))
	testFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d/main_test.go", year, day))

	fetchers.MakeDir(filepath.Dir(solutionFilename))

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

	// note: data is no longer used, but keeping it for future reference of text/template
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
