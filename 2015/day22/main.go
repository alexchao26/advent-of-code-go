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

	ans := wizardSimulator(util.ReadFile("./input.txt"), 50, 500, part)
	fmt.Println("Output:", ans)
}

func wizardSimulator(input string, myHP, myMana, part int) int {
	lines := strings.Split(input, "\n")
	bossHP := cast.ToInt(strings.Split(lines[0], ": ")[1])
	bossDamage := cast.ToInt(strings.Split(lines[1], ": ")[1])

	initState := newBattleState(myHP, myMana, bossHP, bossDamage, [5]int{}, true, 0)

	return simBattle(initState, map[string]int{}, part)
}

// Spell struct is used to generalize all spell types by leveraging zero values.
// The zero value for ints is 0 (which can be added with no effect)
type spell struct {
	name          string // redundant, for debugging
	index         int    // for indexing in an array (which is easily passed by value)
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

type battleState struct {
	myHP            int
	myMana          int
	bossHP          int
	bossDamage      int
	effectDurations [5]int
	isMyTurn        bool
	depth           int // recursive branch depth, for debugging
}

func newBattleState(myHP, myMana, bossHP, bossDamage int, effectDurations [5]int, isMyTurn bool, depth int) battleState {
	return battleState{
		myHP:            myHP,
		myMana:          myMana,
		bossHP:          bossHP,
		bossDamage:      bossDamage,
		effectDurations: effectDurations,
		isMyTurn:        isMyTurn,
		depth:           depth,
	}
}

func (s battleState) hashKey() string {
	return fmt.Sprintf("%d_%d_%d_%v_%v", s.myHP, s.myMana, s.bossHP, s.effectDurations, s.isMyTurn)
}

func simBattle(state battleState, memo map[string]int, part int) (minMana int) {
	// check cache
	hash := state.hashKey()
	if val, ok := memo[hash]; ok {
		return val
	}

	if part == 2 && state.isMyTurn {
		state.myHP--
	}

	// check myHP after a potential part 2 HP loss, if player dies, then return
	// a huge number which will essentially be ignored by a mathy.MinInt comparison
	if state.myHP <= 0 {
		return math.MaxInt32
	}

	// apply any active spell effects
	var myArmor int
	for _, sp := range spellsMap {
		if state.effectDurations[sp.index] > 0 {
			state.effectDurations[sp.index]--
			// many of values will be zero for any given spell
			state.bossHP -= sp.effectDamage
			state.myHP += sp.heal
			myArmor += sp.armorBuff
			state.myMana += sp.manaRecharge
		}
	}

	// check bossHP after effects take place, it could die form poison
	if state.bossHP <= 0 {
		return 0
	}

	// get minMana from the current state, to a player win
	minMana = math.MaxInt32
	if state.isMyTurn {
		// iterate through spells, create a recursive call for each spell that
		// can be called (i.e. its effectDuration index is zero)
		var spellCasted bool
		for _, sp := range spellsMap {
			if state.effectDurations[sp.index] == 0 {
				if state.myMana >= sp.cost {
					spellCasted = true
					// make new durations array & add effect duration for this spell
					newDurations := state.effectDurations
					newDurations[sp.index] += sp.effectLength

					nextState := newBattleState(state.myHP+sp.instantHeal,
						state.myMana-sp.cost,
						state.bossHP-sp.instantDamage,
						state.bossDamage,
						newDurations,
						false,
						state.depth+1,
					)

					castResult := sp.cost + simBattle(nextState, memo, part)

					minMana = mathy.MinInt(minMana, castResult)
				}
			}
		}
		// if cannot cast spell, player loses
		if !spellCasted {
			return math.MaxInt32
		}
	} else {
		// boss's turn, boss attacks w/ a minimum damage of 1
		attackDamage := mathy.MaxInt(1, state.bossDamage-myArmor)

		// recurse w/ next state
		nextState := newBattleState(state.myHP-attackDamage,
			state.myMana,
			state.bossHP,
			state.bossDamage,
			state.effectDurations,
			true,
			state.depth+1,
		)
		bossAttackResult := simBattle(nextState, memo, part)

		minMana = mathy.MinInt(minMana, bossAttackResult)
	}

	// add to memoized to prevent unnecessary recursive branches, then return
	memo[hash] = minMana
	return minMana
}
