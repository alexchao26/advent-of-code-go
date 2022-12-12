package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	// tail follows head logically if in the same row or column
	// if not in same row or column, always moves diagonally
	insts := parseInput(input)

	// start stacked at 0,0
	var head, tail [2]int

	// "normal" grid mapping...
	diffs := map[string][2]int{
		"U": {1, 0},
		"D": {-1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}

	visited := map[[2]int]bool{
		{0, 0}: true,
	}
	for _, inst := range insts {
		for inst.val > 0 {
			// move head
			diff := diffs[inst.dir]
			head[0] += diff[0] // row
			head[1] += diff[1] // col

			// update tail
			// if diff to row or col is > 1

			rowDiff := head[0] - tail[0]
			colDiff := head[1] - tail[1]

			// if either row or col diff is > 1, then that dimension HAS to move
			// additionally, if the other diff is not zero, it needs to be
			//   adjusted to move diagonally
			// note: the nested if blocks screwed me in part 2 because a longer
			//       rope can make coordinates off by 2 rows AND 2 cols
			if mathy.AbsInt(rowDiff) > 1 {
				/*  0 1 2
				H . T
				diff = head - tail = -2
				want to make tail (2) to (1), so add diff / 2

				T . H
				diff = 2 - 0 = 2
				tail (0) + 2/2 = 1, checks out still
				*/
				tail[0] += rowDiff / 2
				// account for diagonal adjustment, same math... add col diff
				if colDiff != 0 {
					tail[1] += colDiff
				}
			} else if mathy.AbsInt(colDiff) > 1 {
				tail[1] += colDiff / 2
				// account for diagonal adjustment, same math... add col diff
				if rowDiff != 0 {
					tail[0] += rowDiff
				}
			}

			// update where the tail has been...
			visited[tail] = true
			inst.val-- // one step at a time
		}
	}

	// return spots TAIL visited at least once, map[[2]int]bool
	return len(visited)
}

type inst struct {
	dir string
	val int
}

func parseInput(input string) (ans []inst) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, inst{
			dir: line[:1],
			val: cast.ToInt(line[2:]),
		})
	}
	return ans
}

func part2(input string) int {
	// oof, quite the refactor...
	insts := parseInput(input)

	rope := initRope(10)

	visited := map[[2]int]bool{}
	for _, inst := range insts {
		for inst.val > 0 {
			rope.moveOneSpace(inst.dir)

			// update where the tail has been...
			visited[rope.tail.coords] = true

			inst.val-- // one step at a time

			fmt.Println(inst, rope, len(visited))
		}
	}

	return len(visited)
}

type node struct {
	coords [2]int // row, col still
	next   *node
}

type rope struct {
	head, tail *node
}

func initRope(length int) rope {
	head := &node{}
	itr := head

	// start at 1 to account for head already being created
	for i := 1; i < length; i++ {
		itr.next = &node{}
		itr = itr.next
	}

	return rope{
		head: head,
		tail: itr,
	}
}

func (r rope) moveOneSpace(dir string) {
	// "normal" grid mapping...
	diffs := map[string][2]int{
		"U": {1, 0},
		"D": {-1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}

	diff := diffs[dir]
	r.head.coords[0] += diff[0]
	r.head.coords[1] += diff[1]

	// update rest of rope too
	r.head.updateTrailer()
}

func (r rope) String() string {
	str := ""
	i := 0
	for itr := r.head; itr != nil; itr = itr.next {
		str += fmt.Sprintf("%d:[%d,%d]->", i, itr.coords[0], itr.coords[1])
		i++
	}
	return str
}

// recursively updates the node behind itself as it follows
func (n *node) updateTrailer() {
	if n.next == nil {
		return
	}

	rowDiff := n.coords[0] - n.next.coords[0]
	colDiff := n.coords[1] - n.next.coords[1]

	// if either row or col diff is > 1, then that dimension HAS to move
	// additionally, if the other diff is not zero, it needs to be
	//   adjusted to move diagonally
	if mathy.AbsInt(rowDiff) > 1 && mathy.AbsInt(colDiff) > 1 {
		n.next.coords[0] += rowDiff / 2
		n.next.coords[1] += colDiff / 2
	} else if mathy.AbsInt(rowDiff) > 1 {
		// see part1 for math logic
		n.next.coords[0] += rowDiff / 2
		n.next.coords[1] += colDiff
	} else if mathy.AbsInt(colDiff) > 1 {
		n.next.coords[1] += colDiff / 2
		n.next.coords[0] += rowDiff
	} else {
		// no need to continue updating children if movement is over
		return
	}

	// go to next node
	n.next.updateTrailer()
}

func reImplPart1(input string) int {
	// oof, quite the refactor...
	insts := parseInput(input)

	rope := initRope(2)

	visited := map[[2]int]bool{}
	for _, inst := range insts {
		for inst.val > 0 {
			rope.moveOneSpace(inst.dir)

			// update where the tail has been...
			visited[rope.tail.coords] = true

			inst.val-- // one step at a time
		}
	}

	// return spots TAIL visited at least once, map[[2]int]bool
	return len(visited)
}
