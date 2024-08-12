package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parts, workflowsMap := parseInput(input)

	ans := 0

	for _, p := range parts {
		currentWorkflowName := "in"
		for currentWorkflowName != "A" && currentWorkflowName != "R" {
			wf := workflowsMap[currentWorkflowName]
			for _, rule := range wf.ruleStrings {
				if !strings.Contains(rule, ":") {
					currentWorkflowName = rule
					break
				}

				colonSplit := strings.Split(rule, ":")
				output := colonSplit[1]

				if strings.Contains(colonSplit[0], "<") {
					conditionalParts := strings.Split(colonSplit[0], "<")
					if p[conditionalParts[0]] < cast.ToInt(conditionalParts[1]) {
						currentWorkflowName = output
						break
					}
				} else if strings.Contains(colonSplit[0], ">") {
					conditionalParts := strings.Split(colonSplit[0], ">")
					if p[conditionalParts[0]] > cast.ToInt(conditionalParts[1]) {
						currentWorkflowName = output
						break
					}
				} else {
					panic("unexpected workflow rule conditional: " + rule)
				}
			}
		}

		if currentWorkflowName == "A" {
			ans += p.sumRatings()
		}
	}

	return ans
}

type part map[string]int

func (p part) sumRatings() int {
	total := 0
	for _, v := range p {
		total += v
	}
	return total
}

type workflow struct {
	name        string
	ruleStrings []string
}

func parseInput(input string) (parts []part, workflowsMap map[string]workflow) {
	inputParts := strings.Split(input, "\n\n")

	workflowsMap = map[string]workflow{}

	// process workflows
	for _, line := range strings.Split(inputParts[0], "\n") {
		lineParts := strings.Split(line, "{")
		wf := workflow{
			name:        lineParts[0],
			ruleStrings: strings.Split(lineParts[1][:len(lineParts[1])-1], ","),
		}
		workflowsMap[wf.name] = wf
	}

	for _, line := range strings.Split(inputParts[1], "\n") {
		withoutBraces := line[1 : len(line)-1]
		p := part{}
		for _, ratingStr := range strings.Split(withoutBraces, ",") {
			ratingParts := strings.Split(ratingStr, "=")
			p[ratingParts[0]] = cast.ToInt(ratingParts[1])
		}
		parts = append(parts, p)
	}

	return parts, workflowsMap
}

func part2(input string) int {
	_, workflowsMap := parseInput(input)

	// 1 to 4000 bounds for each rating...
	boundedParts := map[string][2]int{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}

	return updatePartBoundsAndSplit(boundedParts, workflowsMap, "in", 0)
}

func updatePartBoundsAndSplit(boundedParts map[string][2]int, workflowsMap map[string]workflow, currentWorkflow string, debugDepth int) int {
	if currentWorkflow == "R" {
		return 0
	}
	if currentWorkflow == "A" {
		product := 1
		for _, bounds := range boundedParts {
			product *= bounds[1] - bounds[0] + 1
		}
		return product
	}

	// split based on rules...
	total := 0

	// for each rule,
	// the rule either passes and moves onto a different workflow,
	// or fails and checks the next rule
	// need to sum up both forks
	//
	// passing is handled via recursion, failing is handled via looping to the next rule
	// in both cases the bounds need to be updated
	for _, rule := range workflowsMap[currentWorkflow].ruleStrings {
		// just the next workflow to go after
		if !strings.Contains(rule, ":") {
			nextWorkflowName := rule
			total += updatePartBoundsAndSplit(boundedParts, workflowsMap, nextWorkflowName, debugDepth+1)
			break
		}

		colonSplit := strings.Split(rule, ":")
		nextWorkflowName := colonSplit[1]

		if strings.Contains(colonSplit[0], "<") {
			conditionalParts := strings.Split(colonSplit[0], "<")
			ratingName := conditionalParts[0]
			ratingTestValue := cast.ToInt(conditionalParts[1])

			// fork the part that passes the < conditional
			copyOfBounds := copyBoundedPartsMap(boundedParts)
			copyOfBounds[ratingName] = [2]int{
				copyOfBounds[ratingName][0],
				ratingTestValue - 1,
			}
			// check that the new bounds are still valid
			if copyOfBounds[ratingName][0] <= copyOfBounds[ratingName][1] {
				total += updatePartBoundsAndSplit(copyOfBounds, workflowsMap, nextWorkflowName, debugDepth+1)
			}

			// second fork for failing the conditional, need to update the boundedParts to fail
			boundedParts[ratingName] = [2]int{
				ratingTestValue,
				boundedParts[ratingName][1],
			}
			// check that the new bounds are still valid
			if boundedParts[ratingName][0] > boundedParts[ratingName][1] {
				break
			}
		} else if strings.Contains(colonSplit[0], ">") {
			conditionalParts := strings.Split(colonSplit[0], ">")
			ratingName := conditionalParts[0]
			ratingTestValue := cast.ToInt(conditionalParts[1])

			// fork the part that passes the > conditional
			copyOfBounds := copyBoundedPartsMap(boundedParts)
			copyOfBounds[ratingName] = [2]int{
				ratingTestValue + 1,
				copyOfBounds[ratingName][1],
			}

			// check that the new bounds are still valid before recursing
			if copyOfBounds[ratingName][0] <= copyOfBounds[ratingName][1] {
				total += updatePartBoundsAndSplit(copyOfBounds, workflowsMap, nextWorkflowName, debugDepth+1)
			}

			// second fork for failing the conditional, need to update the boundedParts to fail
			boundedParts[ratingName] = [2]int{
				boundedParts[ratingName][0],
				ratingTestValue,
			}
			// check that the new bounds are still valid
			if boundedParts[ratingName][0] > boundedParts[ratingName][1] {
				break
			}
		} else {
			panic("unexpected workflow rule conditional: " + rule)
		}
	}

	return total
}

func copyBoundedPartsMap(boundedParts map[string][2]int) map[string][2]int {
	cp := map[string][2]int{}
	for k, v := range boundedParts {
		cp[k] = v
	}
	return cp
}
