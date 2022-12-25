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
		ans := part1(input, 2000000)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input, 4000000)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

// x is col, y is row
// unbounded grid...
// In the row where y=2000000, how many positions canNOT contain a beacon?
//
// very naive approach of marking each coord that is visible from some sensor
// then remove all beacons that are on the target row
// then just return the length of the map containing "seen" cells on target row
//
// this is brutally slow, no way this approach works for part 2
func part1(input string, targetRow int) int {
	pairs := parseInput(input)

	blockedCoords := map[[2]int]bool{}
	for _, p := range pairs {
		manhattanDist := mathy.ManhattanDistance(p.beaconRow, p.beaconCol,
			p.sensorRow, p.sensorCol)

		// if target row is reachable, block coords on it...
		blockable := manhattanDist - mathy.AbsInt(p.sensorRow-targetRow)
		if blockable > 0 {
			for i := 0; i <= blockable; i++ {
				// add blocks to map in both left and right directions
				blockedCoords[[2]int{
					targetRow, p.sensorCol - i,
				}] = true
				blockedCoords[[2]int{
					targetRow, p.sensorCol + i,
				}] = true
			}
		}
	}

	// remove any beacons that are present in the input?
	for _, p := range pairs {
		delete(blockedCoords, [2]int{p.beaconRow, p.beaconCol})
	}

	return len(blockedCoords)
}

func part2(input string, coordLimit int) int {
	pairs := parseInput(input)

	sensors := []parsedSensor{}
	for _, p := range pairs {
		sensors = append(sensors, parsedSensor{
			sensorRow: p.sensorRow,
			sensorCol: p.sensorCol,
			manhattanDist: mathy.ManhattanDistance(p.sensorCol, p.sensorRow,
				p.beaconCol, p.beaconRow),
		})
	}

	// search space is too large to iterate over the entire thing and check if
	// SOME sensor can see that location...
	//
	// we can assume that the final resting point will be 1 cell away from the
	// border of a (actually multiple) sensor. this runs under the assumption
	// that there is only one answer
	for _, sensor := range sensors {
		distPlusOne := sensor.manhattanDist + 1

		// checking in this pattern w/ manhattan distance of 1
		//         1
		//        2 3
		//       4 S 5
		//        6B7
		//         8
		for r := -distPlusOne; r <= distPlusOne; r++ {
			targetRow := sensor.sensorRow + r

			if targetRow < 0 {
				continue
			}
			if targetRow > coordLimit {
				break
			}

			// check left and right on the target row
			// zero for first and last r's... then subtract or add it from the
			// sensor's col
			colOffset := distPlusOne - mathy.AbsInt(r)
			colLeft := sensor.sensorCol - colOffset
			colRight := sensor.sensorCol + colOffset

			if colLeft >= 0 && colLeft <= coordLimit &&
				!isReachable(sensors, colLeft, targetRow) {
				return colLeft*4000000 + targetRow
			}
			if colRight >= 0 && colRight <= coordLimit &&
				!isReachable(sensors, colRight, targetRow) {
				return colRight*4000000 + targetRow
			}
		}
	}
	panic("unreachable")
}

type pair struct {
	sensorRow, sensorCol int
	beaconRow, beaconCol int
}

func parseInput(input string) (ans []pair) {
	// Sensor at x=2150774, y=3136587: closest beacon is at x=2561642, y=2914773
	for _, line := range strings.Split(input, "\n") {
		p := pair{}
		_, err := fmt.Sscanf(line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&p.sensorCol, &p.sensorRow, &p.beaconCol, &p.beaconRow)
		if err != nil {
			panic("parsing: " + err.Error())
		}
		ans = append(ans, p)
	}
	return ans
}

type parsedSensor struct {
	sensorRow, sensorCol int
	manhattanDist        int
}

func isReachable(sensors []parsedSensor, c, r int) bool {
	for _, sensor := range sensors {
		// if reachable, break
		if sensor.manhattanDist >= mathy.ManhattanDistance(c, r,
			sensor.sensorCol, sensor.sensorRow) {
			return true
		}

	}
	return false
}
