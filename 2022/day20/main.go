package main

import (
	_ "embed"
	"flag"
	"fmt"
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

const part2DecryptionKey = 811589153

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := mixList(input, 1, 1)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := mixList(input, 811589153, 10)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func mixList(input string, decryptionKey, mixes int) int {
	zeroNode, originalOrder := parseInput(input)

	for i := 0; i < len(originalOrder); i++ {
		originalOrder[i].val *= decryptionKey
	}

	for i := 0; i < mixes; i++ {
		for _, node := range originalOrder {
			node.move(len(originalOrder))
		}
	}

	// sum 3 vals from  value zero, 1000th, 2000th, 3000th away

	return getNodeXStepsAway(zeroNode, 1000, len(originalOrder)).val +
		getNodeXStepsAway(zeroNode, 2000, len(originalOrder)).val +
		getNodeXStepsAway(zeroNode, 3000, len(originalOrder)).val
}

func getNodeXStepsAway(node *llNode, steps int, listLength int) *llNode {
	if steps < 0 {
		panic("negative steps")
	}

	iter := node
	for steps > 0 {
		iter = iter.next
		steps--
	}
	return iter
}

type llNode struct {
	val        int
	prev, next *llNode
}

func (n *llNode) move(totalLength int) {
	steps := n.val
	// fmt.Println("before steps", steps, "total", totalLength)
	steps %= totalLength - 1

	if steps == 0 {
		// fmt.Println("zero steps")
		return
	}

	// find slot to fit into

	if steps < 0 {
		steps += (totalLength - 1)
	}
	// fmt.Println("modded steps", steps)

	oldPrev, oldNext := n.prev, n.next
	oldPrev.next = oldNext
	oldNext.prev = oldPrev

	iter := n
	for steps > 0 {
		// fmt.Println("steps left", steps, "iter", iter)
		iter = iter.next
		if iter == n {
			panic("repeat")
		}
		steps--
	}

	nextPrev, nextNext := iter, iter.next
	// fmt.Println("nextPrev & nextNext", nextPrev, nextNext)
	nextPrev.next = n
	n.prev = nextPrev
	nextNext.prev = n
	n.next = nextNext
}

func parseInput(input string) (zeroNode *llNode, originalOrder []*llNode) {
	nums := []int{}
	for _, line := range strings.Split(input, "\n") {
		nums = append(nums, cast.ToInt(line))
	}

	var head, iter *llNode
	for _, n := range nums {
		node := &llNode{
			val:  n,
			prev: iter,
		}
		if head == nil {
			head = node
			iter = node
		} else {
			iter.next = node
			iter = iter.next
		}

		if iter.val == 0 {
			zeroNode = iter
		}
		originalOrder = append(originalOrder, node)
	}

	head.prev = iter
	iter.next = head

	return zeroNode, originalOrder
}

// for debugging
func listToString(head *llNode, listLength int) string {
	var sb strings.Builder
	for listLength > 0 {
		sb.WriteString(cast.ToString(head.val) + ",")
		head = head.next
		listLength--
	}
	return sb.String()
}

func printList(head *llNode, listLength int) {
	fmt.Println(listToString(head, listLength))
}
