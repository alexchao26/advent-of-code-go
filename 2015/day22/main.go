package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := part1(util.ReadFile("./input.txt"), 50, 500, part)
	fmt.Println("Output:", ans)
}

func part1(input string, myHP, myMana, part int) int {
	bossHP, bossDamage := parseInput(input)
	_, _, _, _ = bossHP, bossDamage, myHP, myMana
	return backtrackBattleSim(myHP, myMana, bossHP, bossDamage, [5]int{}, true, 0, map[string]int{}, part)
}

type spell struct {
	name          string
	index         int
	cost          int
	effectLength  int
	instantDamage int
	instantHeal   int
	effectDamage  int
	heal          int
	armorBuff     int
	manaRecharge  int
}

var spellsMap = map[string]spell{
	"Magic Missile": {
		name:          "Magic Missile",
		index:         0,
		cost:          53,
		instantDamage: 4,
	},
	"Drain": {
		name:          "Drain",
		index:         1,
		cost:          73,
		instantDamage: 2,
		instantHeal:   2,
	},
	"Shield": {
		name:         "Shield",
		index:        2,
		cost:         113,
		effectLength: 6,
		armorBuff:    7, // does not stack for each turn
	},
	"Poison": {
		name:         "Poison",
		index:        3,
		cost:         173,
		effectLength: 6,
		effectDamage: 3,
	},
	"Recharge": {
		name:         "Recharge",
		index:        4,
		cost:         229,
		effectLength: 5,
		manaRecharge: 101,
	},
}

func hashState(myHP, myMana, bossHP int, effectDurations [5]int, isMyTurn bool) string {
	return fmt.Sprintf("%d_%d_%d_%v_%v", myHP, myMana, bossHP, effectDurations, isMyTurn)
}

func backtrackBattleSim(myHP, myMana, bossHP, bossDamage int, effectDurations [5]int, isMyTurn bool, depth int, memo map[string]int, part int) (minMana int) {
	// fmt.Printf("\nDEPTH: %d %v\nMe %d Mana: %d; Boss %d Dmg: %d\n  DURATIONS: %v\n", depth, isMyTurn, myHP, myMana, bossHP, bossDamage, effectDurations)

	hash := hashState(myHP, myMana, bossHP, effectDurations, isMyTurn)
	if val, ok := memo[hash]; ok {
		return val
	}

	if part == 2 && isMyTurn {
		myHP--
	}
	if myHP <= 0 {
		// fmt.Println("BOSS WINS")
		return math.MaxInt32
	}

	// apply any active spell effects
	var myArmor int
	for _, sp := range spellsMap {
		if effectDurations[sp.index] > 0 {
			effectDurations[sp.index]--
			// many of values will be zero for any given spell
			bossHP -= sp.effectDamage
			myHP += sp.heal
			myArmor += sp.armorBuff
			myMana += sp.manaRecharge
		}
	}

	// fmt.Printf("  Post Effects: Me %d Mana: %d Def: %d; Boss %d\n", myHP, myMana, myArmor, bossHP)

	if bossHP <= 0 {
		// fmt.Println("PLAYER WINS")
		return 0
	}

	minMana = math.MaxInt32
	if isMyTurn {
		// iterate through spells, create a recursive call for each spell that
		// can be called (i.e. its effectDuration index is zero)
		var spellCasted bool
		for _, spName := range []string{"Recharge", "Poison", "Shield", "Drain", "Magic Missile"} {
			sp := spellsMap[spName]
			if effectDurations[sp.index] == 0 {
				if myMana >= sp.cost {
					spellCasted = true
					// make new durations array & add effect duration for this spell
					newDurations := effectDurations
					newDurations[sp.index] += sp.effectLength
					// fmt.Printf(" dp %d, casting %s\n", depth, sp.name)
					castResult := sp.cost + backtrackBattleSim(myHP+sp.instantHeal, myMana-sp.cost, bossHP-sp.instantDamage, bossDamage, newDurations, false, depth+1, memo, part)
					// fmt.Printf("  result from casting %s @ depth %d: %d\n", sp.name, depth, castResult)

					minMana = mathy.MinInt(minMana, castResult)
				}
				// } else {
				// 	fmt.Println("    CANNOT cast", sp.name)
			}
		}
		// if cannot cast spell, lose
		if !spellCasted {
			// fmt.Println("  CANNOT CAST SPELL, LOST")
			return math.MaxInt32
		}
	} else {
		// boss attacks, minimum 1 damage
		attackDamage := mathy.MaxInt(1, bossDamage-myArmor)
		// recurse
		bossAttackResult := backtrackBattleSim(myHP-attackDamage, myMana, bossHP, bossDamage, effectDurations, true, depth+1, memo, part)

		minMana = mathy.MinInt(minMana, bossAttackResult)
	}

	// fmt.Println("  from", depth, minMana)
	memo[hash] = minMana
	return minMana
}

func parseInput(input string) (bossHP, bossDamage int) {
	lines := strings.Split(input, "\n")
	bossHP = cast.ToInt(strings.Split(lines[0], ": ")[1])
	bossDamage = cast.ToInt(strings.Split(lines[1], ": ")[1])
	return bossHP, bossDamage
}
