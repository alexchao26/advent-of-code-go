package main

import (
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

	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid", // optional
	}
	var ans int
	for _, entry := range parsed {
		hasAllFields := true
		for _, field := range requiredFields {
			if entry[field] == "" {
				hasAllFields = false
				break
			}
		}
		if hasAllFields {
			ans++
		}
	}

	return ans
}

func part2(input string) int {
	passports := parseInput(input)

	var ans int
	for _, entry := range passports {
		if validateFields(entry) {
			ans++
		}
	}

	return ans
}

func parseInput(input string) []map[string]string {
	var passports []map[string]string

	lines := strings.Split(input, "\n\n")

	for _, l := range lines {
		passportDetails := map[string]string{}

		oneLine := strings.ReplaceAll(l, "\n", " ")
		for _, entry := range strings.Split(oneLine, " ") {
			splitEntry := strings.Split(entry, ":")
			passportDetails[splitEntry[0]] = splitEntry[1]
		}
		passports = append(passports, passportDetails)
	}

	return passports
}

func validateFields(entry map[string]string) bool {
	byr := entry["byr"]
	if byr > "2002" || byr < "1920" || len(byr) != 4 {
		return false
	}

	iyr := entry["iyr"]
	if iyr > "2020" || iyr < "2010" || len(iyr) != 4 {
		return false
	}

	eyr := entry["eyr"]
	if eyr > "2030" || eyr < "2020" || len(eyr) != 4 {
		return false
	}

	hgt := entry["hgt"]
	if hgt == "" {
		return false
	}

	var num int
	var unit string
	fmt.Sscanf(hgt, "%d%s", &num, &unit)

	switch unit {
	case "cm":
		if num < 150 || num > 193 {
			return false
		}
	case "in":
		if num < 59 || num > 76 {
			return false
		}
	default:
		return false
	}

	hcl := entry["hcl"]
	reg := regexp.MustCompile("^#[0-9a-f]{6}$")
	if !reg.Match([]byte(hcl)) {
		return false
	}

	ecl := entry["ecl"]
	reg = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	if !reg.Match([]byte(ecl)) {
		return false
	}

	pid := entry["pid"]
	reg = regexp.MustCompile("^[0-9]{9}$")
	if !reg.Match([]byte(pid)) {
		return false
	}

	return true
}
