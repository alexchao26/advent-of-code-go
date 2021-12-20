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
	nums := parseInput(input)

	base := nums[0]
	for i := 1; i < len(nums); i++ {
		base = addNodes(base, nums[i])
	}

	return base.magnitude()
}

func part2(input string) int {
	nums := parseInput(input)

	var best int
	for i, n1 := range nums {
		for j, n2 := range nums {
			if i == j {
				continue
			}
			// these operations are destructive to the lists so make copies every time
			cp1, cp2 := copyList(n1), copyList(n2)

			n1PlusN2 := addNodes(cp1, cp2)
			mag := n1PlusN2.magnitude()
			if mag > best {
				best = mag
			}
			cp1, cp2 = copyList(n1), copyList(n2)

			n2PlusN1 := addNodes(cp2, cp1)
			mag = n2PlusN1.magnitude()
			if mag > best {
				best = mag
			}
		}
	}

	return best
}

// doubly linked list node
type node struct {
	val        int
	prev, next *node
	depth      int
}

func addNodes(n1, n2 *node) *node {
	// increase the depth on all nodes in n1 and n2 bc they will form a pair
	n := n1
	for n != nil {
		n.depth++
		n = n.next
	}
	n = n2
	for n != nil {
		n.depth++
		n = n.next
	}

	// point last of n1 to first of n2 AND in reverse
	lastNode1 := n1
	for lastNode1.next != nil {
		lastNode1 = lastNode1.next
	}
	lastNode1.next = n2
	n2.prev = lastNode1 // reverse

	n1 = n1.reduce()
	return n1
}

func copyList(n *node) *node {
	var head, last *node

	for p := n; p != nil; p = p.next {
		cp := &node{
			val:   p.val,
			prev:  last,
			next:  nil,
			depth: p.depth,
		}
		if head == nil {
			head = cp
			last = cp
		} else {
			last.next = cp
			last = cp
		}
	}

	return head
}

func (n *node) reduce() (head *node) {
	for pointer := n; pointer != nil; pointer = pointer.next {
		// nested inside 4 pairs means depth is 5 or more
		if pointer.depth >= 5 {
			// should explode

			pairRight := pointer.next // pointer to right node of pair
			if pairRight.depth != pointer.depth {
				panic(fmt.Sprintf("exploding pair should have same depth, got %d and %d", pointer.depth, pairRight.depth))
			}

			// pair will become a node with zero value
			replacement := &node{
				val:   0,
				depth: pointer.depth - 1,
				prev:  pointer.prev,   // may be nil
				next:  pairRight.next, // may be nil
			}

			if pointer.prev != nil {
				pointer.prev.val += pointer.val
				// reassignments to remove old pair
				pointer.prev.next = replacement
			}
			if pairRight.next != nil {
				pairRight.next.val += pairRight.val
				// reassignments to remove old pair
				pairRight.next.prev = replacement
			}

			// recursively call reduce on n again and RETURN to stop further reductions (reset logic)

			// edge case for head is exploding
			if n == pointer {
				return replacement.reduce()
			}

			return n.reduce()
		}
	}

	for pointer := n; pointer != nil; pointer = pointer.next {
		if pointer.val >= 10 {
			// should split

			replacementLeft := &node{
				val:   pointer.val / 2, // integer division will round down
				prev:  pointer.prev,
				next:  nil, // will be replacementRight
				depth: pointer.depth + 1,
			}
			replacementRight := &node{
				val:   pointer.val / 2, // need to round up
				prev:  replacementLeft,
				next:  pointer.next,
				depth: pointer.depth + 1,
			}

			// adjustments to inits
			replacementLeft.next = replacementRight
			if pointer.val%2 == 1 {
				replacementRight.val++
			}

			toLeft, toRight := pointer.prev, pointer.next
			if toLeft != nil {
				toLeft.next = replacementLeft
			}
			if toRight != nil {
				toRight.prev = replacementRight
			}

			// recursively call reduce on n again and RETURN to stop further reductions (reset logic)

			// edge case for head is splitting
			if n == pointer {
				return replacementLeft.reduce()
			}

			return n.reduce()
		}
	}

	// just return self if none of the reduce actions occurred
	return n
}

func (n *node) String() string {
	var sb strings.Builder

	for p := n; p != nil; p = p.next {
		sb.WriteString(fmt.Sprintf("v: %d, depth: %d -> ", p.val, p.depth))
	}

	// this would be prettier, but my brain hurts
	// for p := n; p != nil; p = p.next {
	// 	if p == n {
	// 		sb.WriteString(strings.Repeat("[", p.depth))
	// 		sb.WriteString(cast.ToString(p.val))
	// 		sb.WriteString(",")
	// 	} else {
	// 		if p.depth == p.prev.depth {
	// 			// right of a pair
	// 			sb.WriteString(cast.ToString(p.val))
	// 			sb.WriteString("]")
	// 		} else {

	// 		}

	// 		if p.next == nil  {
	// 			sb.WriteString(strings.Repeat("]", p.prev.depth - p.depth))
	// 		}
	// 	}
	// }

	return sb.String()
}
func (n *node) magnitude() int {
	// recursive calculation, 3*left + 2*right
	// regular numbers are just their number so [2,5] -> 3*2 + 2*5 = 16

	// make copy because calculation works by collapsing the list's pairs
	cp := copyList(n)
	for depth := 4; depth > 0; depth-- {

		for p := cp; p != nil; p = p.next {
			if p.depth == depth && p.next != nil && p.next.depth == depth {
				left, right := p, p.next
				newNode := node{
					val:   3*left.val + 2*right.val,
					prev:  left.prev,
					next:  right.next,
					depth: depth - 1,
				}
				if left == cp {
					cp = &newNode
				} else {
					left.prev.next = &newNode
				}
				if right.next != nil {
					right.next.prev = &newNode
				}
			}
		}
	}

	return cp.val
}

func parseInput(input string) []*node {
	var snailfishNums []*node
	for _, line := range strings.Split(input, "\n") {
		var depth int

		var pointer, head *node

		for _, r := range line {
			switch r {
			case '[':
				depth++
			case ']':
				depth--
			case ',': // do nothing

			default: // all single digit numbers
				newNode := node{
					val:   cast.ToInt(string(r)),
					prev:  pointer,
					next:  nil,
					depth: depth,
				}
				// assign head and pointer if none already
				if pointer == nil {
					head = &newNode
					pointer = &newNode
				} else {
					// otherwise assign pointer's next to new node, reassign pointer to new node
					pointer.next = &newNode
					pointer = &newNode
				}
			}
		}
		snailfishNums = append(snailfishNums, head)
	}
	return snailfishNums
}
