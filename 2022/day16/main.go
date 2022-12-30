package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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

/*
PART 1 and PART 2 ARE IMPLEMENTED COMPLETELY SEPARATELY.
If you're looking for part 2 inspiration you might want to not even look at my part1 code...
*/

func part1(input string) int {
	// 30 min until cave erupts
	// start in room AA
	// creates graph to other rooms...
	// flow rate is zero to start
	// could spend 1 minute moving to BB and 1 more minute opening it. will release pressure of during remaining 28 of 30 minutes. 13 flow * 28 minutes = 364 total pressure lol
	// then can move to CC in min 3, open in min 4, etc etc
	// release max pressure in 30 min?
	graph := makeGraph(input)

	open := map[string]bool{"AA": true}
	for name, rm := range graph {
		if rm.flowRate == 0 {
			open[name] = true
		}
	}

	// "AA" open will not matter really...
	return bfs(graph, "AA", 30, 0, open, map[string]int{})
}

type room struct {
	name        string
	flowRate    int
	connectedTo []string
}

func makeGraph(input string) map[string]room {
	graph := map[string]room{}

	for _, line := range strings.Split(input, "\n") {
		// Valve BB has flow rate=13; tunnels lead to valves CC, AA
		parts := strings.Split(line, "; ")
		rm := room{}
		_, err := fmt.Sscanf(parts[0], "Valve %s has flow rate=%d", &rm.name, &rm.flowRate)
		if err != nil {
			panic("parsing valve name and flow rate" + err.Error())
		}
		connections := strings.Split(parts[1], ", ")
		// update first entry to remove leading string
		connections[0] = connections[0][len(connections[0])-2:]
		rm.connectedTo = connections
		graph[rm.name] = rm
	}

	return graph
}

func bfs(graph map[string]room, currentRoom string, minutesLeft, currentPressure int, open map[string]bool, memo map[string]int) int {
	if minutesLeft == 0 {
		return 0
	}

	key := hash(currentRoom, minutesLeft, open, currentPressure)
	if v, ok := memo[key]; ok {
		return v
	}

	// recursive calls will update this if it is better, then return it
	bestFlow := 0

	// there are two paths to take at a room
	// 1. stay and open the valve
	//    this is only worth doing if the valve is not already on
	// 2. move to a neighboring

	// 1. open current room's valve
	if !open[currentRoom] {
		open[currentRoom] = true
		// totalPressureContribution := (minutesLeft - 1) * graph[currentRoom].flowRate

		newPressure := currentPressure + graph[currentRoom].flowRate

		maybeBest := currentPressure + bfs(graph, currentRoom, minutesLeft-1, newPressure, open, memo)

		bestFlow = mathy.MaxInt(bestFlow, maybeBest)

		// backtrack
		open[currentRoom] = false
	}

	// 2. move to neighbors
	for _, neighbor := range graph[currentRoom].connectedTo {
		maybeBest := currentPressure + bfs(graph, neighbor, minutesLeft-1, currentPressure, open, memo)
		bestFlow = mathy.MaxInt(bestFlow, maybeBest)
	}

	memo[key] = bestFlow

	return bestFlow
}

func hash(currentRoom string, minutesLeft int, open map[string]bool, currentPressure int) string {
	rms := []string{}
	for k := range open {
		rms = append(rms, k)
	}
	sort.Strings(rms)
	return fmt.Sprint(currentRoom, minutesLeft, rms, currentPressure)
}

// PART 2, basically restarting because my part1 seems too far in the opposite direction to reuse

func part2(input string) int {
	// index within [16]int arrays which will be used to track which rooms have been visited
	//   from analysis i know that my input has 15 non-zero pressure rooms
	// every time we visit a room we'll open the valve

	graph := makeGraph(input)

	roomToFlowRate := map[string]int{}
	highestFlowRatePossible := 0 // might be useful for an optimization later

	// sort room names so the arrays/slices later will be in a repeatable order (easier debugging)
	// include starting room "AA" in this list even though it has zero flow rate
	roomNames := []string{"AA"}
	for name, room := range graph {
		if room.flowRate != 0 {
			roomNames = append(roomNames, name)
		}
	}

	// reformat into arrays/slices so we don't have to do room name lookups
	flowRates := make([]int, len(roomNames))

	sort.Strings(roomNames)

	for i, name := range roomNames {
		roomToFlowRate[name] = graph[name].flowRate
		flowRates[i] = roomToFlowRate[name]
		highestFlowRatePossible += graph[name].flowRate
	}

	weightedGraph := makeWeightedGraph(graph, roomNames)

	visitedArrayToHighestPressureTotals := map[[16]bool]int{}
	dfsGetHighestPressureTotalsForEveryVisitedState(26, [16]bool{}, weightedGraph, flowRates, 0, 0, 0, visitedArrayToHighestPressureTotals)

	return highestDisjointPair(visitedArrayToHighestPressureTotals)
}

func makeWeightedGraph(graph map[string]room, roomNames []string) [][]int {
	ans := make([][]int, len(roomNames))
	for i := range ans {
		ans[i] = make([]int, len(roomNames))
	}

	// bfs between every node to make graph

	for startIndex, startName := range roomNames {
		for endIndex, endName := range roomNames {
			if startName == endName {
				continue
			}
			stepsBetweenRooms := bfsDistanceBetweenRooms(graph, startName, endName)
			ans[startIndex][endIndex] = stepsBetweenRooms
			ans[endIndex][startIndex] = stepsBetweenRooms
		}
	}

	return ans
}

func bfsDistanceBetweenRooms(graph map[string]room, startName, endName string) int {
	type node struct {
		name  string
		steps int
	}

	queue := []node{
		{
			name:  startName,
			steps: 0,
		},
	}

	seen := map[string]bool{}
	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]
		if seen[pop.name] {
			continue
		}
		seen[pop.name] = true

		if pop.name == endName {
			return pop.steps
		}
		for _, neighbor := range graph[pop.name].connectedTo {
			queue = append(queue, node{
				name:  neighbor,
				steps: pop.steps + 1,
			})
		}
	}
	// assume all rooms are reachable
	panic("should return from loop")
}

// populates ansArray with the best possible values for visiting a particular set of rooms
func dfsGetHighestPressureTotalsForEveryVisitedState(timeLeft int, visited [16]bool,
	graph [][]int, flowRates []int, currentRoom int, flowRate, totalPressure int,
	ansArray map[[16]bool]int) {
	if timeLeft < 0 {
		panic("negative timeLeft")
	}
	if timeLeft == 0 {
		ansArray[visited] = mathy.MaxInt(ansArray[visited], totalPressure)
		return
	}

	// branch 1: just not moving at all
	dfsGetHighestPressureTotalsForEveryVisitedState(0, visited, graph, flowRates, currentRoom,
		flowRate, totalPressure+flowRate*timeLeft, ansArray)

	// rest of branches: attempt to visit every possible non-visited node
	for roomIndex := range graph {
		hasBeenVisited := visited[roomIndex]
		if hasBeenVisited {
			continue
		}

		// get to room, one more to open valve
		timeToOpenNextValve := graph[currentRoom][roomIndex] + 1
		// not worth visiting if the valve can't be opened in time
		if timeLeft < timeToOpenNextValve {
			continue
		}

		// in Go this makes a full copy of the array, so &nextVisited != &visited
		// this is NOT true for slices ([]bool)
		nextVisited := visited
		nextVisited[roomIndex] = true
		dfsGetHighestPressureTotalsForEveryVisitedState(
			timeLeft-timeToOpenNextValve,
			nextVisited,
			graph,
			flowRates,
			roomIndex,
			flowRate+flowRates[roomIndex],
			totalPressure+flowRate*timeToOpenNextValve,
			ansArray)
	}

}

func highestDisjointPair(visitedArrayToHighestPressureTotals map[[16]bool]int) int {
	type finishingState struct {
		bitmap int
		// leaving this here to explain a simpler solution (without bitmap overkill)
		// visited       [16]bool
		totalPressure int
	}

	allFinishingStates := []finishingState{}
	for visited, totalPressure := range visitedArrayToHighestPressureTotals {
		allFinishingStates = append(allFinishingStates, finishingState{
			bitmap:        convertToBitmap(visited),
			totalPressure: totalPressure,
		})
	}

	// sort in decreasing order
	sort.Slice(allFinishingStates, func(i, j int) bool {
		return allFinishingStates[i].totalPressure > allFinishingStates[j].totalPressure
	})

	bestCombo := -1

	for _, baseFinishingState := range allFinishingStates {
		for _, maybeDisjointFinishingState := range allFinishingStates {
			pressureSum := baseFinishingState.totalPressure + maybeDisjointFinishingState.totalPressure
			// allFinishingStates is sorted in decreasing order so at this point there is no reason
			// to continue checking sums
			if pressureSum < bestCombo {
				break
			}

			// only update baseFinishing state if sets are disjointed (human and elephant can't
			// open the same room's valve)
			//
			// using bit logic is overkill, this could've been replaced with this single for loop:
			// isDisjoint := true
			// for i, wasVisited := range baseFinishingState.visited {
			// 	if wasVisited && maybeDisjointFinishingState.visited[i] {
			// 		isDisjoint = false
			// 		break
			// 	}
			// }
			if baseFinishingState.bitmap&maybeDisjointFinishingState.bitmap == 0 {
				bestCombo = mathy.MaxInt(bestCombo, pressureSum)
			}

		}
	}

	return bestCombo
}

func convertToBitmap(visited [16]bool) int {
	var bitmap int
	for i, wasVisited := range visited {
		if wasVisited {
			bitmap |= 1 << i
		}
	}
	return bitmap
}

// recheck part1 using part2 logic
func part1ViaPart2(input string) int {
	graph := makeGraph(input)

	roomToFlowRate := map[string]int{}
	highestFlowRatePossible := 0 // might be useful for an optimization later

	// sort room names so the arrays/slices later will be in a repeatable order (easier debugging)
	// include starting room "AA" in this list even though it has zero flow rate
	roomNames := []string{"AA"}
	for name, room := range graph {
		if room.flowRate != 0 {
			roomNames = append(roomNames, name)
		}
	}

	// reformat into arrays/slices so we don't have to do room name lookups
	flowRates := make([]int, len(roomNames))

	sort.Strings(roomNames)

	for i, name := range roomNames {
		roomToFlowRate[name] = graph[name].flowRate
		flowRates[i] = roomToFlowRate[name]
		highestFlowRatePossible += graph[name].flowRate
	}

	weightedGraph := makeWeightedGraph(graph, roomNames)

	visitedArrayToHighestPressureTotals := map[[16]bool]int{}
	dfsGetHighestPressureTotalsForEveryVisitedState(30, [16]bool{}, weightedGraph, flowRates, 0, 0, 0, visitedArrayToHighestPressureTotals)

	highest := 0
	for _, val := range visitedArrayToHighestPressureTotals {
		highest = mathy.MaxInt(highest, val)
	}
	return highest
}
