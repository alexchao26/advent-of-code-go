package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathutil"
	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(util.ReadFile("./input.txt"))
		fmt.Println("Output:", ans)
	}
}

type TreeNode struct {
	children []*TreeNode
	metadata []int
}

func part1(input string) int {
	nums := parseInput(input)

	root, _ := makeTree(nums)

	return sumMetadataOnly(root)
}

func part2(input string) int {
	nums := parseInput(input)

	root, _ := makeTree(nums)

	return sumMetadataViaChildrenIndices(root)
}

func parseInput(input string) []int {
	split := strings.Split(input, " ")

	parsed := make([]int, len(split))

	for i, v := range split {
		parsed[i] = mathutil.StrToInt(v)
	}

	return parsed
}

func makeTree(nums []int) (node *TreeNode, recursiveValuesHandled int) {
	if len(nums) == 0 {
		return nil, 0
	}
	if len(nums) == 2 {
		return &TreeNode{nil, []int{}}, 2
	}

	childrenCount := nums[0]
	metadataCount := nums[1]

	newNode := TreeNode{}

	valuesHandled := 2
	for i := 2; childrenCount > 0 || metadataCount > 0; {
		if childrenCount > 0 {
			// recursively make child
			child, subValuesHandled := makeTree(nums[i:])

			newNode.children = append(newNode.children, child)
			valuesHandled += subValuesHandled

			childrenCount--

			i += subValuesHandled
		} else {
			newNode.metadata = append(newNode.metadata, nums[i])
			valuesHandled++

			metadataCount--
			i++
		}
	}

	return &newNode, valuesHandled
}

// part1
func sumMetadataOnly(node *TreeNode) int {
	var sumMetadata int

	for _, v := range node.metadata {
		sumMetadata += v
	}

	for _, child := range node.children {
		sumMetadata += sumMetadataOnly(child)
	}

	return sumMetadata
}

// part2
func sumMetadataViaChildrenIndices(node *TreeNode) int {
	var sumMetadata int
	if len(node.children) == 0 {
		for _, v := range node.metadata {
			sumMetadata += v
		}
		return sumMetadata
	}

	// ONE INDEXED
	for _, valAsChildOneIndex := range node.metadata {
		if valAsChildOneIndex == 0 || valAsChildOneIndex > len(node.children) {
			continue
		}

		sumMetadata += sumMetadataViaChildrenIndices(node.children[valAsChildOneIndex-1])
	}

	return sumMetadata
}
