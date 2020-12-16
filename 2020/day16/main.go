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

	ans := ticketTranslation(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func ticketTranslation(input string, part int) int {
	rules, myTicket, nearbyTickets := parseInput(input)

	var validTickets [][]int // for part 2
	var errorRate int
	for _, ticket := range nearbyTickets {
		isValidTicket := true
		for _, ticketValue := range ticket {
			valuePassesRules := false
			for _, ruleBounds := range rules {
				if valuePassesRule(ticketValue, ruleBounds) {
					valuePassesRules = true
					break
				}
			}
			if !valuePassesRules {
				errorRate += ticketValue
				isValidTicket = false // for part 2
				break
			}
		}
		// filter out valid tickets for part 2
		if isValidTicket {
			validTickets = append(validTickets, ticket)
		}
	}

	if part == 1 {
		return errorRate
	}

	// part 2, figure out which field belongs to which
	fieldNameToIndex := map[string]int{}
	skipTicketIndices := map[int]bool{}
	// run until all the rules are accounted for
	for len(rules) > 0 {
		// iterate over "columns" of the valid tickets matrix
		for ticketValIndex := range validTickets[0] {
			if skipTicketIndices[ticketValIndex] {
				continue
			}
			// run all the rules against each ticket, store which ones pass for
			// all values at this ticket index. if only one rule applies, it
			// must be for this index within a ticket
			var passingNames []string
			for ruleName, ruleBounds := range rules {
				allValuesPassed := true
				// iterate over all tickets and if any fail for this rule, break out
				for _, ticket := range validTickets {
					ticketValue := ticket[ticketValIndex]
					if !valuePassesRule(ticketValue, ruleBounds) {
						allValuesPassed = false
						break
					}
				}

				// append this rule name as one that passed for these values
				if allValuesPassed {
					passingNames = append(passingNames, ruleName)
				}
			}

			// if only one rule passes, assign it to this ticket value index
			// remove it from the rules list
			if len(passingNames) == 1 {
				fieldNameToIndex[passingNames[0]] = ticketValIndex
				// remove the rule from the map b/c we've determined its index
				delete(rules, passingNames[0])
				// remember which indices have already been taken by a rule
				skipTicketIndices[ticketValIndex] = true
			}
		}
	}

	// get final answer by multiplying all ticket details/rules prefixed "departure"
	departureProduct := 1
	for rule, valueIndex := range fieldNameToIndex {
		if strings.HasPrefix(rule, "departure") {
			departureProduct *= myTicket[valueIndex]
		}
	}

	return departureProduct
}

func valuePassesRule(value int, ruleBounds [2][2]int) bool {
	firstBounds := ruleBounds[0]
	secondBounds := ruleBounds[1]
	return ((value >= firstBounds[0] && value <= firstBounds[1]) ||
		(value >= secondBounds[0] && value <= secondBounds[1]))
}

func parseInput(input string) (map[string][2][2]int, []int, [][]int) {
	blocks := strings.Split(input, "\n\n")

	// parse rules from first block
	rules := map[string][2][2]int{}
	for _, rule := range strings.Split(blocks[0], "\n") {
		parts := strings.Split(rule, ": ")
		name := parts[0]

		var r1L, r1H, r2L, r2H int
		fmt.Sscanf(parts[1], "%d-%d or %d-%d", &r1L, &r1H, &r2L, &r2H)
		rules[name] = [2][2]int{
			[2]int{r1L, r1H},
			[2]int{r2L, r2H},
		}
	}

	// my ticket values in second block
	splitTicket := strings.Split(blocks[1], "\n")
	var myTicket []int
	for _, v := range strings.Split(splitTicket[1], ",") {
		myTicket = append(myTicket, mathutil.StrToInt(v))
	}

	// all values for nearby tickets
	var nearbyTickets [][]int
	for _, nearby := range strings.Split(blocks[2], "\n")[1:] {
		var near []int
		for _, v := range strings.Split(nearby, ",") {
			near = append(near, mathutil.StrToInt(v))
		}
		nearbyTickets = append(nearbyTickets, near)
	}

	return rules, myTicket, nearbyTickets
}
