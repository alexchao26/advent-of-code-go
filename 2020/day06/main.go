package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Reading file" + err.Error())
	}

	if part == 1 {
		ans := part1(string(fileBytes))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(string(fileBytes))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	var sum int

	groups := strings.Split(input, "\n\n")
	for _, group := range groups {
		questionsSeen := map[string]bool{}

		people := strings.Split(group, "\n")
		for _, person := range people {
			for _, question := range strings.Split(person, "") {
				questionsSeen[question] = true
			}
		}

		sum += len(questionsSeen)
	}

	return sum
}

func part2(input string) int {
	var sum int

	groups := strings.Split(input, "\n\n")
	for _, group := range groups {
		questionsToCount := map[string]int{}

		people := strings.Split(group, "\n")
		for _, person := range people {
			for _, question := range strings.Split(person, "") {
				questionsToCount[question]++
			}
		}

		for _, count := range questionsToCount {
			if count == len(people) {
				sum++
			}
		}
	}

	return sum
}
