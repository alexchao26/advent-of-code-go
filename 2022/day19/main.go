package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
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
	blueprints := parseInput(input)

	// how many geodes can be opened in 24 minutes?
	sum := 0
	for _, bp := range blueprints {
		st := newState(bp)
		geodesMade := st.calcMostGeodes(0, map[string]int{}, 24, 24)
		// fmt.Println("ID:", bp.id, geodesMade)
		sum += st.blueprint.id * geodesMade
	}

	// total quality of all blueprints, quality = id * (# geodes in 24 min)
	return sum
}

func part2(input string) int {
	blueprints := parseInput(input)
	if len(blueprints) > 3 {
		blueprints = blueprints[:3]
	}

	prod := 1
	for _, bp := range blueprints {
		st := newState(bp)
		geodesMade := st.calcMostGeodes(0, map[string]int{}, 32, 32)
		// fmt.Println(bp.id, geodesMade)
		prod *= geodesMade
	}

	// total quality of all blueprints, quality = id * (# geodes in 24 min)
	return prod
}

type blueprint struct {
	id                                        int
	oreForOreRobot                            int
	oreForClayRobot                           int
	oreForObsidianRobot, clayForObsidianRobot int
	oreForGeodeRobot, obsidianForGeodeRobot   int
}

type state struct {
	blueprint
	ore, clay, obsidian, geode                         int
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
}

func newState(blueprint blueprint) state {
	return state{
		blueprint: blueprint,
		oreRobots: 1,
	}
}

func (s *state) farm() {
	s.ore += s.oreRobots
	s.clay += s.clayRobots
	s.obsidian += s.obsidianRobots
	s.geode += s.geodeRobots
}

func (s *state) hash(time int) string {
	return fmt.Sprint(time, s.ore, s.clay, s.obsidian,
		s.geode, s.oreRobots, s.clayRobots, s.obsidianRobots, s.geodeRobots)
}

// NOT A POINTER METHOD SO A COPY CAN BE MADE
// this is some cheeky Go struct copying, it'd be easier to read if it was just
// directly recreating all the fields
func (s state) copy() state {
	return s
}

func (s *state) calcMostGeodes(time int, memo map[string]int, totalTime int, earliestGeode int) int {
	if time == totalTime {
		return s.geode
	}

	h := s.hash(time)
	if v, ok := memo[h]; ok {
		return v
	}

	if s.geode == 0 && time > earliestGeode {
		return 0
	}

	// factory can try to make any possible robot, will backtrack if necessary
	mostGeodes := s.geode

	// always make geode robots
	if s.ore >= s.oreForGeodeRobot &&
		s.obsidian >= s.obsidianForGeodeRobot {
		cp := s.copy()

		cp.farm()

		cp.ore -= cp.oreForGeodeRobot
		cp.obsidian -= cp.obsidianForGeodeRobot
		cp.geodeRobots++
		if cp.geodeRobots == 1 {
			earliestGeode = mathy.MinInt(earliestGeode, time+1)
		}
		mostGeodes = mathy.MaxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))

		memo[h] = mostGeodes
		return mostGeodes
	}

	if time <= totalTime-16 &&
		s.oreRobots < s.oreForObsidianRobot*2 &&
		s.ore >= s.oreForOreRobot {
		cp := s.copy()
		cp.ore -= cp.oreForOreRobot

		cp.farm()

		cp.oreRobots++
		mostGeodes = mathy.MaxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))
	}
	if time <= totalTime-8 &&
		s.clayRobots < s.clayForObsidianRobot &&
		s.ore >= s.oreForClayRobot {
		cp := s.copy()
		cp.ore -= cp.oreForClayRobot

		cp.farm()

		cp.clayRobots++
		mostGeodes = mathy.MaxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))
	}
	if time <= totalTime-4 &&
		s.obsidianRobots < s.obsidianForGeodeRobot &&
		s.ore >= s.oreForObsidianRobot && s.clay >= s.clayForObsidianRobot {

		cp := s.copy()
		cp.ore -= cp.oreForObsidianRobot
		cp.clay -= cp.clayForObsidianRobot
		cp.farm()

		cp.obsidianRobots++
		mostGeodes = mathy.MaxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))
	}

	// or no factory production this minute
	cp := s.copy()
	cp.ore += cp.oreRobots
	cp.clay += cp.clayRobots
	cp.obsidian += cp.obsidianRobots
	cp.geode += cp.geodeRobots
	mostGeodes = mathy.MaxInt(mostGeodes, cp.calcMostGeodes(time+1, memo, totalTime, earliestGeode))

	memo[h] = mostGeodes
	return mostGeodes
}

func parseInput(input string) (ans []blueprint) {
	// Blueprint 1: Each ore robot costs 3 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 20 clay. Each geode robot costs 2 ore and 12 obsidian.
	for _, line := range strings.Split(input, "\n") {
		bp := blueprint{}
		_, err := fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&bp.id, &bp.oreForOreRobot, &bp.oreForClayRobot, &bp.oreForObsidianRobot,
			&bp.clayForObsidianRobot, &bp.oreForGeodeRobot, &bp.obsidianForGeodeRobot)
		if err != nil {
			panic("parsing: " + err.Error())
		}
		ans = append(ans, bp)
	}
	return ans
}
