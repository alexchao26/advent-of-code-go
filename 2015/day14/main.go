package main

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

	ans := reindeerOlympics(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func reindeerOlympics(input string, part int) int {
	reindeerToDistanceMap := map[string][]int{}

	for _, line := range strings.Split(input, "\n") {
		var name string
		var speed, runTime, restTime int
		_, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &runTime, &restTime)
		if err != nil {
			panic(err)
		}

		// 1-index the distances slice, indices line up with elapsed time
		reindeerToDistanceMap[name] = append(reindeerToDistanceMap[name], 0)

		var dist int
		remainingRunTime := runTime
		remainingRestTime := restTime
		for t := 0; t < 2503; t++ {
			if remainingRunTime > 0 {
				dist += speed
				remainingRunTime--
			} else {
				remainingRestTime--
				if remainingRestTime == 0 {
					remainingRunTime = runTime
					remainingRestTime = restTime
				}
			}
			reindeerToDistanceMap[name] = append(reindeerToDistanceMap[name], dist)
		}
	}

	// for part 1 return the furthest end distance (time 2503 seconds)
	if part == 1 {
		var furthest int
		for _, distSli := range reindeerToDistanceMap {
			furthest = mathutil.MaxInt(distSli[2503], furthest)
		}
		return furthest
	}

	// for part 2, score each second, then find the highest score
	reindeerScores := map[string]int{}
	for sec := 1; sec <= 2503; sec++ {
		var names []string
		var bestDist int
		for name, distanceSli := range reindeerToDistanceMap {
			if distanceSli[sec] > bestDist {
				names = []string{name}
				bestDist = distanceSli[sec]
			} else if distanceSli[sec] == bestDist {
				names = append(names, name)
			}
		}

		for _, name := range names {
			reindeerScores[name]++
		}
	}

	var bestScore int
	for _, v := range reindeerScores {
		bestScore = mathutil.MaxInt(bestScore, v)
	}

	return bestScore
}
