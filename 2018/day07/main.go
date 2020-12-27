package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
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
		ans := part2(util.ReadFile("./input.txt"), 5, 60)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) string {
	graph := parseInputs(input)

	completedSteps := make(map[string]bool)
	var stepsOrder string

	for len(completedSteps) != len(graph) {
		var readySteps []string

		for name, prereqs := range graph {
			if completedSteps[name] {
				continue
			}

			prereqsReady := true
			for _, v := range prereqs {
				if !completedSteps[v] {
					prereqsReady = false
					break
				}
			}
			if prereqsReady {
				readySteps = append(readySteps, name)
			}
		}

		// run the lowest ready step
		sort.Strings(readySteps)
		completedSteps[readySteps[0]] = true
		stepsOrder += readySteps[0]
	}

	return stepsOrder
}

func part2(input string, workers, fudgeTime int) int {
	graph := parseInputs(input)

	stepsCompletedAt := make(map[string]int, len(graph))

	workerLogs := make([][]string, workers)
	for i := 0; i < workers; i++ {
		workerLogs[i] = []string{}
	}

	for time := 0; len(stepsCompletedAt) < len(graph); time++ {
		availableWorkers := []int{}
		for i, w := range workerLogs {
			if time >= len(w) {
				availableWorkers = append(availableWorkers, i)
			}
		}

		if len(availableWorkers) > 0 {
			// get ready steps
			var readySteps []string

			for name, prereqs := range graph {
				if stepsCompletedAt[name] != 0 {
					continue
				}

				prereqsCompletionTimes := make([]int, len(prereqs))
				// might want more data here
				for _, pre := range prereqs {
					if stepsCompletedAt[pre] == 0 {
						prereqsCompletionTimes = append(prereqsCompletionTimes, math.MaxInt32)
					} else {
						prereqsCompletionTimes = append(prereqsCompletionTimes, stepsCompletedAt[pre]+1)
					}
				}

				if len(prereqsCompletionTimes) == 0 {
					readySteps = append(readySteps, name)
				} else {
					earliestScheduleTime := mathy.MaxInt(prereqsCompletionTimes...)
					if earliestScheduleTime <= time {
						readySteps = append(readySteps, name)
					}
				}
			}

			// schedule steps that are ready
			sort.Strings(readySteps)

			for _, step := range readySteps {
				if len(availableWorkers) > 0 {
					for i := timeForStep(step, fudgeTime); i > 0; i-- {
						workerLogs[availableWorkers[0]] = append(workerLogs[availableWorkers[0]], step)
					}
					// set time this step is done at
					stepsCompletedAt[step] = len(workerLogs[availableWorkers[0]])
					// remove worker from available "pool"
					availableWorkers = availableWorkers[1:]
				}
			}
		}

		// fill out empty time elements in any workerLogs
		for i := range workerLogs {
			for len(workerLogs[i]) < time {
				workerLogs[i] = append(workerLogs[i], "-")
			}
		}
	}

	var longestLength int
	for _, w := range workerLogs {
		if len(w) > longestLength {
			longestLength = len(w)
		}
	}
	return longestLength
}

func parseInputs(input string) map[string][]string {
	graphValToPrereqs := make(map[string][]string)
	lines := strings.Split(input, "\n")
	for _, inst := range lines {
		words := strings.Split(inst, " ")
		if len(words) < 8 {
			continue
		}
		graphValToPrereqs[words[7]] = append(graphValToPrereqs[words[7]], words[1])
		if len(graphValToPrereqs[words[1]]) == 0 {
			graphValToPrereqs[words[1]] = []string{}
		}
	}
	return graphValToPrereqs
}

func timeForStep(char string, fudgeFactor int) int {
	byteDiff := []byte(char)[0] - byte('A')
	return fudgeFactor + int(byteDiff) + 1
}
