package main

import (
	"github.com/alexchao26/advent-of-code-go/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadFile("../input.txt")
	lines := strings.Split(input, "\n")

	// sort inputs by time stamp, string sorting is sufficient
	sort.Strings(lines)

	timeEntries := make([]entry, len(lines))
	for i, line := range lines {
		timeEntries[i] = makeEntry(line)
	}

	// find which guard has slept the most total time
	// then find which minute he is asleep at most frequently
	mapIDGuard := make(map[int]*[60]int)
	lastGuardID := timeEntries[0].ID
	for i, timeEntry := range timeEntries {
		if timeEntry.ID != 0 {
			lastGuardID = timeEntry.ID
		} else {
			// if the time entry is awake, then check the next one entry
			// if the next entry is the same day, then assume its the same guard
			// update that guard's stats
			if !timeEntry.awake && i+1 != len(timeEntries) &&
				timeEntries[i+1].day == timeEntry.day {
				endTime := timeEntries[i+1].minute
				startTime := timeEntry.minute
				if mapIDGuard[lastGuardID] == nil {
					mapIDGuard[lastGuardID] = &[60]int{}
				}
				for startTime < endTime {
					mapIDGuard[lastGuardID][startTime]++
					startTime++
				}
			}
		}
	}

	var highestFreq, ID, bestMinute int
	// find the minute they are the asleep the most
	for i, arr := range mapIDGuard {
		for min, freq := range arr {
			if freq > highestFreq {
				bestMinute = min
				highestFreq = freq
				ID = i
			}
		}
	}

	// print ID * time (minute)
	fmt.Printf("Guard %v is asleep the most at minute %v.\ngithub.com/alexchao26/advent-of-code-go answer: %v\n",
		ID,
		bestMinute,
		ID*bestMinute,
	)
}

type entry struct {
	ID, year, month, day, minute int
	awake                        bool
}

func makeEntry(line string) entry {
	fmt.Println(line)
	var e entry
	e.year, _ = strconv.Atoi(line[1:5])
	e.month, _ = strconv.Atoi(line[6:8])
	e.day, _ = strconv.Atoi(line[9:11])

	hour, _ := strconv.Atoi(line[12:14])
	// if started before midnight, zero out minute value
	if hour != 0 {
		e.minute = 0
	} else {
		e.minute, _ = strconv.Atoi(line[15:17])
	}

	// if the instruction is that the guard has fallen asleep, leave awake=false
	if strings.Contains(line, "falls asleep") {
		return e
	}

	e.awake = true
	if strings.Contains(line, "Guard") {
		e.ID, _ = strconv.Atoi(line[strings.Index(line, "#")+1 : strings.Index(line, " begins")])
	}
	return e
}
