package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	ans1, ans2 := rpgSimulator(util.ReadFile("./input.txt"))
	fmt.Printf("Part1: %d\nPart2: %d\n", ans1, ans2)
}

func rpgSimulator(input string) (rpgSimulator, part2 int) {
	bossHP, bossDamage, bossArmor, shopWeapons, shopArmor, shopRings := parseInput(input)

	// all attacks do at least 1 damage
	// damage equal to attack - defenders armor
	// armor is optional, limit 1
	// 0-2 rings allowed
	// player attacks first
	// must buy exactly 1 weapon
	var combinations [][]item
	for weapon := 0; weapon < len(shopWeapons); weapon++ {
		for armor := -1; armor < len(shopArmor); armor++ {
			for ring1 := -1; ring1 < len(shopRings); ring1++ {
				for ring2 := -1; ring2 < len(shopRings); ring2++ {
					comb := []item{shopWeapons[weapon]}
					if armor != -1 {
						comb = append(comb, shopArmor[armor])
					}
					if ring1 != -1 {
						comb = append(comb, shopRings[ring1])
					}
					if ring2 != -1 && ring2 != ring1 {
						comb = append(comb, shopRings[ring2])

					}

					combinations = append(combinations, comb)
				}
			}
		}
	}

	minCost := math.MaxInt32
	var maxCost int
	for _, comb := range combinations {
		myHP := 100
		var myDamage, myArmor, cost int
		for _, it := range comb {
			myDamage += it.damage
			myArmor += it.armor
			cost += it.cost
		}
		playerWins := simulateBattle(bossHP, bossDamage, bossArmor, myHP, myDamage, myArmor)
		if playerWins {
			// part 1, min cost to win
			minCost = mathy.MinInt(minCost, cost)
		} else {
			// part 2, max cost to still lose
			maxCost = mathy.MaxInt(maxCost, cost)
		}
	}

	return minCost, maxCost
}

func simulateBattle(bossHP, bossDamage, bossArmor, myHP, myDamage, myArmor int) (playerWins bool) {
	attackOnBoss := myDamage - bossArmor
	attackOnPlayer := bossDamage - myArmor
	attackOnBoss = mathy.MaxInt(attackOnBoss, 1)
	attackOnPlayer = mathy.MaxInt(attackOnPlayer, 1)
	for bossHP > 0 && myHP > 0 {
		bossHP -= attackOnBoss
		myHP -= attackOnPlayer
	}

	// the boss takes damage first, so if it hit zero or less, then the player
	// won the round (potentially with very little HP left)
	return bossHP <= 0
}

type item struct {
	name                string
	cost, damage, armor int
}

var shop = `Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`

func parseInput(input string) (hp, damage, armor int, shopWeapons, shopArmor, shopRings []item) {
	lines := strings.Split(input, "\n")
	hp = cast.ToInt(strings.Split(lines[0], ": ")[1])
	damage = cast.ToInt(strings.Split(lines[1], ": ")[1])
	armor = cast.ToInt(strings.Split(lines[2], ": ")[1])

	shopBlocks := strings.Split(shop, "\n\n")
	for blockIndex := range shopBlocks {
		for _, line := range strings.Split(shopBlocks[blockIndex], "\n")[1:] {
			it := item{}
			if blockIndex == 2 {
				// gross, have to get rid of whitespace in name for the rings
				line = strings.ReplaceAll(line, "e +", "e+")
			}
			_, err := fmt.Sscanf(line, "%s %d %d %d", &it.name, &it.cost, &it.damage, &it.armor)
			if err != nil {
				panic(err)
			}
			switch blockIndex {
			case 0:
				shopWeapons = append(shopWeapons, it)
			case 1:
				shopArmor = append(shopArmor, it)
			case 2:
				shopRings = append(shopRings, it)
			}
		}
	}

	return hp, damage, armor, shopWeapons, shopArmor, shopRings
}
