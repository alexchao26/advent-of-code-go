package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

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

	ans := pulsePropagation(input, part)
	util.CopyToClipboard(fmt.Sprintf("%v", ans))
	fmt.Println("Output:", ans)
}

func pulsePropagation(input string, part int) int {
	modules := parseInput(input)

	var lowPulses, highPulses int

	buttonPresses := 1000
	if part == 2 {
		// let this cycle infinitely so I can figure out cycle times for part 2
		buttonPresses = math.MaxInt64
	}

	// for part 2:
	// looking at input, rx's only input is &lb, which is a conjunction, so needs to get ALL high signals to send a low to rx
	// lb is fed from four other modules that all need to send high signals:
	// &rz, &lf, &br, &fk
	// figuring out the cycle times of these four then maybe the LCM will be the answer if the input is kind?

	lastCycleForHighPulse := map[string]int{
		"rz": -1,
		"lf": -1,
		"br": -1,
		"fk": -1,
	}

	cycles := []int{}

	for i := 0; i < buttonPresses; i++ {
		if part == 2 && len(cycles) == 4 {
			break
		}

		queue := []pulse{}
		queue = append(queue, pulse{
			isLowPulse:  true,
			source:      "button",
			destination: "broadcaster",
		})

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if p.isLowPulse {
				lowPulses++
			} else {
				highPulses++
			}

			if val, ok := lastCycleForHighPulse[p.source]; ok && !p.isLowPulse {
				// fmt.Println("found for ", p.source, i+1)
				if val == -1 {
					lastCycleForHighPulse[p.source] = i + 1
				} else {
					cycles = append(cycles, (i+1)-val)
				}
			}

			if _, ok := modules[p.destination]; !ok {
				continue
			}

			switch modules[p.destination].moduleType {
			case "broadcaster":
				for _, dest := range modules[p.destination].destinations {
					queue = append(queue, pulse{
						isLowPulse:  p.isLowPulse,
						source:      "broadcaster",
						destination: dest,
					})
				}
			case "flipflop":
				if p.isLowPulse {
					for _, dest := range modules[p.destination].destinations {
						queue = append(queue, pulse{
							// if it was on, it flips off and sends a low pulse
							// if it was off, then sends a high pulse (isLowPulse = false)
							isLowPulse:  modules[p.destination].flipFlopIsOn,
							source:      p.destination,
							destination: dest,
						})
					}
					// flip it
					modules[p.destination].flipFlopIsOn = !modules[p.destination].flipFlopIsOn
				}
			case "conjunction":
				modules[p.destination].conjunctionInputsMapWasLastPulseHigh[p.source] = !p.isLowPulse
				allHigh := true
				for source, wasStrongPulse := range modules[p.destination].conjunctionInputsMapWasLastPulseHigh {
					_ = source
					if !wasStrongPulse {
						allHigh = false
						break
					}
				}

				for _, dest := range modules[p.destination].destinations {
					queue = append(queue, pulse{
						// all high sends a low pulse, otherwise high pulse
						isLowPulse:  allHigh,
						source:      p.destination,
						destination: dest,
					})
				}
			default:
				panic("unexpected module type" + modules[p.destination].moduleType)
			}
		}
	}

	// wow that worked, super generous on the inputs...
	if part == 2 {
		ans := 1
		for _, c := range cycles {
			ans *= c
		}
		return ans
	}

	return lowPulses * highPulses
}

type module struct {
	moduleType                           string
	name                                 string
	flipFlopIsOn                         bool
	conjunctionInputsMapWasLastPulseHigh map[string]bool
	destinations                         []string
}

type pulse struct {
	isLowPulse          bool
	source, destination string
}

func parseInput(input string) (ans map[string]*module) {
	ans = map[string]*module{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")

		mod := module{
			moduleType:                           "",
			flipFlopIsOn:                         false,
			conjunctionInputsMapWasLastPulseHigh: map[string]bool{},
			destinations:                         []string{},
		}

		if parts[0] == "broadcaster" {
			mod.moduleType = "broadcaster"
			mod.name = "broadcaster"
			mod.destinations = strings.Split(parts[1], ", ")
		} else if parts[0][:1] == "%" {
			mod.moduleType = "flipflop"
			mod.name = parts[0][1:]
			mod.destinations = strings.Split(parts[1], ", ")
		} else if parts[0][:1] == "&" {
			mod.moduleType = "conjunction"
			mod.name = parts[0][1:]
			mod.destinations = strings.Split(parts[1], ", ")
		} else {
			panic("unidentified module type: " + line)
		}

		ans[mod.name] = &mod
	}

	// initialize conjunction maps with all their source modules
	for name, module := range ans {
		for _, dest := range module.destinations {

			if _, ok := ans[dest]; !ok {
				continue
			}
			if ans[dest].moduleType == "conjunction" {
				ans[dest].conjunctionInputsMapWasLastPulseHigh[name] = false
			}
		}
	}

	return ans
}
