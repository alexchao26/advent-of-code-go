package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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
		ans := part1(input, 1000)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string, iters int) int {
	boxes := parseInput(input)

	distances := map[*box]map[*box]int{}

	type edge struct {
		b1, b2   *box
		distance int
	}

	allEdges := []edge{}
	for i, b1 := range boxes {
		distances[b1] = map[*box]int{}
		for _, b2 := range boxes[i+1:] {
			allEdges = append(allEdges, edge{
				b1:       b1,
				b2:       b2,
				distance: calcRoughDistance(b1, b2),
			})
		}
	}

	sort.Slice(allEdges, func(i int, j int) bool {
		return allEdges[i].distance < allEdges[j].distance
	})

	circuits := []*circuit{}

	for range iters {
		closestEdge := allEdges[0]
		allEdges = allEdges[1:]

		if closestEdge.b1.circuit == nil && closestEdge.b2.circuit == nil {
			circuits = append(circuits, &circuit{
				boxes: []*box{
					closestEdge.b1,
					closestEdge.b2,
				},
			})
			closestEdge.b1.circuit = circuits[len(circuits)-1]
			closestEdge.b2.circuit = circuits[len(circuits)-1]
		} else if closestEdge.b1.circuit == nil {
			closestEdge.b2.circuit.boxes = append(closestEdge.b2.circuit.boxes, closestEdge.b1)
			closestEdge.b1.circuit = closestEdge.b2.circuit
		} else if closestEdge.b2.circuit == nil {
			closestEdge.b1.circuit.boxes = append(closestEdge.b1.circuit.boxes, closestEdge.b2)
			closestEdge.b2.circuit = closestEdge.b1.circuit
		} else if closestEdge.b1.circuit != closestEdge.b2.circuit {
			// combine two circuits
			oldCircuit := closestEdge.b2.circuit
			for _, box := range closestEdge.b2.circuit.boxes {
				closestEdge.b1.circuit.boxes = append(closestEdge.b1.circuit.boxes, box)
				box.circuit = closestEdge.b1.circuit
			}

			// remove oldCircuit
			for i := range circuits {
				if circuits[i] == oldCircuit {
					circuits[len(circuits)-1], circuits[i] = circuits[i], circuits[len(circuits)-1]
					circuits = circuits[:len(circuits)-1]
					break
				}
			}
		}
	}

	circuitSizes := []int{}
	for _, c := range circuits {
		circuitSizes = append(circuitSizes, len(c.boxes))
	}

	sort.Ints(circuitSizes)

	ans := 1
	for i := range 3 {
		ans *= circuitSizes[len(circuitSizes)-1-i]
	}

	return ans
}

func part2(input string) int {
	boxes := parseInput(input)

	distances := map[*box]map[*box]int{}

	type edge struct {
		b1, b2   *box
		distance int
	}

	allEdges := []edge{}
	for i, b1 := range boxes {
		distances[b1] = map[*box]int{}
		for _, b2 := range boxes[i+1:] {
			allEdges = append(allEdges, edge{
				b1:       b1,
				b2:       b2,
				distance: calcRoughDistance(b1, b2),
			})
		}
	}

	sort.Slice(allEdges, func(i int, j int) bool {
		return allEdges[i].distance < allEdges[j].distance
	})

	circuits := []*circuit{}

	for len(allEdges) > 0 {
		closestEdge := allEdges[0]
		allEdges = allEdges[1:]

		if closestEdge.b1.circuit == nil && closestEdge.b2.circuit == nil {
			circuits = append(circuits, &circuit{
				boxes: []*box{
					closestEdge.b1,
					closestEdge.b2,
				},
			})
			closestEdge.b1.circuit = circuits[len(circuits)-1]
			closestEdge.b2.circuit = circuits[len(circuits)-1]
		} else if closestEdge.b1.circuit == nil {
			closestEdge.b2.circuit.boxes = append(closestEdge.b2.circuit.boxes, closestEdge.b1)
			closestEdge.b1.circuit = closestEdge.b2.circuit
		} else if closestEdge.b2.circuit == nil {
			closestEdge.b1.circuit.boxes = append(closestEdge.b1.circuit.boxes, closestEdge.b2)
			closestEdge.b2.circuit = closestEdge.b1.circuit
		} else if closestEdge.b1.circuit != closestEdge.b2.circuit {
			// combine two circuits
			oldCircuit := closestEdge.b2.circuit
			for _, box := range closestEdge.b2.circuit.boxes {
				closestEdge.b1.circuit.boxes = append(closestEdge.b1.circuit.boxes, box)
				box.circuit = closestEdge.b1.circuit
			}

			// remove oldCircuit
			for i := range circuits {
				if circuits[i] == oldCircuit {
					circuits[len(circuits)-1], circuits[i] = circuits[i], circuits[len(circuits)-1]
					circuits = circuits[:len(circuits)-1]
					break
				}
			}
		}

		// if this is the join that makes the entire circuit the size of all boxes, then break out
		// and return final calc.
		// note that this works in the "no join added" as well because this returns true on the first
		// instance where circuit length is equal to length of all boxes
		if len(closestEdge.b1.circuit.boxes) == len(boxes) {
			return closestEdge.b1.x * closestEdge.b2.x
		}
	}

	panic("should return from loop")
}

type box struct {
	x, y, z int
	circuit *circuit
}

func (b *box) String() string {
	return fmt.Sprintf("x: %d, y: %d, z: %d, circuit %+v", b.x, b.y, b.z, b.circuit)
}

// equivalent to distance
func calcRoughDistance(b1, b2 *box) int {
	x := b1.x - b2.x
	y := b1.y - b2.y
	z := b1.z - b2.z
	return x*x + y*y + z*z
}

type circuit struct {
	boxes []*box
}

func parseInput(input string) (boxes []*box) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ",")
		boxes = append(boxes, &box{
			x: cast.ToInt(parts[0]),
			y: cast.ToInt(parts[1]),
			z: cast.ToInt(parts[2]),
		})
	}
	return boxes
}
