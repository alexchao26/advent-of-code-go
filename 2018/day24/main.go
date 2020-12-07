package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
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
	immuneGroup, infectionGroup := parseInput(input)

	for !(len(immuneGroup) == 0 || len(infectionGroup) == 0) {
		immuneGroup, infectionGroup, _ = battle(immuneGroup, infectionGroup)
	}

	var totalWinningUnits int
	for _, g := range immuneGroup {
		totalWinningUnits += g.units
	}
	for _, g := range infectionGroup {
		totalWinningUnits += g.units
	}

	return totalWinningUnits
}

func part2(input string) int {
	// binary search b/c this is kind of computationally expensive to test
	immuneBoostLower, immuneBoostUpper := 0, math.MaxInt16

	for immuneBoostLower < immuneBoostUpper {
		immuneGroup, infectionGroup := parseInput(input)
		boost := (immuneBoostUpper + immuneBoostLower) / 2

		_, _, immuneSystemWon := runWithImmuneBoost(immuneGroup, infectionGroup, boost)
		// if immune system won, try lower numbers
		if immuneSystemWon {
			immuneBoostUpper = boost
		} else {
			// otherwise boost more
			immuneBoostLower = boost + 1
		}
	}

	// run it back w/ the found immuneBoost to get final
	var winningUnits int

	immuneGroup, infectionGroup := parseInput(input)
	finalImmuneSystem, _, _ := runWithImmuneBoost(immuneGroup, infectionGroup, immuneBoostLower)

	for _, group := range finalImmuneSystem {
		winningUnits += group.units
	}

	return winningUnits
}

type group struct {
	groupType   string // immune or infection
	units       int
	hp          int
	weakTo      map[string]bool
	immuneTo    map[string]bool
	attackPower int
	attackType  string
	initiative  int
	number      int // for debugging if needed
}

func (g *group) effectivePower() int {
	return g.units * g.attackPower
}

func (g *group) attackMultiplier(incomingAttackType string) int {
	if g.weakTo[incomingAttackType] {
		return 2
	}
	if g.immuneTo[incomingAttackType] {
		return 0
	}
	return 1
}

func (g *group) calcDamageToTarget(targetGroup *group) int {
	myEP := g.effectivePower()
	multiplier := targetGroup.attackMultiplier(g.attackType)
	return myEP * multiplier
}

func (g *group) takeDamage(damage int, attackType string) {
	// integer division removes whole units only, per the prompt
	g.units -= g.attackMultiplier(attackType) * damage / g.hp
}

func (g *group) String() string {
	return fmt.Sprintf("{ No.%d; \tHP:%d;\tInitiative:%d\tEP:%d\tUnits:%d\tAttack:%s %d\tweaknesses:%v\timmunities:%v }",
		g.number, g.hp, g.initiative, g.effectivePower(), g.units, g.attackType, g.attackPower, g.weakTo, g.immuneTo)
}

func parseInput(input string) ([]*group, []*group) {
	factions := strings.Split(input, "\n\n")

	immuneLines := strings.Split(factions[0], "\n")[1:]
	infectLines := strings.Split(factions[1], "\n")[1:]

	immuneGroup := makeGroups(immuneLines, "immune")
	infectGroup := makeGroups(infectLines, "infection")

	return immuneGroup, infectGroup
}

func makeGroups(lines []string, groupType string) []*group {
	var groups []*group
	for i, str := range lines {
		g := group{
			groupType: groupType,
			number:    i + 1,
			weakTo:    map[string]bool{},
			immuneTo:  map[string]bool{},
		}
		// units and hit points are at start of string
		fmt.Sscanf(str, "%d units each with %d hit points", &g.units, &g.hp)

		if strings.Contains(str, "(") {
			openIndex := strings.Index(str, "(")
			closeIndex := strings.Index(str, ")")

			affinities := strings.Split(str[openIndex+1:closeIndex], "; ") // w/o parens
			for _, aff := range affinities {
				if strings.Contains(aff, "weak to ") {
					weaknesses := strings.Split(aff[len("weak to "):], ", ")
					for _, w := range weaknesses {
						g.weakTo[w] = true
					}
				}
				if strings.Contains(aff, "immune to") {
					immunities := strings.Split(aff[len("immune to "):], ", ")
					for _, imm := range immunities {
						g.immuneTo[imm] = true
					}
				}
			}
		}

		// the rest of the string is fairly uniform, so this can be generalized
		attackIndex := strings.Index(str, "attack that does ") + len("attack that does ")
		restOfString := strings.Split(str[attackIndex:], " ")
		g.attackPower, _ = strconv.Atoi(restOfString[0])
		g.attackType = restOfString[1]
		g.initiative, _ = strconv.Atoi(restOfString[5])

		groups = append(groups, &g)
	}

	return groups
}

func battle(immune, infection []*group) (immunesAfter []*group, infectionsAfter []*group, isStalemate bool) {
	// target selection, using a slice so it can be easily sorted later
	attackerToTarget := [][2]*group{}
	hasBeenTargetted := map[*group]bool{}

	// sort via decreasing EP, ties broken by highest initiative
	sort.Slice(immune, func(i, j int) bool {
		haveEqualEP := immune[i].effectivePower() == immune[j].effectivePower()
		if haveEqualEP {
			return immune[i].initiative > immune[j].initiative
		}
		return immune[i].effectivePower() > immune[j].effectivePower()
	})
	sort.Slice(infection, func(i, j int) bool {
		haveEqualEP := infection[i].effectivePower() == infection[j].effectivePower()
		if haveEqualEP {
			return infection[i].initiative > infection[j].initiative
		}
		return infection[i].effectivePower() > infection[j].effectivePower()
	})

	for _, immuneGroup := range immune {
		// target = who I'd deal the most damage to, ties broken via higher EP, then higher initiative
		var bestTarget *group
		for _, target := range infection {
			// each unit can only be attacked once
			if !hasBeenTargetted[target] {
				if bestTarget == nil {
					if immuneGroup.calcDamageToTarget(target) != 0 {
						bestTarget = target
					}
				} else {
					damageToBest := immuneGroup.calcDamageToTarget(bestTarget)
					damageToCurrent := immuneGroup.calcDamageToTarget(target)
					if damageToBest < damageToCurrent {
						bestTarget = target
					} else if damageToBest == damageToCurrent {
						// break damage tie on higher target EP first, then initiative
						epOfBest := bestTarget.effectivePower()
						epOfCurrent := target.effectivePower()
						if epOfBest < epOfCurrent {
							bestTarget = target
						} else if epOfBest == epOfCurrent && bestTarget.initiative < target.initiative {
							bestTarget = target
						}
					}
				}
			}
		}

		if bestTarget != nil {
			attackerToTarget = append(attackerToTarget, [2]*group{immuneGroup, bestTarget})
			hasBeenTargetted[bestTarget] = true
		}
	}

	// identical logic
	for _, infectGroup := range infection {
		var bestTarget *group
		for _, target := range immune {
			if !hasBeenTargetted[target] {
				if bestTarget == nil {
					if infectGroup.calcDamageToTarget(target) != 0 {
						bestTarget = target
					}
				} else {
					damageToBest := infectGroup.calcDamageToTarget(bestTarget)
					damageToCurrent := infectGroup.calcDamageToTarget(target)
					if damageToBest < damageToCurrent {
						bestTarget = target
					} else if damageToBest == damageToCurrent {
						epOfBest := bestTarget.effectivePower()
						epOfCurrent := target.effectivePower()
						if epOfBest < epOfCurrent {
							bestTarget = target
						} else if epOfBest == epOfCurrent && bestTarget.initiative < target.initiative {
							bestTarget = target
						}
					}
				}
			}
		}

		if bestTarget != nil {
			attackerToTarget = append(attackerToTarget, [2]*group{infectGroup, bestTarget})
			hasBeenTargetted[bestTarget] = true
		}
	}

	// attack phase, iterate through selections & make attacks
	// highest initiative attacks first
	sort.Slice(attackerToTarget, func(i, j int) bool {
		return attackerToTarget[i][0].initiative > attackerToTarget[j][0].initiative
	})

	isStalemate = true
	for _, attack := range attackerToTarget {
		attacker, defender := attack[0], attack[1]
		// check units != 0 before attacking, they could have been killed off by
		// a previous attack (these groups will be removed at the end)
		if attacker.units > 0 {
			targetUnitsBefore := defender.units
			defender.takeDamage(attacker.effectivePower(), attacker.attackType)
			// if units have died, then it is not a stalemate
			if defender.units != targetUnitsBefore {
				isStalemate = false
			}
		}
	}

	// remove groups that have zero units
	for i := 0; i < len(immune); {
		if immune[i].units <= 0 {
			immune[i] = immune[len(immune)-1]
			immune = immune[:len(immune)-1]
		} else {
			i++
		}
	}

	for i := 0; i < len(infection); {
		if infection[i].units <= 0 {
			infection[i] = infection[len(infection)-1]
			infection = infection[:len(infection)-1]
		} else {
			i++
		}
	}

	return immune, infection, isStalemate
}

// helper function for part 2
func runWithImmuneBoost(immuneGroup []*group, infectionGroup []*group, boost int) (immuneGroupAfter, infectionGroupAfter []*group, immuneWon bool) {
	// boost attacks of all immune groups
	for _, g := range immuneGroup {
		g.attackPower += boost
	}

	var isStalemate bool
	for !(len(immuneGroup) == 0 || len(infectionGroup) == 0) {
		immuneGroup, infectionGroup, isStalemate = battle(immuneGroup, infectionGroup)
		if isStalemate {
			break
		}
	}

	// if immune groups are all dead OR a stalemate has occurred, return false
	if len(immuneGroup) == 0 || isStalemate {
		return nil, nil, false
	}
	return immuneGroup, infectionGroup, true
}
