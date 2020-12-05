package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Reading file: %s", err)
	}

	if part == 1 {
		ans := part1(string(bytes))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(string(bytes))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	var best int
	for _, seatDetails := range strings.Split(input, "\n") {
		if len(seatDetails) == 0 {
			continue
		}
		rowLeft, rowRight := 0, 127
		for _, char := range seatDetails[0:6] {
			if char == 'F' {
				rowRight -= (rowRight - rowLeft + 1) / 2
			} else {
				rowLeft += (rowRight - rowLeft + 1) / 2
			}
		}
		if seatDetails[6] == 'B' {
			rowLeft = rowRight
		}

		colLeft, colRight := 0, 7
		for _, col := range seatDetails[7:9] {
			if col == 'L' {
				colRight -= (colRight - colLeft + 1) / 2
			} else {
				colLeft += (colRight - colLeft + 1) / 2
			}
		}

		if seatDetails[9] == 'R' {
			colLeft = colRight
		}

		id := rowLeft*8 + colLeft
		if id > best {
			best = id
		}
	}

	return best
}

func part2(input string) int {
	allSeatIDs := make(map[int]bool)

	for _, seatDetails := range strings.Split(input, "\n") {
		rowLeft, rowRight := 0, 127
		for _, char := range seatDetails[0:6] {
			if char == 'F' {
				rowRight -= (rowRight - rowLeft + 1) / 2
			} else {
				rowLeft += (rowRight - rowLeft + 1) / 2
			}
		}
		if seatDetails[6] == 'B' {
			rowLeft = rowRight
		}

		colLeft, colRight := 0, 7
		for _, col := range seatDetails[7:9] {
			if col == 'L' {
				colRight -= (colRight - colLeft + 1) / 2
			} else {
				colLeft += (colRight - colLeft + 1) / 2
			}
		}

		if seatDetails[9] == 'R' {
			colLeft = colRight
		}

		id := rowLeft*8 + colLeft
		allSeatIDs[id] = true
	}

	var mySeatID int
	for id := range allSeatIDs {
		if allSeatIDs[id] && allSeatIDs[id+2] && !allSeatIDs[id+1] {
			mySeatID = id + 1
			break
		}
	}

	return mySeatID
}
