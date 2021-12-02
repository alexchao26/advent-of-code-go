package aoc

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"

	"github.com/alexchao26/advent-of-code-go/util"
)

func GetPrompt(day, year int, cookie string) {
	fmt.Printf("fetching for day %d, year %d\n", day, year)

	// make the request
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	body := GetWithAOCCookie(url, cookie)

	// parse the dang html
	prompt := parseHTML(body)

	// write to file
	filename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d/prompt.md", year, day))
	WriteToFile(filename, []byte(prompt))

	fmt.Println("Wrote prompt to file: ", filename)

	fmt.Println("Done!")
}

// uses dfsHTML function once to get the class=day-desc html nodes, then parse
// the text inside of them
func parseHTML(htmlIn []byte) (promptOnly string) {
	strBuilder := strings.Builder{}
	node, _ := html.Parse(bytes.NewReader(htmlIn))

	dayDescNodes := dfsHTML(node, cbFindDayDescClass)
	dayDescNodesMap := map[*html.Node]bool{}

	for _, ddNode := range dayDescNodes {
		dayDescNodesMap[ddNode.(*html.Node)] = true
		dfsHTML(ddNode.(*html.Node), cbParseHTMLText(&strBuilder, dayDescNodesMap))
	}

	return strBuilder.String()
}

// function takes in a node and a callback that is run on each node
// callback returns a slice of interfaces which are returned by dfs
func dfsHTML(node *html.Node, cb func(*html.Node) []interface{}) []interface{} {
	var sli []interface{}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if cb(child) != nil {
			sli = append(sli, child)
		}
		sli = append(sli, dfsHTML(child, cb)...)
	}
	return sli
}

func cbFindDayDescClass(node *html.Node) []interface{} {
	for _, attr := range node.Attr {
		if attr.Key == "class" && attr.Val == "day-desc" {
			// fmt.Println("day-desc node found!")
			return []interface{}{node}
		}
	}

	return nil
}

func cbParseHTMLText(builder *strings.Builder, dayDescNodesMap map[*html.Node]bool) func(*html.Node) []interface{} {
	return func(node *html.Node) []interface{} {
		if node.Type == html.TextNode {
			builder.WriteString(node.Data)
		}
		if node.Parent != nil && dayDescNodesMap[node.Parent] {
			builder.WriteString("\n")
		}
		return nil
	}
}
