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

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// {"actual", ACTUAL_ANSWER, util.ReadFile("input.txt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// {"actual", ACTUAL_ANSWER, util.ReadFile("input.txt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
`

var solutionTemplateString = `package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []int) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		ans = append(ans, mathutil.StrToInt(l))
	}
	return ans
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
