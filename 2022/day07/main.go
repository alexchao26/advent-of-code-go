package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
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

func part1(input string) int {
	root := parseInput(input)

	return sumDirsUnder100000(root)
}

func sumDirsUnder100000(itr *dir) int {
	SizeLimit := 100000

	sum := 0
	if itr.totalSize <= SizeLimit {
		sum += itr.totalSize
	}
	for _, child := range itr.childDirs {
		sum += sumDirsUnder100000(child)
	}
	return sum
}

func part2(input string) int {
	root := parseInput(input)
	totalSapceAvailable := 70000000
	spaceNeeded := 30000000

	// find smallest directory to be deleted that would free up enough space...
	directoryMinSize := spaceNeeded - (totalSapceAvailable - root.totalSize)
	return findSmallestDirToDelete(root, directoryMinSize)
}

func findSmallestDirToDelete(itr *dir, directoryMinSize int) int {
	smallest := math.MaxInt64
	if itr.totalSize >= directoryMinSize {
		smallest = mathy.MinInt(smallest, itr.totalSize)
	}

	for _, childDirs := range itr.childDirs {
		smallest = mathy.MinInt(smallest, findSmallestDirToDelete(childDirs, directoryMinSize))
	}

	return smallest
}

type dir struct {
	name      string
	parentDir *dir
	childDirs map[string]*dir
	files     map[string]int
	totalSize int
}

func parseInput(input string) *dir {
	root := &dir{
		name:      "root",
		childDirs: map[string]*dir{},
	}
	itr := root

	cmds := strings.Split(input, "\n")
	c := 0

	for c < len(cmds) {
		switch cmd := cmds[c]; cmd[0:1] {
		case "$":
			if cmd == "$ ls" {
				// just move on, we will assume we're always in an listing state
				c++
			} else {
				changeDir := strings.Split(cmd, "cd ")[1]
				changeDir = strings.TrimSpace(changeDir)
				if changeDir == ".." {
					itr = itr.parentDir
				} else {
					// if changeDir doesn't exist..
					if _, ok := itr.childDirs[changeDir]; !ok {
						itr.childDirs[changeDir] = &dir{
							name:      changeDir,
							parentDir: itr,
							childDirs: map[string]*dir{},
							files:     map[string]int{}}
					}

					itr = itr.childDirs[changeDir]
				}
				c++
			}
		default:
			// assume we're listing a dir's contents... add it
			if strings.HasPrefix(cmd, "dir") {
				childDirName := cmd[4:]
				if _, ok := itr.childDirs[childDirName]; !ok {
					itr.childDirs[childDirName] = &dir{
						name:      childDirName,
						parentDir: itr,
						childDirs: map[string]*dir{},
						files:     map[string]int{},
					}
				}
			} else {
				// file name
				parts := strings.Split(cmd, " ")
				itr.files[parts[0]] = cast.ToInt(parts[0])
			}
			c++
		}
	}

	populateFileSizes(root)
	return root
}

func populateFileSizes(itr *dir) int {
	totalSize := 0

	for _, childItr := range itr.childDirs {
		totalSize += populateFileSizes(childItr)
	}

	for _, sz := range itr.files {
		totalSize += sz
	}

	itr.totalSize = totalSize

	return totalSize

}
