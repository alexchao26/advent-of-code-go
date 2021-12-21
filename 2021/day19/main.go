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

	ans1, ans2 := part1(input)
	if part == 1 {
		util.CopyToClipboard(fmt.Sprintf("%v", ans1))
		fmt.Println("Output:", ans1)
	} else {
		util.CopyToClipboard(fmt.Sprintf("%v", ans2))
		fmt.Println("Output:", ans2)
	}
}

func part1(input string) (part1, part2 int) {
	scanners := parseInput(input)

	// determine scanner locations by finding scanners that see 12 of the same beacons
	// everything will be relative from scanner 0, so it is "settled" and its abs & rel coords are the same
	settled := []scanner{scanners[0]}
	settled[0].absoluteCoords = settled[0].relativeCoords
	settled[0].fillAbsoluteCoordsMap()

	// create helper functions that can create all the rotated versions of the seen beacons
	// scanner 0 will have KNOWN coordinates (0,0,0)
	// maintain a list of all UNKNOWN scanners (all but 0 at the start)
	undetermined := scanners[1:]
	// iterate while it has a non zero length
	for len(undetermined) > 0 {
		fmt.Printf("progress: %d/%d\n", len(settled), len(scanners))

		for i, undet := range undetermined {
			maybeUpdated, ok := findAbsoluteCoordsForScanner(undet, settled)
			if ok {
				settled = append(settled, maybeUpdated)
				// remove the determined scanner from undetermined list
				copy(undetermined[i:], undetermined[i+1:])
				// undetermined[i] = undetermined[len(undetermined)-1]
				undetermined = undetermined[:len(undetermined)-1]
				// restart checks from start of undetermined list
				break
			}
		}
	}

	allBeacons := map[[3]int]bool{}
	for _, s := range settled {
		for c := range s.absoluteCoordsMap {
			allBeacons[c] = true
		}
	}

	var furthest int
	for i, s1 := range settled {
		for j, s2 := range settled {
			if i == j {
				continue
			}
			manhattanDist := mathy.AbsInt(s1.x-s2.x) + mathy.AbsInt(s1.y-s2.y) + mathy.AbsInt(s1.z-s2.z)
			if manhattanDist > furthest {
				furthest = manhattanDist
			}
		}
	}
	return len(allBeacons), furthest
}

type scanner struct {
	number            int
	x, y, z           int
	relativeCoords    [][3]int
	rotations         [][][3]int
	absoluteCoords    [][3]int
	absoluteCoordsMap map[[3]int]bool
}

func (s *scanner) fillAbsoluteCoordsMap() {
	s.absoluteCoordsMap = map[[3]int]bool{}
	if len(s.absoluteCoords) == 0 {
		panic(fmt.Sprintf("absolute coords not set for scanner %d", s.number))
	}
	for _, ac := range s.absoluteCoords {
		s.absoluteCoordsMap[ac] = true
	}
}

// create the 24 rotations given a slice of coords (3-length arrays)
func (s *scanner) fillRotations() {
	// facing negative x
	posX := s.relativeCoords
	var dir2, dir3, dir4, dir5, dir6 [][3]int
	for _, c := range posX {
		x, y, z := c[0], c[1], c[2]
		dir2 = append(dir2, [3]int{x, -y, -z})
		dir3 = append(dir3, [3]int{x, -z, y})
		dir4 = append(dir4, [3]int{-y, -z, x})
		dir5 = append(dir5, [3]int{-x, -z, -y})
		dir6 = append(dir6, [3]int{y, -z, -x})
	}
	sixRotations := [][][3]int{
		posX, dir2,
		dir3, dir4,
		dir5, dir6,
	}

	// apply 4 rotations around the axis that the scanner is "staring down"
	var finalRotations [][][3]int
	for _, rotation := range sixRotations {
		var r2, r3, r4 [][3]int // r1 is rotation itself
		for _, c := range rotation {
			x, y, z := c[0], c[1], c[2]
			r2 = append(r2, [3]int{-y, x, z})
			r3 = append(r3, [3]int{-x, -y, z})
			r4 = append(r4, [3]int{y, -x, z})
		}
		finalRotations = append(finalRotations, rotation, r2, r3, r4)
	}
	s.rotations = finalRotations
}

func findAbsoluteCoordsForScanner(undet scanner, settled []scanner) (maybeUpdated scanner, didUpdate bool) {
	// for all orientations of the unknown's beacon list
	for _, rotatedCoords := range undet.rotations {
		// for each beacon in known list
		for _, set := range settled {
			for _, absCoord := range set.absoluteCoords {
				// for each beacon in unknown list
				for _, relativeCoord := range rotatedCoords {
					// assume the known and unknown beacon are the same, calculate the absolute coords of the unknown's scanner coords
					// convert all of unknown list to their absolute coords, check against known list
					unsettledAbsoluteCoords := makeAbsoluteCoordsList(absCoord, relativeCoord, rotatedCoords)

					var matchingCount int
					// var matched [][3]int // !
					for _, ac := range unsettledAbsoluteCoords {
						if set.absoluteCoordsMap[ac] {
							// matched = append(matched, ac) // !
							matchingCount++
						}
					}

					// if true return a true or something, modify the scanner param pointer
					if matchingCount >= 12 {
						undet.relativeCoords = rotatedCoords
						undet.absoluteCoords = unsettledAbsoluteCoords
						undet.fillAbsoluteCoordsMap()
						undet.x = absCoord[0] - relativeCoord[0]
						undet.y = absCoord[1] - relativeCoord[1]
						undet.z = absCoord[2] - relativeCoord[2]
						return undet, true
					}
				}
			}
		}
	}

	// did not update
	return undet, false
}

func makeAbsoluteCoordsList(absolute, relative [3]int, relativeCoords [][3]int) [][3]int {
	// assuming absolute and relative are pointing to the same coord
	// generate the all of the abolute coords

	// diff to the scanner's coordinates, then calculate all beacons from scanner's coords
	diff := [3]int{
		absolute[0] - relative[0],
		absolute[1] - relative[1],
		absolute[2] - relative[2],
	}

	var absCoords [][3]int
	for _, c := range relativeCoords {
		absCoords = append(absCoords, [3]int{
			diff[0] + c[0],
			diff[1] + c[1],
			diff[2] + c[2],
		})
	}

	return absCoords
}

func parseInput(input string) (ans []scanner) {
	for _, rawScanner := range strings.Split(input, "\n\n") {
		var number int
		lines := strings.Split(rawScanner, "\n")
		_, err := fmt.Sscanf(lines[0], "--- scanner %d ---", &number)
		if err != nil {
			panic("parsing error " + err.Error())
		}

		var coords [][3]int
		for _, line := range lines[1:] {
			var x, y, z int
			_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
			if err != nil {
				panic("parsing error " + err.Error())
			}
			coords = append(coords, [3]int{x, y, z})
		}

		sc := scanner{
			number:            number,
			x:                 0,
			y:                 0,
			z:                 0,
			relativeCoords:    coords,
			absoluteCoords:    nil,
			absoluteCoordsMap: map[[3]int]bool{},
		}
		sc.fillRotations()
		ans = append(ans, sc)
	}

	return ans
}

// var exampleBeaconsList [][3]int = func() [][3]int {
// 	var coords [][3]int
// 	for _, l := range strings.Split(rawBeaconsList, "\n") {
// 		var x, y, z int
// 		fmt.Sscanf(l, "%d,%d,%d", &x, &y, &z)
// 		coords = append(coords, [3]int{x, y, z})
// 	}
// 	return coords
// }()

// var rawBeaconsList = `-892,524,684
// -876,649,763
// -838,591,734
// -789,900,-551
// -739,-1745,668
// -706,-3180,-659
// -697,-3072,-689
// -689,845,-530
// -687,-1600,576
// -661,-816,-575
// -654,-3158,-753
// -635,-1737,486
// -631,-672,1502
// -624,-1620,1868
// -620,-3212,371
// -618,-824,-621
// -612,-1695,1788
// -601,-1648,-643
// -584,868,-557
// -537,-823,-458
// -532,-1715,1894
// -518,-1681,-600
// -499,-1607,-770
// -485,-357,347
// -470,-3283,303
// -456,-621,1527
// -447,-329,318
// -430,-3130,366
// -413,-627,1469
// -345,-311,381
// -36,-1284,1171
// -27,-1108,-65
// 7,-33,-71
// 12,-2351,-103
// 26,-1119,1091
// 346,-2985,342
// 366,-3059,397
// 377,-2827,367
// 390,-675,-793
// 396,-1931,-563
// 404,-588,-901
// 408,-1815,803
// 423,-701,434
// 432,-2009,850
// 443,580,662
// 455,729,728
// 456,-540,1869
// 459,-707,401
// 465,-695,1988
// 474,580,667
// 496,-1584,1900
// 497,-1838,-617
// 527,-524,1933
// 528,-643,409
// 534,-1912,768
// 544,-627,-890
// 553,345,-567
// 564,392,-477
// 568,-2007,-577
// 605,-1665,1952
// 612,-1593,1893
// 630,319,-379
// 686,-3108,-505
// 776,-3184,-501
// 846,-3110,-434
// 1135,-1161,1235
// 1243,-1093,1063
// 1660,-552,429
// 1693,-557,386
// 1735,-437,1738
// 1749,-1800,1813
// 1772,-405,1572
// 1776,-675,371
// 1779,-442,1789
// 1780,-1548,337
// 1786,-1538,337
// 1847,-1591,415
// 1889,-1729,1762
// 1994,-1805,1792`
