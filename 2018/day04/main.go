package main

import (
	"flag"
	"fmt"
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

	ans := part1(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func part1(input string, part int) int {
	lines := strings.Split(input, "\n")

	// sort inputs by time stamp, string sorting is sufficient
	sort.Strings(lines)

	timeEntries := make([]entry, len(lines))
	for i, line := range lines {
		timeEntries[i] = makeEntry(line)
	}

	// find which guard has slept the most total time
	// then find which minute he is asleep at most frequently
	mapIDGuard := make(map[int]*guard)
	// for part 2, track each guard for the actual minute asleep
	mapIDToMinutesArray := make(map[int]*[60]int)
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
					mapIDGuard[lastGuardID] = &guard{}
				}
				// part 2 parsing
				if mapIDToMinutesArray[lastGuardID] == nil {
					mapIDToMinutesArray[lastGuardID] = &[60]int{}
				}

				mapIDGuard[lastGuardID].totalTimeAsleep += endTime - startTime
				for startTime < endTime {
					mapIDGuard[lastGuardID].minutesAsleep[startTime]++
					mapIDToMinutesArray[lastGuardID][startTime]++
					startTime++
				}
			}
		}
	}

	if part == 1 {
		// who sleeps the most
		var IDOfSleepiestGuard, bestMinute, highestFreq int
		for i, g := range mapIDGuard {
			if IDOfSleepiestGuard == 0 {
				IDOfSleepiestGuard = i
			}
			if g.totalTimeAsleep > mapIDGuard[IDOfSleepiestGuard].totalTimeAsleep {
				IDOfSleepiestGuard = i
			}
		}

		// find the minute they are the asleep the most
		for min, freq := range mapIDGuard[IDOfSleepiestGuard].minutesAsleep {
			if freq > highestFreq {
				bestMinute = min
				highestFreq = freq
			}
		}

		// print ID * time (minute)
		return IDOfSleepiestGuard * bestMinute
	}

	// part 2 stuff
	var highestFreq, ID, bestMinute int
	// find the minute they are the asleep the most
	for i, arr := range mapIDToMinutesArray {
		for min, freq := range arr {
			if freq > highestFreq {
				bestMinute = min
				highestFreq = freq
				ID = i
			}
		}
	}

	// print ID * time (minute)
	return ID * bestMinute
}

type entry struct {
	ID, year, month, day, minute int
	awake                        bool
}

func makeEntry(line string) entry {
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

type guard struct {
	totalTimeAsleep int
	minutesAsleep   [60]int
}
